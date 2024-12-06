# Goloves: A Distributed ID Generation Service

**Goloves** is a distributed ID generation service inspired by Meituan's Leaf. This project aims to provide a high-performance, scalable, and flexible ID generation solution.

---

## Project Structure

```
Goloves/
├── cmd/                    # Application entry points
│   └── idgen/             
│       └── main.go        # Main application entry
├── internal/              # Private application code
│   ├── domain/           # Enterprise business rules
│   │   ├── entity/       # Core business entities
│   │   └── generator/    # ID generation algorithms
│   ├── usecase/          # Application business rules
│   │   └── idgen/        # ID generation use cases
│   ├── repository/       # Data access layer
│   │   └── segment/      # Segment mode storage
│   └── delivery/         # External interfaces
│       ├── http/         # HTTP handlers
│       └── grpc/         # gRPC handlers
├── pkg/                  # Public libraries
│   ├── errors/          # Error handling
│   └── utils/           # Utility functions
├── configs/             # Configuration files
├── api/                 # API definitions
│   ├── http/           # HTTP API specs
│   └── proto/          # gRPC proto files
├── deployments/         # Deployment configurations
│   ├── docker/         # Docker related files
│   └── kubernetes/     # Kubernetes manifests
├── test/               # Additional test files
│   └── benchmark/      # Performance tests
├── Makefile            # Build automation
├── Dockerfile          # Docker build file
├── go.mod              # Go modules file
└── README.md           # Project documentation
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
   git clone https://github.com/Terry-Yuxiang/goloves.git
   cd goloves
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Build the service:
   ```bash
   make build
   ```

4. Run the service:
   ```bash
   ./bin/idgen
   ```