package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type TestLog struct {
	gorm.Model
	/* Idだとduplicate column name: id  */
	Testid  string `gorm:"primaryKey"`
	Content string
	Result  string
	Tag     string
	Date    time.Time
}
