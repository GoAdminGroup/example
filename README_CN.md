# GoAdmin 上手例子

一个运行go-admin的例子。仅供参考，在[这里](http://www.go-admin.cn)了解更多。

以下三种方法。建议go版本大于1.11使用模块加载，同时设置环境变量```GOPROXY=http://goproxy.cn```，版本低于1.11的盆友使用第二种方法。如果本机没有golang环境，可以使用docker。

**如果你没有golang基础，是golang新手的话，建议花几分钟了解一下[golang的依赖包管理机制](https://ms.logger.im/search?q=golang%20%E4%BE%9D%E8%B5%96%E7%AE%A1%E7%90%86)**

如果你是windows用户，那么你需要下载gcc，因为本例子使用的是sqlite数据库，如果你不想使用sqlite数据库，你可以换成mysql，则不需要下载gcc。

劝退：没有计算机基础或基础比较差的请谨慎使用或不要使用orz。

## 安装与运行

### 使用模块go module加载依赖

使用 go module的话，需要先设置环境变量```GO111MODULE```为```on```

#### 第一步

```shell
git clone https://github.com/GoAdminGroup/example.git
```

#### 第二步

```shell
cd example
GO111MODULE=on go run .
```

访问: [http://localhost:9033/admin](http://localhost:9033/admin)

### 使用GOPATH加载依赖

使用 go path的话，需要先设置环境变量```GO111MODULE```为```off```

#### 第一步

```shell
git clone https://github.com/GoAdminGroup/example.git
```

#### 第二步

```shell
go get github.com/kardianos/govendor
cd example
govendor sync
```

如果你在中国，因为各种原因导致用以上步骤进行下载安装依赖有问题，那么你可以直接从这里下载：[vendor.zip](http://file.go-admin.cn/go_admin/vendor/v1_2_9/vendor.zip)

下载完解压到项目文件夹即可。

#### 第三步

```shell
go run .
```

访问: [http://localhost:9033/admin](http://localhost:9033/admin)

### use docker 使用docker

#### 第一步

```shell
git clone https://github.com/GoAdminGroup/example.git
```

#### 第二步

```shell
cd example
docker build -t go-admin-example .
```

#### 第三步

```shell
docker attach $(docker run -p 9033:9033 -it -d go-admin-example /bin/bash -c "cd /go/src/app && GOPROXY=http://goproxy.cn GO111MODULE=on go run .")
```

访问: [http://localhost:9033/admin](http://localhost:9033/admin)

## 文件夹介绍

```
.
├── Dockerfile          Dockerfile
├── Makefile            Makefile命令
├── adm_config.ini      adm配置文件
├── admin.db            数据库文件
├── build               二进制构建目标文件夹
├── config.json         配置文件
├── deploy              部署命令说明
├── go.mod              go.mod
├── go.sum              go.sum
├── html                前端html文件
├── logs                日志存放文件夹
├── main.go             main文件
├── main_test.go        CI测试文件
├── models              ORM模型文件
├── pages               页面控制器
├── tables              数据模型表格文件
├── uploads             图片等上传文件夹
└── vendor              第三方依赖
```

## 开发与部署流程

### 第一步

#### 1. 新建表格

数据库新建表格后，通过执行```make generate```生成数据表格模型文件，修改数据表格模型文件。

#### 2. 新建页面

pages文件夹新建页面控制器文件，如果需要更大程度定制可以html文件夹下新建一个golang tmpl模板文件，然后在main.go中载入。

### 第二步

本地执行```make serve```，查看效果。并增加菜单，与对应权限角色。

### 第三步

编写main_test.go测试文件，本地执行```make test```，测试所有API与页面UI逻辑。

### 第四步

测试没问题后，通过git等版本控制软件发布到线上，触发CI测试。

### 第五步

测试没问题后，将```make build```编译出的二进制文件，发布到生产环境。


**注意：以上windows用户请自行将makefile命令转换为windows下命令**