@echo off
setlocal enabledelayedexpansion

if "%1" == "" (
  echo 请通过命令行参数指定要打包的目录
  exit /b 1
)

set ROOT_DIR=%CD%
set DIST_DIR=%1

if not exist "%DIST_DIR%" (
  echo 未找到指定的目录： %DIST_DIR%
  exit /b 1
)

set PUBLIC_DIR=%ROOT_DIR%\data\public

echo 🔎 正在检查 %PUBLIC_DIR% 目录...
if exist "%PUBLIC_DIR%" (
  echo 🗑️ 正在删除 %PUBLIC_DIR% 目录
  rmdir /s /q "%PUBLIC_DIR%"
  echo ✅ %PUBLIC_DIR% 目录已删除
) else (
  echo ✅ %PUBLIC_DIR% 目录未找到
)

echo 🚚 正在移动 %DIST_DIR% 目录到 data 目录并重命名为 public
move "%DIST_DIR%" "%PUBLIC_DIR%"
echo ✅ 已移动并重命名 %DIST_DIR% 目录

echo 📦 正在打包 public 目录
cd "%ROOT_DIR%\data"
gf pack public "%ROOT_DIR%\internal\packed\public.go" -y
echo ✅ 打包完成
