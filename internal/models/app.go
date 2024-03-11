package models

import (
	"GORAbackend/internal/config"
	"github.com/disintegration/imaging"
	"github.com/labstack/echo/v4"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"
)

type App struct {
	*config.Config
	DB DBInterface
	ServerInterface
}

type DBInterface interface {
	GetPhotoList() ([]Photo, error)
	GetPhoto(id int) (Photo, error)
	DelPhoto(id int) error
	LoadPhoto(Photo) (int64, error)
	Close() error
}

type ServerInterface interface {
	Start(port string) error
}

func (a *App) Stop() {
	a.DB.Close()
}
func (a *App) UploadPhotoHandler(c echo.Context) error {
	// todo предусмотреть 413 Payload Too Large
	// todo предусмотреть 422 неверный формат

	file, err := c.FormFile("image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	name := genFileName()

	pathToOriginal, err := imgToFile(name, file, a.Config.ImgFileStorage+name+".jpg")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	pathToPreview, err := previewToFile(name, pathToOriginal, a.Config.PreviewFileStorage+name+".jpg")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	data := PreparePhoto(name, "jpg", pathToOriginal, pathToPreview)
	data.ID, err = a.DB.LoadPhoto(data)

	return c.JSON(http.StatusOK, data)
}

func (a *App) GetPhotoListHandler(c echo.Context) error {

	photos, err := a.DB.GetPhotoList()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, photos)
}

func (a *App) GetPhotoHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	photo, err := a.DB.GetPhoto(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, photo)

}

func (a *App) DelPhotoHandler(c echo.Context) error {
	ParamID := c.Param("id")
	id, err := strconv.Atoi(ParamID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = a.DB.DelPhoto(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, ParamID)

}

func (a *App) ShowPhotoHandler(c echo.Context) error {
	ParamID := c.Param("id")
	id, err := strconv.Atoi(ParamID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	photo, err := a.DB.GetPhoto(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	image, err := os.Open(photo.ImgPath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer image.Close()

	return c.Stream(http.StatusOK, "image/jpeg", image)
}

func (a *App) ShowPreviewHandler(c echo.Context) error {
	ParamID := c.Param("id")
	id, err := strconv.Atoi(ParamID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	photo, err := a.DB.GetPhoto(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	preview, err := os.Open(photo.PrevPath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer preview.Close()

	return c.Stream(http.StatusOK, "image/jpeg", preview)

}

func imgToFile(name string, file *multipart.FileHeader, path string) (string, error) {

	original, err := file.Open()
	if err != nil {
		return "", err
	}
	defer original.Close()

	imgFile, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer imgFile.Close()

	if _, err = io.Copy(imgFile, original); err != nil {
		return "", err
	}

	return path, nil

}

func previewToFile(name, pathToOriginal, pathPreview string) (string, error) {

	previewFile, err := os.Create(pathPreview)
	if err != nil {
		return "", err
	}
	defer previewFile.Close()

	srcImage, err := imaging.Open(pathToOriginal)
	if err != nil {
		return "", err
	}

	dstImage128 := imaging.Resize(srcImage, 128, 128, imaging.Lanczos)

	err = imaging.Save(dstImage128, pathPreview)
	if err != nil {
		return "", err
	}
	return pathPreview, nil
}

func genFileName() string {
	// todo cделать уникальным
	//это не уникальное значение.
	//Теоретически, при моей частоте процессора 5 GHz период одного такта равен ~0.1 нс,
	// возможно, что инструкция выполнятся 2 раза в одну наносекунду🤡
	//можно сделать хэш

	currentTime := time.Now().UnixNano()
	return strconv.FormatInt(currentTime, 10)
}
