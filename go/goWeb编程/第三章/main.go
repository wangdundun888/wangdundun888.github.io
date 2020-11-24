/*
	理解那些隐藏在框架之下的底层概念和基础设施是非常重要的
	2020/11/23
*/
package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

type Hello struct{}

//在go语言中,一个处理器就是一个拥有ServeHTTP方法的接口
//第一个参数是http.ResponseWriter接口,第二个是http.Request指针
func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world I'm go")
}

type Hi struct{}

func (h *Hi) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hi i'm go!!!")
}

func helloAgain(w http.ResponseWriter, r *http.Request) { //处理器函数
	fmt.Fprint(w, "hello again")
}

func getTime(h http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, time.Now().String(), "\n")
		h(writer, request)
		//如果是handler,则调用serveHTTP方法
		//h.serverHTTP(w,r)
	}
}
func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "hello,%s!\n", p.ByName("name"))
}

func main() {
	//ServeMux是一个HTTP请求多路复用器
	//ServeMux结构包含一个映射,这个映射会将URL映射至相应的处理器
	//HttpRouter是一个高效的轻量级第三方多路复用器
	mux := httprouter.New()
	mux.GET("/hello/:name", hello)
	server := http.Server{
		Addr:    "localhost:9090",
		Handler: mux,
		/*
			Handler: &Hello{},
				如果是Handler:nil 或者干脆省略该参数,则默认使用DefaultServeMux
			DefaultServeMux不仅仅是一个多路复用器,也是一个Handler,是一个特殊的Handler
			它唯一要做的就是根据请求的URL将请求重定向到不同的处理器

			当使用&Hello{}构建一个http.Server时,它代替了原来的多路复用器,此时不会再通过匹配URL来
			将不同的请求发到不同的处理器,而是将所有的请求都发送到&Hello{}处理,所以,此时无论访问什么URL,
			结果都是一样的
			这也是需要使用多路复用器的原因
		*/
	}
	//很多时候我们需要多个处理器去处理不同的URL请求,所以要用默认的DefaultServeMux作为处理器
	//使用http.Handle方法将handler绑定到DefaultServeMux
	//此时如果访问没有绑定的url会404
	http.Handle("/hello", &Hello{})
	http.Handle("/hi", &Hi{})
	/*
		处理器函数:
			与ServeHTTP方法拥有相同的签名
			eg:hello(w http.ResponseWriter,r *http.Request)
		而http.HandleFunc方法可以将处理器函数转换为一个Handler,再将该Handler绑定到DefaultServeMux
		这样可以简化代码,但并不能完全代替处理器
	*/
	http.HandleFunc("helloAgain", helloAgain)

	//串联多个处理器多个处理器函数
	//一个切面思想,类似于spring的aop
	http.HandleFunc("/time", getTime(helloAgain))

	server.ListenAndServe()
}
