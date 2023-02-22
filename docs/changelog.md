# 更新日志

## 1.0.17

- list page 接口增加 ModifyAfter 钩子,用于对数据进行修改


## 1.0.16

- 更新依赖
- cool-tools创建的项目增加容器开发环境

## 1.0.15

- 更新 gf 至 v2.3.1

## 1.0.14

- 更新 gf 至 v2.3.0
- 调整 base 模块中的事务对象为接口定义以匹配 gf 的变更
- 引入 redis 库(gf v2.3.0 版本 redis 拆分为单独的库)
- 增加用户时对密码进行 md5 加密

## 1.0.13

- 增加 pgsql 支持
- 调整部分表字段类型以兼容 pgsql

## 1.0.12

- 更新 gf 版本至 v2.2.6(sql 统计中 total 由 int64 修改为 int)

## 1.0.11

- 修复 dict 模块的 bug

## 1.0.10

- 修复 base/menu/add 不支持数组菜单的问题

## 1.0.9

- GetCfgWithDefault 函数支持环境变量,优先级 配置文件>环境变量>默认值
- 更新 gf 版本至 v2.2.5 修复软删除 bug

## 1.0.8

- 更新依赖包版本
- 调整 cool-tools version 命令输出格式
- 修复 go install 模式安装的 cool-tools docs 命令无法使用的问题

## 1.0.7

- 更新 gf 依赖至 v2.2.4
- cool-tools 增加 -y 支持
- 集成 gf pack

## 1.0.6

- docs 独立为单独 mod,减小主库体积
- 清理部分无用文件,减小主库体积
- 更新依赖
- 主库移除内置的前端
- 增加 make frontend 命令,用于构建前端
- 权限中间件移除部分 Debug 日志
- 清理 cool-tools 模块部分无用文件
- 引入 gf 的 run 命令至 cool-tools 模块
- 引入 gf 的 build 命令至 cool-tools 模块
- 调整 cool-tools 使用 hack 目录下的配置文件

## 1.0.5

- 本次更新主要针对主库开发环境
- Added changelog.md
- 更新主框架开发用前端打包脚本 使用 make public 更新
- 更新主框架支持 remote container 开发, 基于`cool-admin-codespace`镜像
- 调整优化文档发布，更新了一些依赖
