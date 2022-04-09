package main

import (
	"container/list"
	"fmt"
)

type People interface {
	show()
}
type student struct{}

func (s *student) show() {
}
func live() People {
	var stu *student
	return stu
}
func t() []int {
	var s []int
	return s
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	cur := root //用于遍历二叉搜索树
	for cur != nil {
		if cur.Val == key {
			break
		}
		if cur.Val > key {
			if cur.Left != nil { //如果不判断下一个节点是否为空，就像是求长度一样，最后就是nil了，上一次应该就是这个问题了
				cur = cur.Left
			}
		}
		if cur.Val < key {
			if cur.Right != nil { //确保能到达那个节点，而不是到达最后的nil
				cur = cur.Right
			}
		}
	}
	if cur.Val != key { //不相等的情况
		return root
	}
	if cur.Left == nil && cur.Right == nil {
		cur = nil
		return root
	}
	//如果左子树为空,直接让当前节点的右孩子补上他的位置
	if cur.Left == nil {
		cur = cur.Right
		return root
	}
	//右节点为空，只能让左孩子补上了
	if cur.Right == nil {
		cur = cur.Left
		return root
	}
	now := cur
	now = now.Right
	if now.Left == nil {
		cur = now
		return root
	} else {
		for now.Left != nil {
			now = now.Left
		}
		if now.Right == nil {
			cur = now
			now = nil
			return root
		} else {
			cur = now
			now = now.Right
			return root
		}
	}
}
func d(root *TreeNode) *TreeNode {
	left := root.Left
	root = root.Right
	root.Left = left
	fmt.Println(levelOrder(root))
	return root
}
func main() {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   3,
			Left:  &TreeNode{2, nil, nil},
			Right: &TreeNode{4, nil, nil},
		},
		Right: &TreeNode{
			Val:  6,
			Left: nil,
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	x := deleteNode(root, 7)
	fmt.Println(levelOrder(x))
}

//二叉树的层序遍历，先将根节点入队，然后进入循环，每次记录下当前队列的长度，循环这个长度，然后每次循环将队首元素拿出来（队首元素代表一棵子树），再将其左右子树放入到队列中
func levelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	queue := list.New()
	queue.PushBack(root) //根节点入队
	for queue.Len() > 0 {
		temp := make([]int, 0)
		length := queue.Len() //保存当前这一层的长度，依次遍历
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			//使用队列，顺序添加，出队的时候从对头开始
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			temp = append(temp, node.Val)
		}
		ans = append(ans, temp)
		temp = []int{} //清空这一层所存放的数
	}
	return ans
}
