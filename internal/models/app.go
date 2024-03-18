package models

import (
	"GORAbackend/internal/config"
	"github.com/labstack/echo/v4"
	"net/http"
)

type App struct {
	*config.Config
	DB DBInterface
	ServerInterface
}

type DBInterface interface {
	AddCase(c *Case) (CaseId uint, err error)
	DelCase(caseId string) error
	UpdateCase(caseId string, c *Case) error // могу ли я заменить caseId и обновить чужое дело
	GetCases(id string) ([]Case, error)
	Close() error
}

type ServerInterface interface {
	Start(port string) error
}

func (a *App) Stop() {
	a.DB.Close()
}

func (a *App) GetRoutineHandler(c echo.Context) error {
	//apiV1.GET("/routine/:id", a.GetRoutineHandler)

	args := validate(c, "id")
	err := verification(args["id"], c.QueryParam("token"))
	if err != nil {
		return c.NoContent(http.StatusUnauthorized)
	}

	cases, err := a.DB.GetCases(args["id"])
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	return c.JSON(http.StatusOK, cases)
}

func (a *App) AddCase(c echo.Context) error {
	//apiV1.POST("/add/:id", a.AddCase)

	args := validate(c, "id")
	err := verification(args["id"], c.QueryParam("token"))
	if err != nil {
		return c.NoContent(http.StatusUnauthorized)
	}

	UserCase := new(Case)
	if err := c.Bind(UserCase); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	// todo
	UserCase.CaseId, err = a.DB.AddCase(UserCase)

	return c.JSON(http.StatusOK, UserCase)
}

func (a *App) DeleteCase(c echo.Context) error {
	//apiV1.DELETE("/delete/:id/:case_id", a.DeleteCase)

	args := validate(c, "id", "case_id")
	err := verification(args["id"], c.QueryParam("token"))
	if err != nil {
		return c.NoContent(http.StatusUnauthorized)
	}

	err = a.DB.DelCase(args["case_id"])
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	return c.NoContent(http.StatusOK)
}

func (a *App) UpdateCase(c echo.Context) error {
	//apiV1.POST("/update/:id/:case_id", a.UpdateCase)

	args := validate(c, "id", "case_id")
	err := verification(args["id"], c.QueryParam("token"))
	if err != nil {
		return c.NoContent(http.StatusUnauthorized)
	}

	UserCase := new(Case)
	if err := c.Bind(UserCase); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	// todo
	err = a.DB.UpdateCase(args["case_id"], UserCase)

	return c.JSON(http.StatusOK, UserCase)
}

func (a *App) GetSoundHandler(c echo.Context) error {
	// static хранилице
	//apiV1.GET("/sound/:path", a.GetSoundHandler)
	//args := validate(c, "path")

	return c.NoContent(http.StatusOK)
}

func validate(c echo.Context, args ...string) map[string]string {
	m := make(map[string]string)
	for i := range args {
		arg := c.Param(args[i])
		m[args[i]] = arg
	}

	return m
}

func verification(id string, token string) error {
	//todo
	return nil
}

//func (a *App) UploadPhotoHandler(c echo.Context) error {
//	// todo предусмотреть 413 Payload Too Large
//	// todo предусмотреть 422 неверный формат
//
//	file, err := c.FormFile("image")
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, err)
//	}
//
//	name := genFileName()
//
//	pathToOriginal, err := imgToFile(name, file, a.Config.ImgFileStorage+name+".jpg")
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, err)
//	}
//
//	data := PreparePhoto(name, "jpg", pathToOriginal)
//	data.ID, err = a.DB.LoadPhoto(data)
//
//	return c.JSON(http.StatusOK, data)
//}
//
//func (a *App) GetPhotoListHandler(c echo.Context) error {
//
//	photos, err := a.DB.GetPhotoList()
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, err)
//	}
//
//	return c.JSON(http.StatusOK, photos)
//}
//
//func (a *App) GetPhotoHandler(c echo.Context) error {
//	id, err := strconv.Atoi(c.Param("id"))
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, err)
//	}
//
//	photo, err := a.DB.GetPhoto(id)
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, err)
//	}
//
//	return c.JSON(http.StatusOK, photo)
//
//}
//
//func (a *App) DelPhotoHandler(c echo.Context) error {
//	ParamID := c.Param("id")
//	id, err := strconv.Atoi(ParamID)
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, err)
//	}
//
//	err = a.DB.DelPhoto(id)
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, err)
//	}
//
//	return c.JSON(http.StatusOK, ParamID)
//
//}
//
//func (a *App) ShowPhotoHandler(c echo.Context) error {
//	ParamID := c.Param("id")
//	id, err := strconv.Atoi(ParamID)
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, err)
//	}
//
//	photo, err := a.DB.GetPhoto(id)
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, err)
//	}
//
//	image, err := os.Open(photo.ImgPath)
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, err)
//	}
//	defer image.Close()
//
//	return c.Stream(http.StatusOK, "image/jpeg", image)
//}
//
//func (a *App) ShowPreviewHandler(c echo.Context) error {
//	ParamID := c.Param("id")
//	id, err := strconv.Atoi(ParamID)
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, err)
//	}
//
//	return c.JSON(http.StatusOK, id)
//
//}
//
//func imgToFile(name string, file *multipart.FileHeader, path string) (string, error) {
//
//	original, err := file.Open()
//	if err != nil {
//		return "", err
//	}
//	defer original.Close()
//
//	imgFile, err := os.Create(path)
//	if err != nil {
//		return "", err
//	}
//	defer imgFile.Close()
//
//	if _, err = io.Copy(imgFile, original); err != nil {
//		return "", err
//	}
//
//	return path, nil
//
//}
//
//func genFileName() string {
//	// todo cделать уникальным
//	//это не уникальное значение.
//	//Теоретически, при моей частоте процессора 5 GHz период одного такта равен ~0.1 нс,
//	// возможно, что инструкция выполнятся 2 раза в одну наносекунду🤡
//	//можно сделать хэш
//
//	currentTime := time.Now().UnixNano()
//	return strconv.FormatInt(currentTime, 10)
//}
