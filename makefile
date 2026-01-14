# Makefile for emergency-material-system

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