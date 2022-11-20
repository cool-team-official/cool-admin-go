#!/bin/bash

# 出错退出
set -e

# 获取程序根目录
ROOT_DIR=$(pwd)
echo "ROOT_DIR: $ROOT_DIR"
# 如果运行目录不存在Makefile,则退出
if [ ! -f "Makefile" ]; then
    echo "Makefile not found, please run this script in the project root directory"
    echo "请使用 make frontend 命令在项目根目录运行此脚本"
    exit 1
fi
# 下载最新的前端代码并打包
if [ ! -d "$ROOT_DIR/data" ]; then
  echo "data directory not found, creating data directory"
  mkdir -p $ROOT_DIR/data
fi
cd $ROOT_DIR/data
# 如果已经存在前端代码,则删除
if [ -d "$ROOT_DIR/data/cool-admin-vue" ]; then
  echo "cool-admin-vue directory found, deleting cool-admin-vue directory"
  rm -rf $ROOT_DIR/data/cool-admin-vue
fi
# 如果当前为codespace开发环境，则使用git clone,否则使用pgit clone
if [ "$CODESPACES" = "true" ]; then
  echo "Cloning cool-admin-vue from github using git"
  git clone --depth=1 https://github.com/cool-team-official/cool-admin-vue.git
else
  echo "Cloning cool-admin-vue from github use pgit"
    pgit clone --depth=1 https://github.com/cool-team-official/cool-admin-vue.git
fi

# 进入前端代码目录
cd $ROOT_DIR/data/cool-admin-vue
# 替换 src/cool/config/index.ts 中的 mode: "history" 为 mode: "hash"
sed -i 's#mode: "history"#mode: "hash"#g' src/cool/config/index.ts
# 替换 src/cool/config/prod.ts 中的 baseUrl: "/api" 为 baseUrl: "/"
sed -i 's#baseUrl: "/api"#baseUrl: ""#g' src/cool/config/prod.ts

# 替换yarn.lock中的npm镜像地址
sed -i 's#https://registry.npmjs.org/#https://registry.npmmirror.com/#g' yarn.lock
# 安装前端依赖
echo "Installing front-end dependencies"
yarn install
# # 打包前端代码
echo "Building front-end code"
yarn build


# ## 如果存在 data/public 目录,则删除
# cd $ROOT_DIR
# if [ -d "data/public" ]; then
#   echo "$ROOT_DIR/data/public directory found, deleting $ROOT_DIR/data/public directory"
#   rm -rf $ROOT_DIR/data/public
# fi
# ## 移动dist目录到data目录并重命名为public
# echo "Moving dist directory to data directory and renaming it to public"
# mv $ROOT_DIR/data/cool-admin-vue/dist $ROOT_DIR/data/public
# ## gf 打包
# cd $ROOT_DIR/data
# echo "Packaging public directory"
# gf pack public $ROOT_DIR/internal/packed/public.go -y