package main

import "fmt"

//优化方案：对已经排好序的序列，采用二分法查找方法，携带多个元素，数据链表化
//直接插入排序，每一次都把一个数插入到已经排好序的数列里面形成新地排好序的数列，以此反复
func insertSortUp(nums []int) {
	n := len(nums)
	//进行n-1轮迭代
	for i := 1; i <= n-1; i++ {
		deal := nums[i]
		j := i - 1 //带排序的数左边的第一个数的位置
		//如果是第一次比较，比左边的已排好序的第一个数小，那么直接插入chuli
		if deal < nums[j] {
			//一直往左边找，比待排序大的数都往后移动，腾空位给排序插入
			for ; j >= 0 && deal < nums[j]; j-- {
				nums[j+1] = nums[j] //数往后移，给待排序的数流出空位
			}
			nums[j+1] = deal //结束了，待排序的数插入空位
		}
	}
}

func insertSortDown(nums []int) {
	n := len(nums)
	for i := 1; i <= n-1; i++ {
		now := nums[i]
		j := i - 1 //	nums := []int{1, 2, 34, 5, 35}
		//当前的这个数比前一个大，就需要往前面插入，以下是处理过程
		if now > nums[j] {
			//j从前一个数往第一个数遍历
			for ; j >= 0 && now > nums[j]; j-- {
				nums[j+1] = nums[j]
			}
			nums[j+1] = now //前面的已经排好序了，当前可以直接插入
		}
	}
}

func insertSortOtherUp(nums []int) {
	n := len(nums)
	//进行n-1轮迭代
	for i := 1; i <= n-1; i++ {
		deal := nums[i]
		j := i - 1 //带排序的数左边的第一个数的位置
		//如果是第一次比较，比左边的已排好序的第一个数小，那么直接插入chuli
		if deal < nums[j] {
			//一直往左边找，比待排序大的数都往后移动，腾空位给排序插入
			for ; j >= 0 && deal < nums[j]; j-- {
				nums[j+1], nums[j] = nums[j], nums[j+1] //数往后移，给待排序的数流出空位
			}
		}
	}
}
func insertSortOtherDown(nums []int) {
	n := len(nums)
	//进行n-1轮迭代
	for i := 1; i <= n-1; i++ {
		deal := nums[i]
		j := i - 1 //带排序的数左边的第一个数的位置
		//如果是第一次比较，比左边的已排好序的第一个数小，那么直接插入chuli
		if deal > nums[j] {
			//一直往左边找，比待排序大的数都往后移动，腾空位给排序插入
			for ; j >= 0 && deal > nums[j]; j-- {
				nums[j+1], nums[j] = nums[j], nums[j+1] //数往后移，给待排序的数流出空位
			}
		}
	}
}

func main() {
	nums := []int{1, 2, 34, 5, 35}
	/*	insertSortUp(nums)
		fmt.Println(nums)*/
	insertSortOtherUp(nums)
	fmt.Println(nums)
	insertSortOtherDown(nums)
	fmt.Println(nums)
}
