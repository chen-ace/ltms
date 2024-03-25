package slaves

import (
	"llm_training_management_system/pkg/gpu"
	"time"
)

type SlaveNode struct {
	NodeId        string
	Gpus          *gpu.GPUInfo
	LastHeartBeat time.Time
}

var slaves = make(map[string]SlaveNode)

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
