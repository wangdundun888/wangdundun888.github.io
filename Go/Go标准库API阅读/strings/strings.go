/*
	string类型作为go的基本类型并没有方法,所以阅读strings包里的strings.Go
	源码:https://github.com/golang/go/blob/master/src/strings/strings.go
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello world"
	str1 := "hello"
	fmt.Println("字符串str: ", str, "\n字串1: ", str1)
	fmt.Println("str是否包含字串1: ", strings.Contains(str, str1))
	/*
		func Contains(s, substr string) bool
		s是否包含substr,这不是重点,重点是函数引用了Index,
		Index:
			func Index(s, substr string) int
		返回substr第一次在s中出现的位置,如果s中没有substr,返回-1
		Index函数的实现和数组s[i:j]相关,暂不阅读,纯粹实现应该是使用KMP算法
	*/

	//EqualFold,判断两个utf-8编码字符串是否相同,忽视大小写
	fmt.Println(strings.EqualFold("Go", "Go"))

	//func HasPrefix(s, prefix string) bool
	//HasPrefix,判断s中是否包含前缀prefix
	//return len(s) >= len(prefix) && s[0:len(prefix)] == prefix
	// 先判断len(prefix),避免s[0:len(prefix)]越界
	fmt.Println(strings.HasPrefix("hello,world", "hello"))

	//func HasSuffix(s, suffix string) bool
	//HasSuffix,判断s中是否包含后缀suffix,实现参考HasPrefix
	fmt.Println(strings.HasSuffix("hello,world,", "world"))

	//func ContainsRune(s string, r rune) bool
	//s中是否包含utf-8值r
	//rune底层结构是uint32,相当于char,go中没有char只有byte
	//rune是用来存储utf-8编码的单个字符
	fmt.Println(strings.ContainsRune("你好,世界", '界'))

	//func ContainsAny(s, chars string) bool
	//ContainsAny,s中是否包含chars中的任一字符
	fmt.Println(strings.ContainsAny("hello world,你好世界", "test your code"))

	//func Count(s, substr string) int
	//s中含有多少个不重复的substr字串
	fmt.Println(strings.Count("hello,hello,hello,world", "hello"))

	//func Index(s, substr string) int
	//字串substr在s中第一次出现的位置,不存在则返回-1
	fmt.Println(strings.Index("hello,hello,hello,world", ",hello"))

	//func IndexByte(s string, c byte) int
	//返回s中第一次出现c字符的位置,不存在则返回-1
	//c只能是byte类型,不能为rune类型字符,可以使用整型参数,但数值大小不能超过byte
	//注意,非ASCII码占位不止一位,len("我")=3

	//调用了internal/bytealg包的IndexByteString函数
	//func IndexByteString(s string, c byte) int
	//常规的逐一比对,时间复杂度O(n)
	fmt.Println(strings.IndexByte("hello world 你好 世界", 'h'))

	//func IndexRune(s string, r rune) int
	//返回s中第一次出现utf-8编码r的位置,不存在则返回-1
	//r可以是ASCII范围内的字符,注意'你'在s中的位置是13,但'好'不是14,而是16,这个编码方式有关
	//先判断r是否在byte范围,若是,可调用IndexByte,再判断r是否有效
	//如以上都不符合,把r转换为string,然后调用Index
	fmt.Println(strings.IndexRune("hello world, 你好 世界", '好'))

	//func IndexAny(s, chars string) int
	//返回s中包含chars任一字符的第一个位置,都不存在返回-1
	//从s中逐一提取,若S[i]存在与chars,返回i
	fmt.Println(strings.IndexAny("hello world ", "d"))

	/*
		Index系列对应还有LastIndex,顾名思义,从后遍历,返回第一个遇见的那个它!!!
	*/

	//按照一定的格式把s格式化,默认分隔符为' '
	fmt.Println(strings.Title("hello world I love you"))

	//stings.ToLower  ToUpper 转换大小写

	//返回count个s串联的字符串
	//如果count小于0或count个s溢出,引发panic异常
	fmt.Println(strings.Repeat("hello ", 3))

	//func Map(mapping func(rune) rune, s string) string
	//对s中的每个字符都调用mapping(r)处理,如果mapping返回一个负值,则该r会被丢弃
	fmt.Println(strings.Map(
		func(r rune) rune {
			if r >= 65 && r <= 90 {
				return r + 32
			}
			return -1 //丢弃所有非大写字母的字符
		},
		"HELLO WORLD,heekkisf"))

	//func Trim(s, cutset string) string
	//在s的前后,如果字符在cutset里就去除,遇到第一个不在cutset里的字符就不再继续遍历
	fmt.Println(strings.Trim("!!!a!b!!b!!!", "a!"))
	/*
		Trim系列:
			func TrimSpace(s string) string   去除s前后端的空白字符(unicode.isSpace指定)
			func TrimFunc(s string, f func(rune) bool) string 去除前后端满足f(r)的字符
			func TrimLeft(s string, cutset string) string
			func TrimRight(s string, cutset string) string  //left和right组成Trim
			func TrimLeftFunc(s string, f func(rune) bool) string
			func TrimRightFunc(s string, f func(rune) bool) string  去除左或或满足f(r)的字符
			func TrimPrefix(s, prefix string) string 去除s的前缀prefix
			func TrimSuffix(s, suffix string) string 去除s的后缀suffix
	*/
	//TrimPrefix只会去除包含的第一个前缀,如果想要去除所有前缀应配合HasPrefix使用(如下)
	//TrimSuffix同理
	fmt.Println(strings.TrimPrefix("hellohello world", "hello"))
	s := "hellohellohelloworld"
	fmt.Println(s)
	for strings.HasPrefix(s, "hello") {
		s = strings.TrimPrefix(s, "hello")
	}
	fmt.Println(s)

	//func Fields(s string) []string
	//将s按照空白自己分割,返回一个字符切片
	fmt.Println(strings.Fields(" hello world is slice"))
	//func FieldsFunc(s string, f func(rune) bool) []string
	//按照f(rune)的方式进行分割

	//func Split(s, sep string) []string
	//根据sep将s分割为字符串数组
	//如果sep为空,则将s切分为每一个unicode码值的一个字符串
	//如果sep是前缀或者后缀,则返回的的字符数组中,前或后有一个空白字符串""
	fmt.Println(strings.Split("ababa", "a"))

	//func SplitN(s, sep string, n int) []string
	//将s按照sep切成n份
	//n==0:返回一个nil切片,n<0:按照sep切完
	//n>0:且n份,若n>=nums(sep)+1则是切完,若n<nums(sep)+1,则最后一段会包含sep
	fmt.Println(strings.SplitN("ababababa", "b", 4))

	//func SplitAfter(s, sep string) []string
	//SplitAfter会将sep附在每一段的后面(除了最后一段)

	//func Join(a []string, sep string) string
	//将a用sep连接成一个string并返回
}
