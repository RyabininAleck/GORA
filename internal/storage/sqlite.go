package storage

import (
	"GORAbackend/internal/config"
	"GORAbackend/internal/models"
	"database/sql"
	"fmt"
	"os"
)

const (
	sqlLoadPhoto = "INSERT INTO photo (image_path, preview_path, name, type, uploaded) VALUES (?, ?, ?, ?, ?)"
	sqlDelPhoto  = "DELETE FROM photo WHERE id = ?;"
	sqlGetPhotos = "SELECT id, image_path, preview_path, name, type, uploaded FROM photo"
	sqlGetPhoto  = "SELECT id, image_path, preview_path, name, type, uploaded FROM photo WHERE id = ?"
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
	photos := []models.Photo{}

	rows, err := sqlite.Query(sqlGetPhotos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tmp := models.Photo{}
	for rows.Next() {

		err = rows.Scan(&tmp.ID, &tmp.ImgPath, &tmp.PrevPath, &tmp.Name, &tmp.Extension, &tmp.Time)
		if err != nil {
			return nil, err //continue
		}
		photos = append(photos, tmp)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return photos, nil
}
func (sqlite *SqliteStorage) GetPhoto(id int) (models.Photo, error) {
	p := models.Photo{}

	row := sqlite.QueryRow(sqlGetPhoto, id)
	err := row.Scan(&p.ID, &p.ImgPath, &p.PrevPath, &p.Name, &p.Extension, &p.Time)
	if err != nil {
		return p, err
	}
	return p, nil
}
func (sqlite *SqliteStorage) DelPhoto(id int) error {
	photo, err := sqlite.GetPhoto(id)
	if err != nil {
		return err
	}

	stmt, err := sqlite.Prepare("DELETE FROM photo WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	err = os.Remove(photo.ImgPath)
	if err != nil {
		fmt.Println("Ошибка при удалении файла:", err)
		return err
	}

	err = os.Remove(photo.PrevPath)
	if err != nil {
		fmt.Println("Ошибка при удалении файла:", err)
		return err
	}

	return nil
}
func (sqlite *SqliteStorage) LoadPhoto(p models.Photo) (int64, error) {

	stmt, err := sqlite.Prepare(sqlLoadPhoto)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(p.ImgPath, p.PrevPath, p.Name, p.Extension, p.Time)
	if err != nil {
		return 0, err
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastInsertID, nil
}
