package main

/*
	学习在go中如何生成随机数
	源码地址:https://github.com/golang/go/blob/master/src/math/rand/rand.go?name=release#31
*/
//Source代表均匀分布 随机数的范围是[0,1<<63)
type Source interface {
	Int63() int64
	Seed(seed int64)
}

func main() {

}
