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
// 时间复杂度：O(n)，其中 n 为序列长度。每个位置至多被遍历两次。
// 空间复杂度：O(1)。只需要常数的空间存放若干变量。
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

### [977. 有序数组的平方](https://leetcode-cn.com/problems/squares-of-a-sorted-array/)

> 给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。

```go
package main
func sortedSquares(nums []int) []int {
	// 双指针法，数组的最大值一定是第一个或者最后一个，不可能是其他值
	//一个指向头，一个指向尾
	// 如果A[i] * A[i] < A[j] * A[j] 那么result[k--] = A[j] * A[j];
	// 如果A[i] * A[i] >= A[j] * A[j] 那么result[k--] = A[i] * A[i];
	// 头大于尾，新数组的最后一个位置就是头，此时头向后移动，比较第二个和最后一个
	// 头小于尾，新数组的最后一个位置就是尾，尾向前移动一个，比较第一个和倒数第二个
	l := len(nums)
	i,j,k := 0,l-1,l-1
	//i指向数组的最左边，j指向最右边
	//k是新数组的索引
	res := make([]int, l)
	for i <= j { //这里的i<=j，是要处理两个元素
		lm,jm := nums[i]*nums[i], nums[j]*nums[j]
		if lm > jm {
			res[k] = lm
			i ++
		} else {
			res[k] = jm
			j --
		}
		k--//更新k的值，下次继续赋值，这里的k是从最后开始的
	}
	return res
	//时间复杂度是O(n)
}
```



## 长度最小的子数组

### 模板

```
// 模板
for () {
    // 将新进来的右边的数据，计算进来
    // 更新数据

    // 判断窗口数据是否不满足要求了
    while (窗口数据不满要求 && left < arrSize) {
        // 移除left数据，更新窗口数据
        left++;    
    }
    right++;
}
```



### [209. 长度最小的子数组](https://leetcode-cn.com/problems/minimum-size-subarray-sum/)

> 给定一个含有 n 个正整数的数组和一个正整数 target 。
>
> 找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。

#### 暴力破解

```go
package main
import "math"
func minSubArrayLen(target int, nums []int) int {
    length := len(nums)
    if length == 0 {
        return 0
    }
    res := math.MaxInt32
    for i := 0; i < length;i++ {
        sum := 0
        for j := i; j <length; j++ {
            sum += nums[j]
            if sum >= target {
                res = min(ans,j - i + 1)
                break
            }
        }
    }
    if ans == math.MaxInt32 {
        return 0
    }
    return res
}
func min(a,b int) int {
    if a < b {
        return a
    } 
    return b
}
//时间复杂度：O(n*2)
//空间复杂度: O(1)
```

#### 滑动窗口

```go
func minSubArrayLen(target int, nums[int]) int {
    length := len(nums)
    if length == 0 {
        return 0
    }
    ans := math.MaxInt32//设置为最大值，是为了确保第一次比较的时候ans能成为较小
    //的那个或者是不存在，如果设置为0什么的第一次比较会丧失真的那个
    start, end := 0, 0
    sum := 0 
    for end < length {
        sum += nums[end]
        for sum >= target {//找到符合的子数组，更新子数组的最小长度
            ans = min(ans,end-start+1)
            sum -= nums[start]//这里相减是因为start要右移了
            start++//start右移
        }
        end ++ //sum<target,end后移继续寻找
    }
    if ans == math.MathInt32 {
        return 0
    }
    return ans
}
func min(a,b int) int {
    if a < b {
        return a
    }
   	return b
}
时间复杂度：O(n)，其中 n 是数组的长度。指针 start 和 end 最多各移动 n 次。
空间复杂度：O(1)
```

### [904. 水果成篮](https://leetcode-cn.com/problems/fruit-into-baskets/)

> 你正在探访一家农场，农场从左到右种植了一排果树。这些树用一个整数数组 fruits 表示，其中 fruits[i] 是第 i 棵树上的水果 种类 。
>
> 你想要尽可能多地收集水果。然而，农场的主人设定了一些严格的规矩，你必须按照要求采摘水果：
>
>     你只有 两个 篮子，并且每个篮子只能装 单一类型 的水果。每个篮子能够装的水果总量没有限制。
>     你可以选择任意一棵树开始采摘，你必须从 每棵 树（包括开始采摘的树）上 恰好摘一个水果 。采摘的水果应当符合篮子中的水果类型。每采摘一次，你将会向右移动到下一棵树，并继续采摘。
>     一旦你走到某棵树前，但水果不符合篮子的水果类型，那么就必须停止采摘。
>
> 给你一个整数数组 fruits ，返回你可以收集的水果的 最大 数目。

- 问题转化为寻找最长的子序列，但是类型最多包含两种，例如2,3,2,2,3,4那么最长的就是从2开始到最后一个3，只包含两种类型

```go
package main
func totalFruit(fruits []int) int {
    if len(fruits) == 0 {
        return 0
    }
    //问题是求数组中包含两个值的最大连续子序列
    hashM := map[int]int{}
    //用map记录不同类型的水果篮子，以及水果篮子中水果的个数
    res := 0
    //记录每次得到的局部最大长度
    for start,end := 0,0; end < len(fruits); end++ {
        hashM[fruits[end]]++//这种类型的水果数量+1
        //遍历水果数组
        for start <= end && len(hashM) > 2 {
            //此时水果的种类>2，说明盘子中进入了第三种水果
            hashM[fruits[start]]--
            if hashM[fruits[start]] == 0 {
                delete(hashM,fruits[start])
            }
            start++
        }
        if res < end-start+1 {
            //将res更新为最大值
            res = end-start+1
        }
    }
    return res
}
    //时间复杂度：O(N)，其中 N 是 tree 的长度。
    //空间复杂度：O(N)。
```



### [76. 最小覆盖子串](https://leetcode-cn.com/problems/minimum-window-substring/)

> 给你一个字符串 `s` 、一个字符串 `t` 。返回 `s` 中涵盖 `t` 所有字符的最小子串。如果 `s` 中不存在涵盖 `t` 所有字符的子串，则返回空字符串 `""` 。

```go
func minWindow(s string, t string) string {
    if s == "" || t == "" {
        return ""
    }
    tMap,cnt := map[byte]int{},map[byte]int{}//tMap用来存储t的每个元素以及相应的个数，cnt用来存储s中已有的元素以及出现的次数
    for i := 0; i < len(t); i++ {
        tMap[t[i]]++//将t的每个元素及个数依次赋给tMap
    }
    sLen := len(s)
    length := math.MaxInt32
    ansL,ansR := -1,-1
    //这里是将匿名函数赋值给变量
    check := func() bool {
        for k,v := range tMap {
            if cnt[k] < v {
                //判断k是否存在于cnt当中，并且个数是否小于tMap中的k的个数
                return false
            }
        }
        return true //只要上面的不出现false，就是true
    }
    for left,right := 0,0;right < sLen; right++ {//遍历字符串s
        if right < sLen && tMap[s[right]] > 0 {//tMap中存在s[right]再添加，也就是cnt只存储tMap中存在的
            cnt[s[right]]++//将扩展指针指着的元素加入到cnt当中
        }
        for check() && left <= right {//缩小子串的范围
            if right-left+1 < length {//取最小的
                length = right-left + 1
                ansL,ansR = left,left + length//更新最终的l，r
            }
            //去除重复项，缩小范围
            if _,ok := tMap[s[left]]; ok {//如果tMap中存在s[left]
                cnt[s[left]] -= 1//cnt当中的s[left] --
            }
            left++//left向右移
        }
    }
    if ansL == -1 {
        return ""
    }
    return s[ansL:ansR] 
}
```



##  螺旋矩阵 II

### [59. 螺旋矩阵 II](https://leetcode-cn.com/problems/spiral-matrix-ii/)

> 给你一个正整数 `n` ，生成一个包含 `1` 到 `n*2` 所有元素，且元素按顺时针顺序螺旋排列的 `n x n` 正方形矩阵 `matrix` 。

```go
func generateMatrix(n int) [][]int {
    num := 1
    high, low := 0,n-1
    left, right := 0,n-1
    tar := n*n
    matrix := make([][]int,n)
    for i := 0; i < n; i++ {
        matrix[i] = make([]int, n)
    }
    for num <= tar {
        for i := left; i <= right; i++ {
            matrix[high][i] = num
            num++
        }
        high++
        for i := high; i <= low; i++ {
            matrix[i][right] = num
            num++
        } 
        right--
        for i := right; i >= left; i-- {
            matrix[low][i] = num
            num++
        }
        low--
        for i := low;i >= high; i-- {
            matrix[i][left] = num
            num++
        }
        left++
    } 
    return matrix
}
//时间复杂度是O(mn)
//空间复杂度是O(1)
```

### [54. 螺旋矩阵](https://leetcode-cn.com/problems/spiral-matrix/)

> 给你一个 `m` 行 `n` 列的矩阵 `matrix` ，请按照 **顺时针螺旋顺序** ，返回矩阵中的所有元素。

- 注：最大值在右下角，排列是按照从左到右，从上到下递增排列的

```go
func spiralOrder(matrix [][]int) []int {
	n := len(matrix)//列
	m := len(matrix[0])//行
	index := 0//返回的数组的下标
	end := m*n//一共m*n个数
	low,high := 0,n-1//上边界、下边界
	left,right := 0,m-1//左边界、右边界
	res := make([]int,end)//需要返回的数组
	for index < end  {
		//最上排从左到右
		for i := left; i <= right; i++ {
			if index < end {//这一步必须判断，在外层的for index < end 这一步，进入之后，index还会增长
				res[index] = matrix[high][i]
				index++
			}
		}
		high++//减少一行
        //最右列从上到下
		for i := high;i <= low; i++ {
			if index < end {
				res[index] = matrix[i][right]
				index++
			}
		}
		right--//减少一列
        //最下排，从右往左
		for i := right;i >= left; i-- {
			if index < end {
				res[index] = matrix[low][i]
				index++
			}
		}
		low--
        //最左列，从下往上
		for i := low; i >= high; i-- {
			if index < end {
				res[index] = matrix[i][left]
				index++
			}
		}
		left++
	}
	return res
}
```

### [剑指 Offer 29. 顺时针打印矩阵](https://leetcode-cn.com/problems/shun-shi-zhen-da-yin-ju-zhen-lcof/)

> 输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。

```go
func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return nil
	}
	n := len(matrix[0])
	if n == 0 {
		return nil
	}
	index := 0
	end := m * n
	high,low := 0,m-1
	left,right := 0,n-1
	res := make([]int,end)
	for index < end  {
		//最上排从左到右
		for i := left; i <= right; i++ {
			if index < end {
				res[index] = matrix[high][i]
				index++
			}
		}
		high++
		for i := high;i <= low; i++ {
			if index < end {
				res[index] = matrix[i][right]
				index++
			}
		}
		right--
		for i := right;i >= left; i-- {
			if index < end {
				res[index] = matrix[low][i]
				index++
			}
		}
		low--
		for i := low; i >= high; i-- {
			if index < end {
				res[index] = matrix[i][left]
				index++
			}
		}
		left++
	}
	return res
}
```

### [189.轮转数组]()

> 给你一个数组，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。

```go
package main
func rotate(nums []int, k int)  {
	length := len(nums)
	k = k % length
	ans := append(nums[length-k:],nums[:length-k]...)
	copy(nums,ans)
}
```

### [36.有效的数独](https://leetcode-cn.com/problems/valid-sudoku/solution/you-xiao-de-shu-du-by-leetcode-solution-50m6/)

> 请你判断一个9x9的数独是否有效。根据规则 ，验证已经填入的数字是否有效即可。 

```go
package main
func isValidSudoku(board [][]byte) bool {
	var rows, columns [9][9]int//九行九列
	var subboxs [3][3][9]int //横着数三个块，竖着也是三个，然后每一个块有九个数字
	for i,row := range board {//此时row是切片，代表着每一行
		for j,c := range row {//从第一行，第一列开始
			if c == '.' {
				continue
			}
			index := c-'1'//index表示索引，因为数组索引是从0开始的，所以所有的数字都要-1
			rows[i][index]++
			columns[j][index]++
			subboxs[i/3][j/3][index]++
			if rows[i][index] > 1 || columns[j][index] > 1 || subboxs[i/3][j/3][index]>1 {
				return false
			}
		}
	}
	return true
}
```



































