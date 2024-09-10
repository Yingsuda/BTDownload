package main

import (
	"BTDownload/bencode"
	"fmt"
)

type Bint struct {
	val int
}

type Bfloat64 struct {
	val float64
}

type cacheEnable interface {
	Bint | Bfloat64
}

type cache[bn cacheEnable] struct {
	vals []bn
}

func (c *cache[bn]) Set(num int, val bn) {
	c.vals = append(c.vals, val)
}

func main() {
	obj := new(bencode.Bobject)
	bencode.SetObjValue(obj, "123")
	fmt.Println(bencode.GetObjValue[string](obj))
}
