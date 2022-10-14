#!/bin/sh
# Install cool tools
# Path: scripts/install-cool-tools.sh
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
DOWNLOAD_URL="https://download.fastgit.org/cool-team-official/cool-admin-go/releases/latest/download/cool-tools_${GOOS}_${GOARCH}"
echo "下载地址为:$DOWNLOAD_URL"
curl -L $DOWNLOAD_URL -o ./cool-tools
# 赋予执行权限
chmod +x ./cool-tools
# 执行 install
./cool-tools install
# 删除下载的文件
# rm -rf ./cool-tools