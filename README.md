## go-locust
参考python的locust压测系统，所实现的golang版的分布式压测系统。  
## 功能特性
支持平台化管理压测任务  
支持分布式压测  
支持水平扩展服务器节点  
基于golang协程并发  


-------

# 包错误问题

go sdk版本在go1.13和go1.14版本使用go mod管理依赖包中有etcd时会出现异常，无法正常拉取etcd包。

错误如下：

go.etcd.io/etcd/clientv3 tested by
go.etcd.io/etcd/clientv3.test imports
github.com/coreos/etcd/auth imports
github.com/coreos/etcd/mvcc/backend imports
github.com/coreos/bbolt: github.com/coreos/bbolt@v1.3.5: parsing go.mod:
module declares its path as: go.etcd.io/bbolt
but was required as: github.com/coreos/bbolt

引起以上的原因主要是etcd中使用的bbolt和grpc版本冲突引起

# 解决方法

删除原来已生成得go.mod和go.sum

go mod init

go mod edit -replace github.com/coreos/bbolt@v1.3.4=go.etcd.io/bbolt@v1.3.4

go mod edit -replace google.golang.org/grpc@v1.29.1=google.golang.org/grpc@v1.26.0

go mod tidy

------------

# 包问题

go: finding module for package google.golang.org/grpc/naming
loadtest-locust/db imports
github.com/coreos/etcd/clientv3 tested by
github.com/coreos/etcd/clientv3.test imports
github.com/coreos/etcd/integration imports

github.com/coreos/etcd/proxy/grpcproxy imports
google.golang.org/grpc/naming: module google.golang.org/grpc@latest found (v1.51.0), but does not contain package google.golang.org/grpc/naming

原因：

etcd clientv3包的版本依赖冲突导致（实际是etcd依赖的grpc断代）

详情：ETCD中使用的旧版本gRPC库与最新版本的gRPC库不兼容，需要在go.mod中将gRPC替换为 v1.29.1（etcd v3.4.9+要求grpc v1.29.1之前的，从下一个版本开始断代）版本的即可。

在你项目的go.mod中添加

replace google.golang.org/grpc => google.golang.org/grpc v1.29.1

注意：grpc的版本要看etcd的版本，比如etcd v3.3.20 的 release 版本要求 grpc 的版本是 v1.26.0 之前的。总之，etcd的断代，不是etcd本身，而是grpc断代了。


对了，etcd最新的v3版本（etcd clientv3与etcd/client/v3是不一样的），这个v3的接口是和最新的grpc兼容的，所以不需要再像上面这么麻烦了，新项目直接使用etcd/client/v3即可。（所以一把全升级了，go build能过，新的三方、二方库用的最新版etcd）

-----------

问题： undefined: balancer.PickOptions

导致这个问题的原因是grpc的版本，go.mod文件中：

google.golang.org/grpc v1.27.0

而v1.26.0是可以的，修改go.mod中的版本为：

google.golang.org/grpc v1.26.0

执行以下命令：

go get google.golang.org/grpc@v1.26.0