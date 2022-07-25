package common

import (
	"fmt"
	"log"
	"net/url"
	"server/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	host := viper.GetString("datasource.url")
	port := viper.GetString("datasource.port")
	dbname := viper.GetString("datasource.dbname")
	charset := viper.GetString("datasource.charset")
	local := viper.GetString("datasource.loc")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=%s", username, password, host, port, dbname, charset, url.QueryEscape(local))
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic("链接数据库失败:" + err.Error())
	}
	log.Println("链接数据库成功！")
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}
func GetDB() *gorm.DB {
	return DB
}
