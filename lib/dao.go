package lib

import (
	"chat/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	err error
)

func init() {
	var dsn = config.Username + ":" + config.Password + "@tcp(" + config.Host + ":" +
		config.Port + ")/" + config.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {

	}
}
