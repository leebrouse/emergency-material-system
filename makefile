# ----------------------
# REST 和 gRPC 代码生成相关
# ----------------------

.PHONY: help gen-openapi gen-proto clean clean-bin deploy-up deploy-down deploy-restart deploy-ps deploy-logs deploy-all

help: ## 显示帮助信息
	@echo "可用命令:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-15s %s\n", $$1, $$2}'

gen-openapi: ## 生成 OpenAPI 代码 (从 YAML 生成 Go 类型和服务器接口)
	@echo "生成 OpenAPI 代码..."
	./backend/script/genopenapi.sh

gen-proto: ## 生成 gRPC 代码 (从 .proto 文件生成 Go gRPC 代码)
	@echo "生成 gRPC 代码..."
	./backend/script/genproto.sh

gen-all: gen-openapi gen-proto ## 生成所有代码 (OpenAPI + gRPC)

clean: clean-bin ## 清理生成的文件
	@echo "清理生成的代码文件..."
	rm -rf ./backend/internal/common/genopenapi/
	rm -rf ./backend/internal/common/genproto/

# ----------------------
# 服务编译相关
# ----------------------

.PHONY: build-auth build-stock build-dispatch build-statistics build-logistics build-bin clean-bin

build-auth: ## 编译 auth 服务
	@echo "编译 auth 服务..."
	cd backend/internal/auth && go build -o auth main.go

build-stock: ## 编译 stock 服务
	@echo "编译 stock 服务..."
	cd backend/internal/stock && go build -o stock main.go

build-dispatch: ## 编译 dispatch 服务
	@echo "编译 dispatch 服务..."
	cd backend/internal/dispatch && go build -o dispatch main.go

build-statistics: ## 编译 statistics 服务
	@echo "编译 statistics 服务..."
	cd backend/internal/statistics && go build -o statistics main.go

build-logistics: ## 编译 logistics 服务
	@echo "编译 logistics 服务..."
	cd backend/internal/logistics && go build -o logistics main.go

build-bin: build-auth build-stock build-dispatch build-statistics build-logistics ## 编译所有服务

clean-bin: ## 清理编译后的二进制文件
	@echo "清理编译文件..."
	rm -f backend/internal/auth/auth
	rm -f backend/internal/stock/stock
	rm -f backend/internal/dispatch/dispatch
	rm -f backend/internal/statistics/statistics
	rm -f backend/internal/logistics/logistics

# ----------------------
# 部署相关 (Docker Compose)
# ----------------------
DOCKER_COMPOSE_FILE := ./backend/deploy/docker-compose.yaml

deploy-up: ## 启动所有服务 (后台运行)
	@echo "启动 Docker 容器..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d

deploy-down: ## 停止并删除所有容器
	@echo "停止 Docker 容器..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) down

deploy-restart: deploy-down deploy-up ## 重启所有服务

deploy-ps: ## 查看服务运行状态
	docker-compose -f $(DOCKER_COMPOSE_FILE) ps

deploy-logs: ## 查看服务日志
	docker-compose -f $(DOCKER_COMPOSE_FILE) logs -f

deploy-all: build-bin ## 编译并启动所有服务
	@echo "编译并启动所有服务..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) up --build -d

# ----------------------
# 测试相关
# ----------------------
test:
	cd backend/test && \
	k6 run k6_load_test.js
