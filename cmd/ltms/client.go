package main

import (
	"llm_training_management_system/pkg/gpu"
	"llm_training_management_system/rpcs"
	"log"
	"net/rpc"
	"time"
)

func main() {
	// 获取GPU信息
	//if err != nil {
	//	fmt.Println(err)
	//}
	client, err := rpc.Dial("tcp", "localhost:9332")
	if err != nil {
		log.Fatal("Dialing:", err)
	}

	// 注册RPC命令

	go SendHeartBeat(err, client)

}

//func handleOrder(order *rpcs.LLMOrder) (string, error) {
//	// 从服务器获取命令
//	// 执行命令
//	// 返回执行结果
//
//}

func SendHeartBeat(err error, client *rpc.Client) {
	for {
		gpuInfo, _ := gpu.GetNvidiaInfo()
		request := &rpcs.Request{Message: "Ping", GpuInfo: gpuInfo, NodeId: "node1"}
		response := new(rpcs.Response)
		err = client.Call("HeartbeatService.Beat", request, response)
		if err != nil {
			log.Println("Call error:", err)
		}
		time.Sleep(5 * time.Second) // 每5s发送一次心跳
		log.Printf("Server response: %s", response.Message)
	}
}
