package main

import (
	"llm_training_management_system/rpcs"
	"log"
	"net/rpc"
	"time"
)

func main() {
	// 获取GPU信息
	//gpuInfo, err := GetNvidiaInfo()
	//if err != nil {
	//	fmt.Println(err)
	//}
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Dialing:", err)
	}

	request := &rpcs.Request{Message: "Ping"}
	response := new(rpcs.Response)

	for {
		err = client.Call("HeartbeatService.Beat", request, response)
		if err != nil {
			log.Fatal("Call error:", err)
		}
		log.Printf("Server response: %s", response.Message)
		time.Sleep(3 * time.Second) // 每分钟发送一次心跳
	}
}
