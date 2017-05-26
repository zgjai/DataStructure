package concurrentmap

/**
 * 对外暴露的接口
 */
type ConcurrentMap interface {
	// 用于返回并发量
	Concurrency() int
	
	// 向map中存储一个k-v对，若k已存在，则会覆盖
	Put(key string, element interface{}) (bool, error)
	
	// 从map中获取一个k对应的v
	Get(key string) (interface{}, error)
	
	// 删除一个k-v对
	Delete(key string) (bool, error)
	
	// 获取当前map中k-v对数量
	Len() uint64
}

/**
 * 内部实现的数据结构
 */
type myConcurrentMap struct {
	// 并发量，同时也是segments的长度
	concurrency int
	// 整个map由多个散列段构成，形成分段锁，降低锁粒度，利于并发读写
	segments []Segment
	// 当前map中k-v对的总数
	total uint64
}

func NewConcurrentMap()  {
	
}