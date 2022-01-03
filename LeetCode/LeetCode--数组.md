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
	return l 
    // 或者是return mid + 1，寻找的是>=target值的第一个元素的下标
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

### [27. 移除元素](https://leetcode-cn.com/problems/remove-element/)

> 给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
>
> 不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
>
> 元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。



```go
func removeElement(nums []int, val int) int {
	// 暴力破解
	l := len(nums)
	for i := 0 ;i < l ; i ++ {//遍历数组元素
		if nums[i] == val {//匹配成功，发现需要移除的元素
			for j := i ;j < l-1 ; j ++ {//循环更新数组
				nums[j] = nums[j+1]//这里要保证的是最后一个元素不超出范围，但是又能够
			}
			i --//下标i以后的元素向前移动一位，为了使下次循环指向的元素是正确的，所以需要更新i
				//i--是考虑到val可能出现多次，如果有多次，在第一次之后，不进行i--
				//那么之后i会更新到下一个，可是没更新前的i被后面那个相同的值替代了，
			l --//移除一位元素，长度要减一
		}

	}
	return l
}
func removeElement(nums []int, val int) int {
	//双指针法（快慢指针）：快指针负责遍历数组，寻找和val不相等的元素，慢指针用来记录新数组的长度，下面的写法是利用了Go的for range语法糖，没有明显的体现快指针++
	var res int = 0
	for _,m := range nums {
		if m != val {//不需要移除
			nums[res] = m//赋值即可
			res ++//慢指针 +1
		}
	}
	return res
}
func removeElement(nums []int, val int) int {
    var  res = 0
    l := len(nums)
    for i:= 0;i < l;i++ {
        if nums[i] != val {
            nums[res] = nums[i]
            res ++
        }
    }
    return res
}
```

### [26. 删除有序数组中的重复项](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/)

> 给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。
>
> 不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成

```go
func removeDuplicates(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	slow := 0
	for fast:= 1; fast < l; fast++ { 
    //i作为快指针，这里从1开始，因为快慢指针在最一开始肯定是相等的
		if nums[fast] != nums[slow] {
    //快慢指针不相等，慢指针向后移动一位，然后快指针赋值
			slow ++
			nums[slow] = nums[fast]
		}
	}
	return slow+1
	//时间复杂度：O(n)，其中 n是数组的长度。快指针和慢指针最多各移动n次。
	//空间复杂度：O(1)。只需要使用常数的额外空间。
}
```

### [283. 移动零](https://leetcode-cn.com/problems/move-zeroes/)

> 给定一个数组 `nums`，编写一个函数将所有 `0` 移动到数组的末尾，同时保持非零元素的相对顺序。

- 示例 

  ```
  输入: [0,1,0,3,12]
  输出: [1,3,12,0,0]
  ```

- 说明
  1. 必须在原数组上操作，不能拷贝额外的数组。
  2. 尽量减少操作次数。

```go
// 时间复杂度：O(n)O(n)O(n)，其中 nnn 为序列长度。每个位置至多被遍历两次。
// 空间复杂度：O(1)O(1)O(1)。只需要常数的空间存放若干变量。
package main
func moveZeroes1(nums []int)  {
	l := len(nums)
	for fast,slow := 0,0; fast < l; fast ++ {
		if nums[fast] != 0 {//fast是快指针，用来寻找不为0的元素
			//找到的话就交换快慢指针对应的元素，然后让慢指针++，这里的慢指针是负责存储结果的     
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow ++
		}
	}
}
//另一种写法，值得借鉴
func moveZeroes(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}
//另一种写法
func moveZeroes(nums []int) {
// 这里是先利用移除元素的办法，将所有非0元素移除，顺带记录0的个数，然后再循环将最后的几个数置0
	slow := 0
	for _,fast := range nums {
		if fast != 0 {
			nums[slow] = fast
			slow ++
		}
	}
	for slow < len(nums) {
		nums[slow] = 0
		slow ++
	}
}
```

### [844. 比较含退格的字符串](https://leetcode-cn.com/problems/backspace-string-compare/)

> 给定 `s` 和 `t` 两个字符串，当它们分别被输入到空白的文本编辑器后，请你判断二者是否相等。`#` 代表退格字符。
>
> 如果相等，返回 `true` ；否则，返回 `false` 。
>
> **注意：**如果对空文本输入退格字符，文本继续为空。

- 由于 # 号只会消除左边的一个字符，所以对右边的字符无影响，所以我们选择从后往前遍历 SSS，TTT 字符串。

```go
（用的是栈）
func build (str string) string {
	var s []byte
	for i:= range str {
		if str[i] != '#' {
			s = append(s, str[i])//入栈
		} else if len(s) > 0 {
			s = s[:len(s) - 1]//去除栈顶元素
		}
	}
	return string(s)//类型转换之后再返回
	//时间复杂度：O(N+M)，其中 N 和 M 分别为字符串 S 和 T 的长度。我们需要遍历两字符串各一次。
	//空间复杂度：O(N+M)，其中 N 和 M 分别为字符串 S 和 T 的长度。主要为还原出的字符串的开销。
}
func backspaceCompare(s string, t string) bool {
	return build(s) == build(t)
}
//双指针法（不是快慢指针）
func backspaceCompare(s string , t string) bool {
	skipS,skipT := 0,0
	i ,j:= len(s)-1,len(t)-1
	for i >=0 || j >= 0 {
        //前面的两个for循环是找出普通的字符进行比较，当前字符是#或者#的数量>0都可以向前推进，然后再根据现在i和j的长度，将它们作为索引比较对应的字符
        //意思是当前是#或者#数量>0都不用比较，因为可以做退格处理，此时也要注意更新长度,比较完#相关的之后再开始比较剩下的字符了
        //循环的条件当然是两个的长度>=0
		for i >=0 {
			if s[i] == '#' { //当前字符为#,数量+1，指针向前移动一位
				skipS++
				i--
			} else if skipS > 0 { //当前字符不是#，#数量-1，继续向前移动
				skipS--
				i--
			} else {//普通字符，直接退出当前循环，去下一轮
				break
			}
		}
		for j >= 0 {
			if t[j] == '#' {
				skipT++
				j--
			} else if skipT > 0 {
				skipT--
				j--
			} else {
				break
			}
		}
		if i >= 0 && j >= 0 { //两个的长度都>=0,依次比较
			if s[i] != t[j] {
				return false//不相等直接返回false
			}
		} else if i >= 0 || j >= 0 {//有一个小于0，直接返回false
			return false
		}
		i--
		j--
	}
	return true
}
//时间复杂度是O(m+n)
//空间复杂度是O(1)
```



## 有序数组的平方





## 长度最小的子数组





##  螺旋矩阵 II



