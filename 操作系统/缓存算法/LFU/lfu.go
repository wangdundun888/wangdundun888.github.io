package LFU

import "math"

const (
	MAX_CACHE_NUMBER = 100
)

//方式1:数组+hashmap
//插入和访问数据都是O(1),淘汰数据则是O(n)
type LFUCache interface {
	Set(k, v string)
	Get(k string) string
}

type myLFUCache struct {
	Data   ValueData
	KeyMap map[string]uint64
	Total  uint64
	MaxLen uint64
}

type ValueData struct {
	value []string
	cnt   []int
}

func NewMLC(size uint64) *myLFUCache {
	if size == 0 {
		size = MAX_CACHE_NUMBER
	}
	mlc := &myLFUCache{
		Data: ValueData{
			value: make([]string, size),
			cnt:   make([]int, size),
		},
		KeyMap: make(map[string]uint64),
		Total:  0,
		MaxLen: size,
	}
	return mlc
}

func (mlc *myLFUCache) Set(k, v string) {
	index, exist := mlc.KeyMap[k]
	if exist {
		mlc.Data.value[index] = v
		return
	}
	if mlc.Total >= mlc.MaxLen {
		//淘汰并放入
		cnt := math.MaxInt32
		index = 0
		for i, num := range mlc.Data.cnt {
			if num < cnt {
				cnt = num
				index = uint64(i)
			}
		}
		var deleteKey string
		for key, value := range mlc.KeyMap {
			if value == index {
				deleteKey = key
				break
			}
		}
		delete(mlc.KeyMap, deleteKey)
		mlc.KeyMap[k] = index
		mlc.Data.value[index] = v
		mlc.Data.cnt[index] = 0
		return
	}
	mlc.KeyMap[k] = mlc.Total
	index = mlc.Total
	mlc.Data.value[index] = v
	mlc.Total++
}

func (mlc *myLFUCache) Get(k string) string {
	index, exist := mlc.KeyMap[k]
	if !exist {
		return ""
	}
	value := mlc.Data.value[index]
	mlc.Data.cnt[index]++
	return value
}

//方式二:最小堆+hashmap
//
