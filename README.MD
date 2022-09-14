# CoolAdminGo

[![Go Reference](https://pkg.go.dev/badge/github.com/cool-team-official/cool-admin-go.svg)](https://pkg.go.dev/github.com/cool-team-official/cool-admin-go)
[![Build Release](https://github.com/cool-team-official/cool-admin-go/actions/workflows/release.yml/badge.svg)](https://github.com/cool-team-official/cool-admin-go/actions/workflows/release.yml)

CoolAdmin 的 Go 版本后端

开发中，未完成

- [docs](https://pkg.go.dev/github.com/cool-team-official/cool-admin-go) - 开发文档
- [github](https://github.com/cool-team-official/cool-admin-go) - cool-admin-go 仓库
- [release](https://github.com/cool-team-official/cool-admin-go/releases) - cool-tools 编译版本下载
- [issues](https://github.com/cool-team-official/cool-admin-go/issues) - 问答区
- [CoolAdmin](https://cool-js.com) - CoolAdmin 官网
- [GoFrame](https://goframe.org) - GoFrame 官网

## TODO

- [x] 基础 CURD 结构搭建
- [x] model 层快速生成
- [x] service 层快速生成
- [x] controller 层快速生成
- [x] 中间件错误状态及错误码与 node 版本统一
- [x] 全局错误码统一
- [x] controller 层接口输入与 node 版本统一
- [ ] 优化开发工具安装体验
- [x] service 层逻辑补全
- [ ] base模块逻辑补全
- [ ] 其他项目需求

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

如果提示找不到命令,一般是因为没有将 GOPATH 环境变量配置到系统环境变量中,我们可以进入 GOPATH 目录下的 bin 目录下执行命令来将`cool-tools`安装到系统中

```bash
cd $(go env GOPATH)/bin

./cool-tools install
```

### 2.创建新项目

```bash
cool-tools init hello-world
```

### 3.安装`gf`工具

这一步需要有 wget 命令,晚点再调整成 curl 命令

```bash
cd hello-world
make cli
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
