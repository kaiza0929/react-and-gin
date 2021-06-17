package model

import (
	"github.com/jinzhu/gorm"
)

type TestSuggestion struct {
	gorm.Model
	SuggestionId string `gorm:"primaryKey"`
	TestId       string
	Content      string
}
