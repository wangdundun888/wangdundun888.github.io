package FIFO

import "container/list"

// 基于双向链表

const (
	CACHE_NUMBER = 100
)

type FIFOCache interface {
	Set(k, v string)
	Get(k string) string
}

type myFifoCache struct {
	Data   *list.List
	MaxLen uint64
}

type Pair struct {
	key   string
	value string
}

func NewMFC(size uint64) *myFifoCache {
	mfc := &myFifoCache{
		Data: list.New(),
	}
	if size <= 0 {
		size = CACHE_NUMBER
	}
	mfc.MaxLen = size
	return mfc
}

//如果想提高访问效率,可以使用hashmap来保存key在链表中的位置
//get的时候,直接在hashmap中查找返回
//set的时候,
func (fc *myFifoCache) Set(k, v string) {
	l := fc.Data
	for e := l.Front(); e != nil; e = e.Next() {
		pair := e.Value.(Pair)
		if pair.key == k {
			pair.value = v
			e.Value = pair
			return
		}
	}
	if uint64(l.Len()) >= fc.MaxLen {
		e := l.Front()
		l.Remove(e)
		newE := Pair{k, v}
		l.PushBack(newE)
		return
	}
	newE := Pair{k, v}
	l.PushBack(newE)
}

func (fc *myFifoCache) Get(k string) string {
	l := fc.Data
	for e := l.Front(); e != nil; e = e.Next() {
		pair := e.Value.(Pair)
		if pair.key == k {
			v := pair.value
			return v
		}
	}
	return ""
}
