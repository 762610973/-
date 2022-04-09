package main

import "fmt"

//todo:最大堆，进行上浮操作：将元素放到末尾进行上浮
//todo：最小堆，进行下沉操作，首先将堆顶元素pop，然后将最尾部的节点放到根节点上，接着对根节点进行下沉操作

//一个最大堆，一棵完全二叉树
//最大堆要求节点元素都不小于其左右孩子
//O(nlogn)

type heap struct {
	//堆的大小
	size int
	//使用内部的数组来模拟树
	//一个节点下标为i，那么父亲节点的下标为(i-1)/2
	//一个节点下标为i，那么左儿子的下标为2i+1，右儿子下标为2i+2
	array []int
}

//初始化一个堆
func newHeap(array []int) *heap {
	h := new(heap)
	h.array = array
	return h
}

//最大堆插入元素
func (h *heap) push(x int) {
	//堆中没有元素，使元素成为顶点后退出
	if h.size == 0 {
		h.array[0] = x
		h.size++
		return
	}
	//i是要插入节点的下标
	i := h.size
	//如果下标存在
	//将小的值x一直上浮
	for i > 0 {
		//parent为钙元素节点的下标
		parent := (i - 1) / 2
		//如果插入的值小于等于父节点，那么可以直接退出循环，因为父亲仍然是最大的
		if x <= h.array[parent] {
			break
		}
		//否则将父亲节点与该节点互换，然后向上翻转，将最大的元素一直网上推
		h.array[i] = h.array[parent]
		//令i等于父亲节点的值
		i = parent
	}
	//将要插入的值x放在不会再翻转的位置
	h.array[i] = x
	//堆数量加1
	h.size++
}

//最大堆移除根节点元素，也就是最大的元素
func (h *heap) pop() int {
	if h.size == 0 {
		return -1
	}
	//取出根节点元素
	ret := h.array[0]
	//因为根节点要被删除，将最有一个节点放到根节点的位置上
	h.size--
	x := h.array[h.size]  //将最后一个元素的值拿出来
	h.array[h.size] = ret //将移除的元素放在最后一个元素的位置上
	//对根节点进行向下翻转，小的值一直下沉，维持最大堆的特征
	i := 0
	for {
		left := 2*i + 1
		right := 2*i + 2
		//左儿子下标超出了，表示没有左子树，那么右子树也没有，这个时候上面的操作已经完成了，直接返回
		if left >= h.size {
			break
		}
		//有右子树，拿到两个子节点中较大节点的下标
		//不满足条件直接在左子树上进行，完全不关心是否含有右子树或者两个节点的大小
		if right < h.size && h.array[right] > h.array[left] {
			left = right
		}
		//父亲节点的值都大于等于两个儿子较大的那个，不需要向下继续翻转，返回即可
		if x >= h.array[left] {
			break
		}
		//将较大的儿子与父亲交换，位置这个最大堆的特征
		h.array[i] = h.array[left]
		//继续往下操作,更新i
		i = left
	}
	//将最后一个元素的值x在不会再翻转的位置
	h.array[i] = x
	return ret
}

func main() {
	list := []int{1, 65, 5, 6, 12, 54, 45, 756, 2, 322, 132}
	//构建最大堆
	h := newHeap(list)
	fmt.Println(h)
	//这样可以保证不用append
	for _, v := range list {
		h.push(v)
	}
	//此时只是一个树形结构
	fmt.Println(h.array)
	//将对元素移除
	for range list {
		h.pop()
	}
	fmt.Println(list)
}

// 这种堆排序，不再每次都将元素添加到尾部，然后上浮翻转，而是在混乱堆的基础上，从底部向上逐层进行下沉操作，下沉操作比较的次数会减少。步骤如下：
//
//    先对最底部的所有非叶子节点进行下沉，即这些非叶子节点与它们的儿子节点比较，较大的儿子和父亲交换位置。
//    接着从次二层开始的非叶子节点重复这个操作，直到到达根节点最大堆就构建好了。
//
// 从底部开始，向上推进，所以这种堆排序又叫自底向上的堆排序。

func HeapSort(array []int) {
	// 堆的元素数量
	count := len(array)

	// 最底层的叶子节点下标，该节点位置不定，但是该叶子节点右边的节点都是叶子节点
	start := count/2 + 1

	// 最后的元素下标
	end := count - 1

	// 从最底层开始，逐一对节点进行下沉
	for start >= 0 {
		sift(array, start, count)
		start-- // 表示左偏移一个节点，如果该层没有节点了，那么表示到了上一层的最右边
	}

	// 下沉结束了，现在要来排序了
	// 元素大于2个的最大堆才可以移除
	for end > 0 {
		// 将堆顶元素与堆尾元素互换，表示移除最大堆元素
		array[end], array[0] = array[0], array[end]
		// 对堆顶进行下沉操作
		sift(array, 0, end)
		// 一直移除堆顶元素
		end--
	}
}

// 下沉操作，需要下沉的元素时 array[start]，参数 count 只要用来判断是否到底堆底，使得下沉结束
func sift(array []int, start, count int) {
	// 父亲节点
	root := start

	// 左儿子
	child := root*2 + 1

	// 如果有下一代
	for child < count {
		// 右儿子比左儿子大，那么要翻转的儿子改为右儿子
		if count-child > 1 && array[child] < array[child+1] {
			child++
		}

		// 父亲节点比儿子小，那么将父亲和儿子位置交换
		if array[root] < array[child] {
			array[root], array[child] = array[child], array[root]
			// 继续往下沉
			root = child
			child = root*2 + 1
		} else {
			return
		}
	}
}
