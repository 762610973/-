# LeetCode--数组

## 二分查找

**建议**：在理解二分查找的时候一定要手动的画图，模拟整个过程，这样才能理解二分超以及类似的题

### [704.二分查找](https://leetcode-cn.com/problems/binary-search/)

> 给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。

```go
package main
func search(nums []int, target int) int {
	l,r := 0,len(nums)-1
	for l <= r {
		mid := l + (r-l) / 2
		if  target > nums[mid] {
			l = mid + 1 //目标值在中间值的右侧，更新l
		} else if target < nums[mid]{
			r = mid - 1 //目标值在中间值的左侧，更新r
		} else {
			return mid
		}
	}
	return -1
}
//时间复杂度为O(log2(n)),空间复杂度为O(1)
```

### [35. 搜索插入位置](https://leetcode-cn.com/problems/search-insert-position/)

> 1、给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
>
>  2、必须使用时间复杂度为O(logn)的算法

```go
package main
func searchInsert(nums []int, target int) int {
	l,r := 0,len(nums)-1
	for l <= r {
		mid := l + (r-l) >> 1
		if target == nums[mid] {
			return mid //相等直接返回mid（当前元素的索引下标）
		} else if target > nums[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l //或者是return mid + 1，寻找的是>=target值的第一个元素的下标
    // 本题的关键点是返回值的取值
```

### [34. 在排序数组中查找元素的第一个和最后一个位置](https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/)

> 给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
>
> 如果数组中不存在目标值 target，返回 [-1, -1]。
>

```go
func searchFirst(nums []int,target int) int {
	left ,right := 0,len(nums) - 1
	for left <= right {
		mid := left + (right-left) >> 1
		if target < nums[mid]{
			right = mid - 1
		}else if target > nums[mid]{
			left = mid + 1
		}else{
			if mid == 0 || nums[mid-1]!=target {
				return mid
			}
// 如果这个数是第一个元素，即下标为0，那么可以直接return
// 如果这个元素的前一个位置不等于target,那么也可以直接return
// 例如1，2，3,3...假设现在到了第二个三的位置，现在是不能return的，因为他的前一个位置还有3
// 此时还要继续进行循环，如果到了第1个3的位置，他的前一个元素不等于3了，那么就可以直接return了
			right = mid - 1//不相等的话，往前判断，对比查找最后一个元素，这里是要让right往前
// 可以理解这样target<=nums[mid]，tright = mid - 1,可以这样考虑
// 1，2，3，3,3,4,5，第一次正好是中间的3，我们要找的是第一次出现的，肯定是往左找，所以小于等于目标值，都要往前找
		}
	}
	return -1
}
// 查找最后一次出现的位置
func searchOut(nums []int,target int) int {
	left ,right := 0,len(nums) - 1
	for left <= right {
		mid := left + (right-left) >> 1
		if target < nums[mid]{
			right = mid - 1
		}else if target > nums[mid]{
			left = mid + 1
		}else{
			if mid == len(nums)-1 || nums[mid+1]!=target{
				return mid//mid是最后一个或者不是下一个不是目标值
			}
			left = mid + 1
		}
	}
	return -1
}
//这种写法最有意思的是最后一步else的处理方法，很值得学习
func searchRange(nums []int, target int) []int {
	return []int{searchFirst(nums ,target),searchOut(nums,target)}
}
```

### [69. Sqrt(x)](https://leetcode-cn.com/problems/sqrtx/)

> 给你一个非负整数 x ，计算并返回 x 的 算术平方根 。
>
> 由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。
>
> 注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。
>

```go
func mySqrt(x int) int {
	left,right:= 0,x
	for left <= right {
		mid := left+(right-left)>>1
		res := mid * mid
		if res == x {
			return mid
		}else if res < x {
			left = mid + 1
		}else {
			right = mid - 1
		}
	}
	// 寻找的是小于等于根号x的最后一个数，所以返回right
	return right
}
```



### [367. 有效的完全平方数](https://leetcode-cn.com/problems/valid-perfect-square/)

> 给定一个 正整数 num ，编写一个函数，如果 num 是一个完全平方数，则返回 true ，否则返回 false 。
>
> 进阶：不要 使用任何内置的库函数，如  sqrt。
>

```go
package main
//二分法解决即可
func isPerfectSquare(num int) bool {
	if num == 0 || num == 1 {
		return true
	}
	left,right := 1,num>>1
	for left <= right {
		mid := left + (right-left)/2
		temp := mid*mid
		if temp < num {
			left = mid + 1
		}else if temp > num {
			right = mid - 1
		}else {
			return true
		}
	}
	return false
}
//这是时间复杂度比较高的方法，作用不大
func isPerfectSquare(num int) bool {
	for x := 0 ; x <= num ;x ++{
		res := x*x
		if res == num {
			return true
		}
	}
	return false
}
```



## 移除元素





## 有序数组的平方





## 长度最小的子数组





##  螺旋矩阵 II



