# quickutil4go
Utils for golang

### Change Log

#### v0.1 
+ logutil
+ httputil
+ fileutil
+ dbutil
+ cryptoutil
+ environmentutil

### How to begin
+ idea 安装go
+ 下载goroot，https://golang.org/，设置环境变量
+ 设置goproxy，idea->preference->go->go module 设置GOPROXY=https://mirrors.aliyun.com/goproxy/
+ 第一次启动，go mod tidy，如果有问题可尝试 go clean --modcache
+ 新增依赖1，go get github.com/jmoiron/sqlx
+ 新增依赖2，go mod download