package orm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func InitDB() *gorm.DB {
	host, isHostSet := os.LookupEnv("MYSQL_HOST")
	port, isPortSet := os.LookupEnv("MYSQL_PORT")
	username, isUsernameSet := os.LookupEnv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DB_NAME")
	if !isHostSet {
		host = "localhost"
	}
	if !isPortSet {
		port = "3306"
	}
	if !isUsernameSet {
		username = "root"
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbName)
	log.Printf("DSN = %s\n", dsn)
	x, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败")
	}
	log.Println("数据库连接成功")

	x.AutoMigrate(&User{})
	return x
}
