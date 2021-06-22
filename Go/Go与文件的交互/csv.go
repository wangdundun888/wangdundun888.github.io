package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

func main() {
	//csv全称Comma-Separated Values 逗号分隔值文件格式
	//csvFile, err := os.Create("posts.csv")
	//if err != nil {
	//	panic(err)
	//}
	//defer csvFile.Close()
	//
	//allPosts := []Post{
	//	Post{Id: 1, Content: "hello world!", Author: "Wdc"},
	//	Post{Id: 2, Content: "hello world again!", Author: "Wdc"},
	//}
	//
	////调用csv.NewWriter创建一个csv写入器,参数是一个*File
	//writer := csv.NewWriter(csvFile)
	//
	//for _, post := range allPosts {
	//	line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
	//	//通过写入器的Write方法,将一个字符串切片写入
	//	err := writer.Write(line)
	//	if err != nil {
	//		panic(err)
	//	}
	//}
	////保证缓冲区的数据都写入文件内
	//writer.Flush()

	file, err := os.Open("test2.csv")
	if err != nil {
		_, _ = os.Create("test2.csv")

	} else {
		fileInfo, _ := os.Stat(file.Name())
		winFileAttr := fileInfo.Sys().(*syscall.Win32FileAttributeData)
		creatTime := SecondToTime(winFileAttr.CreationTime.Nanoseconds() / 1e9)
		nowTime := time.Now()
		fmt.Println("文件创建时间：", creatTime.Unix())
		fmt.Println("文件创建时间：", creatTime)
		fmt.Println("现在的时间:", nowTime)
		fmt.Println("现在的时间:", nowTime.Unix())
		diff := nowTime.Unix() - creatTime.Unix()
		fmt.Println("相差时间为", diff/(60*60), "小时")
		defer file.Close()
	}

	////调用csv.NewReader创建一个读取器
	//reader := csv.NewReader(file)
	////FieldsPerRecord 每条记录的读取字段数量
	////如果该值为正值,那么当读取器从CSV文件读取的字段数量少于这个值时,Go会抛出一个错误
	////如果该值为0,则读取器会将第一条记录的字段数量作为该值的值
	////如果该值为负数,读取时,即使记录少了某些字段,读取也不会中断
	//reader.FieldsPerRecord = -1
	//record, err := reader.ReadAll()
	//if err != nil {
	//	panic(err)
	//}
	//for _, item := range record {
	//	id,_:= strconv.Atoi(item[0])
	//	fmt.Println(id, " ", item[1])
	//}

}

func SecondToTime(sec int64) time.Time {
	return time.Unix(sec, 0)
}
