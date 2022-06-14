# web编程
## 跨域问题
同源策略（Same origin policy）：所谓同源（即指在同一个域）就是两个页面具有相同的协议
* DOM 同源策略：禁止对不同源页面 DOM 进行操作。这里主要场景是 iframe 跨域的情况，不同域名的 iframe 是限制互相访问的。
* XMLHttpRequest 同源策略：禁止使用 XHR 对象向不同源的服务器地址发起 HTTP 请求。
同源策略是浏览器最核心也最基本的安全功能，同时产生跨域拦截问题
跨域：当一个请求url的协议、域名、端口三者之间任意一个与当前页面url不同即为跨域

CORS（Cross-origin resource sharing，跨域资源共享）是一个 W3C 标准，定义了在必须访问跨域资源时，浏览器与服务器应该如何沟通。
* 基本思想：使用自定义的 HTTP 头部让浏览器与服务器进行沟通，从而决定请求或响应是应该成功，还是应该失败。CORS 需要浏览器和服务器同时支持
### 简单请求
简单请求的定义：
* 请求方法为 HEAD、GET、POST中的一种
* HTTP头信息不超过以下几种：
```
Accept
Accept-Language
Content-Language、Last-Event-ID
Content-Type（只限于三个值application/x-www-form-urlencoded、multipart/form-data、text/plain）
```

对于简单请求，浏览器回自动在请求的头部添加一个 Origin 字段来说明本次请求来自哪个源（协议 + 域名 + 端口），服务端则通过这个值判断是否接收本次请求。如果 Origin 在许可范围内，则服务器返回的响应会多出几个头信息
```
Access-Control-Allow-Credentials: true
Access-Control-Allow-Headers: Content-Type, Content-Length
Access-Control-Allow-Origin: *
Content-Type: text/html; charset=utf-8
```
### 非简单请求
非简单请求是那种对服务器有特殊要求的请求，比如请求方法是 PUT 或 DELETE ，或者 Content-Type 字段的类型是 application/json
非简单请求的CORS请求，会在正式通信之前，增加一次HTTP查询请求，称为"预检"请求（preflight），预检请求其实就是我们常说的 OPTIONS 请求，表示这个请求是用来询问的。头信息里面，关键字段 Origin ，表示请求来自哪个源，除 Origin 字段，"预检"请求的头信息包括两个特殊字段
```
//该字段是必须的，用来列出浏览器的CORS请求会用到哪些HTTP方法
Access-Control-Request-Method
//该字段是一个逗号分隔的字符串，指定浏览器CORS请求会额外发送的头信息字段.
Access-Control-Request-Headers
```
浏览器先询问服务器，当前网页所在的域名是否在服务器的许可名单之中，以及可以使用哪些HTTP动词和头信息字段。只有得到肯定答复，浏览器才会发出正式的 XMLHttpRequest 请求，否则就报错
### 配置CORS
```
//该字段是必须的。它的值要么是请求时Origin字段的值，要么是一个*，表示接受任意域名的请求
Access-Control-Allow-Origin

//该字段必需，它的值是逗号分隔的一个字符串，表明服务器支持的所有跨域请求的方法。注意，返回的是所有支持的方法，而不单是浏览器请求的那个方法。这是为了避免多次"预检"请求。
Access-Control-Allow-Methods

//如果浏览器请求包括Access-Control-Request-Headers字段，则Access-Control-Allow-Headers字段是必需的。它也是一个逗号分隔的字符串，表明服务器支持的所有头信息字段，不限于浏览器在"预检"中请求的字段
Access-Control-Allow-Headers

//该字段可选。CORS请求时，XMLHttpRequest对象的response只能拿到6个基本字段：Cache-Control、Content-Language、Content-Type、Expires、Last-Modified、Pragma。如果想拿到其他字段，就必须在Access-Control-Expose-Headers里面指定。
Access-Control-Expose-Headers

//该字段可选。它的值是一个布尔值，表示是否允许发送Cookie。默认情况下，Cookie不包括在CORS请求之中。设为true，即表示服务器明确许可，Cookie可以包含在请求中，一起发给服务器。这个值也只能设为 true，如果服务器不要浏览器发送Cookie，删除该字段即可
Access-Control-Allow-Credentials

//该字段可选，用来指定本次预检请求的有效期，单位为秒，在此期间，不用发出另一条预检请求
Access-Control-Max-Age
```
### Go配置
以Gin框架为例，配置处理跨域的中间件
```
func Cors(context *gin.Context) {
	method := context.Request.Method
	// 必须，接受指定域的请求，可以使用*不加以限制，但不安全；存在一个问题是不允许XMLHttpRequest携带Cookie，所以要实现通配的话可以采用动态获取Origin
	//context.Header("Access-Control-Allow-Origin", "*")
	
	//动态获取Origin
	context.Header("Access-Control-Allow-Origin", context.GetHeader("Origin"))

	// 必须，设置服务器支持的所有跨域请求的方法
	context.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
	
	// 服务器支持的所有头信息字段，不限于浏览器在"预检"中请求的字段
	context.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Token")
	
	// 可选，设置XMLHttpRequest的响应对象能拿到的额外字段
	context.Header("Access-Control-Expose-Headers", "Access-Control-Allow-Headers, Token")
	
	// 可选，是否允许后续请求携带认证信息Cookir，该值只能是true，不需要则不设置
	context.Header("Access-Control-Allow-Credentials", "true")
	
	// 放行所有OPTIONS方法
	if method == "OPTIONS" {
		context.AbortWithStatus(http.StatusNoContent)
		return
	}
	context.Next()
}

```
## Cookie
无状态协议：不必保存状态，减少了服务器的 CPU 及内存资源消耗，协议足够简单，使用场景广泛
HTTP ：是无状态协议，不记录之前发生过的请求和响应，也因此无法根据历史状态信息处理当前请求
Cookie 技术：通过在 HTTP 请求和响应报文中写入 Cookie 信息来控制客户端的状态
* 服务端在响应报文内添加一个叫做 Set-Cookie 的首部字段信息，客户端接收到后会保存 Cookie
* 客户端向服务器发送请求时，就会自动在请求报文中加入保存的 Cookie 值

| 首部字段名  |   首部类型   |             说明              |
| ---------- | ----------- | ---------------------------- |
| Set-Cookie | 响应首部字段 | 开始状态管理所使用的Cookie信息 |
| Cookie     | 请求首部字段 | 服务器接收到的Cookie信息       |
### Set-Cookie
Set-Cookie 的字段值
* expires：设置到期的具体时间点，
* MaxAge： Cookie 的有效时长（单位是秒），考虑到默认时区问题，推荐 MaxAge
* domain：Cookie适用对象的域名（若不指定则默认为创建Cookie的服务器域名），创建则可做到与结尾匹配一致
* secure：仅在HTTPS安全通信时才会发送Cookie
* HttpOnly：加以限制，使Cookie不能被JavaScript脚本访问
### Cookie
在请求报文中添加该字段后，就相当于告诉服务器客户端想要获得 HTTP 状态管理支持。接收到多个Cookie时，同样可以以多个Cookie形式发送。
### Session
某些 Web 页面只想让特定的人浏览，或者干脆仅本人可见，为达到这个目标，需要添加认证功能。HTTP/1.1 实用的认证包括 BASIC认证、DIGEST认证、SSL客户端认证、FormBase认证等，由于使用上的便利性和安全性问题，前两种几乎不适用，SSL客户端认证则由于导入及费用问题未得到普及，目前常用的是最后一种：基于表单的认证
基于表单的认证方法并不是在HTTP协议中定义的，而是由客户端通过表单向服务器提交登录信息，然后由服务器安装自定义的实现方式进行验证，不同的应用使用的验证方式多有不同，但多数情况下，是基于用户输入的用户ID（通常是任意字符串或邮件地址）和密码等登录信息进行认证
鉴于 HTTP 是无状态协议，之前已认证成功额用户状态无法保留，因此一般使用 Cookie 来管理 Session(会话)
1. 客户端把用户ID和密码等登录信息放入报文的实体部分，通常是以POST方法把请求发送给服务器
2. 服务器会发放用以识别用户的Session ID。通过验证从客户端发送过来的登录信息进行身份认证，然后把用户的认证状态与Session ID绑定后记录在服务器端。向客户端返回响应时，会在首部字段Set-Cookie内写入Session ID（如PHPSESSID=028a8c…）
3. 服务器会发放用以识别用户的Session ID。通过验证从客户端发送过来的登录信息进行身份认证，然后把用户的认证状态与Session ID绑定后记录在服务器端。向客户端返回响应时，会在首部字段Set-Cookie内写入Session ID（如PHPSESSID=028a8c…）
### Go配置
设置：
```
expiration := time.Now()
expiration := expiration.AddDate(1, 0, 0)
cookie := http.Cookie{
    Name: "username", 
    Value: "zuolan", 
    Expires: expiration
}
http.SetCookie(writer, &Cookie)

```
读取：
```
cookie, _ := r.Cookie("username")
for _, cookie := range r.Cookies() {    
    fmt.Fprint(w, cookie.Name)
}
```


