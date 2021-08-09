package main

import "testing"

type testpair struct{
	value [][]int
	sweetnessMax int
	output int
}

var testCases = []testpair{
	{[][]int{{1, 3}, {2, 2}, {3, 1}},4,8},

	{[][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}},10,91},

	{[][]int{{35, 14}, {57, 99}, {70, 48}, {50, 70}, {59, 24}, {48, 72},
	{27, 48}, {50, 89}, {91, 9}, {87, 66}, {74, 58}, {52, 29},
	{10, 19}, {11, 87}, {56, 71}, {83, 67}, {73, 31}, {41, 58},
	{26, 39}, {100, 99}, {96, 51}, {33, 34}, {43, 23}, {22, 41},
	{89, 28}, {43, 19}, {87, 56}, {30, 95}, {54, 93}, {81, 98},
	{84, 26}, {51, 52}, {21, 16}},270,26517},
}

func TestMaxCandies(t *testing.T){
	for i, pair := range testCases{
		output:= maxCandies(pair.value, pair.sweetnessMax)
		if output != pair.output{
			t.Errorf("For test#%d expected %d, got %d",i+1,pair.output,output)
		}
	}
}

