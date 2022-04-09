package main

import "fmt"

//希尔排序，也称为缩小增量排序
//规模中等的情况下使用希尔排序，但是大规模还是要使用快速排序，归并排序或者堆排序
func shellSortUp(nums []int) {
	n := len(nums)
	//增量每次减半，直到步长为1
	for step := n / 2; step >= 1; step /= 2 {
		//开始插入排序，每一轮的步长为step
		for i := step; i < n; i += step {
			for j := i - step; j >= 0; j -= step {
				//满足就交换元素,这里的排序采用交换元素的方法
				if nums[j+step] < nums[j] {
					nums[j], nums[j+step] = nums[j+step], nums[j]
					continue
				}
				break
			}
		}
	}
}

//易于理解的希尔排序降序写法
func shellSortDown(nums []int) {
	n := len(nums)
	//增量每次减半，直到步长为1
	for step := n / 2; step >= 1; step /= 2 {
		//开始插入排序，每一轮的步长为step
		for i := step; i < n; i += step {
			now := nums[i]
			j := i - step
			//比前面的大就交换
			if now > nums[j] {
				for ; j >= 0 && now > nums[j]; j -= step {
					nums[j+step], nums[j] = nums[j], nums[j+step]
				}
			}

		}
	}
}
func shellSortOtherUp(nums []int) {
	n := len(nums)
	//增量每次减半，直到步长为1
	for step := n / 2; step >= 1; step /= 2 {
		//开始插入排序，每一轮的步长为step
		for i := step; i < n; i += step {
			now := nums[i]
			j := i - step
			//满足就交换元素,这里的排序采用交换元素的方法
			if now < nums[j] {
				for ; j >= 0 && now < nums[j]; j -= step {
					nums[j+step] = nums[j]
				}
				nums[j+step] = now
			}
		}
	}
}
func main() {
	t := []int{1, 324, 5, 12, 5, 99, 5, 2, 16, 2, 5, 34, 6, 3}
	/*shellSortUp(t)
	fmt.Println(t)*/
	/*	shellSortOtherUp(t)
		fmt.Println(t)*/
	shellSortDown(t)
	fmt.Println(t)
}
