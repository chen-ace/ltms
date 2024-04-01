package ltms_config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type MysqlConfig struct {
	HOST   string `mapstructure:"mysql_host"`
	PORT   int    `mapstructure:"mysql_port"`
	DBNAME string `mapstructure:"mysql_db_name"`
	USER   string `mapstructure:"mysql_username"`
	PASS   string `mapstructure:"mysql_password"`
}

type RpcConfig struct {
	HOST string `mapstructure:"rpc_server_host"`
	PORT int    `mapstructure:"rpc_server_port"`
}

type HttpConfig struct {
	HOST    string `mapstructure:"http_server_host"`
	PORT    int    `mapstructure:"http_server_port"`
	DataDir string `mapstructure:"http_server_data_dir"`
}

func ReadServerConfig() (MysqlConfig, RpcConfig, HttpConfig) {
	viper.SetConfigType("yaml")
	viper.SetConfigName("server_config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/ltms/")
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
		fmt.Println("No .env file found, using environment variables or default values", err)
	}
	var mysqlConfig MysqlConfig
	if err := viper.UnmarshalKey("server_db", &mysqlConfig); err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	var rpcConfig RpcConfig
	if err := viper.UnmarshalKey("rpc_server", &rpcConfig); err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	var httpConfig HttpConfig
	if err := viper.UnmarshalKey("http_server", &httpConfig); err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	log.Println(viper.AllSettings())
	return mysqlConfig, rpcConfig, httpConfig
}
