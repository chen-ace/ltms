package main

import (
	"fmt"
	"llm_training_management_system/internal/models"
	"llm_training_management_system/pkg/gpu"
	"llm_training_management_system/pkg/ltms_config"
	"llm_training_management_system/rpcs"
	"log"
	"net/rpc"
	"time"
)

var clientConfig ltms_config.ClientConfig

func main() {
	clientConfig = ltms_config.ReadClientConfig()
	client, err := rpc.Dial("tcp", fmt.Sprintf("%s:%d", clientConfig.MasterHost, clientConfig.MasterPort))
	if err != nil {
		log.Fatal("Dialing:", err)
	}

	// 注册RPC命令

	go SendHeartBeat(err, client)

}

func handleOrder(order *[]models.LLMOrder) (string, error) {
	// 从服务器获取命令
	// 执行命令
	// 返回执行结果
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
