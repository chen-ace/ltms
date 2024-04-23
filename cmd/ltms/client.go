package main

import (
	"errors"
	"fmt"
	"llm_training_management_system/internal/models"
	"llm_training_management_system/internal/router"
	"llm_training_management_system/pkg/gpu"
	"llm_training_management_system/pkg/ltms_config"
	"llm_training_management_system/rpcs"
	"log"
	"net/rpc"
	"os/exec"
	"strings"
	"time"
)

var clientConfig ltms_config.ClientConfig
var r router.Router

func main() {
	clientConfig = ltms_config.ReadClientConfig()
	client, err := rpc.Dial("tcp", fmt.Sprintf("%s:%d", clientConfig.MasterHost, clientConfig.MasterPort))
	if err != nil {
		log.Fatal("Dialing:", err)
	}

	// 注册RPC命令
	r = router.Router{}
	r.Register("Hello", handleHello)
	r.Register("Echo", handleEcho)

	go SendHeartBeat(err, client)

}

// handleHello 处理服务端返回的Hello命令
func handleHello(request router.Request) router.Response {
	log.Println("Hello")
	return router.Response{}
}

// handleEcho 处理服务端返回的Echo命令
func handleEcho(request router.Request) router.Response {
	log.Println("Message from master : ", request.Get("message"))
	return router.Response{}
}

// handleShell 处理服务端返回的shell命令
func handleShell(name string, args []string) (string, bool) {
	// 执行shell命令
	result, err := executeShellCommand(name, args)
	if err != nil {
		log.Println("Error executing shell command:", err)
		return "", false
	}
	return result, true
}

// executeShellCommand 执行shell命令
func executeShellCommand(name string, args []string) (string, error) {
	cmd := exec.Command(name, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return strings.TrimSuffix(string(output), "\n"), nil
}

// handleOrder 处理服务端返回的命令
func handleOrder(orders *[]models.LLMOrder) (string, error) {
	// 从服务器获取命令
	// 执行命令
	for _, order := range *orders {
		if order.IsShell {
			log.Println("执行shell命令：", order.Name, " ", order.Args)
			response, ok := handleShell(order.Name, order.Args)
			if ok {
				log.Println("shell命令执行成功，返回为：", response)
				return response, nil
			} else {
				log.Println("shell命令执行失败")
				return "", errors.New("shell命令执行失败")
			}
		} else {
			log.Println("执行RPC命令：", order.Name, " ", order.Data)
			response, ok := r.Call(order.Name, order.Data)
			if ok {
				log.Println("命令执行成功，返回为：", response)
				return fmt.Sprintf("%v", response), nil
			} else {
				log.Println("命令执行失败")
				return "", errors.New("命令执行失败")
			}
		}
	}
	return "", nil
}

func SendHeartBeat(err error, client *rpc.Client) {
	for {
		gpuInfo, _ := gpu.GetNvidiaInfo()
		request := &rpcs.Request{Message: "Ping",
			GpuInfo:  gpuInfo,
			NodeId:   clientConfig.NodeId,
			NodeRank: clientConfig.NodeRank}
		response := new(rpcs.Response)
		err = client.Call("HeartbeatService.Beat", request, response)
		if err != nil {
			log.Println("Call error:", err)
		}
		handleOrder(response.Orders)
		time.Sleep(5 * time.Second) // 每5s发送一次心跳
		log.Printf("Server response: %s", response.Message)
	}
}
