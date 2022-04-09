# LeetCode--链表

### [203. 移除链表元素](https://leetcode-cn.com/problems/remove-linked-list-elements/)

> 给你一个链表的头节点 `head` 和一个整数 `val` ，请你删除链表中所有满足 `Node.val == val` 的节点，并返回 **新的头节点** 。

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    dummyHead := &ListNode {
        Val:0,
        Next:head,
    }//设置一个虚拟的哑结点，指向头结点，可以统一处理所有的节点
    cur := dummyHead//cur用来遍历
    for cur != nil && cur.Next != nil {
        if cur.Next.Val == val {
            cur.Next = cur.Next.Next
        } else {
            cur = cur.Next
        }
    }
    return dummyHead.Next
}
   // 时间复杂度：O(n)O，其中 n 是链表的长度。需要遍历链表一次。
    //空间复杂度：O(1)。
```

### [707. 设计链表](https://leetcode-cn.com/problems/design-linked-list/)

> 设计链表的实现。您可以选择使用单链表或双链表。单链表中的节点应该具有两个属性：val 和 next。val 是当前节点的值，next 是指向下一个节点的指针/引用。如果要使用双向链表，则还需要一个属性 prev 以指示链表中的上一个节点。假设链表中的所有节点都是 0-index 的。
>
> 在链表类中实现这些功能：
>
> ```go
> get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
> addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
> addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
> addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
> deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。
> ```
>

```go
package main

type Node struct {
	val int
	next *Node
}

type MyLinkedList struct {
	head *Node
}

func Constructor() MyLinkedList {
	head := &Node{
		-1,
		nil,
	}
	return MyLinkedList{
		head,
	}
}
// 带头结点的链表
func (this *MyLinkedList) Get(index int) int {
	p := this.head
	i := 0
	for p != nil && i <= index {
		p = p.next
		i++
	}
	if p == nil || index < 0 {
		return -1
	}
	return p.val
}
func (this *MyLinkedList) AddAtHead(val int) {
	q := &Node{val,this.head.next}
	this.head.next = q
}
func (this *MyLinkedList) AddAtTail(val int) {
	//先找到链表的第一个元素
	p := this.head
	for p.next != nil {
		p = p.next
	}
	p.next = &Node{val,nil}
}
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	p := this.head
	i := 0
	//找到第i-1个
	for p.next != nil && i < index {
		p = p.next
		i++
	}
	if p.next == nil && index > i {
		return
	}
	q := &Node{val,nil}
	q.next = p.next
	p.next = q
}
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 {
		return
	}
	p := this.head
	i := 0
	for p != nil && i < index {
		p = p.next
		i++
	}
	if p.next == nil && index >= i {
		return
	}
	q := p.next
	p.next = q.next
	q.next = nil
}



/*package main
type Node struct {
	val int
	next *Node
}
type MyLinkedList struct {
	head *Node
	length int
}
func Constructor() MyLinkedList {
	return MyLinkedList{&Node{}, 0}
}
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.length {
		return -1
	}
	p := this.head
	for i := 0; i <= index; i++ {
		p = p.next
	}
	return p.val
}
func (this *MyLinkedList) AddAtHead(val int) {
	this.head.next = &Node{val,this.head.next}
	this.length++
}
func (this *MyLinkedList) AddAtTail(val int) {
	p := this.head
	for p.next != nil {
		p = p.next
	}
	p.next = &Node{val,nil}
	this.length++
}
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	if index <= 0 {
		this.AddAtHead(val)
		return
	} else if index == this.length {
		this.AddAtTail(val)
		return
	} else if index > this.length {
		return
	}
	p := this.head
	for i:= 0; i < index; i++ {
		p = p.next
	}
	q := &Node{val,nil}
	q.next = p.next
	p.next = q
	this.length++
}
func (this *MyLinkedList) DeleteAtIndex(index int)  {
	if index < 0 || index >= this.length {
		return
	}
	p := this.head
	for i := 0; i < index; i++ {
		p = p.next
	}
	q := p.next
	p.next = q.next
	q.next = nil
	this.length--
}*/
/*package main

//带长度带尾指针的单链表

type Node struct {
	val int
	next *Node
}
type MyLinkedList struct {
	head *Node
	tail *Node
	length int
}
func Constructor() MyLinkedList {
	p := &Node{}
	return MyLinkedList{p,p,0}
}
func (this *MyLinkedList)Get(index int) int {
	if index < 0 || index >= this.length {
		return -1
	}
	p := this.tail
	for i:= 0; i <= index; i++ {
		p = p.next
	}
	return p.val
}
func (this *MyLinkedList) AddAtHead (val int) {
	p := this.head
	p.next = &Node{val,p.next}
	if p.next.next == nil {
		this.tail = p.next
	}
	this.length++
}
// 将值为 val 的节点追加到链表的最后一个元素。

func (this *MyLinkedList) AddAtTail(val int)  {
	// 先找到链表的最后一个元素
	p := this.tail
	p.next = &Node{val, nil}
	this.tail = this.tail.next//this.tail.next 就是p.next
	this.length++
}
// 在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。

func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	if index <= 0 {
		this.AddAtHead(val)
		return
	}else if index == this.length {
		this.AddAtTail(val)
		return
	} else if index > this.length {
		return
	}

	p := this.head
	// 在第index个节点之前添加值，则需要找到第index-1个结点的位置
	for i := 0; i < index; i++ {
		p = p.next
	}
	q := &Node{val, nil}
	q.next = p.next
	p.next = q
	this.length++
}

// 如果索引 index 有效，则删除链表中的第 index 个节点

func (this *MyLinkedList) DeleteAtIndex(index int)  {
	if index < 0 || index >= this.length {
		return
	}
	p := this.head

	// 删除第index个结点，则需要找到第index-1个结点
	for i := 0; i < index; i++ {
		p = p.next
	}
	q := p.next
	p.next = q.next
	q.next = nil
	if this.length-1 == index {
		this.tail = p
	}
	this.length--
}*/

/*package main
//循环双链表

type MyLinkedList struct {
	dummy *Node
}

type Node struct {
	Val  int
	Next *Node
	Pre  *Node
}

//仅保存哑节点，pre-> rear, next-> head

func Constructor() MyLinkedList {
	rear := &Node{
		Val:  -1,
		Next: nil,
		Pre:  nil,
	}
	rear.Next = rear
	rear.Pre = rear
	return MyLinkedList{rear}
}


func (this *MyLinkedList) Get(index int) int {
	head := this.dummy.Next
	//head == this, 遍历完全
	for head != this.dummy && index > 0 {
		index--
		head = head.Next
	}
	//否则, head == this, 索引无效
	if 0 != index {
		return -1
	}
	return head.Val
}

func (this *MyLinkedList) AddAtHead (val int) {
	dummy := this.dummy
	node := &Node{
		Val: val,
		//head.Next指向原头节点
		Next: dummy.Next,
		//head.Pre 指向哑节点
		Pre: dummy,
	}

	//更新原头节点
	dummy.Next.Pre = node
	//更新哑节点
	dummy.Next = node
	//以上两步不能反
}

func (this *MyLinkedList) AddAtTail(val int) {
	dummy := this.dummy
	rear := &Node{
		Val: val,
		//rear.Next = dummy(哑节点)
		Next: dummy,
		//rear.Pre = ori_rear
		Pre: dummy.Pre,
	}

	//ori_rear.Next = rear
	dummy.Pre.Next = rear
	//update dummy
	dummy.Pre = rear
	//以上两步不能反
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	head := this.dummy.Next
	//head = MyLinkedList[index]
	for head != this.dummy && index > 0 {
		head = head.Next
		index--
	}
	node := &Node{
		Val: val,
		//node.Next = MyLinkedList[index]
		Next: head,
		//node.Pre = MyLinkedList[index-1]
		Pre: head.Pre,
	}
	//MyLinkedList[index-1].Next = node
	head.Pre.Next = node
	//MyLinkedList[index].Pre = node
	head.Pre = node
	//以上两步不能反
}


func (this *MyLinkedList) DeleteAtIndex(index int) {
	//链表为空
	if this.dummy.Next == this.dummy {
		return
	}
	head := this.dummy.Next
	//head = MyLinkedList[index]
	for head.Next != this.dummy && index > 0 {
		head = head.Next
		index--
	}
	//验证index有效
	/*if index == 0 {
		//MyLinkedList[index].Pre = index[index-2]
		head.Next.Pre = head.Pre
		//MyLinedList[index-2].Next = index[index]
		head.Pre.Next = head.Next
		//以上两步顺序无所谓
	}
}
*/



/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
```

### [206. 反转链表](https://leetcode-cn.com/problems/reverse-linked-list/)

> 给你单链表的头节点 `head` ，请你反转链表，并返回反转后的链表。

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 //双指针法
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
    // var pre *ListNode = nil //pre的初始值必须是空节点
	cur := head
	for cur != nil {
		tmp := cur.Next //保存cur的下一个节点，因为接下来要改变cur->Next
		cur.Next = pre //翻转操作
		//更新pre和cur指针
        pre = cur 
		cur = tmp //cur称为cur.Next，就是让cur向下一个移动
	}
	return pre
}
//  时间复杂度：O(n)，其中 n 是链表的长度。需要遍历链表一次。
//  空间复杂度：O(1)。
//  递归法
func reverseList(head *ListNode) *ListNode {
    return help(nil, head)
}

func help(pre, head *ListNode) *ListNode {
    if head == nil {
        return pre
    }
    tmp := head.Next //临时存储
    head.Next = pre
    return help(head, tmp)
}
//时间复杂度：O(n)，其中 n 是链表的长度。需要对链表的每个节点进行反转操作。
//空间复杂度：O(n)，其中 n 是链表的长度。空间复杂度主要取决于递归调用的栈空间，最多为 n 层。
```

### [24. 两两交换链表中的节点](https://leetcode-cn.com/problems/swap-nodes-in-pairs/)

> 给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。

```go
package main

func swapPairs(head *ListNode) *ListNode {
	//每次交换的是temp后面的两个节点
	dummyHead := &ListNode{0, head}
	temp := dummyHead
	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next
		node2 := temp.Next.Next
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		temp = node1
	}
	return dummyHead.Next
}
/*
   时间复杂度：O(n)，其中 n 是链表的节点数量。需要对每个节点进行更新指针的操作。
   空间复杂度：O(1))。
*/

/*func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    newHead := head.Next
    head.Next = swapPairs(newHead.Next)
    newHead.Next = head
    return newHead
}*/
/*
时间复杂度：O(n)，其中 n 是链表的节点数量。需要对每个节点进行更新指针的操作。
空间复杂度：O(n)，其中 n 是链表的节点数量。空间复杂度主要取决于递归调用的栈空间。
*/
```



### [19. 删除链表的倒数第 N 个结点](https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/)

> 给你一个链表，删除链表的倒数第 `n` 个结点，并且返回链表的头结点

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//双指针法
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{0,head}
	cur := head
	prev := dummyHead
	i := 1
	for cur != nil {
		cur = cur.Next
		if i > n {
			prev = prev.Next
		}
		i++
	}
	prev.Next = prev.Next.Next
	return dummyHead.Next
}
//时间复杂度：O(L)，其中 L 是链表的长度。
//空间复杂度：O(1)。

//计算链表长度
func getLength(head *ListNode)(length int) {
    cur := head
    for cur != nil {
        cur = cur.Next
        length++
    }
    return
}
func removeNthfromEnd(head *ListNode,n int) *ListNode {
    length := getLength(head)
    dummy := &ListNode{0,head}
    cur := dummy
    for i := 0;i < length-n;i++{
        cur = cur.Next
    }
    cur.Next = cur.Next.Next
    return dummy.Next
}
    //时间复杂度：O(L)，其中 LLL 是链表的长度。
    //空间复杂度：O(1)。

//栈
func removeNthfromEnd(head *ListNode,n int) *ListNode {
    nodes := []*ListNode{}
    dummy := &ListNode{0,head}
    for node := dummy;node != nil;node = node.Next{
        nodes = append(nodes,node)
    }
    prev := nodes[len(nodes)-1-n]
    prev.Next = prev.Next.Next
    return dummy.Next
}
//时间复杂度：O(L)，其中 L 是链表的长度。
//空间复杂度：O(L)，其中 L 是链表的长度。主要为栈的开销。

```



### [链表相交](https://leetcode-cn.com/problems/intersection-of-two-linked-lists-lcci/solution/mian-shi-ti-0207-lian-biao-xiang-jiao-sh-b8hn/)

> 给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null 。 

```go
package main
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lengthA,lengthB := 0,0
	curA,curB := headA,headB
	for curA != nil {
		lengthA++
		curA = curA.Next
	}
	for curB != nil {
		lengthB++
		curB = curB.Next
	}
	step := 0
	var fast,slow *ListNode = nil,nil
	if lengthA > lengthB {
		step = lengthA-lengthB
		fast,slow = headA,headB
	} else {
		step = lengthB-lengthA
		fast,slow = headB,headA
	}

	for i := 0; i < step; i++ {
		fast = fast.Next
	}
	for fast != slow  {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
//时间复杂度：O(n + m)
//空间复杂度：O(1)

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a,b := headA,headB
	for a != b {
		if a != nil {
			a = a.Next
		} else {
			a = headB
		}
		if b != nil {
			b = b.Next
		} else {
			b = headA
		}
	}
	return a
}
/*
时间复杂度 O(a+b) ： 最差情况下（即 ∣a−b∣=1|a - b| = 1∣a−b∣=1 , c=0c = 0c=0 ），此时需遍历 a+ba + ba+b 个节点。
空间复杂度 O(1)：节点指针 A , B 使用常数大小的额外空间。
*/
```



### [环形链表](https://leetcode-cn.com/problems/linked-list-cycle/solution/huan-xing-lian-biao-by-leetcode-solution/)

> 给你一个链表的头节点 head ，判断链表中是否有环。

```go
package main
func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil {
		return false
	}
	fast,slow := head.Next,head
	for fast != slow {
		//如果fast指向最后一个元素，那么不是环形链表，此时head.Next == nil, return false
		//如果fast == nil,直接return false
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true

}
/*
时间复杂度：O(N)，其中 N 是链表中的节点数。
    当链表中不存在环时，快指针将先于慢指针到达链表尾部，链表中每个节点至多被访问两次。
    当链表中存在环时，每一轮移动后，快慢指针的距离将减小一。而初始距离为环的长度，因此至多移动 NNN 轮。
空间复杂度：O(1)。我们只使用了两个指针的额外空间
*/


func hasCycle(head *ListNode) bool {
	//seen := map[*ListNode]int{}//这个后面的{}必须加,value是什么类型都可以
	var seenTest = make(map[*ListNode]int,0)
	for head != nil {
		if _, ok := seenTest[head]; ok {
			return true
		}
		seenTest[head] = 1
		head = head.Next
	}
	return false
}
func hasCycle(head *ListNode) bool {
	seen := make(map[*ListNode]struct{},0)
	seen := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := seen[head]; ok {
			return true
		}
		seen[head] = struct{}{}//struct{}是一种类型
		head = head.Next
	}
	return false
}

/*时间复杂度：O(N)，其中 N 是链表中的节点数。最坏情况下我们需要遍历每个节点一次。
空间复杂度：O(N)，其中 N 是链表中的节点数。主要为哈希表的开销，最坏情况下我们需要将每个节点插入到哈希表中一次。
*/

//环形链表2

package main
func detectCycle(head *ListNode) *ListNode {
    slow,fast := head,head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		//快慢指针相遇，此时conghead和相遇点，同时查找直至相遇
		if slow == fast {
			for slow != head {
				slow = slow.Next
				head = head.Next
			}
			return head
			}
		}
	return nil
}
```

#### [142. 环形链表 II](https://leetcode-cn.com/problems/linked-list-cycle-ii/)

> 给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
>
> 如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
>
> 不允许修改 链表。

```go
package main
func detectCycle(head *ListNode) *ListNode {
    slow,fast := head,head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		//快慢指针相遇，此时从head和相遇点，同时查找直至相遇，此时相遇的节点就是环的入口节点
		if slow == fast {
			for slow != head {
				slow = slow.Next
				head = head.Next
			}
			return head
			}
		}
	return nil
}
```

### [21. 合并两个有序链表](https://leetcode-cn.com/problems/merge-two-sorted-lists/)

> 将两个升序链表合并为一个新的 **升序** 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 

```go
//递归法，很巧妙
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil {
        return list2
    } else if list2 == nil {
        return list1
    } else if list1.Val < list2.Val {
        list1.Next = mergeTwoLists(list1.Next,list2)
        return list1
    } else {
        lsit2.Next = mergeTwoLists(list1,list2.Next)
        return list2
    }
}
// 时间复杂度O(m+n)
// 空间复杂度O(m+n)
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	ans := &ListNode{}//要返回的结果
	cur := ans//用于遍历
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1 //list1的Val小于list2的，将当前结点放入结果链表中
			list1 = list1.Next //节点后移
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next //用于遍历的指针后移
	}
    //考虑到两个链表长度不一样，所以要接着判断
	if list1 != nil {
		cur.Next = list1
	} else {
		cur.Next  = list2
	}
	return ans.Next
}
//时间复杂度 O(m+n)
//空间复杂度 O(1)
```

### [链表求和]()

> 给定两个用链表表示的整数，每个节点包含一个数位，这些数为是反向存放的，也就是个位排在链表首部，编写函数对这两个整数求和，并用链表形式返回结果

```go
package main
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	node := &ListNode{}
	result := node
	t := 0 // 大于10的要进1
	for l1 != nil || l2 != nil {
		num := t//这一步比较巧妙
		if l1 != nil {
			num += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			num += l2.Val
			l2 = l2.Next
		}
		// 处理节点
		node.Next = &ListNode{Val: num % 10} // 位数对应的数字
		node = node.Next
		t = num / 10 // 逢十进一
	}
	if t > 0 { // 进1的要处理干净，补位
		node.Next = &ListNode{t,nil}
	}
	return result.Next
}
```



