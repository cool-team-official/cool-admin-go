#!/bin/bash
# ============================================================
# 发布前版本号同步脚本(多模块 Monorepo 专用)
# 统一把以下内容更新到指定版本:
#   1. cool-tools/internal/cmd/version.go 中的 binVersion
#   2. 仓库根 go.mod 与所有子模块 go.mod 中对
#      github.com/cool-team-official/cool-admin-go/... 的内部依赖版本
# 用法: bash scripts/bump-version.sh v1.5.12
# 注意:执行前请确保工作区干净;执行后需人工 review 并提交
# ============================================================
set -euo pipefail

if [ $# -ne 1 ]; then
    echo "Usage: bash scripts/bump-version.sh <version>"
    echo "Example: bash scripts/bump-version.sh v1.5.12"
    exit 1
fi
version=$1

# 校验版本号格式 vMAJOR.MINOR.PATCH
if ! echo "$version" | grep -qE '^v[0-9]+\.[0-9]+\.[0-9]+$'; then
    echo "Invalid version: $version, must be vMAJOR.MINOR.PATCH, e.g. v1.5.12"
    exit 1
fi

# 仓库根目录
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
cd "$ROOT"

# 1. 更新 cool-tools 的 binVersion
VERSION_FILE="cool-tools/internal/cmd/version.go"
if [ -f "$VERSION_FILE" ]; then
    if [ "$(uname)" == "Darwin" ]; then
        sed -i '' -E "s/binVersion := \"v[0-9]+\.[0-9]+\.[0-9]+\"/binVersion := \"$version\"/" "$VERSION_FILE"
    else
        sed -i -E "s/binVersion := \"v[0-9]+\.[0-9]+\.[0-9]+\"/binVersion := \"$version\"/" "$VERSION_FILE"
    fi
    echo "updated: $VERSION_FILE -> $version"
else
    echo "WARN: $VERSION_FILE not found, skip."
fi

# 2. 同步所有 go.mod 中的内部模块依赖版本
#    形如: github.com/cool-team-official/cool-admin-go/cool v1.5.10
#    只会命中 require 行的版本号,不会误改 module 声明行
FILES="go.mod $(find . -name go.mod -not -path './.devcontainer/*')"
for f in $FILES; do
    if [ "$(uname)" == "Darwin" ]; then
        sed -i '' -E \
            "s#(github\.com/cool-team-official/cool-admin-go(/[A-Za-z0-9_./-]+)? )v[0-9]+\.[0-9]+\.[0-9]+#\1$version#g" \
            "$f"
    else
        sed -i -E \
            "s#(github\.com/cool-team-official/cool-admin-go(/[A-Za-z0-9_./-]+)? )v[0-9]+\.[0-9]+\.[0-9]+#\1$version#g" \
            "$f"
    fi
    echo "updated: $f"
done

echo "Done. Please review 'git diff' and commit."
