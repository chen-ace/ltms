package slaves

import (
	"llm_training_management_system/pkg/gpu"
	"llm_training_management_system/rpcs"
	"time"
)

type SlaveNode struct {
	NodeId        string
	Gpus          *gpu.GPUInfo
	LastHeartBeat time.Time
}

type taskElement struct {
}

var slaves = make(map[string]SlaveNode)
var tasks = make(map[string][]*rpcs.LLMOrder)

func HeartBeat(nodeId string, gpu *gpu.GPUInfo) {
	node := SlaveNode{
		NodeId:        nodeId,
		Gpus:          gpu,
		LastHeartBeat: time.Now(),
	}
	slaves[nodeId] = node
}

func ListAllSlaves() []SlaveNode {
	var values []SlaveNode
	for _, node := range slaves {
		values = append(values, node)
	}
	return values
}

func PushLLMOrder(nodeId string, order *rpcs.LLMOrder) {
	tasks[nodeId] = append(tasks[nodeId], order)
}
