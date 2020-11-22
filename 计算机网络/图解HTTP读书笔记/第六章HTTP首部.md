## HTTP首部
### 1 通用首部字段
``` 
    通用首部字段:在请求报文或响应报文都可用

    1 Cache-Control 
        控制缓存的行为
        参数可选,多个指令之间用逗号分割
        例: Cache-Control:private,max-age=0,no-cache
    2 Connection
        控制不再转发给代理的首部字段
        管理持久连接
        eg:Connection:close
    3 Date
        表明创建HTTP报文的日期和时间
        eg:Date : Tue Jul 03 05:05:59 2020
    4 Pragma
        HTTP/1.1之前版本的遗留字段,只用在客服端发送的请求中
        要求所有的中间服务器不返回缓存的资源
        eg:Pragma:no-cache
    5 Trailer
        事先说明在报文主体后记录了哪些首部字段
        eg:Trailer:Expires
    6 Transfer=Encoding
        规定了传输报文主体时采用的编码方式
        eg:Transfer-Encoding:chunked
    7 Upgrade
        用于检测HTTP协议及其他协议是否可使用更高的版本进行通信
        eg:Upgrade:TLS/1.0
    8 Via
        为了追踪客户端与服务器之间的请求和响应报文的传输路径
            经过代理或网关时,会先在首部字段Via中附加该服务器的信息
        还可以避免请求回环的发生,经过代理时附加该首部字段内容
    9 Warning 
        从HTTP/1.0的响应首部(Retry-After)演变过来的
        会告知用户一些与缓存相关的问题的警告
```
### 2 请求首部字段
```
    请求首部字段:从客户端往服务器端发送请求报文中所使用的字段

    1 Accept
        通知服务器用户代理能够处理的媒体类型及媒体类型的相对优先级
        可使用type/subtype这种形式,一次性指定多种媒体类型
        eg:Accept:text/html,application/xhtml+xml;q=0
        q为优先级,默认值为1.0,用分号分隔
    2 Accept-Charset
        通知服务器用户代理支持的字符集及字符集的相对优先顺序
        可以一次性指定多种字符集,q为权重优先级
        eg:Accept-Encoding:iso-8859-5
    3 Accept-Encoding
        告知服务器用户代理支持的内容编码及内容编码的优先级顺序
        eg:Accept-Encoding:gzip
    4 Accept-Language
        告知服务器代理能够处理的自然语言集,以及自然语言集的相对优先级
        可以一次指定多种自然语言集
        eg:Accept-Language:zh-cn
    5 Authorization
        告知服务器用户代理的认证信息
        eg:Authorization:Basic dsflaFAdfaAF21==
    6 Expect
        告知服务器,期待出现某种特定行为
        eg:Expect:100-continue
    7 From
        告知服务器使用用户代理的用户的电子邮件地址
        eg:From:glut@edu.com
    8 Host
        告知服务器,请求的资源所处的互联网主机名和端口号
        在 HTTP/1.1规范内是唯一一个必须被包含在请求内的首部字段
        因为可能有多个虚拟主机运行在同一个IP上
    9 If-xxx   
        暂时先跳过...
    14 Max-Forwards
        最大转发次数,当该值为0,不再进行转发,而是直接返回响应
    15 Proxy-Authorization
        与5 Authorization相似,多了个Proxy,是指收到从代理服务器发来的认证质询
    16 Range
        获取部分资源的范围请求
        eg:Range:bytes=5001-10000
        服务器如是可以处理该字段的请求,则返回206 Partial Content的响应,若无法处理请求,
        会返回200 OK的响应及全部资源
    17 Referer
        告知服务器请求的原始资源的URI
        服务器查看Referer就能知道请求的URI是从哪个Web页面发起的
    18 User-Agent
        讲创建请求的浏览器和用户代理名称等信息传达给服务器
```
### 3 响应首部字段
```
    响应首部字段:由服务器端向客户端返回响应报文中所使用的字段

    1 Accept-Ranges
        告知客户端服务器能否处理范围请求
        两个取值,bytes:可以处理,none:不能处理
    2 Age
        告知客户端,源服务器在多久前创建了响应
        eg:Age:600    单位为秒
        代理创建响应时必须加上该首部字段
    3 ETag
        告知客户端实体标识,是一种可将资源以字符串形式做唯一性标识的方式
        强ETag值:无论实体发生多么细微的变化都会改变其值
        弱ETag值:发生了根本改变,产生差异时才会改变
    4 Location
        将响应接收方引导至某个与请求URI位置不同的资源
        通常配合3XX Redirection响应
    5 Proxy-Authenticate
        把由代理服务器所要求的认证信息发送给客户端
    6 Retry-After
        告知客户端应该在多久之后再次发送请求
        eg:Retry-After:120 秒
    7 Server
        告知客户端当前服务器上安装的HTTP服务器应用程序的信息
        eg:Server:Apache/2.2.17(Unix)
```
### 4 实体首部字段
```
    实体首部字段:包含在请求报文和响应报文中的实体部分所使用的首部

    1 Allow
        用于通知客户端能够支持Request-URI指定资源的所有HTTP方法
        当服务器接收到不支持的HTTP方法时,会以状态码405 Method Not Allowed
        返回,并把支持的所有HTTP方法写入该首部字段
        eg:Allow:GET,HEAD
    2 Content-Encoding
        告知客户端服务器对实体的主体部分选用的内容编码方式
        eg:Content-Encoding:gzip
    3 Content-Language
        eg:Content-Language:zh-CN
    4 Content-Length
        Content-Length:15000
        单位是字节,若实体主体已经进行内容编码,即使用了Content-Encoding字段
        则该字段不能用
    5 Content-Location
        给出与报文主体部分相对应的URI
    6 Content-MD5
        客户端会对接收的报文主体执行相同的MD5算法,然后检查报文是否保持完整
    7 Content-Range
        eg:Content-Range:bytes 5001-10000/10000
        表示当前发送部分及整个实体的大小
    8 Content-Type
        eg:Content-Type:text/html;charset=UTF-8
    9 Expires  
        将资源失效的日期告知客户端
    10 Last-Modified
        指明资源最终修改的时间
```
### 5 为Cookie服务的首部字段
```
    1 Set-Cookie
        响应首部字段,诸多属性,以PropertyName=Value方式设置
    2 Cookie
        请求首部字段
```