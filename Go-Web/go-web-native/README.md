## 《GoWeb编程》2017版笔记

[TOC]



### Go与Web应用

#### 1.1 使用Go构建Web应用

- go的可拓展Web应用
  - 有两种方式可以对性能进行扩展
    - 垂直扩展，提升单台设备的CPU数量或性能
      - 一个GoWeb应用只需要使用一个操作系统线程（OS thread），可以通过调度来高效地运行数十万个goroutine
      - go应用被编译为不包含任何动态依赖关系的静态二进制文件
    - 水平扩展，通过增加计算机的数量来提升性能
- go的模块化Web应用
  - 通过复用模块化的组件来降低软件开发所需的费用
  - go可通过它的接口机制对行为进行描述，以此来实现动态类型匹配（dynamic typing）
  - go的函数可以接受接口作为参数，所有的类型都实现了空接口，还实现了函数式编程的特性，包括函数类型、使用函数作为值以及闭包
  - go常被用于创建微服务，微服务分为系统级服务和应用级服务
- go的可维护Web应用
  - 拥有一个易于维护的代码库（codebase）
  - go的包管理系统
  - go的工程化工具：gofmt、godoc等
- go的高性能Web应用
  - 高性能的意义
    - 能够短时间内处理大量请求
    - 服务器能快速对客户端进行响应，并让终端用户（end user）快速执行操作
  - go程序会被编译为本地码（native code）
  - goroutine对并发编程提供了原生的支持

#### 1.2 Web应用的工作原理

- 从狭义的角度解析什么是Web应用：它会对客户端发送的HTTP请求做出响应，并通过HTTP响应将HTML回传至客户端
  - 从上面定义可以看成：一个Web服务器就是一个Web应用
- 完整定义**Web应用**（为了便于分辨和理解）：
  - 这个程序必须向发送命令请求的客户端返回HTML，而客户端则会向用户展示渲染后的HTML
  - 这个程序在向客户端传送数据时必须使用HTTP协议
  - 在这个定义上，如果一个程序不是向用户渲染并展示HTML，而是向其他程序返回某种非HTML格式的数据，那么这个程序就是一个为其他程序提供服务的**Web服务**

#### 1.3 HTTP简介

- HTTP是万维网的应用层通信协议
- 至今经过了三次迭代修改，第一版为 HTTP 0.9，第二版为 HTTP 1.0，第三版为 HTTP 1.1 ，第四版 HTTP 2.0
- 什么是HTTP
  - HTTP是一种 *无状态* 、*由文本构成*  的请求-响应（ request-response ）协议，这种协议使用的是客户端-服务器（client-server）计算模型
  - 在HTTP协议中，客户端也被称作 **用户代理（user-agent）**而服务器则通常会被称为Web服务器
  - HTTP是以**纯文本方式**而不是二进制方式发送和接收协议数据的

#### 1.4 Web应用的诞生

- *通用网关接口（CGI）*是在早期尝试动态生成HTML内容的技术之一
- CGI 是一个简单的接口，允许Web服务器与一个独立允许的进程对接，通过CGI与服务器进行对接的程序通常被称为 *CGI 程序*
  - 向 CGI 程序传递输入参数是通过设置环境变量来完成的
  - CGI 程序在运行之后将向标准输出（stand output）返回结果，而服务器则会将这些结果传送至客户端
- *服务器端包含（server-side includes, SSI）*技术也是 CGI 的同期技术
  - 这种技术允许开发者在HTML文件里面包含一些指令 （directive）
  - 当客户端请求一个HTML文件的时候，服务器在返回这个文件之前，会先执行文件中包含的指令，并将文件中出现指令的位置替换成这些指令的执行结果
  - SSI 最常见的用法是在 HTML 文件中包含其他被频繁使用的文件，又或者将整个网站都会出现的页面首部（Header）以及尾部（Footer）的代码段嵌入 HTML 文件中
  - SSI 技术的最终演化结果就是在 HTML 里面包含更为复杂的代码，并使用更为强大的解析器（interpreter），后面演化成 PHP、ASP、JSP 等一下Web模板引擎技术

#### 1.5 HTTP请求

-  HTTP请求由一系列文本行组成，按以下顺序排列：

  1. 请求行( request-line )
  2. 零个或任意多个请求首部( header )
  3. 一个空行
  4. 可选的报文主体( body )

  ```http
  GET /Protocols/rfc2616.html HTTP/1.1
  Host: www.w3.org
  User-Agent: Mozilla/5.0
  (empty line)
  ```

  请求行中的第一个单词为*请求方法( request method )*，之后跟着的是*统一资源标识符( URI )*以及所用的          *HTTP版本* 。位于请求行之后的两个文本行为请求的*首部*。

  *注意*，这个报文的最后一行为空行，即使报文的主体部分为空，**这个空行也必须存在**，报文是否包含主体则需要根据请求使用的方法而定。

- 请求方法

  注：在 HTTP 1.1 要求必须实现的只有 GET 方法和 HEAD 方法，其他方法都是可选的，包括 POST

  - 各方法说明
    - HEAD--与GET方法的作用相似，唯一不同在于这个方法不要求服务器返回报文的主体，这个方法通常用于在不获取报文主体的情况下，取得响应的首部
    - GET--命令服务器返回指定的资源
    - POST--命令服务器将报文主体中的数据传递给URI指定的资源，至于服务器具体会对这些数据执行什么动作则取决于服务器本身
    - PUT--命令服务器将报文主体中的数据传递给URI指定的资源，如果URI指定的位置上已经有数据存在，那么使用报文主体中的数据去代替已有的数据，如果资源尚未存在，那么在URI指定的位置上新创建一个资源
    - DELETE--命令服务器删除URI指定的资源
    - TRACE--命令服务器返回请求本身，通过这个方法，客户端可以知道介于它和服务器之间的其他服务器是如何处理请求的
    - OPTIONS--命令服务器返回它支持的HTTP方法列表
    - CONNECT--命令服务器与客户端建立一个网络连接，这个方法通常用于设置SSL隧道以启动HTTPS功能
    - PATCH--命令服务器使用报文主体中的数据对URI指定的资源进行修改
  - 安全的请求方法：一个HTTP方法只要求服务器提供信息而不会对服务器的状态做任何修改，那这个方法就是安全的，如GET、HEAD、OPTIONS 和 TRACE，相对的POST、PUT、DELETE则是不安全的
  - 幂等的请求方法：一个HTTP方法在使用相同的数据进行第二次调用时，不会对服务器的状态造成任何改变，则这个方法就是幂等的，所有的安全的方法都是幂等的，除此之外，PUT和DELETE也是幂等的（虽然它不安全）
  - 关于浏览器对请求方法的支持
    - HTML 不支持除GET和POST以外的其他HTTP方法
    - 用户可以使用XMLHttpRequest（XHR）来获得PUT方法和DELETE方法的支持，XHR是一系列浏览器API，这些API基本由js包裹（XHR其实就是名为 XMLHttpRequest 的**浏览器对象**），这个技术支持 XML格式、JSON以及纯文本格式的请求和响应

- 请求首部

  HTTP请求方法定义了发送请求的客户端想要执行的动作，而请求的首部则记录了与请求本身以及客户端有关的信息。请求的首部由任意多个用冒号分隔的纯文本键值对组成，最后以回车（CR）和换行（LF）结尾。

  大多数HTTP请求首部都是可选的，宿主（Host）首部字段是HTTP 1.1唯一强制要求的首部。根据请求使用的方法，如果请求的报文中包含有可选的主体，那么请求的首部还需要带有内容长度（Content-Length）字段或者传输编码（Transfer-Encoding）字段。

  | 首部字段       | 作用描述                                                     |
  | -------------- | ------------------------------------------------------------ |
  | Accept         | 客户端在HTTP响应中能够接收的内容类型。比如说，客户端可以通过Accept：         text／html这个首部，告知服务器自己希望在响应的主体中收到HTML类型的内容 |
  | Accept-Charset | 客户端要求服务器使用的字符集编码。比如说，客户端可以通过 Accept－ Charset：utf－8这个首部，告知服务器自己希望响应的主体使用UTF－8字符集 |
  | Authorization  | 这个首部用于向服务器发送基本的身份验证证书                   |
  | Cookie         | 客户端应该在这个首部中把服务器之前设置的所有 cookie回传给服务器。比如说，如果服务器之前在浏览器上设置了3个 cookie，那么 Cookie首部字段将在一个字符串里面包含这3个 cookie，并使用分号对这些 cookie进行分隔。以下是一个 Cookie首部示例: Cookie: my_first_cookie=hello; my_second_cookie=world |
  | Content-Length | 请求主体的字节长度                                           |
  | Content-Type   | 当请求包含主体的时候，这个首部用于记录主体内容的类型。在发送POST或PUT请求时，内容的类型默认为`x－ww－form－ urlen－ coded`，但是在上传文件时，内容的类型应该设置为 `multipart／form－data`（上传文件这一操作可以通过将 input标签的类型设置为 file来实现） |
  | Host           | 服务器的名字以及端口号。如果这个首部没有记录服务器的端口号，就表示服务器使用的是80端口 |
  | Referrer       | 发起请求的页面所在的地址                                     |
  | User-Agent     | 对发起请求的客户端进行描述                                   |


#### 1.6 HTTP响应

​	HTTP响应报文是对HTTP请求报文的回复。包括：

 	1. 一个状态行
 	2. 零个或任意数量的响应首部
 	3. 一个空行
 	4. 一个可选的报文主体

```http
200 OK
Date: Sat,22 Nov 2014 12:58:58 GMT
Server: Apache/2
	Last-Modified: Thu, 28 Aug 2014 21:01:33 GMT
Content-Length:33115
Content-Type:text/html;charset=iso-8859-1

<!DOCTYPE html ....>
```



​	响应状态码

| 状态码类型 | 作用描述                                                     |
| ---------- | ------------------------------------------------------------ |
| 1XX        | 情报状态码。服务器通过这些状态码来告知客户端，自己已经收到了客户端发送的请求，并且已经对请求进行了处理 |
| 2XX        | 成功状态码。这些状态码说明服务器已经接收到了客户端发送的请求，并且已经成功地对请求进行了处理。这类状态码的标准响应为“200 OK” |
| 3XX        | 重定向状态码。这些状态码表示服务器已经接收到了客户端发送的请求，并且已经成功处理了请求，但为了完成请求指定的动作，客户端还需要再做一些其他工作。这类状态码大多用于实现URL重定向 |
| 4XX        | 客户端错误状态码。这类状态码说明客户端发送的请求出现了某些问题，在这一类型的状态码中，最常见的就是 "404 Not Found"了，这个状态码表示服务器无法从请求指定的URL中找到客户端想要的资源 |
| 5XX        | 服务器错误状态码:当服务器因为某些原因而无法正确地处理请求时，服务器就会使用这类状态码来通知客户端。在这一类状态码中， 最常见的就是“500 Internal Server Error" 状态码了 |



​	响应首部

响应的首部和请求的首部一样，由任意多个用冒号分隔的纯文本键值对组成，同样以回车（CR）和换行（LF）结尾。

| 首部字段         | 作用描述                                                     |
| ---------------- | ------------------------------------------------------------ |
| Allow            | 告知客户端，服务器支持哪些请求方法                           |
| Content-Length   | 响应主体的字节长度                                           |
| Content-Type     | 如果响应包含可选的主体，那么这个首部记录的就是主体内容的类型 |
| Date             | 以格林尼治标准时间(GMT)格式记录的当前时间                    |
| Location         | 这个首部仅在重定向时使用，它会告知客户端下来应该向哪个URL发送请求 |
| Server           | 返回响应的服务器的城名                                       |
| Set-Cookie       | 在客户端里面设置一个cookie。一个响应里面可以包含多个Set-Cookie首部 |
| VWW-Authenticate | 服务器通过这个首部来告知客户端，在Athorization请求首部中应该提供哪种类型的身份验证信息。服务器常常会把这个首部与“401 Unauthorzed"状态行一同发送。除此之外，这个首部还会向服务器许可的认证授权模式(schema)提供验证信息（challenge information）(比如RFC 2617描述的基本和摘要访问认证模式） |



#### 1.7 URI

在1994年发表的 RFC 1630 中对统一资源标识符（URI）进行了定义，Tim Berners-Lee在这篇RFC中，描述了一种使用字符串表示资源名字的方法，以及一种使用字符串表示资源所在位置的方法，前一种是**统一资源名称（URN）**，后一种是**统一资源定位符（URL）**。URI包含URN和URL，两者有相似的语法和格式。

> URI的一般格式为：
>
> <方案名称>:<分层部分>[  ？<查询参数> ]空格[ # <片段>]

URI中的方案名称(scheme name)记录了URI正在使用的方案，它定义了URI其余部分的结构。在大多数情况下只会使用HTTP方案。

URI的分层部分( hierarchical part ) 包含了资源的识别信息，这些信息会以分层的方式进行组织。如果分层部分以双斜线(//)开头，那么说明它包含了可选的用户信息，这些信息将以@符号结尾，后跟分层路径。不带用户信息的分层部分就是个单纯的路径， 每个路径都由一连串的分段(segment)组成，各个分段之间使用单斜线 **(/)** 分隔。

在URI的各个部分当中，**只有“方案名称”和“分层部分”是必需**的。以问号 **(?) ** 为前缀的查询参数(query) 是可选的，这些参数用于包含无法使用分层方式表示的其他信息。多个查询参数会被组织成一连串的键值对，各个键值对之间使用 **&** 符号分隔。

URI的另一个可选部分为片段( fragment ),片段使用井号 **(#)** 作为前缀，它可以对URI定义的资源中的次级资源( secondary resource) 进行标识。当URI包含查询参数时，URI的片段将被放到查询参数之后。因为URI的片段是由客户端负责处理的，所以Web议览器在将URI发送给服务器之前，一般都会先把URI中的片段移除掉。如果程序员想要取得URI片段，那么可以通过JavaScript或者某个HTTP客户端库，将URI片段包含在一个GET请求里面。

HTTP方案的URI示例：

```http
http://sausheng:password@www.example.com/docs/file?name=sausheong&location=singapore#summary
```

因为每个 URL 都是一个单独的字符串，所以 URL 里面是不能够包含空格的，**问号**和**井号**等符号在URL中具有特殊含义，所以这些符号是不能用于其他用途的，为了避开这些限制， **URL编码**（又称作百分号编码） 就出现了。

RFC 3986定义了 URL中的保留字符以及非保留字符，所有保留字符都需要进行 URL 编码： URL编码会把保留字符转换成该字符在 ASCII 编码中对应的字节值（ byte value ），接着把这些字节值表示为一个两位长的十六进制数字，最后再在这个十六进制数字的签名加上一个百分号（%）。 如：空格的 ASCII 字节值为32，也就是十六进制中的20，经过 URL 编码处理就成了 **%20**

#### 1.8 HTTP/2 简介

​	HTTP/2是一种二进制协议：让HTTP/2的语法分析变得高效，但无法继续通过 telnet 等应用程序直接发送 HTTP/2 报文进行调试。

​	HTTP/2是完全多路复用的，这以为这多个请求和响应可以在同一时间内使用同一个连接。HTTP/2还会对首部进行压缩以减少需要传送的数据量，并允许服务器将响应**推送（push）**到客户端

​	尽管 HTTP/2 对协议的通信性能进行了优化，但它并没有对HTTP协议本身的语法进行修改，**方法和状态码**等还是与HTTP 1.1时一样。

​	在Go 1.6中，用户在使用 HTTPS 时，将自动使用 HTTP/2

#### 1.9 Web应用的各个组成部分

​	Web应用就是一个执行以下任务的程序：

 	1. 通过HTTP协议，以HTTP请求报文的形式获取客户端输入
 	2. 对HTTP请求报文进行处理，并执行必要的操作
 	3. 生成HTML，并以HTTP响应报文的形式将其返回给客户端

Web应用被分成处理器（handler）和模板引擎（template engine）

##### 处理器

用MVC模式来讲，处理器既是控制器（controller），也是模型（model），在理想的MVC模型中，控制器应该只包含路由（routing）代码以及HTTP报文的解包和打包逻辑，而模型应该包含应用的逻辑以及数据。

> 模型-视图-控制器（Model-View-Controller，MVC）模式是编写Web应用时常用的模式，其中模型用于表示底层的数据，而视图则以可视化的方式向用户展示模型，至于控制器则会根据用户的输入对模型进行修改。

Web应用的本质是一个通过HTTP协议与用户互动的程序，只要能够实现这种互动，程序本身可以使用任何一种模式开发，甚至不使用模式也可以的。如可以通过控制器和模型进行合并，然后由处理器执行执行所有操作并向客户端返回响应的做法不仅可行，而且也十分合理。

##### 模板引擎

通过HTTP响应报文回传给客户端的HTML是由模板转换而成的，模板里可能包含HTML也可能不会。模板引擎通过模板和数据来生成最终的HTML。

模板分为**静态模板**和**动态模板**

- 静态模板是一些夹杂着占位符的HTML，静态模板引擎通过将静态模板中的占位符替换成响应的数据来生成最终的HTML，也成为无逻辑模板，CTemplate和Mustache都属于静态模板引擎。
- 动态模板除了包含HTML和占位符外，还包含一些编程语言结构，如条件语句。迭代语句、变量、JSP、ASP、ERB等都属于动态模板引擎

*处理器负责接收HTTP请求并处理它们*

*模板引擎负责生成HTML，这些HTML之后会作为HTTP响应的其中一部分被回传至客户端*



### ChitChat论坛

ChitChat关键特性：在这个论坛里面，用户可以注册账号，并在登录之后发表新帖子又或者回复已有的帖子；未注册用户可以查看帖子，但无法发表帖子或是回复帖子。

#### 2.1 应用设计

​	服务器会向客户端提供HTML页面，并通过页面的超链接向客户端表明请求的格式以及被请求的数据，而客户端则会在发送请求时向服务器提供相应的数据。

​	请求的格式通常是由应用自行决定的，比如ChitChat的请求使用的是以下格式：

​						`http://<服务器名><处理器名>?<参数>`

处理器的名字是按层级进行划分的，比如，对 /thread/read 这个处理器名字来说， thread 是被调用的模块，而 read 则是这个模块中负责读取帖子内容的处理器。该应用的参数会以URL查询的形式传递给处理器，而处理器则会根据这些参数对请求进行处理。

![服务器在典型Web应用中的工作流程](服务器在典型Web应用中的工作流程.jpg)

#### 2.2 数据模型

ChitChat数据模型：

- User——表示论坛的用户信息
- Session——表示论坛用户当前的登录会话
- Thread——表示论坛里面的帖子，每一个帖子都记录了多个论坛用户之间的对话
- Post——表示用户在帖子里面添加的回复

![Web应用访问数据存储系统的流程](Web应用访问数据存储系统的流程.jpg)

#### 2.3 请求的接收与处理

请求的接收和处理是所有Web应用的核心，Web应用的工作流程应该为：

1. 客户端将请求发送到服务器的一个URL上
2. 服务器的多路复用器将接收到的请求重定向到正确的处理器，然后由该处理器对请求进行处理
3. 处理器处理请求并执行必要的动作
4. 处理器调用模板引擎，生成相应的HTML并将其返回给客户端

##### 多路复用器（multiplexer）

`net/http`标准库提供了一个默认的多路复用器，这个多路复用器可以通过调用`NewServeMux`函数来创建：

`mux := http.NewServeMux()`

为了将发送至根URL的请求重定向到处理器，程序使用了`HandleFunc`函数：

`mux.HandleFunc("/",index)`

`HandleFunc`函数接收一个URL和一个处理器的名字作为参数，并将针对给定URL的请求转发至指定的处理器进行处理。此外，所有处理器都接受一个`ResponseWriter`和一个指向`Request`结构的**指针**作为参数，并且所有请求参数都可以通过访问`Request`结构得到。

注意：尽管处理器和处理器函数提供的最终结果是一样的，但他们实际上并不相同。

##### 服务静态文件

程序使用`FileServer`函数创建了一个能够为指定目录中的静态文件服务的处理器，并将这个处理器传递给了多路复用器的`Handle`函数，此外，程序还使用`StripPrefix`函数去移除请求URL中的指定前缀：

```go
files := http.FileServer(http.Dir("/public"))
mux.Handle("/static/",http.StripPrefix("/static/",files))
```

当服务器接收到一个以`/static/`开头的URL请求时，以上两行代码会移除URL中的`/static`字符串，然后在public目录中查找被请求的文件。比如，当服务器接收到一个针对文件`http://localhost/static/css/bootstrap.min.css`请求时，它将会在`public`目录中查找以下文件：

`<application root>/css/bootstrap.min.css`

当服务器成功找到这个文件后，会把它返回给客户端

##### 创建处理器函数

处理器函数实际上就是一个接受 `ResponseWriter` 和 `Request` 指针作为参数的Go函数

```go
func index(w http.ResponseWriter, r *http.Request) {
    files := []string{"templates/layout.html",
    				  "templates/navbar.html",
    				  "templates/index.html",}
    templates := template.Must(template.ParseFiles(files...))
    threads, err := data.Threads();if err == nil {
        templates.ExecuteTemplate(w, "layout", threads)
    }
}
```

index 函数负责生成 HTML 并将其写入 ResponseWriter 中。对 Go 语言来说，用户只需要把位于相同目录下的所有文件都设置成同一个包，那么这些文件就会与包中的其他文件分享彼此的定义。用户也可以把文件放到其他独立的包里面，然后通过导入（import）这些包来使用它们。

##### 使用cookie进行访问控制

当一个用户成功登陆以后，服务器必须在后续的请求中标示出这是一个已登录的用户，为了做到这一点，服务器会在响应的首部中写入一个 cookie ，而客户端在接收这个 cookie 之后则会把它存储到浏览器里面。

```go
func authenticate(w http.ResponseWriter,r *http.Request) {
    r.ParseForm()
    user,_ := data.UserByEmail(r.PostFormValue("password")) {
        session := user.CreateSession()
        cookie := http.Cookie{
            Name:"_cookie",
            Value:session.Uuid,
            HttpOnly:true,
        }
        http.SetCookie(w, &cookie)
        http.Redirect(w, r, "/", 302)
    } else {
        http.Redirect(w, r, "/login", 302)
    }
}
```

data.UserByEmail 函数通过给定的电子邮件地址获取与之对应的 User 结构，而data.Encrypt 函数则用于加密给定的字符串。

在验证用户身份的时候，程序必须先确保用户是真实存在的，并且提交给处理器的密码在加密之后跟存储在数据库里面的已加密用户密码完全一致。

CreateSession 方法创建一个 Session 结构，如下：

```go
type Session struct {
    Id int
    Uuid string
    Email string
    UserId int
    CreatedAt time.Time
}
```

Uuid 字段存储的是一个随机生成的唯一 ID ，这个 ID 是实现会话机制的核心，服务器会通过cookie 把这个 ID 存储到浏览器中，并把 Session 结构记录的各项信息存储到数据库中。

在创建了 Session 结构后，又创建了 Cookie 结构

```go
cookie := http.Cookie{
    Name:"_cookie",
    Value:session.Uuid,
    HttpOnly:true,
}
```

cookie 的名字是随意设置的，而 cookie 的值则是将要被存储到浏览器里面的唯一 ID 。因为程序没有给 cookie 设置过期时间，所以这个 cookie 就成了一个会话 cookie ，它将在浏览器关闭时自动被移除。程序将 HttpOnly 字段的值设置成了 true ，这意味这个 cookie 只能通过HTTP或者 HTTPS 访问，但却无法通过 JavaScript 等非 HTTP API 进行访问。

设置完cookie后,将其添加到响应的首部:

```go
http.SetCookie(writer, &cookie)
```

然后在处理器函数里面检查当前访问的用户是否已经登录,因此创建一个名为 session 的工具(util)函数,并在各个处理器函数里面复用它。

```go
func session(w http.ResponeWriter,r *http.Request) (sess data.Session,err error) {
    cookie, err := r.Cookie("_cookie")
    if err == nil {
        session = data.Session(Uuid: cookie.Value)
        if ok, _ := sess.Check(); !ok {
            err = errors.New("Invaild session")
        }
    }
    return
}
```

从请求中取出 cookie ，session 函数使用了以下代码：

```go
cookie, err := r.Cookie("_cookie")
```

如果 cookie 不存在,俺么很明显用户并未登录,相反,如果 cookie 存在,那么 session 函数将继续进行第二项检查--访问数据库并核实会话的唯一 ID 是否存在.

#### 2.4 使用模板生成HTML响应

index 处理器函数里面的大部分代码都是用来为客户端生成 HTML 的.函数把每个用到的模板文件都放到了 Go 的切片里.切片指定的三个 HTML 文件都包含了特定的嵌入命令,这些命令被成为动作 (action) ,动作会在HTML文件里面会被 `{{` 符号和 `}}` 符号包围

程序会调用 ParseFiles 函数对这些模板文件进行语法分析,并创建出相应的模板,程序使用了 Must 函数去包围 ParseFiles 函数的执行结果,这样当 ParseFiles 返回错误的时候, Must 函数就会向用户返回相应的错误报告.

模板文件和模板一一对应的做法可以给开发带来方便.

程序通过调用 ExecuteTemplate 函数,执行 (execute) 已经经过语法分析的 layout 模板.执行模板意味着把模板文件中的内容和来自其他渠道的数据进行合并,然后生成最终的 HTML 内容.

##### 整理代码

```go
func generateHTML(w http.ResponseWriter, data interface{}, fn ...string){
    var files []string
    for _, file := range fn {
        files = append(files, fmt.Sprintf("templates/%s.html", file))
    }
    templates := template.Must(template.ParseFiles(files...))
    templates.ExecuteTemplate(w, "layout", data)
}
```

generateHTML 函数接受一个 ResponseWriter、一些数据及一系列模板文件作为参数.

一个空接口就是一个空集合,这意味这任何类型都可以成为一个空接口,也就是说任何类型的值都可以传递给函数作为参数.

generateHTML 函数的最后一个参数以3个点(...) 开头,表示函数是一个可变参数函数,这意味着这个函数可以在最后的可变参数中接受零个或任意多个值作为参数.在 Go 语言里面,可变参数必须是可变参数函数的最后一个参数.

![模板引擎通过合并数据和模板来生成HTML](模板引擎通过合并数据和模板来生成HTML.jpg)

#### 2.5 安装PostgreSQL

Ubuntu: `sudo apt-get install postgresql postgresql-contrib `

登入Postgres账号: `sudo su postgres`

使用 `createuser` 命令创建一个账号: `createuser --interactive <role name>`

使用 `createdb` 命令创建以你账号名字命名的数据库:

`createdb <YOUR ACCOUNT NAME>`

postgreSQL 导入 sql 文件：执行  psql -d 数据库名 -h ip地址 -p 数据库端口 -U 用户名 -f 文件地址

#### 2.6 连接数据库

![通过结构模型连接数据库和处理器](通过结构模型连接数据库和处理器.jpg)

#### 2.7 Web 应用运作流程回顾

1. 客户端向服务器发送请求
2. 多路复用器接收到请求,并将其重定向到正确的处理器
3. 处理器对请求进行处理
4. 在需要访问数据库的情况下,处理器会使用一个或多个数据结构,这些数据结构都是根据数据库中的数据建模而来的
5. 当处理器调用与数据结构有关的函数或者方法时,这些数据结背后的模型会与数据库进行连接,并执行相应的操作.
6. 当请求处理完毕后,处理器会调用模板引擎,有时候还会向模板引擎传递一些通过模型获取到的数据
7. 模板引擎会对模板文件进行语法分析并创建相应的模板,而这些模板又会与处理器传递的数据一起合并生成最终的HTML
8. 生成的HTML会作为响应的一部分回传至客户端.

![Web应用工作流程概览](Web应用工作流程概览.jpg)

#### 小结

- 请求的接收和处理是所有Web应用的核心
- 多路复用器会将 HTTP 请求重定向到正确的处理器进行处理，针对静态文件的请求也是如此
- 处理器函数是一种接受 ResponseWriter 和 Request 指针作为参数的 Go 函数
- cookie 可以用作一种访问控制机制
- 对模板文件以及数据进行语法分析会产生相应的 HTML ，这些 HTML 会被用作返回给浏览器的响应数据
- 通过使用 sql 包以及相应的 SQL 语句，用户可以将数据持久地存储在关系数据库中