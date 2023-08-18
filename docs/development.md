# 开发环境

[返回目录](README.md)

::: warning 注意
推荐使用 Linux 或 MacOS 进行开发，Windows 下可使用 WSL2。

Linux 及 WSL2 下推荐使用 root 用户进行开发.
:::

## Node.js 环境

官网下载地址：[https://nodejs.org/en/download/](https://nodejs.org/en/download/)

一般选择 LTS 版本即可。

MacOS 下可使用 Homebrew 进行安装 nvm：

```bash
brew install nvm
```

nvm 是 node 版本管理工具，可以通过`nvm install <version>` 安装指定版本，使用 `nvm use  <version>` 切换版本。

或者直接下载 pkg 安装包进行安装。

Linux 下可使用以下脚本进行安装：

```bash
wget -O nodejs-install.sh https://cool-team-official.github.io/cool-admin-go/scripts/nodejs-install.sh \
&& chmod +x nodejs-install.sh \
&& ./nodejs-install.sh 18.12.0
```

脚本文件内容如下:

<<< @/docs/.vuepress/public/scripts/nodejs-install.sh

::: tip
安装完成后，可使用`node -v`查看版本号，使用`npm -v`查看 npm 版本号。
为提高依赖包下载速度，可使用`npm config set registry https://registry.npmmirror.com`切换到淘宝镜像。
新版本的 node 已经集成了 yarn,需激活`corepack`,可使用 `corepack enable`命令激活。激活后可使用`yarn -v`查看版本号。

Linux 安装脚本已完成镜像切换及 corepack 激活。
:::

## Go 环境

官网下载地址：[https://go.dev/dl/](https://go.dev/dl/)

一般选择最新版本即可。

MacOS 下可使用 Homebrew 进行安装：

```bash
brew install go
```

或者直接下载 pkg 安装包进行安装。

Linux 下可使用以下脚本进行安装：

```bash
wget -O golang-install.sh https://cool-team-official.github.io/cool-admin-go/scripts/golang-install.sh \
&& chmod +x golang-install.sh \
&& ./golang-install.sh 1.19.3
```

脚本文件内容如下:

<<< @/docs/.vuepress/public/scripts/golang-install.sh

::: tip
安装完成后，可使用`go version`查看版本号。
为提高依赖下载速度，推荐配置`goproxy`，可使用`go env -w GOPROXY=https://goproxy.cn,direct`切换到 goproxy.cn 镜像。
:::

## VSCode

官网下载地址：[https://code.visualstudio.com/](https://code.visualstudio.com/)

一般选择最新版本即可。

推荐安装以下插件：

- [Go](https://marketplace.visualstudio.com/items?itemName=golang.go)
- [Vetur](https://marketplace.visualstudio.com/items?itemName=octref.vetur)
- [ESLint](https://marketplace.visualstudio.com/items?itemName=dbaeumer.vscode-eslint)
- [Prettier](https://marketplace.visualstudio.com/items?itemName=esbenp.prettier-vscode)
- [EditorConfig](https://marketplace.visualstudio.com/items?itemName=EditorConfig.EditorConfig)
- [GitLens](https://marketplace.visualstudio.com/items?itemName=eamodio.gitlens)

## Docker

云原生时代，Docker 已经成为开发者必备的工具之一。

开发过程中，我们将使用 Docker 进行数据库管理，以及打包测试。

官网下载地址：[https://www.docker.com/products/docker-desktop](https://www.docker.com/products/docker-desktop)

一般选择最新版本即可。

配置 Docker 镜像加速器：

```bash
# Linux
sudo mkdir -p /etc/docker
sudo tee /etc/docker/daemon.json <<-'EOF'
{
  "registry-mirrors": ["https://registry.docker-cn.com"]
}
EOF
sudo systemctl daemon-reload
sudo systemctl restart docker

```

MacOS 及 Windows 下可在 Docker Desktop 的设置中配置。
