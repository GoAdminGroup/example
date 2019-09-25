# go-admin example

a example show how to run go-admin.

一个运行go-admin的例子。

Following two ways to run the code.

以下两种方法。建议go版本大于1.11使用模块加载，同时设置环境变量```GOPROXY=http://goproxy.cn```，版本低于1.11的盆友使用第二种方法。

## Way 1: use go module 使用模块加载依赖

### step 1

```shell
git clone https://github.com/GoAdminGroup/example.git
```

### step 2

```shell
GO111MODULE=on go run main.go
```

## Way 2: use gopath 使用GOPATH加载依赖

### step 1

```shell
cd $GOPATH && git clone https://github.com/GoAdminGroup/example.git
```

### step 2

```shell
go get github.com/kardianos/govendor
govendor sync
```

如果你在中国，因为各种原因导致用以上步骤进行下载安装依赖有问题，那么你可以直接从这里下载：[vendor.zip](https://gitee.com/cg33/go-admin/attach_files/279551/download)

下载完解压到项目文件夹即可。

### step 3

```shell
go run main.go
```

