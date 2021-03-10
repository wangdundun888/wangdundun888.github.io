package LRU

import "container/list"

//lru可以像lfu一样设置相同的数据结构
//在数据结构中的计数改为一个时间戳值,
//淘汰时选择时间戳值最大的淘汰

//本次实现使用双链表+hashmap

const (
	CACHA_MAX_NUMBER = 100
)

type LRUCache interface {
	Set(k, v string)
	Get(k string) string
}

type myLruCache struct {
	List   *list.List
	Map    map[string]*list.Element
	MaxLen uint64
}

func NewMLC(size uint64) *myLruCache {
	if size == 0 {
		size = CACHA_MAX_NUMBER
	}
	mlc := &myLruCache{
		List:   list.New(),
		Map:    make(map[string]*list.Element),
		MaxLen: size,
	}
	return mlc
}

func (mlc *myLruCache) Set(k, v string) {
	e, exist := mlc.Map[k]
	if exist {
		mlc.List.Remove(e)
		mlc.List.PushFront(e.Value.(string))
		return
	}
	if uint64(mlc.List.Len()) >= mlc.MaxLen {
		e = mlc.List.Back()
		mlc.List.Remove(e)
		deleteKey := ""
		for k1, e1 := range mlc.Map {
			if e1 == e {
				deleteKey = k1
				break
			}
		}
		delete(mlc.Map, deleteKey)
		newElement := mlc.List.PushFront(v)
		mlc.Map[k] = newElement
		return
	}
	newElement := mlc.List.PushFront(v)
	mlc.Map[k] = newElement
}

func (mlc *myLruCache) Get(k string) string {
	e, exist := mlc.Map[k]
	if !exist {
		return ""
	}
	v := e.Value.(string)
	mlc.List.Remove(e)
	delete(mlc.Map, k)
	newE := mlc.List.PushFront(v)
	mlc.Map[k] = newE
	return v
}
