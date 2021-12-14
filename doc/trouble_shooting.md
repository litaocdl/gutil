## Install go module
go get -v golang.org/x/tools/cmd/godoc 


## timeout

```
go get: module golang.org/x/tools/cmd/godoc: Get "https://proxy.golang.org/golang.org/x/tools/cmd/godoc/@v/list": dial tcp 142.251.43.17:443: i/o timeout
```

```
go env -w GOPROXY=https://goproxy.cn
```

## Start a godoc local

```
godoc -http :8000
```