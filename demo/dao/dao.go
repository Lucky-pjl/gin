package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init(){
	DB, _ = gorm.Open("mysql", "root:@(localhost)/test?charset=utf8mb4&parseTime=True&loc=Local")
}

