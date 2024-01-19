package main

import (
	"fmt"
	"sort"
)

// 定义一个pair
type Pair struct {
	From   string
	To     string
	IsUsed bool
}

// 定义一个pair的slice
type Pairs []Pair

func (p Pairs) Len() int {
	return len(p)
}

// 按照字典序排序
// 1. 先按照from排序
// 2. 如果from相同，按照to排序
func (p Pairs) Less(i, j int) bool {
	if p[i].From < p[j].From {
		return true
	} else if p[i].From == p[j].From {
		return p[i].To < p[j].To
	} else {
		return false
	}
}

func (p Pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// 332. 重新安排行程
func findItinerary(tickets [][]string) []string {
	if len(tickets) == 0 {
		return nil
	}

	// 1. 构建slice
	pairs := make([]Pair, len(tickets))
	for i, ticket := range tickets {
		pairs[i] = Pair{
			From:   ticket[0],
			To:     ticket[1],
			IsUsed: false,
		}
	}

	// 2. 排序
	sort.Sort(Pairs(pairs))

	res := make([]string, 0)
	var dfs func(string) bool
	dfs = func(from string) bool {
		if len(res) == len(tickets) {
			res = append(res, from)
			return true
		}

		for i, pair := range pairs {
			if pair.From == from && !pair.IsUsed {
				pairs[i].IsUsed = true
				res = append(res, pair.From)
				if dfs(pair.To) {
					return true
				}
				pairs[i].IsUsed = false
				res = res[:len(res)-1]
			}
		}
		return false
	}
	dfs("JFK")

	return res
}

func main() {
	// [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
	fmt.Println(findItinerary([][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}))

}
