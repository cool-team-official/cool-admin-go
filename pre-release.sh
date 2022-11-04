#!/bin/bash
# 版本发布前的脚本,设置版本号,更新打包资源,更新相关依赖
# Usage: pre-release.sh [version]

# 出错时终止执行
set -e

# 读取版本号
if [ -z "$1" ]; then
    echo "Usage: pack-resource.sh [version]"
    exit 1
fi
version=$1

# 替换版本号 cool-tools/internal/cmd/version.go 中的 binVersion = "xxxx"
if [ "$(uname)" == "Darwin" ]; then
    sed -i '' -e "s/binVersion := \".*\"/binVersion := \"$version\"/g" cool-tools/internal/cmd/version.go
else
    sed -i -e "s/binVersion := \".*\"/binVersion := \"$version\"/g" cool-tools/internal/cmd/version.go
fi
# sed -i '' "s/binVersion := \".*\"/binVersion := \"$version\"/g" cool-tools/internal/cmd/version.go

# 进入脚本所在目录
cd "$(dirname "$0")"

# 进入cool-tools目录
cd cool-tools

# 如果当前环境为github codespace, 则使用https方式拉取, 否则使用ssh方式拉取
if [ -n "$CODESPACES" ]; then
    echo "github codespace detected, using https to pull"
    make pack.template-simple
    make pack.docs
else
    echo "github codespace not detected, use ssh"
    make pack.template-simple.ssh
    make pack.docs.ssh
fi
