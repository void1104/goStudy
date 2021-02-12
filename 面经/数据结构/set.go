package main

import "fmt"

/**
以空结构体作为map的value来实现，空结构体不占内存
*/

type Empty struct{}

var empty Empty

// set类型
type Set struct {
	m map[int]Empty
}

// 添加元素
func (s *Set) Add(val int) {
	s.m[val] = empty // 使用一个empty单例作为所有键的值
}

// 删除元素
func (s *Set) Remove(val int) {
	delete(s.m, val)
}

// 获取长度
func (s *Set) Size() int {
	return len(s.m)
}

// 清空set
func (s *Set) Clear() {
	s.m = make(map[int]Empty)
}

// 查看某个元素是否存在
func (s *Set) Exist(val int) (ok bool) {
	_, ok = s.m[val]
	return
}

// 遍历set
func (s *Set) Traverse() {
	for v := range s.m {
		fmt.Println(v)
	}
}
