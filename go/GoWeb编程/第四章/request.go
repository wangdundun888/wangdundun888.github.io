package main

import (
	"fmt"
	"io/ioutil"
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

func getHeader(writer http.ResponseWriter, request *http.Request) {
	header := request.Header
	//获取request的头部
	//header可以当作一个[string][]string使用
	for key, value := range header {
		fmt.Println(key, ": ", value)
	}
	fmt.Fprint(writer, header)
}
func getBody(writer http.ResponseWriter, request *http.Request) {
	length := request.ContentLength
	body := make([]byte, length)
	//用Body的read方法,把Body读入一个byte数组
	//长度就是request.ContentLength
	request.Body.Read(body)
	fmt.Println(string(body))
	fmt.Fprint(writer, string(body))
}
func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "hello go")
}

func getForm(writer http.ResponseWriter, request *http.Request) {
	//获取Form字段前,要先调用ParseForm进行语法分析
	request.ParseForm()
	//Form字段是一个map[string][]string映射
	//键是字符串,值是字符串切片
	fmt.Println("form字段: ", request.Form)
	//你好 [Zhangsan] 访问name属性值方式之一
	fmt.Fprintf(writer, "你好 %s\n", request.Form["name"])
	//你好 Zhangsan	方式之二 该方法调用时会自动调用ParseForm或ParseMultipartForm方法
	fmt.Fprintf(writer, "你好 %s\n", request.FormValue("name"))
}

func getPostForm(writer http.ResponseWriter, request *http.Request) {
	//获得PostForm字段需先调用ParseForm进行语法解析
	request.ParseForm()
	//PostForm和Form字段都是一个映射
	//不同的是,PostForm字段不包含URL键值对,只包含键的表单值
	//且PostForm只支持application/x-www-form-urlencoded编码
	fmt.Println("PostForm字段: ", request.PostForm)
}

func getMultipartForm(writer http.ResponseWriter, request *http.Request) {
	//获取MultipartForm字段前要先调用此方法
	//参数是maxMemory是指从表单里获取maxMemory字节的数据
	request.ParseMultipartForm(1024)
	//MultipartForm和PostForm字段一样,只包含表键值对,不包含URL键值对
	//同时,MultipartForm不是一个映射,而是一个包含两个映射的结构
	//第一个映射和Form字段一样
	//第二个映射是记录用户上传的文件
	//该字段用来获取multipart/form-data编码的表单数据
	//使用multipart/form-data编码时,表单数据存储到MultipartForm字段
	fmt.Println("MultipartForm字段: ", request.MultipartForm)
}

func uploadFile(write http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(1024)
	/*
		根据键值,即  <input type="file" name="fileName" /> name属性 从
		MultipartForm字段的File字段取出文件头fileHeader
		跟FormValue类似,也可以使用
			file,_,err := request.FormFile("fileName")
		提前,无须调用ParseMultipartForm方法,同时第二个参数是返回fileHeader
	*/
	fileHeader := request.MultipartForm.File["fileName"][0]
	file, err := fileHeader.Open()
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintf(write, string(data))
		}
	}
}

func main() {
	serve := http.Server{
		Addr: "localhost:9090",
	}
	http.HandleFunc("/", index)
	http.HandleFunc("/header", getHeader)
	http.HandleFunc("/body", getBody)

	http.HandleFunc("/form", getForm)
	http.HandleFunc("/postForm", getPostForm)
	http.HandleFunc("/multipartForm", getMultipartForm)

	http.HandleFunc("/file", uploadFile)

	serve.ListenAndServe()
}
