package storage

import (
	"GORAbackend/internal/config"
	"GORAbackend/internal/models"
	"database/sql"
)

type SqliteStorage struct {
	*sql.DB
}

func GetStorage(cfg *config.Config) *SqliteStorage {
	// Открываем соединение с базой данных SQLite
	db, err := sql.Open("sqlite3", cfg.DataSourcePath)
	if err != nil {
		panic(err)
	}
	// todo	defer db.Close()

	return &SqliteStorage{db}
}

func (sqlite *SqliteStorage) GetPhotoList() ([]models.Photo, error) {

	return []models.Photo{}, nil
}
func (sqlite *SqliteStorage) GetPhoto() (models.Photo, error) {
	return models.Photo{}, nil
}
func (sqlite *SqliteStorage) DelPhoto(id uint) error {
	return nil
}
func (sqlite *SqliteStorage) LoadPhoto(models.Photo) error {
	return nil
}
