# 关于channel的一些思考

    昨晚睡前想到了使用goroutdine(以下简写为 gr)做一个小实验,即多个gr通过channel发送数据,一个gr接收数据:
```
    var wg sync.WaitGroup //配合sync.WaitGroup让主线程等待其他gr结束后退出
    wg.Add(10) //wg.Add函数接收一个整型参数,参数代表额外的gr个数
    wg.Done() //配合wg.Add使用,每个gr结束后都要使用wg.Done函数
    wg.Wait()  //阻塞主线程,等待其他gr退出后执行后面语句
```
### 实验一 多个发送gr VS 一个gr
```
    根据初步实验,代码如下:
    package main
    import (
        "fmt"
        "sync"
    )
    const MAXGR = 20
    func main() {
        var wg sync.WaitGroup
        wg.Add(MAXGR+1) //+1为接收gr个数
        test := make(chan int)
        //发送MAXGR个数,创建MAXGR个发送gr
        for i:=0;i<MAXGR;i++{
            go func(num int){
                defer wg.Done();
                defer fmt.Printf("发送线程%d,已经退出\n",num)
                fmt.Printf("%d 已发送\n",num)
                test <- num
            }(i)
        }
        //创建一个接收
        go func(){
            defer wg.Done()
            defer fmt.Printf("B:接收线程已经退出\n")
            for i:= range test{
                fmt.Printf("%d 已经收到\n",i)
            }
        }()
        wg.Wait()
    }
    经调试后发现,出现bug,提示:deadlock,经过分析与查询资料,发现是channel使用方式错误,具体原因如下:
        wg.Add(MAXGR+1)和wg.Wait()配合使主线程等待MAXGR+1个gr结束后退出,发送线程发送后Done了MAXGR个gr,
        而接收gr使用了 i:=range channel 的方式接收数据,需要调用channel关闭(即调用close函数)后才会结束自身,
        所以造成了死锁。解决的方法就是发送数据完毕后关闭通道,但由于是多个gr发送数据,而这些gr之间彼此又没有通信,
        使用的又是无缓冲的通道,所以关闭通道的时机不好把握,如果随意关闭,那么向一个已关闭的通道发送数据会引发Panic
        异常。
    解决方法,暂时还没考虑,我先进行了另外一个实验。
```
### 实验二 一个发送gr VS 多个接收gr
```
    类似于第七章里的work包,创建gr池:
    package main

    import (
        "fmt"
        "sync"
    )

    const MAXGR = 11
    func main() {

        var wg sync.WaitGroup
        wg.Add(MAXGR+1)
        test := make(chan int)

        for i:=0;i<MAXGR;i++{
            go func(num int){
                defer wg.Done()
                defer fmt.Printf("接收线程已经退出 BBBBB %d \n",num)
                for j:= range test {
                    fmt.Printf("%d 已经收到\n", j)
                }
            }(i)
        }


        go func(){
            defer wg.Done();
            for i:=0;i<MAXGR;i++{
                test <- i
                fmt.Printf("数组 %d 已经发送AAAAAAAAAAAAA\n",i)
            }
            close(test)
        }()

        wg.Wait()
    }
    经过测试,是没有问题的,一个gr发送数据,可以很好把控channel的关闭时机,channel关闭后,
    所有的接收gr会自动关闭,所以关键是channel的关闭时机!!!
```
### 总结
```
    剩下的"一对一"实验也是和"一对多"一样比较好把控,"多对多"和"多对一"实验一样,要把握好channel的关闭时机,
    为此,我想到了一个办法:
        创建一个监听gr,传入发送gr的个数作为参数,每个发送gr结束时,和监听gr通过另一个通道通信,
        监听gr通过统计发送gr结束的个数,如与总个数相等,则判定发送数据结束,关闭发送数据通道,同时关闭自身,
        代码如下:
            package main

            import (
                "fmt"
                "sync"
            )
            const MAXGR = 20
            func main() {
                var wg sync.WaitGroup
                wg.Add(MAXGR+2)
                test := make(chan int)
                info := make(chan int)
                //发送MAXGR个数,创建MAXGR个发送gr
                for i:=0;i<MAXGR;i++{
                    go func(num int){
                        defer wg.Done();
                        defer fmt.Printf("发送线程%d,已经退出\n",num)
                        fmt.Printf("%d 已发送\n",num)
                        test <- num
                        info <- 1 //通知监控线程
                    }(i)
                }

                go func(maxGr int){
                    i := 0
                    defer wg.Done()
                    for {
                        i += <- info
                        if i == maxGr {
                            close(test)
                            close(info)
                            return
                        }
                    }
                }(MAXGR)

                //创建一个接收
                go func(){
                    defer wg.Done()
                    defer fmt.Printf("B:接收线程已经退出\n")
                    for i:= range test{
                        fmt.Printf("%d 已经收到\n",i)
                    }
                }()
                wg.Wait()
            }
    总的来说,在gr使用并发,要把握好channel的应用.
```