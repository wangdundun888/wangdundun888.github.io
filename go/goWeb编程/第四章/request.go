package main

import (
	"fmt"
	"net/http"
)

/*//Request结构
type Request struct{
		//Method指定HTTP方法(GET POST PUT等
		//对于客户端,""空字符串代表GET
		Method string

		//URL在服务端表示被请求的URI,在客户端表示要访问的URL
		//在服务端,URL字段是解析请求行的URI得到的
		URL *url.URL

		//接收到的请求的协议版本 HTTP/1.1 or HTTP/2
		Proto string
		ProtoMajor int  //1
		ProtoMinor int //1

		//Header字段用来表示HTTP请求的头域
		//格式如下:
		//	Header = map[string][]string{
		//		"Accept-Encoding": {"gzip, deflate"},
		//		"Accept-Language": {"en-us"},
		//		"Connection": {"keep-alive"},
		//	}
		Header http.Header

		//请求的主体
		//对于客户端,Body为nil代表请求没有主体
		//Client的Transport字段会负责调用Body的Close方法
		//在服务端,Body字段总是非nil的,但在没有主体时,读取Body会立刻返回EOF
		//Server会关闭请求的主体,ServerHTTP处理器不需要关闭Body字段
		Body io.ReadCloser

		//GetBody 返回一个Body的副本
		GetBody func()(io.ReadCloser,error)

		//记录主体的长度
		//值为-1时,代表着长度位置
		//>=0时,代表可以从Body读取的长度
		//对于客户端,如果值为0且Body不为nil,代表Body长度未知
		ContentLength int64

		//按照从最外到最里的顺序列出传输编码
		//如果为空,表示"identity"编码
		TransferEncoding []string

		//是否需要关闭连接
		//对于服务端,是在回复请求之后
		//对于客户端,实在发送请求和得到回复之后
		Close  bool

		//在服务端,Host指定URL会在其上寻找资源的主机
		//Host的格式可以是:"host:port"
		//在客户端,如果该字段为"",Request.Write方法会使用URL字段的Host
		Host string

		//解析好的表单数据,包括URL字段的query参数和POST或PUT的表单数据
		//该字段只有调用ParseForm之后才有效
		//在客户端,会忽略请求中的本字段而使用Body替代
		Form url.Values

		//解析好的POST或PUT的表单数据
		//与Form相似,同样需要在ParseForm调用后生效,在客户端用Body字段替代
		PostForm url.Values

		//解析好的多部件表单,包括上传的文件
		//该字段只有在调用ParseMultipartForm后才有效
		//在客户端同样被Body字段替代
		MultipartForm *multipart.Form

		//指定了会在请求主体发送的额外的头域
		//读取Body返回EOF后,该字段才会更新,此时才可以访问本字段
		Trailer http.Header

		//被客户端发送到服务端的请求的请求行中未修改的请求URI
		RequestURI string

		//请求的来源地址
		RemoteAddr string

		//暂时不想看
		TLS *tls.ConnectionState

		//
		Cancel <-chan struct{}

		//客户端重定向时使用
		Response *http.Response

		//
		ctx context.Context
}*/

func main() {
	serve := http.Server{
		Addr: "localhost:9090",
	}
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		header := request.Header
		fmt.Println(request.Body)
		for key, value := range header {
			fmt.Println(key, ": ", value)
		}
		fmt.Fprint(writer, "hello go web")
	})
	serve.ListenAndServe()
}
