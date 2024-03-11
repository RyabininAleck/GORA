package models

import "time"

type Photo struct {
	Name      string
	Data      time.Time
	Extension string
	Img       []byte
}
