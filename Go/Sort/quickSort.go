package main

import "fmt"

func quickSort(array []int, begin, end int) {
	if begin < end {
		//进行切分
		loc := partition(array, begin, end)
		//对左部分进行快排
		quickSort(array, begin, loc-1)
		//对右部分进行快排
		quickSort(array, loc+1, end)

	}
}

// 小规模数组采用直接插入效率最好，当快速排序递归部分进入小数组范围，可以切换成直接插入排序
// quickSort1(array []int,begin,end int)
func quickSort1(array []int, begin, end int) {
	if begin < end {
		//当数组小于4时直接使用插入排序
		if end-begin <= 4 {
			insertSortUp(array[begin : end+1])
			return
		}
	}
	//进行切分
	loc := partition(array, begin, end)
	//对左部分进行快排
	quickSort1(array, begin, loc-1)
	//对右部分进行快排
	quickSort1(array, loc+1, end)
}
func QuickSort2(array []int, begin, end int) {
	if begin < end {
		// 三向切分函数，返回左边和右边下标
		lt, gt := partition3(array, begin, end)
		// 从lt到gt的部分是三切分的中间数列
		// 左边三向快排
		QuickSort2(array, begin, lt-1)
		// 右边三向快排
		QuickSort2(array, gt+1, end)
	}
}

// 切分函数，并返回切分元素的下标
func partition3(array []int, begin, end int) (int, int) {
	lt := begin       // 左下标从第一位开始
	gt := end         // 右下标是数组的最后一位
	i := begin + 1    // 中间下标，从第二位开始
	v := array[begin] // 基准数

	// 以中间坐标为准
	for i <= gt {
		if array[i] > v { // 大于基准数，那么交换，右指针左移
			array[i], array[gt] = array[gt], array[i]
			gt--
		} else if array[i] < v { // 小于基准数，那么交换，左指针右移
			array[i], array[lt] = array[lt], array[i]
			lt++
			i++
		} else {
			i++
		}
	}

	return lt, gt
}

// quickSort3 伪递归快速排序,使用伪递归减少程序栈空间占用，使得空间复杂度为O(logn)
func quickSort3(array []int, begin, end int) {
	for begin < end {
		// 切分
		loc := partition(array, begin, end)
		//哪边元素少先排哪边
		if loc-begin < end-loc {
			//先排左边
			quickSort3(array, begin, loc-1)
			begin = loc + 1
		} else {
			quickSort3(array, loc+1, end)
			end = loc - 1
		}
	}
}

// partition 切分函数
func partition(array []int, begin, end int) int {
	i := begin + 1 //将array[begin]作为基准数，因此从array[begin+1]开始与基准数比较
	j := end       //array[end]是数组的最后一位
	//没重合之前
	for i < j {
		if array[i] > array[begin] {
			array[i], array[j] = array[j], array[i] //交换
			j--
		} else {
			i++
		}
	}
	//此时i到j之间是分好组的,还需要处理begin
	if array[i] >= array[begin] {
		//这里必须是 >=，然后进行下一步的交换元素，i--之后，那个array[i]就是小于的了
		i--
	}
	array[begin], array[i] = array[i], array[begin] //交换位置
	return i
}

func main() {
	s := []int{1, 2, 6, 3, 4, 764, 64, 64, 64, 64, 3, 23, 7}
	quickSort(s, 0, len(s)-1)
	fmt.Println(s)
}
