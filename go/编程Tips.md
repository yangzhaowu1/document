# 编程Tips
## json
```
type Person struct {
    // 指定json序列化/反序列化时使用小写name
	Name   string `json:"name"` 
	
	//// 指定json序列化/反序列化时忽略此字段
	Weight float64 `json:"-"`
	 
	//当 struct 中的字段没有值时,默认输出字段的类型零值;想要在序列序列化时忽略这些没有值的字段时，可以在对应字段添加omitempty
	Hobby []string `json:"hobby,omitempty"` 
	
    //匿名嵌套Profile时序列化后的json串为单层的,即序列化后site字段与name字段无异
    Profile
    
    //具名嵌套：多层
    Profile1 Profile
    
    //定义字段tag:多层
    Profile `json:"profile"`
    
    //嵌套的结构体为空值时，不会忽略该字段
    Profile `json:"profile,omitempty"`
    
    //嵌套的结构体为空值时，会忽略该字段
    *Profile `json:"profile,omitempty"`
}

type Profile struct {
	Website string `json:"site"`
	Slogan  string `json:"slogan"`
}
```
### 不修改忽略空值字段
需要json序列化User，但是不想把密码也序列化，又不想修改User结构体，这个时候我们就可以使用创建另外一个结构体PublicUser匿名嵌套原User，同时指定Password字段为匿名结构体指针类型，并添加omitempty
```
type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type PublicUser struct {
	*User             // 匿名嵌套
	Password *struct{} `json:"password,omitempty"`
}

func omitPasswordDemo() {
	u1 := User{
		Name:     "七米",
		Password: "123456",
	}
	b, err := json.Marshal(PublicUser{User: &u1})
	if err != nil {
		fmt.Printf("json.Marshal u1 failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)  // str:{"name":"七米"}
}
```
### 字符串格式的数字
前端在传递来的json数据中可能会使用字符串类型的数字，这个时候可以在结构体tag中添加string来告诉json包从字符串中解析相应字段的数
```
type Card struct {
	ID    int64   `json:"id,string"`    // 添加string tag
	Score float64 `json:"score,string"` // 添加string tag
}

func intAndStringDemo() {
	jsonStr1 := `{"id": "1234567","score": "88.50"}`
	var c1 Card
	if err := json.Unmarshal([]byte(jsonStr1), &c1); err != nil {
		fmt.Printf("json.Unmarsha jsonStr1 failed, err:%v\n", err)
		return
	}
	fmt.Printf("c1:%#v\n", c1) // c1:main.Card{ID:1234567, Score:88.5}
}
```
### 整数变浮点数
在 JSON 协议中是没有整型和浮点型之分的，它们统称为number。json字符串中的数字经过Go语言中的json包反序列化之后都会成为float64类型，这种场景下如果想更合理的处理数字就需要使用decoder去反序列化
```
func decoderDemo() {
	// map[string]interface{} -> json string
	var m = make(map[string]interface{}, 1)
	m["count"] = 1 // int
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("marshal failed, err:%v\n", err)
	}
	fmt.Printf("str:%#v\n", string(b))
	// json string -> map[string]interface{}
	var m2 map[string]interface{}
	// 使用decoder方式反序列化，指定使用number类型
	decoder := json.NewDecoder(bytes.NewReader(b))
	decoder.UseNumber()
	err = decoder.Decode(&m2)
	if err != nil {
		fmt.Printf("unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("value:%v\n", m2["count"]) // 1
	fmt.Printf("type:%T\n", m2["count"])  // json.Number
	// 将m2["count"]转为json.Number之后调用Int64()方法获得int64类型的值
	count, err := m2["count"].(json.Number).Int64()
	if err != nil {
		fmt.Printf("parse to int64 failed, err:%v\n", err)
		return
	}
	fmt.Printf("type:%T\n", int(count)) // int
}
```
### 匿名结构体添加字段
使用内嵌结构体能够扩展结构体的字段，但有时候我们没有必要单独定义新的结构体，可以使用匿名结构体简化操作
```
type UserInfo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func anonymousStructDemo() {
	u1 := UserInfo{
		ID:   123456,
		Name: "七米",
	}
	// 使用匿名结构体内嵌User并添加额外字段Token
	b, err := json.Marshal(struct {
		*UserInfo
		Token string `json:"token"`
	}{
		&u1,
		"91je3a4s72d1da96h",
	})
	if err != nil {
		fmt.Printf("json.Marsha failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)
	// str:{"id":123456,"name":"七米","token":"91je3a4s72d1da96h"}
}
```

### 组合多个结构体
使用匿名结构体来组合多个结构体来序列化与反序列化数据
```
type Comment struct {
	Content string
}

type Image struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

func anonymousStructDemo2() {
	c1 := Comment{
		Content: "永远不要高估自己",
	}
	i1 := Image{
		Title: "赞赏码",
		URL:   "https://www.liwenzhou.com/images/zanshang_qr.jpg",
	}
	// struct -> json string
	b, err := json.Marshal(struct {
		*Comment
		*Image
	}{&c1, &i1})
	if err != nil {
		fmt.Printf("json.Marshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)
	// json string -> struct
	jsonStr := `{"Content":"永远不要高估自己","title":"赞赏码","url":"https://www.liwenzhou.com/images/zanshang_qr.jpg"}`
	var (
		c2 Comment
		i2 Image
	)
	if err := json.Unmarshal([]byte(jsonStr), &struct {
		*Comment
		*Image
	}{&c2, &i2}); err != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("c2:%#v i2:%#v\n", c2, i2)
}

str:{"Content":"永远不要高估自己","title":"赞赏码","url":"https://www.liwenzhou.com/images/zanshang_qr.jpg"}
c2:main.Comment{Content:"永远不要高估自己"} i2:main.Image{Title:"赞赏码", URL:"https://www.liwenzhou.com/images/zanshang_qr.jpg"}
```
### 处理不确定层级
如果json串没有固定的格式导致不好定义与其相对应的结构体时，我们可以使用json.RawMessage原始字节数据保存下来。
```
type sendMsg struct {
	User string `json:"user"`
	Msg  string `json:"msg"`
}

func rawMessageDemo() {
	jsonStr := `{"sendMsg":{"user":"q1mi","msg":"永远不要高估自己"},"say":"Hello"}`
	// 定义一个map，value类型为json.RawMessage，方便后续更灵活地处理
	var data map[string]json.RawMessage
	if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
		fmt.Printf("json.Unmarshal jsonStr failed, err:%v\n", err)
		return
	}
	var msg sendMsg
	if err := json.Unmarshal(data["sendMsg"], &msg); err != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("msg:%#v\n", msg)
	// msg:main.sendMsg{User:"q1mi", Msg:"永远不要高估自己"}
}
```
## 文件操作

## 十三、Tips

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