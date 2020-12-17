package main

//import (
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"log"
//	"net/http"
//	"strconv"
//	"time"
//)
//
//func uploadFile(c *gin.Context){
//	//该方式适合single即单个文件上传
//	file,_ := c.FormFile("fileName")
//
//	//这里有雷,string(time.Now().Unix())是一个乱码
//	//结果并不是像我理解的那样,把一个int64的值转化为字符串
//	//比如string(123456789)应该转化为"123456789"
//	//解释:string方法就是把一个整数值当作ASCII编码转化为一个字符串,超过范围就会乱码
//	//所以string(123456789)是不对的, strconv.Itoa方法可以将一个int值转为一个字符串
//	//但如果需要将int64甚至int128转为字符串,则此方法不可用
//	//strconv.FormatInt(int64,base)
//	//该函数将一个int64的值转化为一个base进制的字符串
//	dst := "d:\\"+ strconv.FormatInt(time.Now().Unix(),10)+".bat"
//
//	log.Println(file.Filename)
//
//	//存储在d盘下,名字以时间命名
//	c.SaveUploadedFile(file,dst)
//
//	c.String(http.StatusOK,fmt.Sprintf("'%s' uploaded!\n'%s'",file.Filename,dst))
//}
//
//func uploadFiles(c *gin.Context){
//	form,_ := c.MultipartForm()
//	//该方法的参数为type=file表单的name属性,此处name="upload[]"
//	//多个file狂
//	files,_ := form.File["upload[]"]
//
//	for _,file := range files {
//		log.Println(file.Filename)
//	}
//	c.String(http.StatusOK,fmt.Sprintf("uploaded %d files",len(files)))
//}
//
//func main() {
//	r := gin.Default()
//
//	//设置内存限制,默认为32MiB
//	//1KB = 1000MB
//	//1KiB = 2^10 = 1024 B
//	r.MaxMultipartMemory = 8 << 20 // 8MiB
//
//	r.POST("/file",uploadFile)
//
//	r.POST("/files",uploadFiles)
//
//	r.Run(":9090")
//}
