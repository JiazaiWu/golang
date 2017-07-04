package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	mMap := make(map[int]int);
	//golang range will get index and value
	for index, val := range nums {
		toFind:= target - val;
		//notice how to use map in the go
		value, ok := mMap[toFind];
		if (ok) {
			return []int{value, index};
		}
		mMap[val] = index;
	}
	//should not be here
	return nil;
}

func main() {
	arg := []int{5, 3, 1, 2, 4};
	ret := twoSum(arg, 6);
	fmt.Println(ret);
}
