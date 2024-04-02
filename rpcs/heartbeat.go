package rpcs

import (
	"llm_training_management_system/internal/models"
	"llm_training_management_system/pkg/gpu"
	"llm_training_management_system/pkg/slaves"
	"log"
)

// HeartbeatService 定义了心跳服务的接口。
type HeartbeatService struct{}

// Request 是心跳请求的结构体。
type Request struct {
	Message  string
	GpuInfo  gpu.GPUInfo
	NodeId   string
	NodeRank int
}

// Response 是心跳响应的结构体。
type Response struct {
	Message string
	Orders  *[]models.LLMOrder
}

// Beat 是HeartbeatService的方法，用于处理心跳请求。
func (h *HeartbeatService) Beat(req *Request, res *Response) error {
	log.Println("收到心跳信息，来自", req.NodeId)
	llmOrders := slaves.HeartBeat(req.NodeId, &req.GpuInfo)

	res.Message = "Pong" // 简单地回应"Pong"作为心跳响应
	res.Orders = &llmOrders
	return nil
}
