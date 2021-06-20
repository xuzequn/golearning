package main

import (
	"errors"
	"fmt"
	"sync"
)

func Add(values []interface{}, val interface{}, index int) ([]interface{}, error) {

	if index < 0 || index > len(values) {
		return nil, errors.New("index not allow")
	}

	// res := []interface{}{}

	//先放 1234
	// 子切片
	res := values[:index]
	// for i := 0; i < index; i++ {
	// 	v := values[i]
	// 	res = append(res, v)
	// }
	// 再放新数
	res = append(res, val)

	// 再放后面的数
	// for i := index; i < len(values); i++ {
	// 	v := values[i]
	// 	res = append(res, v)
	// }
	// … 其实是go的一种语法糖
	// 子切片不会带来性能上的好处，无法扩展切片长度
	res = append(res, values[index:]...)

	return res, nil
}

func Delete(values []interface{}, index int) []interface{} {
	if index < 0 || index > len(values) {
		return values
	}
	res := []interface{}{}
	for i := 0; i < index; i++ {
		v := values[i]
		res = append(res, v)
	}
	for i := index + 1; i < len(values); i++ {
		v := values[i]
		res = append(res, v)
	}
	return res

}

type HashSet interface {
	Set(key string)
	Size() int
	Exist(key string) bool
}

type hashset struct {
	m map[string]interface{}
}

func (h *hashset) Set(key string) {
	// h.m[key] = ""

	h.m[key] = struct{}{} //初始化
}

func (h *hashset) Size() int {
	// 获得mao中的键值对的个数
	return len(h.m)
	// len 可用于数组， 切片 map
	// cap 可用于数组，切片
	// for range 可以用于数组， 切片 map

}

func (h *hashset) Exist(key string) bool {
	_, ok := h.m[key]

	return ok
}

// 组合方式定义安全set，装饰器模式的应用, 在原本的hashset 上增加线程安全的锁
type safeset struct {
	HashSet              // 被委托的 delegate
	mutex   sync.RWMutex // 读写锁
}

// type safesetV2 struct {
// 	m     map[string]interface{}
// 	mutex sync.RWMutex // 读写锁
// }

func (s *safeset) Size() int {
	s.mutex.RLock()
	defer s.mutex.RUnlock() // 类似于 finally，执行完return 前最后一个表达式之后 执行defer，之后再执行return
	return s.HashSet.Size()

	// size := s.HashSet.Size()
	// s.mutex.RUnlock()
	// return size
}

func (s *safeset) Set(key string) {
	// h.m[key] = ""
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.HashSet.Set(key) //初始化
}

func (s *safeset) Exist(key string) bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.HashSet.Exist(key)
}

func main() {
	values := []interface{}{1, 2, 3, 4, 5}
	newValues, err := Add(values, 4, 4)
	// newValues := Delete(values, 3)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	// 1,2,3,4,4,5
	for i, v := range newValues {
		fmt.Printf("下标：%d, 值: %d \n", i, v)

	}

}
