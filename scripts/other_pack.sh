#!/bin/sh

if [ -z "$1" ]; then
  echo "请通过命令行参数指定要打包的目录"
  exit 1
fi

ROOT_DIR=$(pwd)
DIST_DIR="$1"

if [ ! -d "$DIST_DIR" ]; then
  echo "未找到指定的目录：$DIST_DIR"
  exit 1
fi

# 如果存在 data/public 目录,则删除
public_dir="$ROOT_DIR/data/public"

echo "🔎 正在检查 $public_dir 目录..."
if [ -d "$public_dir" ]; then
  echo "🗑️ 正在删除 $public_dir 目录"
  rm -rf "$public_dir"
  echo "✅ $public_dir 目录已删除"
else
  echo "✅ $public_dir 目录未找到"
fi

# 移动指定目录到data目录并重命名为public
echo "🚚 正在移动 $DIST_DIR 目录到 data 目录并重命名为 public"
mv "$DIST_DIR" "$public_dir"
echo "✅ 已移动并重命名 $DIST_DIR 目录"

# gf 打包
cd "$ROOT_DIR/data"
echo "📦 正在打包 public 目录"
gf pack public "$ROOT_DIR/internal/packed/public.go" -y
echo "✅ 打包完成"
