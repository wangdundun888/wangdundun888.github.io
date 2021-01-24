### go自定义排序

    package main

    import (
    "fmt"
    "sort"
    )
    //学习自定义排序
    type product struct {
        id  	int
        name    string
        price   float64
    }
    type products []product
    
    //该方法是返回一个数组切片的长度,所以签名应该是一个数组
    func(p products)Len()int{
        return len(p)
    }
    func(p products)Swap(i,j int){
        p[i],p[j] = p[j],p[i]
    }
    func(p products)Less(i,j int)bool{
        //按照id排序
        //>:降序  < :升序
        //return  p[i].id > p[j].id
        //按照price排序
        //>:降序  < :升序
        return  p[i].price > p[j].price
    }
    
    
    func main() {
        p := &products{
        {1,"亲嘴烧",0.5},
        {3,"热狗",1},
        {2,"鸡蛋",6.5},
        }
        fmt.Println(p)
        //调用sort.Sort()需要实现data接口,一共Len(),Swap(),Less()三个方法
        sort.Sort(p)
        fmt.Println(p)
    }

### 把整数数组排成最小的数

    输入一个正整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。例如输入数组{3，32，321}，则打印出这三个数字能排成的最小数字为321323。

    思路:
        对于两个数字a和b,重新定义加法,a+b=ab,b+a=ba,例如32+3=323.

        所以,在数组中a1,a2,a3,...,ai,a1+aj(j=[2,i])必定小于aj+a1,所以a1一定是在第一个排列,a2...往下亦是如此.
        最后把数组按照上面的加法排序,再把这个数组加起来转为一个字符串返回即可.

    package main
    import (
        "sort"
        "strconv"
    )
    
    type Strings  []string
    
    
    func(s Strings)Len()int{
     return len(s)
    }
    func(s Strings)Swap(i,j int){
      s[i],s[j] = s[j],s[i]
    }
    func(s Strings)Less(i,j int)bool{
        a := s[i] + s[j]
        b := s[j] + s[i]
        a1,_ := strconv.Atoi(a)
        b1,_ := strconv.Atoi(b)
        if a1 < b1 {
        return true
    }
    return  false
    }
    func PrintMinNumber( numbers []int ) string {
        s := make([]string,0)
        for _,v := range numbers{
        s = append(s,strconv.Itoa(v))
     }
        sort.Sort(Strings(s))
        ans := ""
        for _,v := range s{
            ans = ans + v
        }
        return ans
    }