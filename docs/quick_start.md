# 快速开始

[返回目录](README.md)

## 工程目录

一个合理的工程目录可以让开发更加高效，`CoolAdminGo`项目推荐的工程目录如下：

```bash
# 假定项目名称为 cool-study

cool-study/
├── backend # 后端代码
├── cool-study # 主库，存放生产部署相关脚本等
├── frontend # 前端代码
└── mobile # 移动端代码
```

## 创建后端项目

```bash
# 创建工程目录
mkdir cool-study
# 进入工程目录
cd cool-study
# 创建后端代码目录
cool-tools init backend
# 进入后端代码目录
cd backend
# 安装依赖
go mod tidy
# 开发模式运行后端项目
gf run backend
```

## 创建前端项目

```bash
# 进入工程目录
cd cool-study
# 拉取前端代码
git clone https://github.com/cool-team-official/cool-admin-vue frontend
# 如果网络不好，可以使用国内镜像
git clone https://gitee.com/cool-team-official/cool-admin-vue frontend
# 进入前端代码目录
cd frontend
# 安装依赖
yarn
```
