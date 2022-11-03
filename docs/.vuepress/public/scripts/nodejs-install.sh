#!/bin/bash
# Install Node.js on linux x86_64

# 出错时停止
set -e

# 判断当前操作系统是否为Linux,如果不是则退出
if [ "$(uname)" != "Linux" ]; then
    echo "Error: This script only supports Linux."
    exit 1
fi

# 获取当前CPU架构
ARCH=$(uname -m)

# 判断当前CPU架构是否为x86_64,如果不是则退出
if [ "$ARCH" != "x86_64" ]; then
    echo "Error: This script only supports x86_64."
    exit 1
fi

# 判断当前是否为root,如果不是则退出
if [ "$(id -u)" != "0" ]; then
    echo "Error: This script must be run as root."
    echo "You can use 'sudo su' command switch to root "
    exit 1
fi

# 获取第一个参数为版本号
VERSION=$1

# 校验版本号
if [ -z "$VERSION" ]; then
    echo "Usage: $0 VERSION"
    echo "Example: $0 18.12.0"
    echo "You can visit https://nodejs.org/en/download/ to find version"
    exit 1
fi

# 校验版本号格式
if ! echo "$VERSION" | grep -qE '^[0-9]+\.[0-9]+\.[0-9]+$'; then
    echo "Invalid version: $VERSION, should be like 1.2.3"
    exit 1
fi

# 下载安装包
wget https://nodejs.org/dist/v$VERSION/node-v$VERSION-linux-x64.tar.xz

# 解压
tar -xvf node-v$VERSION-linux-x64.tar.xz

# 移动到 /usr/local
mv node-v$VERSION-linux-x64 /usr/local/node

# 添加到环境变量
echo 'export PATH=$PATH:/usr/local/node/bin' >>/etc/profile

# 使环境变量生效
source /etc/profile
# 删除安装包
rm node-v$VERSION-linux-x64.tar.xz
# 查看版本
node -v
npm -v
# 激活yarn
corepack enable
# 配置淘宝镜像
npm config set registry https://registry.npmmirror.com
# 查看配置
npm config list
# 提示安装成功
echo "Node.js $VERSION installed successfully."
# 提示用户重启终端
echo "Please restart your terminal to make the PATH changes effective."
echo "You can use 'source /etc/profile' command to make the PATH changes effective immediately."
