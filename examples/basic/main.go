package main

import (
	"fmt"

	"github.com/zekrotja/sop"
)

func main() {
	r := sop.Wrap([][]int{{1, 2, 3}, {4, 5}, {6, 7, 8}, {9}, {10}})

	rf := sop.Flat(r).
		Shuffle().
		Sort(func(p, q, i int) bool {
			return p > q
		}).
		Filter(func(v, i int) bool {
			return v%2 == 0
		})

	rs := sop.Map(rf,
		func(v int, i int) string {
			return fmt.Sprintf("%d", v)
		}).
		Aggregate(func(a, b string) string {
			return a + ", " + b
		})

	fmt.Println(rs)
}
