# 设置 Go 编译器参数
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
BINARY_NAME=golves

# Docker 相关变量
DOCKER_IMAGE=golves
DOCKER_TAG=latest

# 默认目标
.DEFAULT_GOAL := help

.PHONY: all build test clean run docker-build docker-run help

all: clean build test ## 清理、构建和测试

build: ## 构建应用
	$(GOBUILD) -o $(BINARY_NAME) ./cmd/main.go

test: ## 运行测试
	$(GOTEST) -v ./...

clean: ## 清理构建文件
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run: build ## 构建并运行应用
	./$(BINARY_NAME)

docker-build: ## 构建 Docker 镜像
	docker build -t $(DOCKER_IMAGE):$(DOCKER_TAG) .

docker-run: ## 运行 Docker 容器
	docker run -p 8080:8080 $(DOCKER_IMAGE):$(DOCKER_TAG)

lint: ## 运行代码检查
	golangci-lint run

mock: ## 生成 mock 文件
	mockgen -source=internal/domain/repository/repository.go -destination=internal/domain/repository/mock/repository_mock.go

migrate: ## 运行数据库迁移
	go run cmd/migrate/main.go

docker-compose-up: ## 启动所有依赖服务
	docker compose up -d

docker-compose-down: ## 停止所有依赖服务
	docker compose down

dev: docker-compose-up ## 启动开发环境
	@echo "启动开发环境..."
	$(MAKE) run

help: ## 显示帮助信息
	@echo "使用方法:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
