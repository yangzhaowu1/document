# 编程Tips
# 十三、Tips

1、数组是值传递，切片是引用传递，函数调用时若要修改，使用切片

2、切片赋值

```
x := []int{1, 2, 3}
y := x 		//y获得是切片引用，修改x、y之一，另一个也会修改
```

3、map遍历是顺序不固定

4、recover必须在defer函数中运行

```
func main() {
    defer func() {
        recover()
    }()
    panic(1)
}
```

5、Goroutine是协作式抢占调度，Goroutine本身不会主动放弃CPU

6、defer在函数退出时才能执行，在for执行defer会导致资源延迟释放

```go
func main() {
    for i := 0; i < 5; i++ {
        f, err := os.Open("/path/to/file")
        if err != nil {
            log.Fatal(err)
        }
        defer f.Close()
    }
}

解决的方法可以在for中构造一个局部函数，在局部函数内部执行defer：

func main() {
    for i := 0; i < 5; i++ {
        func() {
            f, err := os.Open("/path/to/file")
            if err != nil {
                log.Fatal(err)
            }
            defer f.Close()
        }()
    }
}
```

7、禁止main函数退出

```
func main() {
    defer func() { for {} }()
}

func main() {
    defer func() { select {} }()
}

func main() {
    defer func() { <-make(chan bool) }()
}
```

6、channel用于同步时，可以用无类型匿名结构体:避免额外的内存消耗

```go
c2 := make(chan struct{})
    go func() {
        fmt.Println("c2")
        c2 <- struct{}{} // struct{}部分是类型, {}表示对应的结构体值
    }()
 <-c2
```