.PHONY: help
help: ## 查看帮助
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'



.PHONY: cli
cli: ## 安装gf-cli
	@set -e; \
	wget -O gf https://github.com/gogf/gf/releases/latest/download/gf_$(shell go env GOOS)_$(shell go env GOARCH) && \
	chmod +x gf && \
	./gf install  && \
	rm ./gf


.PHONY: front
front: ## 下载最新cool-admin-vue,并调整参数编译
	bash scripts/frontend.sh

.PHONY: docs
docs: ## 打开pkgsite文档
	@set -e; \
	go install golang.org/x/pkgsite/cmd/pkgsite@latest;\
	echo "http://localhost:6060/github.com/cool-team-official/cool-admin-go";\
	pkgsite -http=localhost:6060

.PHONY: init
init: ## 初始化环境变量
	bash scripts/init.sh

.PHONY: dev
dev: ## 启动开发环境
	gf run main.go

.PHONY: clean
clean: ## 清理项目,用于删除开发容器及存储卷,需在本地开发环境执行
	@echo "清理项目"
	@bash ./scripts/clean.sh
	@echo "清理完成"
	
# 启动mysql
.PHONY: mysql-up
mysql-up: ## 启动mysql
	@echo "启动mysql"
	@docker-compose -f ./docker-compose.yml up -d mysql

# 停止mysql
.PHONY: mysql-down
mysql-down: ## 停止mysql
	@echo "停止mysql"
	@docker-compose -f ./docker-compose.yml down mysql

# 备份mysql
.PHONY: mysql-backup
mysql-backup: ## 备份mysql
	@echo "备份mysql"
	@bash ./scripts/mysql-backup.sh

# 启动redis
.PHONY: redis-up
redis-up: ## 启动redis
	@echo "启动redis"
	@docker-compose -f ./docker-compose.yml up -d redis

# 停止redis
.PHONY: redis-down
redis-down: ## 停止redis
	@echo "停止redis"
	@docker-compose -f ./docker-compose.yml down redis