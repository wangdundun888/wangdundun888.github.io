# 简单的HTTP协议
### 1
```
    HTTP协议用于客户端和服务器端之间的通信
```
### 2
```
    请求报文由①请求方法、②请求URI、③协议版本、④可选的请求首部字段和⑤内容实体构成的
    例：
        ①POST  ②/form/entry ③HTTP/1.1
        HOST：hacke.jp                |
        CONNECTION:keep-alive         | 
        CONTENT-TYPE:application/xx   | ④
        CONTENT-LENGTH:16             |
        ⑤name=uneo&age=37
    响应报文由协议版本、状态码、解释状态码的原因短语、可选的响应首部字段和内容实体组成
```
### 3.HTTP是不保存状态的协议
```
    无状态协议(stateless):协议对于发送过的请求或响应都不做持久化处理
    目的是为了更快地处理大量事务
    后期引入了cookie技术管理状态
```
### 4.请求URI定位资源
```
    HTTP协议使用URI定位互联网上的资源
    扩展:
        URI(Uniform Resource Identifier 统一资源标识符)
        URL(Uniform Resource Locator 统一资源定位符)
        URN(Uniform Resource Name 统一资源名称)
        URI包含URL和URN,URL是URI的子集
        URI类比于身份证号码,独一无二
        URL类比于名字,例如张三
```
### 5.HTTP方法
```
    ① GET:获取资源
    ② POST:传输实体主体
    ③ PUT:传输文件 //存在安全性问题,一般不用,除非配合验证机制
    ④ HEAD:获得报文首部 //和GET方法一样,但不返回报文主体,常用于确认URI的有效性等
    ⑤ DELETE:删除文件 //与PUT方法相似
    ⑥ OPTIONS:询问服务器端支持的方法
    ⑦ TRACE:追踪路径
    ⑧ CONNECT:要求用隧道协议连接代理
```
### 6.持久连接节省通信量
```
    只要任意一端没有明确提出断开连接,则保持TCP连接状态
```
### 7.管线化
```
    并行发送多个请求,不需要一个接着一个等待响应
```
### 8.使用cookie的状态管理
```
    通过在请求和响应报文中写入cookie信息来控制客户端的状态
    过程:
        服务器端发送的响应报文中有set-cookie的首部字段信息
        -> 通知客户端保存cookie
        -> 客户端下次发送请求时,在请求报文中写入cookie
        -> 服务器端根据cookie进行对比验证
```