package main

//import (
//	"github.com/gin-gonic/gin"
//	"net/http"
//)
//
//func index(c *gin.Context){
//	c.String(http.StatusOK,"hello Grouping Routers")
//}
//
//func main() {
//	r := gin.Default()
//
//	//通过/v1/hello访问
//	v1 := r.Group("/v1")
//	v1.GET("/hello",index)
//
//	//通过/v2/hello访问
//	v2 := r.Group("/v2")
//	v2.GET("/hello",index)
//
//	r.Run(":9090")
//}
