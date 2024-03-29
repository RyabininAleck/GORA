package config

// тут могло бы быть обращение к .env или к любому конфиг файлу.
// А можно вытягивать параметры из команды запуска (go run --config=config.json)
// но чтобы не переписывать сделаем мок
type Config struct {
	Port               string
	DataSourcePath     string
	ImgFileStorage     string
	PreviewFileStorage string
}

func GetConfig() *Config {
	return &Config{
		Port:               ":5123",
		DataSourcePath:     "./storage/sqlite.db",
		ImgFileStorage:     "storage/images/",
		PreviewFileStorage: "storage/preview/",
	}
}
