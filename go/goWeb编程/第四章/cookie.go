package main

import (
	"fmt"
	"net/http"
)

/*
type Cookie struct {
	Name string
	Value string

	//可选
	Path string

	//可选
	Domain string

	//可选
	//没有设置该字段的cookie通常称为会话cookie或临时cookie,
	//这种cookie会在浏览器关闭时自动被移除
	//设置了该字段的称为持久cookie
	//这种cookie会一直存在直到过期或手动移除
	Expires time.Time

	Expires 和 MaxAge都可以用来设置cookie的过期时间
	其中,Expires在指定cookie应该在什么时候过期
	而MaxAge则指明cookie在创建后存活多少秒
	HTTP/1.1废弃了Expires,推荐使用MaxAge,但几乎所有的浏览器都支持Expires

	//仅可读的cookie
	RawExpires string

	//MaxAge=0表示未设置'Max-Age'属性
	//MaxAge<0表示立即删除cookie,等价于'Max-Age:0'
	//MaxAge>0表示存在'Max-Age'属性,单位为秒
	MaxAge int

	Secure bool

	HttpOnly bool

	SameSite http.SameSite

	Raw string

	//未经过解析的"属性-值"对的原始文本
	Unpared []string
}*/

func setCookis(writer http.ResponseWriter, request *http.Request){
	c1 := http.Cookie{
		//发现Name中的字符串不可以有空格,否则写入失败,不知道为什么
		//貌似是因为cookie支持的原因,尽量使用下划线而不使用其他符号
		Name: "first_cookie",
		Value: "GO cookie",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name: "second_cookie",
		Value: "GO cookie",
		HttpOnly: true,
	}
	c3 := http.Cookie{
		Name: "third_cookie",
		Value: "GO cookie",
		HttpOnly: true,
	}
	//设置cookie的三种方法
	writer.Header().Set("Set-Cookie",c1.String())
	writer.Header().Add("Set-Cookie",c2.String())
	http.SetCookie(writer,&c3)
}

func getCookis(writer http.ResponseWriter, request *http.Request){
	//从请求首部获得cookie
	//该方法返回一个切片
	//如果想要获得单独的键值对格式的cookie,需要自行进行语法分
	//或使用例外一种方式
	c := request.Header["Cookie"]

	//获取单个cookie
	c1 ,err := request.Cookie("first_cookie")
	if err == nil {
		fmt.Println(c1)//获得一个first_cookie="GO cookie"字符串
	}
	//该方法和request.Header["Cookie"]获取的完全相同
	cc := request.Cookies()
	fmt.Println(cc)
	fmt.Fprintln(writer,c)
	fmt.Fprintln(writer,cc)
}


func main() {
	serve := http.Server{
		Addr: "localhost:9090",
		Handler: nil,
	}
	http.HandleFunc("/setcookie",setCookis)
	http.HandleFunc("/getcookie",getCookis)
	serve.ListenAndServe()
}
