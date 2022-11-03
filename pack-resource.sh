#!/bin/bash
# 更新cool-tools中打包的资源文件
# Usage: pack-resource.sh [version]

# 出错时终止执行
set -e

# 读取版本号
if [ -z "$1" ]; then
    echo "Usage: pack-resource.sh [version]"
    exit 1
fi
version=$1-dev

# 替换版本号 cool-tools/internal/cmd/version.go 中的 binVersion = "xxxx"
sed -i '' "s/binVersion := \".*\"/binVersion := \"$version\"/g" cool-tools/internal/cmd/version.go

# 进入脚本所在目录
cd "$(dirname "$0")"

# 进入cool-tools目录
cd cool-tools

# 更新资源文件
make pack.template-simpl.ssh
make pack.docs.ssh