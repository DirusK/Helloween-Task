package main

import (
	"fmt"
	"sort"
)

func maxCandies(sweetnessTypes [][]int, sweetnessMax int) int {
	output := 0
	sorting(sweetnessTypes)
	for i := 0; sweetnessMax > 0; i++ {
		if sweetnessMax > sweetnessTypes[i][0] {
			sweetnessMax -= sweetnessTypes[i][0]
			output += sweetnessTypes[i][0] * sweetnessTypes[i][1]
		} else {
			output += sweetnessMax * sweetnessTypes[i][1]
			break
		}
	}
	return output
}

func sorting(sweetnessTypes [][]int) {
	sort.Slice(sweetnessTypes, func(i, j int) bool {
		return sweetnessTypes[i][1] > sweetnessTypes[j][1]
	})
}

func main() {
	fmt.Println("TEST-1")
	sweetnessTypes := [][]int{{1, 3}, {2, 2}, {3, 1}}
	fmt.Println("Output: ", maxCandies(sweetnessTypes, 4))

	fmt.Println("TEST-2")
	sweetnessTypes = [][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}}
	fmt.Println("Output: ", maxCandies(sweetnessTypes, 10))

	fmt.Println("TEST-3")
	sweetnessTypes = [][]int{{35, 14}, {57, 99}, {70, 48}, {50, 70}, {59, 24}, {48, 72},
		{27, 48}, {50, 89}, {91, 9}, {87, 66}, {74, 58}, {52, 29},
		{10, 19}, {11, 87}, {56, 71}, {83, 67}, {73, 31}, {41, 58},
		{26, 39}, {100, 99}, {96, 51}, {33, 34}, {43, 23}, {22, 41},
		{89, 28}, {43, 19}, {87, 56}, {30, 95}, {54, 93}, {81, 98},
		{84, 26}, {51, 52}, {21, 16}}
	fmt.Println("Output: ", maxCandies(sweetnessTypes, 270))
}
