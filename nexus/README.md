
# golang private server

使用Athens搭建golang私服，但Athens私服无管理界面，而nexus3集成了go,maven,docker等资源的私服功能

![image](https://user-images.githubusercontent.com/1940588/75411600-c4afb680-595a-11ea-9c5c-c4d5e5fa79c1.png)

## startup

```bash
docker-compose up
```

1. 访问 http://localhost:3001/ 用户名:admin 密码: 见`nexus-data/admin.password`文件。
1. 新建type为proxy的go repository（添加5个加速镜像）

    - 阿里云 https://mirrors.aliyun.com/goproxy/
    - 七牛 Goproxy 中国  https://goproxy.cn
    - 全球 CDN 加速  https://goproxy.io
    - jfrog   https://gocenter.io
    - Athens https://athens.azurefd.net

1. 新建type为group的go repository，名字为goproxy。

    group版的golang repository可以从proxy go repository下载依赖并缓存到本地，将左侧Avaliable中可用的repository加入到右侧的Menbers中，
    这样就可以从 `http://localhost:3001/repository/goproxy/` 中直接下载依赖，nexus会自动帮我们从proxy go repository中下载依赖。

1. 设置golang代理

    完成上诉步骤之后，还需要设置环境变量启用golang的代理功能，不同操作系统的设置方式可自行修改，将变量GO111MODULE设置为on，
    GOPROXY设置为私服的地 `http://localhost:3001/repository/goproxy/` ，
    若是遇到401 Unauthorized的问题，应该是需要进行nexus3的用户验证，可以直接在代理地址中加入用户名密码，
    例如 `http://username:password@localhost:3001/repository/goproxy/` 。

    若是不想自行搭建私服，也可使用一些现成的镜像站。

1. 更多帮助请见[sonatype go-repositories](https://help.sonatype.com/repomanager3/formats/go-repositories)

![image](https://user-images.githubusercontent.com/1940588/75412237-d09c7800-595c-11ea-8717-65a4b0beef10.png)

## FAQ

1. 默认用户名密码 [sonatype/nexus3](https://hub.docker.com/r/sonatype/nexus3/)

    Default user is `admin` and the uniquely generated password can be found in the `admin.password` file inside the volume.

## GOPROXY

众所周知，国内网络访问国外资源经常会出现不稳定的情况。 Go 生态系统中有着许多中国 Gopher 们无法获取的模块，比如最著名的 golang.org/x/...。并且在中国大陆从 GitHub 获取模块的速度也有点慢。

### 设置代理

在 Linux 或 macOS 上面，需要运行下面命令（或者，可以把以下命令写到 .bashrc 或 .bash_profile 文件中）：

```bash
# 启用 Go Modules 功能
$ go env -w GO111MODULE=on

# 配置 GOPROXY 环境变量，以下三选一

# 1. 官方
$ go env -w  GOPROXY=https://goproxy.io

# 2. 七牛 CDN
$ go env -w  GOPROXY=https://goproxy.cn

# 3. 阿里云
$ go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/

# 确认一下：

$ go env | grep GOPROXY
GOPROXY="https://goproxy.io"
```

在 Windows 上，需要运行下面命令：

```bash
# 启用 Go Modules 功能
$env:GO111MODULE="on"

# 配置 GOPROXY 环境变量，以下三选一，首推阿里云

# 1. 阿里云
$env:GOPROXY="https://mirrors.aliyun.com/goproxy/"

# 2. 官方
$env:GOPROXY="https://goproxy.io"

# 3. 七牛 CDN
$env:GOPROXY="https://goproxy.cn"

测试一下
$ time go get -v golang.org/x/tour

```

本地如果有模块缓存，可以使用命令清空 `go clean --modcache` 。

私有模块

如果你使用的 Go 版本 >=1.13, 你可以通过设置 GOPRIVATE 环境变量来控制哪些私有仓库和依赖 (公司内部仓库) 不通过 proxy 来拉取，直接走本地，设置如下：

```bash
# Go version >= 1.13
go env -w GOPROXY=https://goproxy.io,direct
# 设置不走 proxy 的私有仓库，多个用逗号相隔
go env -w GOPRIVATE=*.corp.example.com
```

## docker 启动 nexus 出现的问题

1. `docker run -d -p 8082:8082 -v /home/nexus-data/:/nexus-data/ --name nexus3 sonatype/nexus3:3.21.1`

    - 启动nexus时，没有权限操作宿主机文件夹 `chmod 777 /home/nexus-data`

## docker 导出导入镜像

1. 导出 `docker save sonatype/nexus3 -o sonatype-nexus3-3.21.1.tar`
1. 导出 `docker save gomods/athens -o gomods-athens-v0.7.2.tar`
1. 导入 `docker load -i sonatype-nexus3-3.21.1.tar`
1. 导入 `docker load -i gomods-athens-v0.7.2.tar`
1. 打标 `docker tag {nexus3 imageID} sonatype/nexus3:3.21.1`
1. 打标 `docker tag {athens imageID} gomods/athens:v0.7.2`
1. 后台启动 `docker-compose up -d`
1. 删除 `docker-compose rm -fsv`
