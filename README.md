# go-admin example

a example show how to run go-admin

一个运行go-admin的例子

## use go module 使用模块加载依赖

### step 1

```shell
git clone https://github.com/GoAdminGroup/example.git
```

### step 2

```shell
GO111MODULE=on go run main.go
```

## use gopath 使用GOPATH加载依赖

### step 1

```shell
git clone https://github.com/GoAdminGroup/example.git
```

### step 2

```shell
govendor sync
```

如果你在中国，下载依赖有问题，那么你可以直接从这里下载：[vendor.zip](https://gitee.com/cg33/go-admin/attach_files/279551/download)

下载完解压到项目文件夹即可。

### step 3

```shell
go run main.go
```

