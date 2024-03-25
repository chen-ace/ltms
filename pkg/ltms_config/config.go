package ltms_config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type MysqlConfig struct {
	HOST   string
	PORT   int
	DBNAME string
	USER   string
	PASS   string
}

type RpcConfig struct {
	HOST string
	PORT int
}

type HttpConfig struct {
	HOST    string
	PORT    int
	DataDir string
}

func ReadConfig() (MysqlConfig, RpcConfig, HttpConfig) {
	initViper()
	// 设置默认值
	viper.SetDefault("MYSQL_HOST", "localhost")
	viper.SetDefault("MYSQL_PORT", 3306)
	viper.SetDefault("MYSQL_DB_NAME", "ltms")
	viper.SetDefault("MYSQL_USERNAME", "root")
	viper.SetDefault("MYSQL_PASSWORD", "")

	viper.SetDefault("RPC_SERVER_HOST", "0.0.0.0")
	viper.SetDefault("RPC_SERVER_PORT", 9332)

	viper.SetDefault("HTTP_SERVER_HOST", "0.0.0.0")
	viper.SetDefault("HTTP_SERVER_PORT", 9582)
	viper.SetDefault("HTTP_SERVER_DATA_DIR", "./")
	// 读取环境变量（如果.env文件中不存在该变量）
	viper.AutomaticEnv()
	// 读取.env文件，如果文件不存在则忽略错误
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("No .env file found, using environment variables or default values")
	}
	mysqlConfig := MysqlConfig{
		HOST:   viper.GetString("MYSQL_HOST"),
		PORT:   viper.GetInt("MYSQL_PORT"),
		DBNAME: viper.GetString("MYSQL_DB_NAME"),
		USER:   viper.GetString("MYSQL_USERNAME"),
		PASS:   viper.GetString("MYSQL_PASSWORD"),
	}

	rpcConfig := RpcConfig{
		HOST: viper.GetString("RPC_SERVER_HOST"),
		PORT: viper.GetInt("RPC_SERVER_PORT"),
	}

	httpConfig := HttpConfig{
		HOST:    viper.GetString("HTTP_SERVER_HOST"),
		PORT:    viper.GetInt("HTTP_SERVER_PORT"),
		DataDir: viper.GetString("HTTP_SERVER_DATA_DIR"),
	}

	log.Println(viper.AllSettings())
	return mysqlConfig, rpcConfig, httpConfig
}

func initViper() {
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")

}
