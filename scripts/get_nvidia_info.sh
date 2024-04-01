#!/bin/bash

if ! command -v nvidia-smi &> /dev/null
then
    echo "ERROR DRIVER_NOT_FOUND: NVIDIA驱动未安装。"
    exit 1
fi

# 使用nvidia-smi命令获取驱动信息
driver_info=$(nvidia-smi | grep "Driver Version")

# 使用awk提取版本号
driver_version=$(echo $driver_info | awk '{print $3}')
cuda_version=$(echo $driver_info| awk '{print $9}')

echo "NVIDIA 驱动版本号为: $driver_version"
echo "CUDA 版本号为: $cuda_version"
echo "GPU列表为:"
nvidia-smi | grep -Po '(?<=\| )[^|]*(?= \|.+\|.+\|)' | awk 'NR > 3 && (NR-4) % 3 == 0'
echo "========================="
echo "GPU显存占用为列表为:"
nvidia-smi | grep -Po '[0-9]+MiB / [0-9]+MiB'
echo "========================="
echo "GPU功率列表为:"
nvidia-smi | grep -Po '[0-9]+W / [0-9]+W'