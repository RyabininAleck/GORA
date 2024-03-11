package models

import "GORAbackend/internal/config"

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
