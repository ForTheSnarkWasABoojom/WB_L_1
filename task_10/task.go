package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	arr := [8]float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	groups := make(map[float64][]float64)
	step := 10.0
	for _, temp := range arr {
		groupKey := math.Floor(temp/step) * step
		groups[groupKey] = append(groups[groupKey], temp)
	}

	sortKeys := make([]float64, 0, len(groups))

	for k := range groups {
		sortKeys = append(sortKeys, k)
	}
	sort.Float64s(sortKeys)

	for _, k := range sortKeys {
		fmt.Println(k, groups[k])
	}
}
