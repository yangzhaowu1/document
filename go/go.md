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

 2、命令

| 命令 | 说明                 |
| :---- | :--------- |
| download  | 下载依赖包    |
| edit      | 编辑go.mod    |
| graph     | 打印模块依赖图    |
| tidy      | 拉取缺少的模块，移除不用的模块    |
| vendor    | 将依赖复制到vendor下     |
| why       | 解释为什么需要依赖    |
| graph     | 打印模块依赖图    |


(1)例
```
$ go env -w GO111MODULE=on
$ go mod init github.com/yangzhaowu1/document
$ go mod tidy
```

（2）go.mod 提供了module, require、replace和exclude 四个命令

- module 语句指定包的名字（路径）
- require 语句指定的依赖项模块
- replace 语句可以替换依赖项模块
- exclude 语句可以忽略依赖项模块

（3）go get升级
- 运行 go get -u 将会升级到最新的次要版本或者修订版本(x.y.z, z是修订版本号， y是次要版本号)
- 运行 go get -u=patch 将会升级到最新的修订版本
- 运行 go get package@version 将会升级到指定的版本号version
- 运行go get如果有版本的更改，那么go.mod文件也会更改
- 使用replace替换无法直接获取的package


（4） 由于某些已知的原因，并不是所有的package都能成功下载，比如：golang.org下的包。
modules 可以通过在 go.mod 文件中使用 replace 指令替换成github上对应的库，比如：

```
replace (
    golang.org/x/crypto v0.0.0-20190313024323-a1f597ede03a => github.com/golang/crypto v0.0.0-20190313024323-a1f597ede03a
)
```



