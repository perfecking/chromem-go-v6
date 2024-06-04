package chromem

import (
	"container/list"
)

// LRUCache 结构体定义
type LRUCache struct {
	capacity int
	cache    map[string]*list.Element
	lruList  *list.List
}

// CacheItem 用于保存缓存中的键值对
type CacheItem struct {
	key   string
	value interface{}
}

// NewLRUCache 初始化LRU缓存
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[string]*list.Element),
		lruList:  list.New(),
	}
}

// Get 从缓存中获取元素，如果存在则将其移到最近使用的队列尾部
func (c *LRUCache) Get(key string) (interface{}, bool) {
	if elem, ok := c.cache[key]; ok {
		c.lruList.MoveToBack(elem)
		return elem.Value.(*CacheItem).value, true
	}
	return nil, false
}

// Put 将元素放入缓存，如果已存在则更新，同时确保不超过最大容量
func (c *LRUCache) Put(key string, value interface{}) {
	if elem, ok := c.cache[key]; ok {
		// 更新已存在的元素
		elem.Value.(*CacheItem).value = value
		c.lruList.MoveToBack(elem)
	} else {
		// 新增元素
		newItem := &CacheItem{key, value}
		elem := c.lruList.PushBack(newItem)
		c.cache[key] = elem
	}
}

// Delete 从缓存中删除指定键的元素
func (c *LRUCache) Delete(key string) {
	if elem, ok := c.cache[key]; ok {
		delete(c.cache, key)
		c.lruList.Remove(elem)
	}
}

func (c *LRUCache) Evict() []string {
    var keys []string
    // 检查是否超过容量，如果是，则删除最老的元素
    for len(c.cache) > c.capacity {
        oldest := c.lruList.Front()
        if oldest != nil {
            delete(c.cache, oldest.Value.(*CacheItem).key)
            c.lruList.Remove(oldest)

            keys = append(keys, oldest.Value.(*CacheItem).key)
        }
    }
    return keys
}
