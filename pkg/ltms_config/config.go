package ltms_config

import (
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

type ClientConfig struct {
	MasterHost string `mapstructure:"master_host"`
	MasterPort int    `mapstructure:"master_port"`
	NodeId     string `mapstructure:"node_id"`
	NodeRank   int    `mapstructure:"node_rank"`
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
		log.Println("配置文件读取失败，请检查配置文件是否存在。配置文件应当放于/etc/ltms/目录下，或者程序所在目录下。")
	}
	var mysqlConfig MysqlConfig
	if err := viper.UnmarshalKey("server_db", &mysqlConfig); err != nil {
		log.Println("配置文件解析失败，请检查配置文件格式是否正确。")
	}

	var rpcConfig RpcConfig
	if err := viper.UnmarshalKey("rpc_server", &rpcConfig); err != nil {
		log.Println("配置文件解析失败，请检查配置文件格式是否正确。")
	}

	var httpConfig HttpConfig
	if err := viper.UnmarshalKey("http_server", &httpConfig); err != nil {
		log.Println("配置文件解析失败，请检查配置文件格式是否正确。")
	}

	log.Println("配置文件解析完成，使用如下配置：", viper.AllSettings())
	return mysqlConfig, rpcConfig, httpConfig
}

func ReadClientConfig() ClientConfig {
	viper.SetConfigType("yaml")
	viper.SetConfigName("client_config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/ltms/")
	// 设置默认值
	viper.SetDefault("master_host", "localhost")
	viper.SetDefault("master_port", 9332)
	// 读取环境变量（如果.env文件中不存在该变量）
	viper.AutomaticEnv()
	// 读取.env文件，如果文件不存在则忽略错误
	if err := viper.ReadInConfig(); err != nil {
		log.Println("配置文件读取失败，请检查配置文件是否存在。配置文件应当放于/etc/ltms/目录下，或者程序所在目录下。")
	}
	var clientConfig ClientConfig
	if err := viper.UnmarshalKey("client", &clientConfig); err != nil {
		log.Println("配置文件解析失败，请检查配置文件格式是否正确。")
	}
	log.Println("配置文件解析完成，使用如下配置：", viper.AllSettings())
	return clientConfig
}
