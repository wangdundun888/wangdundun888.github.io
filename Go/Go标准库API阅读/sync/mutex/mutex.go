package main

/*
	参考:https://www.cnblogs.com/luozhiyun/p/14157542.html

	以及源码: https://github.com/wangdundun888/go/blob/master/src/sync/mutex.go#L126
*/

type Mutex struct {
	//state有32为，低三位分别表示Locked、Woken、Starving三种状态
	//剩下29位可以表示waiter的数量，即有1<<(32-3)个等待者
	state int32
	sema  uint32
}

func (m *Mutex) Lock() {
	/*
		加锁的时候尝试能不能通过CAS直接获得锁，如果可以,直接获得锁,并置state加锁位为1
		如果不行则调用 Mutex.lockSlow()方法
	*/
}

func (m *Mutex) lockSlow() {

}

func main() {

}
