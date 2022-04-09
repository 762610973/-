/*
package main

import "fmt"

//归并排序，时间复杂度是O(nlogn),稳定

func mergeSort(array []int, begin, end int) {
	if end-begin <= 1 {
		return
	}
	mid := (begin + end) >> 1
	//对左边的进行排序
	mergeSort(array, begin, mid)
	//对右边的进行排序
	mergeSort(array, mid, end)
	//现在这个数组的左右两边已经排好序了，开始进行合并
	merge(array, begin, mid, end)

}
func merge(array []int, begin, mid, end int) {
	//申请额外的空间来合并两个有序数组
	//这两个数组是有序数组，array[begin,mid),array[mid,end)
	result := make([]int, 0)
	l, r := begin, mid+1
	for l <= mid && r <= end {
		if array[l] < array[r] {
			result = append(result, array[l])
			l++
		} else if array[l] > array[r] {
			result = append(result, array[r])
			r++
		} else {
			result = append(result, array[l])
			l++
			r++
		}

	}
	if l <= mid {
		result = append(result, array[l:mid+1]...)
	}
	if r <= end {
		result = append(result, array[r:end+1]...)
	}

	for i, item := range result {
		array[begin+i] = item
	}
	return
}

*/
package main

import (
	"fmt"
	"math"
)

func mergeSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	mergeSort(arr, start, mid)
	mergeSort(arr, mid+1, end)
	merge(arr, start, mid, end)
}

func merge(arr []int, start, mid, end int) {
	var tempArr = make([]int, 0, len(arr))
	var s1, s2 = start, mid + 1
	for s1 <= mid && s2 <= end {
		if arr[s1] > arr[s2] {
			tempArr = append(tempArr, arr[s2])
			s2++
		} else if arr[s1] < arr[s2] {
			tempArr = append(tempArr, arr[s1])
			s1++
		} else {
			//相等的情况要添加两个值，之前犯错了，只添加了一个
			tempArr = append(tempArr, arr[s1], arr[s2])
			s1++
			s2++
		}
	}
	//剩下的接到后面
	if s1 <= mid {
		tempArr = append(tempArr, arr[s1:mid+1]...)
	}
	if s2 <= end {
		tempArr = append(tempArr, arr[s2:end+1]...)
	}
	for pos, item := range tempArr {
		arr[start+pos] = item
	}
}

func main() {
	test := []int{1, 451, 61, 6, 3, 25, 3, 436, 2}
	//fmt.Println(test[1:2])
	mergeSortLowToHigh(test)
	fmt.Println(test)
}
func mergeSortLowToHigh(arr []int) {
	var length = len(arr)
	//步长从1开始，1,2,4,8
	for size := 1; size <= length; size += size {
		//这里的步长是2*size，一个和一个合并，下一轮就是两个和两个合并了
		for i := 0; i+size < length; i += 2 * size { //对[i,i+size-1]和[i+size,i+2*size-1]进行归并
			//这里的第二个数组的右边界是左边界加上步长值，这个步长值是i+2*size-1，一定要减一的
			//不减一就是从i的下一个位置走2*size个位置，肯定是不行的
			//同时right-left+1应该是size的整倍数
			merge(arr, i, i+size-1, int(math.Min(float64(i+2*size-1), float64(length-1))))
			// arr left mid right	如果i+2*size>n了，越界了，就取n-1
		}
	}
}

/*func merge(arr []int, left, mid, right int) {
	// 将要合并的部分做个拷贝
	var tmp = make([]int, right-left+1)
	for i, j := left, 0; i <= right; i++ {
		tmp[j] = arr[i]
		j++
	}
	// i做为左半部分的指针   j作为右半部分的指针
	var i, j int = left, mid+1
	for k := left; k <= right; k++ {
		if i > mid { // 左半部分 已经合入完了，将右半部分剩下的 全部合入
			arr[k] = tmp[j-left]
			j++
		} else if j > right { // 右半部分 已经合入完了，将左半部分剩下的 全部合入
			arr[k] = tmp[i-left]
			i++
		} else if tmp[i-left] > tmp[j-left] {
			arr[k] = tmp[j-left]
			j++
		} else {
			arr[k] = tmp[i-left]
			i++
		}
	}
}*/
