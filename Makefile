.PHONY: help
help: ## 查看帮助
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


.PHONY: cool-tools
cool-tools: ## 安装cool-tools
	@echo "install cool-tools"
	@go install github.com/cool-team-official/cool-admin-go/cool-tools@latest

# 安装gf
.PHONY: gf
gf:
	@echo "install gf"
	@set -e; \
	wget -O gf https://github.com/gogf/gf/releases/latest/download/gf_$(shell go env GOOS)_$(shell go env GOARCH) && \
	chmod +x gf && \
	./gf install -y && \
	rm ./gf

.PHONY: dev
dev: ## 开发模式启动
	@echo "开发模式启动"
	@cool-tools run main.go

.PHONY: install
install: ## 安装依赖
	@echo "安装依赖"
	@go mod tidy

.PHONY: build
build: ## 编译为二进制文件
	@echo "编译为二进制文件"
	@cool-tools build
# 复制前端文件到指定目录并打包
.PHONY: build.public
build.public:
	@set -e; \
	rm -rf $(ROOT_DIR)/temp/public;\
	mkdir -p $(ROOT_DIR)/temp/public;\
	cp -r ../cool-admin-vue/dist/* $(ROOT_DIR)/temp/public;\
	gf pack ./temp/public ./internal/packed/public.go -p resource/public
