package main

import "fmt"

//选择排序，每次从待排序的数据元素中选出最小的一个元素，存放在序列的起始
func selectSortUp(nums []int) {
	length := len(nums)
	//进行n-1轮迭代
	for i := 0; i < length-1; i++ {
		min := nums[i] //每一轮都将当前值设为最小值
		minIndex := i
		for j := i + 1; j < length; j++ {
			if nums[j] < min {
				min = nums[j]
				minIndex = j
			}
		}
		//这一轮找到的最小元素的下标不等于最开始的下标，就交换元素
		if i != minIndex {
			nums[i], nums[minIndex] = nums[minIndex], nums[i]
		}
	}
}
func selectSortDown(nums []int) {
	length := len(nums)
	for i := 0; i < length-1; i++ {
		max := nums[i]
		maxIndex := i
		//j从i+1到最后一个数
		for j := i + 1; j < length; j++ {
			if nums[j] > max {
				max = nums[j]
				maxIndex = j
			}
		}
		if i != maxIndex {
			nums[i], nums[maxIndex] = nums[maxIndex], nums[i]
		}
	}
}

// 对算法进行改进
// 每一轮除了寻找到最小值，还要寻找到最大值，然后分别和前面和后面的元素交换
// 优化后的选择排序还是很慢，不建议使用
func selectGoodSortUp(nums []int) {
	n := len(nums)
	//每次可以找到最大值和最小值，因此只循环一半就行了
	for i := 0; i < n/2; i++ {
		minIndex := i
		maxIndex := i
		//在这一轮迭代中找到最大值和最小值的下标
		for j := i + 1; j < n-i; j++ {
			//知道哦啊哦最大值的下标
			if nums[j] > nums[maxIndex] {
				maxIndex = j
			}
			//找到最小值下标
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		//最大值是当前元素，并且最小值不是最尾的元素
		if maxIndex == i && minIndex != n-1 {
			//将最大值和最尾的元素交换
			nums[n-i-1], nums[maxIndex] = nums[maxIndex], nums[n-i-1]
			//将最小的元素放到最开头
			nums[i], nums[minIndex] = nums[minIndex], nums[i]
			//如果最大值在开头，最小值在结尾，直接交换
		} else if maxIndex == i && minIndex == n-i-1 {
			nums[minIndex], nums[maxIndex] = nums[maxIndex], nums[minIndex]
		} else {
			//否则将最小值放在开头，再将最大值放在结尾
			nums[i], nums[minIndex] = nums[minIndex], nums[i]
			nums[n-i-1], nums[maxIndex] = nums[maxIndex], nums[n-i-1]
		}
	}
}
func main() {
	nums := []int{1, 2, 4, 5, 7, 4, 6, 2, 6, 999, 444}
	selectGoodSortUp(nums)
	fmt.Println(nums)
	selectSortUp(nums)
	fmt.Println(nums)
	selectSortDown(nums)
	fmt.Println(nums)
}
