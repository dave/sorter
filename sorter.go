package sorter

import "sort"

func New(len int, swap func(i, j int), less func(i, j int) bool) sort.Interface {
	return st{len: len, swap: swap, less: less}
}

type st struct {
	len  int
	swap func(i, j int)
	less func(i, j int) bool
}

func (s st) Len() int { return s.len }

func (s st) Swap(i, j int) { s.swap(i, j) }

func (s st) Less(i, j int) bool { return s.less(i, j) }
