#!/bin/bash
###############################################################################
# Install Golang 安装Golang,仅限linux系统.                                      #
# Author: LiDong                                                              #
# Email: cnlidong@live.cn                                                     #
# Date: 2022-08-15                                                            #
###############################################################################
# 出错退出
set -e
# 进入脚本所在目录

# 获取第一个参数设为版本号
VERSION=$1
# 判断版本号参数是否为空
if [ -z "$VERSION" ]; then
    echo "Usage: $0 VERSION [mirror]"
    echo "Example: $0 1.19.2 https://mirrors.aliyun.com/golang"
    echo "You can visit https://go.dev/dl/ to find version"
    exit 1
fi
echo "version: $VERSION"
# 获取第二个参数为镜像地址前缀
PREFIX=$2
# 判断镜像地址前缀参数是否为空
if [ -z "$PREFIX" ]; then
    PREFIX="https://go.dev/dl"
fi
echo "prefix: $PREFIX"
# 判断当前用户是否为root，不是则退出
if [ "$(id -u)" != "0" ]; then
    echo "Error: This script must be run as root."
    echo "You can use 'sudo su' command switch to root "
    exit 1
fi
# 安装Golang
echo "Install Golang..."
# 获取当前操作系统
OS=$(uname)
# 判断是否是Linux系统,如果不是则退出
if [ $OS != "Linux" ]; then
    echo "Not Linux system, exit..."
    exit
else
    OS="linux"
fi
# 获取CPU类型
ARCH=$(uname -m)
# 转换CPU类型为go env arch格式
if [ $ARCH = "x86_64" ]; then
    ARCH="amd64"
elif [ $ARCH = "i686" ]; then
    ARCH="386"
elif [ $ARCH = "armv6l" ]; then
    ARCH="armv6l"
elif [ $ARCH = "aarch64" ]; then
    ARCH="arm64"
else
    echo "Not support CPU, exit..."
    exit
fi
# 安装Golang
echo "Download Golang..."
echo $PREFIX/go${VERSION}.$OS-$ARCH.tar.gz
curl -L $PREFIX/go${VERSION}.$OS-$ARCH.tar.gz -o /tmp/go${VERSION}.$OS-$ARCH.tar.gz
tar -C /usr/local -xzf /tmp/go${VERSION}.$OS-$ARCH.tar.gz
# 删除临时文件
rm -f /tmp/go${VERSION}.$OS-$ARCH.tar.gz
# 配置环境变量
echo 'export PATH=$PATH:/usr/local/node/bin' >>/etc/profile

echo 'export PATH=$PATH:/usr/local/go/bin' >>/etc/profile

source /etc/profile

echo 'export PATH=$PATH:$(go env GOPATH)/bin' >>/etc/profile
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct

source /etc/profile

# 安装成功
echo "Install Golang success!"
# 查看Golang版本
go version
# 提示重启终端
echo "Please restart the terminal to take effect!"
echo "安装成功,请重启终端使其生效!"
echo "You can use 'source /etc/profile' command to make the PATH changes effective immediately."
echo "你可以使用 'source /etc/profile' 命令使其立即生效."