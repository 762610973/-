package main

import "fmt"

//升序排列
func bubbleSortUp(nums []int) {
	length := len(nums)
	didSwap := false
	//这里做一下优化，如果这一轮中没有进行交换，那么已经排好序了
	for i := 0; i < length-1; i++ {
		//n个数字，只需进行n-1轮排序
		for j := 0; j < length-1-i; j++ {
			//每一轮都会得到最大值再减去i即可
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				didSwap = true
			}
		}
		if !didSwap {
			return
		}
	}
}
func bubbleSortDown(nums []int) {
	length := len(nums)
	didSwap := false
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-1-i; j++ {
			if nums[j] < nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				didSwap = true
			}
		}
		if !didSwap {
			return
		}
	}
}
func main() {
	nums := []int{7, 6, 10, 4, 3, 2}
	bubbleSortUp(nums)
	fmt.Println(nums)
	bubbleSortDown(nums)
	fmt.Println(nums)
}
