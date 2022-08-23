# CoolAdminGo

CoolAdmin 的 Go 版本后端

开发中，未完成

## TODO

- [x] 基础 CURD 结构搭建
- [x] model 层快速生成
- [x] service 层快速生成
- [x] controller 层快速生成
- [x] 中间件错误状态及错误码与 node 版本统一
- [x] 全局错误码统一
- [x] controller 层接口输入与 node 版本统一
- [] service 层逻辑补全
- [] 其他项目需求

## 使用说明

首先下载安装 Go 语言环境，如果没有安装，请自行安装，如果已经安装，请自行配置环境变量

### 1.安装`cool-tools`开发工具

安装命令

```bash
go install github.com/cool-team-official/cool-admin-go/cool-tools@latest
```

验证

```bash
cool-tools version
```

### 2.创建新项目

```bash
cool-tools init hello-world
```

### 3.安装`gf`工具

```bash
cd hello-world
make cli.install
```

### 4.安装依赖

```bash
go mod tidy
```

### 5.启动开发环境

```bash
gf run main.go
```

enjoy it!
