package service

import (
	"gin_practice/model"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type TestLogService struct{}

/* レシーバー(メソッドが構造体に属していることを示す) */
func (testLogService TestLogService) FindTestLogs() []model.TestLog {
	var db, err = gorm.Open("sqlite3", "db.sqlite")
	if err != nil {
		panic(err)
	}
	var testLogs []model.TestLog
	db.Order("created_at desc").Find(&testLogs)
	db.Close()
	return testLogs
}
