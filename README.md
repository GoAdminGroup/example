# GoAdmin Example

[中文说明](./README_CN.md)

A example show how to run go-admin. Just for reference, [here](http://www.go-admin.cn/en) to know more.

Following three ways to run the code.

If you are Windows user, [go-sqlite-dirver](https://github.com/mattn/go-sqlite3) require to download the gcc to make it work.

## use go module

To use go module, you should set GO111MODULE=on first.

### step 1

```shell
git clone https://github.com/jvcss/go.git
```

### step 2

```shell
cd go
GO111MODULE=on go run .
```

visit: [http://localhost:9033/admin](http://localhost:9033/admin)

## use gopath

To use go path, you should set GO111MODULE=off first.

### step 1

```shell
git clone https://github.com/jvcss/go.git
```

### step 2

```shell
go get github.com/jvcss/govendor
cd example
govendor sync
```

### step 3

```shell
go run .
```

visit: [http://localhost:9033/admin](http://localhost:9033/admin)

## use docker

### step 1

```shell
git clone https://github.com/jvcss/go.git
```

### step 2

```shell
cd go
docker build -t go-admin .
```

### step 3

```shell
docker attach $(docker run -p 9033:9033 -it -d go-admin /bin/bash -c "cd /go/src/app && GOPROXY=http://goproxy.cn GO111MODULE=on go run .")
```

visit: [http://localhost:9033/admin](http://localhost:9033/admin)
