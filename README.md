# Goloves: A Distributed ID Generation Service

**Goloves** is a distributed ID generation service inspired by Meituan's Leaf. This project aims to provide a high-performance, scalable, and flexible ID generation solution.

---

## Project Structure

```
Goloves/
├── cmd/                    # 启动服务的主程序
│   └── main.go
├── config/                 # 配置文件及加载
│   └── config.go
├── internal/               # 内部逻辑
│   ├── segment/            # Segment 模式实现
│   ├── snowflake/          # Snowflake 模式实现
│   ├── rpc/                # RPC 服务接口
│   └── util/               # 工具函数
├── tests/                  # 测试用例
├── docs/                   # 文档和说明
└── go.mod                  # 依赖管理
```
---

## Features

- **Segment Mode**: ID generation based on database segments for high performance and flexibility.
- **Snowflake Mode**: Twitter's Snowflake algorithm for globally unique IDs.
- **High Performance**: Supports QPS up to 50,000 with latency under 1ms (TP999).
- **Scalability**: Designed to scale horizontally with distributed deployments.
- **gRPC Support**: Provides a gRPC interface for ID generation.

---

## Getting Started

### Prerequisites

- Go 1.20+ installed
- MySQL database for Segment mode
- Optional: Docker for testing and deployment

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your_username/goloves.git
   cd goloves