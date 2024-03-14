package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

// 定义结构体来存储GPU信息
type GPUInfo struct {
	DriverVersion string
	CUDAVersion   string
	GPUs          []GPU
}

type GPU struct {
	ID       string
	Model    string
	MemUsage string
	Power    string
}

// 函数执行脚本并返回GPUInfo结构体
func getNvidiaInfo() (*GPUInfo, error) {
	// 执行脚本
	cmd := exec.Command("/bin/bash", "get_nvidia_info.sh")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	// 准备正则表达式
	gpuNameRegex := regexp.MustCompile(`\d+\s+(NVIDIA.+)On`)
	memUsageRegex := regexp.MustCompile(`(\d+)MiB\s+/\s+(\d+)MiB`)
	powerUsageRegex := regexp.MustCompile(`(\d+)W\s+/\s+(\d+)W`)

	info := &GPUInfo{}
	info.GPUs = make([]GPU, 0)
	fmt.Println("LOADING , PLEASE WAIT~~~")
	// 读取输出
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "ERROR") {
			return nil, nil
		}

		// 解析驱动版本号
		if strings.Contains(line, "NVIDIA 驱动版本号为:") {
			info.DriverVersion = strings.Split(line, ": ")[1]
		}
		// 解析CUDA版本号
		if strings.Contains(line, "CUDA 版本号为:") {
			info.CUDAVersion = strings.Split(line, ": ")[1]
		}
		// 解析GPU列表
		if matches := gpuNameRegex.FindStringSubmatch(line); matches != nil {
			gpu := GPU{
				ID:    strings.Fields(matches[0])[0],
				Model: matches[1],
			}
			info.GPUs = append(info.GPUs, gpu)
		}
		// 解析显存占用和功率
		for i, gpu := range info.GPUs {
			if strings.HasPrefix(line, gpu.ID) {
				if matches := memUsageRegex.FindStringSubmatch(line); matches != nil {
					info.GPUs[i].MemUsage = matches[1] + "/" + matches[2]
				}
				if matches := powerUsageRegex.FindStringSubmatch(line); matches != nil {
					info.GPUs[i].Power = matches[1] + "/" + matches[2]
				}
			}
		}
	}

	err = cmd.Wait()
	if err != nil {
		return nil, err
	}

	return info, nil
}

func main() {
	// 获取GPU信息
	gpuInfo, err := getNvidiaInfo()
	if err != nil {
		fmt.Println("环境故障，无法获取")
	}
	if gpuInfo == nil {
		fmt.Printf("环境未安装！")
	} else {
		// 打印结构体内容
		fmt.Printf("驱动版本号: %s\n", gpuInfo.DriverVersion)
		fmt.Printf("CUDA版本号: %s\n", gpuInfo.CUDAVersion)
		for _, gpu := range gpuInfo.GPUs {
			fmt.Printf("GPU ID: %s, 型号: %s, 显存占用: %s, 功率: %s\n",
				gpu.ID, gpu.Model, gpu.MemUsage, gpu.Power)
		}
	}

}
