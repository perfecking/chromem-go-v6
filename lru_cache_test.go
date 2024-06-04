package chromem

import (
    "fmt"
    "testing"
)


func TestNewLRUCache(t *testing.T) {
    cache := NewLRUCache(3)

    cache.Put("1", "one")
    cache.Put("2", "two")
    cache.Put("3", "three")

    cache.Put("4", "four")      // 这将使得键1对应的项成为最旧的，因为下次访问时它会被移至末尾

    fmt.Println(cache.Get("5")) // 返回: <nil>, false，因为5不在缓存中
    fmt.Println(cache.Get("6")) // 返回: <nil>, false，同上

    cache.Put("5", "five")      // 这将使得键1对应的项成为最旧的，因为下次访问时它会被移至末尾

    // 再次获取1，检查是否已被删除
    fmt.Println(cache.Get("1")) // 返回: <nil>, false，因为添加4后，1已经被淘汰了
    fmt.Println(cache.Get("2")) // 返回: <nil>, false，因为添加4后，1已经被淘汰了

}