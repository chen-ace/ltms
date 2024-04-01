package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kardianos/service"
	"gorm.io/gorm"
	"llm_training_management_system/internal/orm"
	"llm_training_management_system/pkg/dataset"
	"llm_training_management_system/pkg/ltms_config"
	"llm_training_management_system/pkg/slaves"
	"llm_training_management_system/rpcs"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

var db *gorm.DB
var userOrm *orm.UserOrm

type HttpResponse struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Result  string      `json:"result"`
	Data    interface{} `json:"data"`
}

var ErrorResponse = HttpResponse{"服务异常，请重试", 500, "", nil}

func ResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // 执行后续的处理函数

		// 从 gin.Context 中获取设置的响应数据
		err, exists := c.Get("err")
		if exists && err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse)
			log.Printf("出现错误,err=%s\n", err)
			return
		}

		resp, _ := c.Get("resp")
		data, _ := c.Get("data")

		var respStr string
		if resp != nil {
			respStr = resp.(string)
		}
		// 设置统一的 JSON 响应格式
		c.JSON(http.StatusOK, HttpResponse{
			Code:    http.StatusOK,
			Message: "",
			Result:  respStr,
			Data:    data,
		})
	}

}

type program struct{}

func (p *program) Start(s service.Service) error {
	// 在这里启动你的服务
	go p.run()
	return nil
}

var dbConfig ltms_config.MysqlConfig
var rpcConfig ltms_config.RpcConfig
var httpConfig ltms_config.HttpConfig

func (p *program) run() {
	// 你的主服务代码放在这里
	//db = orm.InitDB(dbConfig)
	//userOrm = orm.NewUserOrm(db)
	// 启动RPC服务
	go startRPCServer(&rpcConfig)
	log.Println("RPC服务启动成功，监听端口:", rpcConfig.PORT)
	// 启动HTTP服务
	startHttpServer(&httpConfig)
}

func (p *program) Stop(s service.Service) error {
	// 在这里停止你的服务
	fmt.Println("服务正在停止")
	return nil
}

func startHttpServer(config *ltms_config.HttpConfig) {
	r := setupRouter(config)
	// Listen and Server in 0.0.0.0:9080
	httpServerError := r.Run(fmt.Sprintf("%s:%d", config.HOST, config.PORT))
	if httpServerError != nil {
		log.Fatalf("服务启动失败，err=%s\n", httpServerError)
		return
	}
	log.Println("HTTP服务启动成功，监听端口:", "9080")
}

func setupRouter(config *ltms_config.HttpConfig) *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	r.MaxMultipartMemory = 5 << 10 << 10 << 10 // 5 GiB(每左移10位表示*1024）
	r.Use(ResponseMiddleware())
	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.Set("resp", "pong")
	})

	dataset.BindRouter(config, r)
	slaves.BindRouter(config, r)

	return r
}

func startRPCServer(config *ltms_config.RpcConfig) {
	beatService := new(rpcs.HeartbeatService)
	rpc.Register(beatService)
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.HOST, config.PORT))
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	log.Println("RPC is running")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Accept error:", err)
		}
		go rpc.ServeConn(conn)
	}
}

func main() {
	dbConfig, rpcConfig, httpConfig = ltms_config.ReadServerConfig()

	svcConfig := &service.Config{
		Name:        "ltmsd",
		DisplayName: "LTMS Service",
		Description: "LLM Training Management System Service",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) > 1 {
		err = service.Control(s, os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	err = s.Run()
	if err != nil {
		log.Fatal(err)
	}

}
