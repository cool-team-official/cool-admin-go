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
- [x] 优化开发工具安装体验
- [x] service 层逻辑补全
- [x] base 模块逻辑补全 感谢 @imhuso
- [ ] 开发文档编制
- [ ] 数据库驱动补充
- [ ] 文件上传驱动补充
- [ ] 其他项目需求

## 使用说明

首先下载安装 Go 语言环境，如果没有安装，请自行安装，如果已经安装，请自行配置环境变量

### 1.安装`cool-tools`开发工具

从 github 下载

```bash
wget -O cool-tools https://github.com/cool-team-official/cool-admin-go/releases/latest/download/cool-tools_$(go env GOOS)_$(go env GOARCH) && chmod +x cool-tools && ./cool-tools install -y && rm ./cool-tools
```

从镜像下载

```bash
wget -O cool-tools https://gh.hjmcloud.cn/github.com/cool-team-official/cool-admin-go/releases/latest/download/cool-tools_$(go env GOOS)_$(go env GOARCH) && chmod +x cool-tools && ./cool-tools install -y && rm ./cool-tools
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

从 github 下载

```bash
wget -O gf https://github.com/gogf/gf/releases/latest/download/gf_$(go env GOOS)_$(go env GOARCH) && chmod +x gf && ./gf install -y && rm ./gf
```

使用镜像下载

```
wget -O gf https://gh.hjmcloud.cn/github.com/gogf/gf/releases/latest/download/gf_$(go env GOOS)_$(go env GOARCH) && chmod +x gf && ./gf install -y && rm ./gf
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
