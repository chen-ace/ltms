# LTMS: LLM Training Management System

[![Go](https://github.com/chen-ace/ltms/actions/workflows/go.yml/badge.svg)](https://github.com/chen-ace/ltms/actions/workflows/go.yml)

## 简介

LTMS 是一个开源项目，旨在为大型语言模型训练提供集群管理解决方案。它支持训练集群的管理、集群状态监控、训练数据集的管理以及训练任务的提交和调度。

## 特性

- **集群管理**: 简化集群的配置和扩展。
- **状态查看**: 实时监控集群的状态和性能指标。
- **数据集管理**: 高效管理和分发训练数据集。
- **任务提交**: 用户友好的界面用于提交和跟踪训练任务。

## 快速开始

1. 克隆仓库：
    ```bash
    git clone
    ```
2. 构建项目：
    ```bash
   cd ltms 
   make
    ```
3. 构建结果将在 `build` 文件夹中生成两个文件：`ltmsd` (服务端) 和 `ltms` (客户端)。

## 安装与卸载系统服务

你可以使用以下命令将 `ltmsd` 安装为系统服务：

```bash
./ltmsd install
```
如果你想卸载该服务，可以使用以下命令：
```bash
./ltmsd uninstall
```
如果你不想安装为系统服务，你也可以直接启动 ltmsd：
```bash
./ltmsd
```

## 管理系统服务

安装服务后，你可以使用 `systemctl` 命令来启动、停止或查看服务状态。例如：
```bash
sudo systemctl start ltmsd  #启动服务
sudo systemctl stop ltmsd   #停止服务
sudo systemctl status ltmsd #查看服务状态
journalctl -f -u ltmsd      #查看日志
```
## 服务端配置

配置可以参考项目中的 `server_config.yaml`和`client_config.yaml` 文件。程序在启动时会按照以下顺序寻找该配置文件：

1. 程序所在的目录
2. `/etc/ltms`

请确保在这些位置之一有有效的配置文件，以便程序能够正确运行。

## 文档

详细文档请参见 docs 文件夹。

## 贡献

我们欢迎所有形式的贡献，包括错误报告、功能请求和代码提交。请阅读 CONTRIBUTING.md 了解如何开始。

## 许可证

本项目采用 MIT 许可证。
