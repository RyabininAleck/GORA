package main

import (
	"GORAbackend/internal/config"
	"GORAbackend/internal/models"
	"GORAbackend/internal/server"
	"GORAbackend/internal/storage"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	cfg := config.GetConfig()
	srv := server.GetServer(cfg)
	db := storage.GetStorage(cfg)
	App := models.App{
		Config:          cfg,
		DB:              db,
		ServerInterface: srv,
	}

	//todo if err
	App.Start(cfg.Port)

}
