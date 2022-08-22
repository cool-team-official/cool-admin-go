# Cool Admin Go Simple

## 开发工具安装

### 安装 gf 工具

如果已经安装不再重复安装

```bash
make cli.install
```

### 升级 gf 工具

不管是否安装都下载最新版本

```bash
make cli
```

### 安装 cool-tools

如果已经安装不再重复安装

```bash
make tools.install
```

### 升级 cool-tools

不管是否安装都下载最新版本

```bash
make tools
```

## 开发模式启动

先执行 `go mod tidy` 来更新模块依赖

```bash
go mod tidy
```

更新完依赖后, 可以执行 `gf run main.go` 来启动开发模式

```bash
gf run main.go
```
