# go-admin example

A example show how to run go-admin. Just for reference, [here](http://www.go-admin.cn/en) to know more.

一个运行go-admin的例子。仅供参考，在[这里](http://www.go-admin.cn)了解更多。

Following three ways to run the code.

以下三种方法。建议go版本大于1.11使用模块加载，同时设置环境变量```GOPROXY=http://goproxy.cn```，版本低于1.11的盆友使用第二种方法。如果本机没有golang环境，可以使用docker。

## use go module 使用模块加载依赖

### step 1

```shell
git clone https://github.com/GoAdminGroup/example.git
```

### step 2

```shell
cd example
GO111MODULE=on go run main.go
```

visit: [http://localhost:9033/admin](http://localhost:9033/admin)

## use gopath 使用GOPATH加载依赖

### step 1

```shell
git clone https://github.com/GoAdminGroup/example.git
```

### step 2

```shell
go get github.com/kardianos/govendor
cd example
govendor sync
```

如果你在中国，因为各种原因导致用以上步骤进行下载安装依赖有问题，那么你可以直接从这里下载：[vendor.zip](http://file.go-admin.cn/go_admin/vendor/v1_0_0_alpha/vendor.zip)

下载完解压到项目文件夹即可。

### step 3

```shell
go run main.go
```

visit: [http://localhost:9033/admin](http://localhost:9033/admin)

## use docker 使用docker

### step 1

```shell
git clone https://github.com/GoAdminGroup/example.git
```

### step 2

```shell
cd example
docker build -t go-admin-example .
```

### step 3

```shell
docker run -p 9033:9033 -it -d go-admin-example /bin/bash -c "cd /go/src/app && GOPROXY=http://goproxy.cn GO111MODULE=on go run main.go"
```

visit: [http://localhost:9033/admin](http://localhost:9033/admin)