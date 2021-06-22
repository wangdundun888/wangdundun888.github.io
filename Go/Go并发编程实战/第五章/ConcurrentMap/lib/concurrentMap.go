package lib

import (
	"math"
	"sync/atomic"
)

type ConcurrentMap interface {
	//用于返回并发量
	Concurrency() int
	//推送一个键值对
	//element不能为nil
	//重复推同个key会发生覆盖
	//bool表示是否加入成功
	Put(key string, element interface{}) (bool, error)
	//根据key获取value
	Get(key string) interface{}
	//根据key删除键值对
	//bool表示删除成功与否
	Delete(key string) bool
	//返回当前字典中键值对的数量
	Len() uint64
}

type myConcurrentMap struct {
	concurrency int
	segments    []Segment
	total       uint64
}

func NewConcurrentMap(
	concurrency int,
	pairRedistributor PairRedistributor) (ConcurrentMap, error) {
	if concurrency <= 0 {
		return nil, newIllegalParameterError("concurrency is too small")
	}
	if concurrency > MAX_CONCURRENCY {
		return nil, newIllegalParameterError("concurrency is too large")
	}
	cmap := &myConcurrentMap{}
	cmap.concurrency = concurrency
	cmap.segments = make([]Segment, concurrency)
	for i := 0; i < concurrency; i++ {
		cmap.segments[i] =
			newSegment(DEFAULT_BUCKET_NUMBER, pairRedistributor)
	}
	return cmap, nil
}

func (cmap *myConcurrentMap) Concurrency() int {
	return cmap.concurrency
}

func (cmap *myConcurrentMap) Put(key string, element interface{}) (bool, error) {
	p, err := newPair(key, element)
	if err != nil {
		return false, err
	}
	s := cmap.findSegment(p.Hash())
	ok, err := s.Put(p)
	if ok {
		atomic.AddUint64(&cmap.total, 1)
	}
	return ok, err
}

//段定位算法,主要是让键值对在segment中分布得更广更均匀
func (cmap *myConcurrentMap) findSegment(keyHash uint64) Segment {
	if cmap.concurrency == 1 {
		return cmap.segments[0]
	}
	var keyHash32 uint32
	if keyHash > math.MaxUint32 {
		keyHash32 = uint32(keyHash >> 32)
	} else {
		keyHash32 = uint32(keyHash)
	}
	return cmap.segments[int(keyHash32>>16)%(cmap.concurrency-1)]
}

func (cmap *myConcurrentMap) Get(key string) interface{} {
	keyHash := hash(key)
	s := cmap.findSegment(keyHash)
	pair := s.GetWithHash(key, keyHash)
	if pair == nil {
		return nil
	}
	return pair.Element()
}

func (cmap *myConcurrentMap) Len() uint64 {
	return atomic.LoadUint64(&cmap.total)
}

func (cmap *myConcurrentMap) Delete(key string) bool {
	s := cmap.findSegment(hash(key))
	if s.Delete(key) {
		atomic.AddUint64(&cmap.total, ^uint64((0)))
	}
	return false
}
