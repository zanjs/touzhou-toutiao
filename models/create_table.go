package models

import (
	"anla.io/taizhou-ir/db"
)

//CreateTable is init db table
func CreateTable() error {
	gorm.MysqlConn().AutoMigrate(&User{},
		&AppInfo{},
		&Article{},
		&ArticlePic{},
		&Category{},
		&Comment{})
	return nil
}
