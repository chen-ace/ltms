# LTMS: LLM Training Management System

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

## 文档

详细文档请参见 docs 文件夹。

## 贡献

我们欢迎所有形式的贡献，包括错误报告、功能请求和代码提交。请阅读 CONTRIBUTING.md 了解如何开始。

## 许可证

本项目采用 MIT 许可证。
