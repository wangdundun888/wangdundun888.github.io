package main

//func main() {
//	//禁用控制台颜色,因为写入log文件不需要
//	gin.DisableConsoleColor()
//
//	//创建文件,并让在此写入log
//	f,_ := os.Create("gin.log")
//	gin.DefaultWriter = io.MultiWriter(f)
//
//	//如果需要将日志写入文件的同时在控制台输出,则如下
//	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
//
//
//	r := gin.Default()
//
//	r.GET("/", func(c *gin.Context) {
//		c.String(http.StatusOK,"hello gin write log now")
//	})
//
//	r.Run(":9090")
//}
