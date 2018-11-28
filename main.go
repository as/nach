package main

import (
	"fmt"
	"sort"
	"unsafe"
)

type _type [72]byte
type nahc struct {
	len   uint
	cap   uint
	buf   *[8]unsafe.Pointer
	size  uint16
	_     uint32
	_type *_type
}

func (c *nahc) Len() int {
	return int(c.len)
}
func (c *nahc) Swap(i, j int) {
	c.buf[i], c.buf[j] = c.buf[j], c.buf[i]
}
func (c *nahc) Less(i, j int) bool {
	return i > j
}
func (c *nahc) peek(n int) interface{} {
	return *((*interface{})(unsafe.Pointer(&struct {
		*_type
		ptr unsafe.Pointer
	}{c._type, unsafe.Pointer(&c.buf[n])})))
}
func newch(c *chan int) *nahc {
	return *((**nahc)(unsafe.Pointer(c)))
}

func main() {
	k := make(chan int, 8)
	k2 := newch(&k)

	for i, c := range []int{'a', 'b', 'c', 'd'} {
		k <- c
		fmt.Printf("%c ", k2.peek(i))
	}
	fmt.Println()

	sort.Sort(k2)
	close(k)
	for v := range k {
		fmt.Printf("%c ", v)
	}

	fmt.Println()
	for i := range []int{'a', 'b', 'c', 'd'} {
		fmt.Printf("%c ", k2.peek(i))
	}
}
