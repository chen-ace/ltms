package rpcs

import "llm_training_management_system/pkg/gpu"

// HeartbeatService 定义了心跳服务的接口。
type HeartbeatService struct{}

// Request 是心跳请求的结构体。
type Request struct {
	Message string
	GpuInfo gpu.GPUInfo
}

// Response 是心跳响应的结构体。
type Response struct {
	Message string
}

// Beat 是HeartbeatService的方法，用于处理心跳请求。
func (h *HeartbeatService) Beat(req *Request, res *Response) error {
	res.Message = "Pong" // 简单地回应"Pong"作为心跳响应
	return nil
}
