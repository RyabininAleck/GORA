package models

import (
	"GORAbackend/internal/config"
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
	GetPhoto() (Photo, error)
	DelPhoto(id uint) error
	LoadPhoto(Photo) error
}

type ServerInterface interface {
	Start(port string) error
}

func (a *App) UploadPhotoHandler(c echo.Context) error {
	// todo предусмотреть 413 Payload Too Large

	file, err := c.FormFile("image")
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	name := genFileName()       // todo cделать уникальным
	err = imgToFile(file, name) //todo name, generate name
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, "uploadPhotoHandler")
}

func (a *App) GetPhotoListHandler(c echo.Context) error {
	return c.String(http.StatusOK, "getPhotoListHandler")
}

func (a *App) GetPhotoHandler(c echo.Context) error {
	id := c.Param("id")

	// Делаем что-то с полученным идентификатором, например, возвращаем его в ответе
	return c.String(http.StatusOK, "getPhotoHandler: Photo ID: "+id)

}

func (a *App) DelPhotoHandler(c echo.Context) error {
	id := c.Param("id")

	// Делаем что-то с полученным идентификатором, например, возвращаем его в ответе
	return c.String(http.StatusOK, "delPhotoHandler: Photo ID: "+id)

}
func imgToFile(file *multipart.FileHeader, name string) error {

	// Открываем файл, который был загружен
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Открываем целевой файл для записи
	dst, err := os.Create("storage/images/" + name + ".jpg")
	if err != nil {
		return err
	}
	defer dst.Close()

	// Копируем содержимое файла из тела запроса в целевой файл
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	return nil
}

func genFileName() string {
	//это не уникальное значение.
	//Теоретически, при моей частоте процессора 5 GHz период одного такта равен ~0.1 нс,
	// возможно, что инструкция выполнятс 2 раза в одну наносекунду

	currentTime := time.Now().UnixNano()
	return strconv.FormatInt(currentTime, 10)
}
