package slaves

import (
	"container/list"
	"llm_training_management_system/internal/models"
	"llm_training_management_system/pkg/gpu"
	//"llm_training_management_system/rpcs"
	"sync"
	"time"
)

var slaves = make(map[string]SlaveNode)

// HeartBeat 处理Slave的心跳信息
func HeartBeat(nodeId string, gpu *gpu.GPUInfo) []models.LLMOrder {
	node := SlaveNode{
		NodeId:        nodeId,
		Gpus:          gpu,
		LastHeartBeat: time.Now(),
	}
	// 此处可能发生并发问题，但是概率不大，即使出现，也不会对系统造成影响，所以为了性能，此处不加锁
	slaves[nodeId] = node
	return GetLLMOrders(nodeId)
}

// ListAllSlaves 此处可能出现脏读，但影响不大，因此为了性能，此处不加读写锁
func ListAllSlaves() []SlaveNode {
	var values []SlaveNode
	for _, node := range slaves {
		values = append(values, node)
	}
	return values
}

// 影响系统，因此此处使用严格的并发控制
var tasks = sync.Map{}

type SlaveNode struct {
	NodeId        string
	Gpus          *gpu.GPUInfo
	LastHeartBeat time.Time
}

type taskElement struct {
	mu    *sync.RWMutex
	tasks list.List
}

// PushLLMOrder 不保证顺序。
// 需要有序执行命令，请参考 PushLLMOrders 保证偏序
func PushLLMOrder(nodeId string, order *models.LLMOrder) {
	PushLLMOrders(nodeId, []*models.LLMOrder{order})
}

// PushLLMOrders 只保证偏序，不保证全序
func PushLLMOrders(nodeId string, orders []*models.LLMOrder) {
	element, _ := tasks.LoadOrStore(nodeId, &taskElement{
		tasks: *list.New(),
		mu:    new(sync.RWMutex),
	})

	task := element.(*taskElement)
	task.mu.Lock()
	for i := range orders {
		task.tasks.PushBack(orders[i])
	}
	task.mu.Unlock()
}

// GetLLMOrders 获取某个节点的所有累积的命令
func GetLLMOrders(nodeId string) []models.LLMOrder {
	result := []models.LLMOrder{}
	element, ok := tasks.Load(nodeId)
	if !ok {
		return result
	}
	task := element.(*taskElement)
	task.mu.Lock()
	defer task.mu.Unlock()
	for e := task.tasks.Front(); e != nil; e = e.Next() {
		result = append(result, e.Value.(models.LLMOrder))
	}
	return result
}
