package main

import (
	"image"
	"sync"
)

/**
进阶知识点
	- 延迟一个开销很大的初始化操作到真正用到它的时候再执行是一个很好的实践
	- 因为预先初始化一个变量（比如在init函数中完成初始化）会增加程序的启动延时，
	  而且有可能实际执行过程中这个变量没有用上，那这个初始化操作就不是必须要做的。
*/

var icons map[string]image.Image

func loadIcon(name string) (res image.Image) { return res }

func loadIcons() {
	icons = map[string]image.Image{
		"left":  loadIcon("left.png"),
		"up":    loadIcon("up.png"),
		"right": loadIcon("right.png"),
		"down":  loadIcon("down.png"),
	}
}

/**
- Icon被多个goroutine调用时不是并发安全的
	- 现代的编译器和CPU可能会在保证每个goroutine都满足串行一致的基础上自由地重排访问内存的顺序
	- 这种情况下，即使判断额了icons不是nil也不意味着变量初始化完成了。
*/
func Icon(name string) image.Image {
	if icons == nil {
		loadIcons()
	}
	return icons[name]
}

/**
- 能想到的一个办法就是加互斥锁
	- 但这样又会引发性能问题，GO中sync包中提供了一个针对一次性初始化问题的解决方案-sync.Once
*/

var icons1 map[string]image.Image

var loadIconsOnce sync.Once

func loadIcons1() {
	icons = map[string]image.Image{
		"left":  loadIcon("left.png"),
		"up":    loadIcon("up.png"),
		"right": loadIcon("right.png"),
		"down":  loadIcon("down.png"),
	}
}

/**
- Icons1是并发安全的
- sync.Once其实内部包含一个互斥锁和一个布尔值
	- 互斥锁保证布尔值和数据的安全
	- 而布尔值用来记录初始化是否完成。
	- 这样就能保证初始化操作时是并发安全的并且初始化操作也不会被执行多次。
*/
func Icon1(name string) image.Image {
	loadIconsOnce.Do(loadIcons)
	return icons[name]
}
