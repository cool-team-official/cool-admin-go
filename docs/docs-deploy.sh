#!/usr/bin/env sh
# This script deploys the documentation to the gh-pages branch.
# 发布文档到 https://cooladmingo.github.io

# 确保脚本抛出遇到的错误
set -e
# 检测是否存在 package.json,如果不存在，说明运行目录不对
if [ ! -f "package.json" ]; then
    echo "package.json not found, please run this script in the root directory of the project"
    exit 1
fi


# 生成静态文件
npm run docs:build

# 进入生成的文件夹
cd docs/.vuepress/dist

# 如果是发布到自定义域名
# echo 'www.example.com' > CNAME

# 获取当前时间
now=$(date "+%Y.%m.%d-%H.%M.%S")
echo "${now}" > version.txt


git init
git add -A
git commit -m 'deploy'

# 如果当前运行在 github codespace 中, 则使用 https 方式提交.否则使用 ssh 方式提交
if [ -n "$CODESPACES" ]; then
    echo "github codespace detected, using https to push"
    git push -f https://github.com/cool-team-official/cool-admin-go.git master:gh-pages
else
    echo "github codespace not detected, use ssh"
    git push -f git@github.com:cool-team-official/cool-admin-go.git master:gh-pages
fi

# 如果发布到 https://<USERNAME>.github.io
# git push -f git@github.com:cooladmingo/cooladmingo.github.io.git master:gh-pages
# git push -f https://github.com/cool-team-official/cool-admin-go.git master:gh-pages

# 如果发布到 https://<USERNAME>.github.io/<REPO>
# git push -f git@github.com:<USERNAME>/<REPO>.git master:gh-pages

cd -