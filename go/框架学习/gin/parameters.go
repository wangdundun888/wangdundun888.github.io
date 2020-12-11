package main

//import (
//	"github.com/gin-gonic/gin"
//	"net/http"
//)
//
//func index(c *gin.Context){
//	c.JSON(200,gin.H{
//		"message" : "pong",
//	})
//}
//
//func paramExample(c *gin.Context){
//	name := c.Param("name")
//	c.String(http.StatusOK,"hello %s\n",name)
//}
//
//func paramExample1(c *gin.Context){
//	name := c.Param("name")
//	action := c.Param("action")
//	message := name + " is " + action
//	c.String(http.StatusOK,message)
//}
//
//func welcome(c *gin.Context){
//
//	//该方式用来查询url键值对
//
//	//DefaultQuery方法提供一个默认值
//	//当查询不到指定值时用默认值代替
//	firstname := c.DefaultQuery("firstname","Guest")
//	//该方法是 c.Request.URL.Query().Get("lastname")的缩写
//	lastname := c.Query("lastname")
//
//	c.String(http.StatusOK,"hello %s %s",firstname,lastname)
//}
//
//func postForm(c *gin.Context){
//	//同DefaultQuery
//	//假如name没有进行填写,为空,也不会用"Guest"代替
//	//除非name属性不在表单中提交
//	name := c.DefaultPostForm("name","Guest")
//	message := c.PostForm("message")
//
//	c.String(http.StatusOK,"hello %s \nyou said : %s",name,message)
//
//}
//
//
//
//func main() {
//	r := gin.Default()
//
//	r.GET("/",index)
//
//	//这种方式可以获取路径参数 : 冒号方式
//	//可以匹配/user/bob,但不会匹配/user 或 /user/ 会404
//	//  /user/user/user也是404
//	r.GET("/user/:name",paramExample)
//
//	//使用*号获取路径参数 不同与冒号,星号可为空
//	//可以匹配/user/john/send 和 /user/john/
//	//如果没有路由匹配/user/john ,那么此路径会被重定向到/user/john/
//	r.GET("/user/:name/*action",paramExample1)
//
//	//查询url键值对
//	r.GET("/welcome",welcome)
//
//	//查询post表单
//	r.POST("/post_form",postForm)
//
//	//如果参数作为映射提交,那么可以使用QueryMap方法和PostFormMap方法获取
//
//	//Run方法的参数为空时,默认监控:8080
//	//可使用字符串形式"localhost:9090"或":9090"监控自定义窗口
//	r.Run(":9090")
//}
