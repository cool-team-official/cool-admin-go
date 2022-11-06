ROOT_DIR    = $(shell pwd)
NAMESPACE   = "default"
DEPLOY_NAME = "template-single"
DOCKER_NAME = "cool-admin-go-simple"

# 安装cool-tools
.PHONY: cool-tools
cool-tools:
	@echo "install cool-tools"
	@go install github.com/cool-team-official/cool-admin-go/cool-tools@master

# 安装gf
.PHONY: gf
gf:
	@echo "install gf"
	@set -e; \
	wget -O gf https://github.com/gogf/gf/releases/latest/download/gf_$(shell go env GOOS)_$(shell go env GOARCH) && \
	chmod +x gf && \
	./gf install -y && \
	rm ./gf

# 开发模式启动
.PHONY: run
run:
	@echo "开发模式启动"
	@go run main.go

# 安装依赖
.PHONY: install
install:
	@echo "安装依赖"
	@go mod tidy

# 编译为二进制文件
.PHONY: build
build:
	@echo "编译为二进制文件"
	@gf build
# 复制前端文件到指定目录并打包
.PHONY: build.public
build.public:
	@set -e; \
	rm -rf $(ROOT_DIR)/temp/public;\
	mkdir -p $(ROOT_DIR)/temp/public;\
	cp -r ../cool-admin-vue/dist/* $(ROOT_DIR)/temp/public;\
	gf pack ./temp/public ./internal/packed/public.go -p resource/public
