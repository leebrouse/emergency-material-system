# ----------------------
# REST 和 gRPC 代码生成相关
# ----------------------

.PHONY: help gen-openapi gen-proto clean

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

clean: ## 清理生成的文件
	@echo "清理生成的代码文件..."
	rm -rf ./backend/internal/common/genopenapi/
	rm -rf ./backend/internal/common/genproto/

# ----------------------
# 服务启动相关
# ----------------------

.PHONY: run-auth run-stock run-dispatch run-statistics run-logistics run-all

run-auth: ## 启动 auth 服务
	cd backend/internal/auth && go run main.go

run-stock: ## 启动 stock 服务
	cd backend/internal/stock && go run main.go

run-dispatch: ## 启动 dispatch 服务
	cd backend/internal/dispatch && go run main.go

run-statistics: ## 启动 statistics 服务
	cd backend/internal/statistics && go run main.go

run-logistics: ## 启动 logistics 服务
	cd backend/internal/logistics && go run main.go

# 一次性后台启动所有服务
run-all: ## 一次性启动所有后端服务
	$(MAKE) run-auth & \
	$(MAKE) run-stock & \
	$(MAKE) run-dispatch & \
	$(MAKE) run-statistics & \
	$(MAKE) run-logistics & \
	wait

# ----------------------
# 测试相关
# ----------------------
test:
	cd backend/test && \
	k6 run k6_load_test.js
