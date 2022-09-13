#!/bin/sh
set -e
## 获取第一个参数为tag
TAG=$1
if [ -z "$TAG" ]; then
    echo "Usage: $0 <tag>"
    exit 1
fi
# 判断是否是tag
if ! echo "$TAG" | grep -q '^v[0-9]\+\.[0-9]\+\.[0-9]\+$'; then
    echo "Invalid tag: $TAG, must be in the form of vMAJOR.MINOR.PATCH,example: v1.0.0"
    exit 1
fi


# if ! [[ $TAG =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]; then
#     echo "Invalid tag: $TAG, must be in the form of vMAJOR.MINOR.PATCH,example: v1.0.0"
#     exit 1
# fi

# 检查当前是否为master分支
if [ "$(git rev-parse --abbrev-ref HEAD)" != "master" ]; then
    echo "You must be on the master branch to release"
    exit 1
fi
# 检查是否有未提交的修改
if ! git diff-index --quiet HEAD --; then
    echo "You have uncommitted changes, please commit or stash them"
    exit 1
fi
# 检查本地与远程是否有同步问题
if ! git diff-index --quiet --cached HEAD --; then
    echo "Your local repository is out of sync with remote, please sync first"
    exit 1
fi
# 推送到master
git push origin master
# 创建tag
git tag cool/$TAG
git tag cool-tools/$TAG
git tag modules/base/$TAG
git tag modules/demo/$TAG
git tag modules/dict/$TAG
git tag modules/space/$TAG
git tag contrib/drivers/sqlite/$TAG
git tag contrib/drivers/mysql/$TAG
git tag contrib/drivers/mssql/$TAG
git tag $TAG
# 提交tag
git push origin cool/$TAG
git push origin cool-tools/$TAG
git push origin modules/base/$TAG
git push origin modules/demo/$TAG
git push origin modules/dict/$TAG
git push origin modules/space/$TAG
git push origin contrib/drivers/sqlite/$TAG
git push origin contrib/drivers/mysql/$TAG
git push origin contrib/drivers/mssql/$TAG
git push origin $TAG
# 显示最新的tag
git describe --tags --abbrev=0

