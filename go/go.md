#包管理工具
## go path
- 代码开发必须在go path src目录下，不然，就有问题。
- 依赖手动管理
- 依赖包没有版本可言
## go vendor
- 解决了包依赖，一个配置文件就管理
- 依赖包全都下载到项目vendor下，每个项目都把有一份。拉取项目时,开始怀疑人生

##go mod
go modules 是 golang 1.11 新加的特性
 
 1、GO111MODULE 有三个值：off, on和auto（默认值）
- GO111MODULE=off，go命令行将不会支持module功能，寻找依赖包的方式将会沿用旧版本那种通过vendor目录或者GOPATH模式来查找。
- GO111MODULE=on，go命令行会使用modules，而一点也不会去GOPATH目录下查找。
- GO111MODULE=auto，默认值，go命令行将会根据当前目录来决定是否启用module功能。这种情况下可以分为两种情形：
```
$ go env -w GO111MODULE=on
```

