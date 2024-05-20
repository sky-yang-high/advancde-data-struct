package main

import (
	"fmt"
	"math"
)

func main() {
	prices := []int{7, 2, 10, 1, 6, 8, 5, 7, 3, 8}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	buy, sell := math.MinInt32, 0
	for _, p := range prices {
		buy = max(buy, 0-p)
		sell = max(sell, buy+p)
	}
	return sell
}
