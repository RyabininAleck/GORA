package main

import (
	"GORAbackend/internal/config"
	"GORAbackend/internal/models"
	"GORAbackend/internal/server"
	"GORAbackend/internal/storage"
	"GORAbackend/internal/storage/migration"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	App := &models.App{}
	cfg := config.GetConfig()
	db := storage.GetStorage(cfg)
	migration.Migration(cfg)
	srv := server.GetServer(cfg, App)

	App.Config = cfg
	App.DB = db
	App.ServerInterface = srv

	err := App.Start(cfg.Port)
	if err != nil {
		App.Stop()
	}

}
