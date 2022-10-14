#!/bin/sh
# Install gf tools
# Path: scripts/install-gf.sh
# Date: 2022-10-14
# Author: LiDong

set -e
# 获取GOOS
GOOS=$(go env GOOS)
echo "当前操作系统为:$GOOS"
# 获取GOARCH
GOARCH=$(go env GOARCH)
echo "当前操作系统架构为:$GOARCH"
# 拼接下载地址
DOWNLOAD_URL="https://download.fastgit.org/gogf/gf/releases/latest/download/gf_${GOOS}_${GOARCH}"
echo "下载地址为:$DOWNLOAD_URL"
curl -L $DOWNLOAD_URL -o ./gf
# 赋予执行权限
chmod +x ./gf
# 执行 install
./gf install
# 删除下载的文件
# rm -rf ./gf