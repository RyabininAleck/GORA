package migration

import (
	"GORAbackend/internal/config"
	"GORAbackend/internal/storage"
	"os"
)

// можно через github.com/golang-migrate/migrate/cmd/migrate
func Migration(cfg *config.Config) {
	db := storage.GetStorage(cfg)
	// Читаем содержимое файла запроса
	query, err := os.ReadFile("internal/storage/migration/q.sql")
	if err != nil {
		panic(err)
	}

	// Выполняем запрос
	_, err = db.Exec(string(query))
	if err != nil {
		panic(err)
	}
}
