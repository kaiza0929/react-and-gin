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
	Date    time.Time
}
