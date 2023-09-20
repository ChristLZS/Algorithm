package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// 方法一：使用map统计每个元素出现的次数，然后对map的key进行排序，取前k个元素
func topKFrequent(nums []int, k int) []int {
	// 统计每个元素出现的次数
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	// 将map的key放入数组中
	res := make([]int, 0)
	for k := range m {
		res = append(res, k)
	}

	// 对数组进行排序
	sort.Slice(res, func(i, j int) bool {
		return m[res[i]] > m[res[j]]
	})

	return res[:k]
}

///////////////////////////////////////////////////////////////

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	fmt.Println(topKFrequent2([]int{1, 1, 1, 2, 2, 3}, 2))
}

///////////////////////////////////////////////////////////////

// 方法二：使用小顶堆
// 1. 统计每个元素出现的次数
// 2. 将统计结果放入小顶堆中，如果堆的大小大于k，则将堆顶元素弹出
// 3. 将堆中的元素放入数组中，然后对数组进行排序
func topKFrequent2(nums []int, k int) []int {
	// 统计每个元素出现的次数
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	// 将统计结果放入小顶堆中
	h := &IntHeap{}
	for k, v := range m {
		heap.Push(h, &Item{k, v})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	// 将堆中的元素放入数组中
	res := make([]int, 0)
	for h.Len() > 0 {
		res = append(res, heap.Pop(h).(*Item).value)
	}

	// 对数组进行排序
	sort.Ints(res)

	return res[:k]
}

type Item struct {
	value int
	count int
}

type IntHeap []*Item

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i].count < h[j].count
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(*Item))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]

	return x
}
