package main

import (
	"fmt"
	"net/http"
)

/*type Response struct {
	// 状态码 例如"200 OK"
	Status string

	//状态码的数字部分 例如 200
	StatusCode int

	//协议版本号 例如"HTTP/1.1"
	Proto string

	ProtoMajor int

	ProtoMinor int

	//Header保管头域的键值对
	//如果有多个头的键相同,对应的值逗号串联起来保存
	Header http.Header

	//回复的主体
	Body io.ReadCloser

	//记录相关内容的长度
	//-1代表长度未知
	ContentLength int64

	//传输编码
	TransferEncoding []string

	Close bool

	//额外的头域键值对
	Trailer http.Header

	//获取此回复的请求
	Request *http.Request

	TLS *tls.ConnectionState
}*/

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "hello go")
}

func writeExample(writer http.ResponseWriter, request *http.Request) {
	str := `<html>
	<head><title>Go web Programming</title></head>
	<body><h1>Hello Go</h1></body>
	</html>`
	//Write方法接收一个字节数组作为参数
	//并将数组中的字节写入HTTP响应的主体中
	//如果没有为首部设置相应的内容类型,则由写入的前512个字节决定
	writer.Write([]byte(str))
}

func writeHeaderExample(write http.ResponseWriter, request *http.Request) {
	//WriteHeader设置响应的返回状态码
	//该方法调用后,就不能对响应的首部做任何的写入操作
	//用户调用Write之前如果没有调用该方法,则默认使用200OK作为响应的状态码
	write.WriteHeader(501)
	fmt.Fprintln(write, "No such service try next door")
}

func headerExample(write http.ResponseWriter, request *http.Request) {
	//Header方法返回一个由首部组成的映射
	write.Header().Set("Location", "http://www.bing.com")
	//WriteHeader写入之后就不能对头部进行写入,所以需要先写入Location
	write.WriteHeader(302)
}

func main() {
	/*
		关于ResponseWriter
			http.response是一个非导出结构,所以只能使用ResponseWrite结构
			是一个带有结构指针的接口
			包含有Write WriteHeader Header三个方法
	*/
	serve := http.Server{
		Addr:    "localhost:9090",
		Handler: nil,
	}
	http.HandleFunc("/", index)
	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", writeHeaderExample)
	http.HandleFunc("/header", headerExample)
	serve.ListenAndServe()
}
