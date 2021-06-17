package service

import (
	"gin_practice/model"

	"github.com/jinzhu/gorm"
)

type TestSuggestionService struct{}

func (testSuggetionService TestSuggestionService) CreateTestSuggestion() string {
	return "ボタンを連打するのはいかがでしょうか"
}

func (TestSuggestionService TestSuggestionService) FindTestSuggestions() []model.TestSuggestion {

	var db, err = gorm.Open("sqlite3", "db.sqlite")
	if err != nil {
		panic(err)
	}

	var testSuggestions []model.TestSuggestion
	db.Order("created_at desc").Find(&testSuggestions)
	db.Close()

	return testSuggestions

}
