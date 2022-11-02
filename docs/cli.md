# 开发工具

[返回目录](README.md)

::: warning 注意
以下部分命令需您已经安装了 Go 语言环境，如果没有安装，请自行安装，如果已经安装，请自行配置环境变量
:::

## cool-tools

cool-tools 是一个用于快速生成`CoolAdminGo`项目的脚手架工具，可用于快速生成项目、模块、页面、接口等。

### `cool-tools`安装

Linux, Mac 可使用以下命令安装：

从 github 下载

```bash
wget -O cool-tools \
https://github.com/cool-team-official/cool-admin-go/releases/latest/download/cool-tools_$(go env GOOS)_$(go env GOARCH) \
&& chmod +x cool-tools \
&& ./cool-tools install \
&& rm ./cool-tools
```

从镜像下载

```bash
wget -O cool-tools \
https://gh.hjmcloud.cn/github.com/cool-team-official/cool-admin-go/releases/latest/download/cool-tools_$(go env GOOS)_$(go env GOARCH) \
&& chmod +x cool-tools \
&& ./cool-tools install \
&& rm ./cool-tools
```

验证

```bash
cool-tools version
```

Windows 可以直接下载编译后的可执行文件，下载地址：[releases](https://github.com/cool-team-official/cool-admin-go/releases),选择对应的版本下载。下载后复制到`PATH`环境变量中的目录下即可。

## gf

`GoFrame`框架提供了功能强大的`gf`命令行开发辅助工具，是框架发展的一个重要组成部分。

### `gf`安装

Linux, Mac 可使用以下命令安装：

从 github 下载

```bash
wget -O gf  \
https://github.com/gogf/gf/releases/latest/download/gf_$(go env GOOS)_$(go env GOARCH) \
&& chmod +x gf \
&& ./gf install \
&& rm ./gf
```

使用镜像下载

```
wget -O gf \
https://gh.hjmcloud.cn/github.com/gogf/gf/releases/latest/download/gf_$(go env GOOS)_$(go env GOARCH) \
&& chmod +x gf \
&& ./gf install \
&& rm ./gf
```
