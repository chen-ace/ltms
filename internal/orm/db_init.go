package orm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"llm_training_management_system/pkg/ltms_config"
	"log"
)

func InitDB(config ltms_config.MysqlConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.USER, config.PASS, config.HOST, config.PORT, config.DBNAME)
	log.Printf("DSN = %s\n", dsn)
	x, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败")
	}
	log.Println("数据库连接成功")

	x.AutoMigrate(&User{})
	return x
}
