package gpu

import (
	"encoding/xml"
	"fmt"
	"os/exec"
	"regexp"
)

// GPU 结构体定义
type GPU struct {
	ID       string `xml:"minor_number"`
	Model    string `xml:"product_name"`
	MemUsage string `xml:"fb_memory_usage>used"`
	MemTotal string `xml:"fb_memory_usage>total"`
	MemFree  string `xml:"fb_memory_usage>free"`
	Power    string `xml:"gpu_power_readings>power_draw"`
	GPUUtil  string `xml:"utilization>gpu_util"`
}

// NvidiaSmi 结构体定义，用于匹配 nvidia-smi 输出的 XML 结构
type NvidiaSmi struct {
	XMLName       xml.Name `xml:"nvidia_smi_log"`
	DriverVersion string   `xml:"driver_version"`
	CUDAVersion   string   `xml:"cuda_version"`
	GPU           []GPU    `xml:"gpu"`
}

// GPUInfo 结构体定义
type GPUInfo struct {
	DriverVersion string
	CUDAVersion   string
	GPUs          []GPU
}

func GetNvidiaInfo() (GPUInfo, error) {
	cmd := exec.Command("nvidia-smi", "-q", "-x")
	output, err := cmd.Output()
	if err != nil {
		return GPUInfo{}, fmt.Errorf("nvidia-smi执行失败: %w", err)
	}

	// 使用正则表达式匹配驱动版本和CUDA版本
	driverVersionRegex := regexp.MustCompile(`<driver_version>(.*)</driver_version>`)
	cudaVersionRegex := regexp.MustCompile(`<cuda_version>(.*)</cuda_version>`)

	driverVersionMatch := driverVersionRegex.FindStringSubmatch(string(output))
	cudaVersionMatch := cudaVersionRegex.FindStringSubmatch(string(output))

	if driverVersionMatch == nil || cudaVersionMatch == nil {
		return GPUInfo{}, fmt.Errorf("unable to find driver or CUDA version in the output")
	}

	var smiOutput NvidiaSmi
	if err := xml.Unmarshal(output, &smiOutput); err != nil {
		return GPUInfo{}, fmt.Errorf("error unmarshalling XML: %w", err)
	}

	return GPUInfo{
		DriverVersion: smiOutput.DriverVersion,
		CUDAVersion:   smiOutput.CUDAVersion,
		GPUs:          smiOutput.GPU,
	}, nil
}
