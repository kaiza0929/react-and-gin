package service

import (
	"gin_practice/model"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type InitService struct{}

func (initService InitService) Init() {

	var db, err = gorm.Open("sqlite3", "db.sqlite")
	if err != nil {
		panic(err)
	}

	/* マイグレート(テーブルがなければ自動作成) */
	/* 複数のモデルを引数に取れるか? */
	db.AutoMigrate(&model.TestLog{})
	defer db.Close()

}
