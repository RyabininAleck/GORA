package models

import "time"

type Photo struct {
	ID        int64
	Name      string
	Time      int64
	Extension string
	ImgPath   string
	PrevPath  string
}

func PreparePhoto(name string, Extension string, pathToOriginal string, pathToPreview string) Photo {

	return Photo{0, name, time.Now().Unix(), Extension, pathToOriginal, pathToPreview}
}
