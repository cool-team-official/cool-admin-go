#!/bin/sh

set -e
# 获取脚本所在目录
DIR=$(
    cd "$(dirname "$0")"
    pwd
)

find $DIR -name "go.mod" | while read line; do
    # 获取go.mod文件所在目录
    MOD_DIR=$(dirname $line)
    echo "MOD_DIR: $MOD_DIR"
    cd $MOD_DIR
    go mod tidy
done



