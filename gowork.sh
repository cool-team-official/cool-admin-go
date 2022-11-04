#!/bin/bash
# 在指定目录初始化一个go 的 workspace, 并将所有子模块加入 go workspace
# 如果指定的目录不存在, 则创建目录
# 如果未指定目录, 则使用当前目录
# Usage: gowork.sh [dir]
# Author: LiDong
# Date: 2022-11-04

# 设置出错退出
set -e

# 设置默认目录
dir="."

# 如果有参数, 则使用参数作为目录
if [ $# -gt 0 ]; then
    dir=$1
fi

# 如果目录不存在, 则创建目录
if [ ! -d $dir ]; then
    mkdir -p $dir
fi

# 进入目录
cd $dir

# 如果目录下没有go.work, 则创建
if [ ! -f go.work ]; then
    go work init
fi

# 将目录下的所有子模块加入go workspace
go work add -r .
