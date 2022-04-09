package main

import "fmt"

type node struct {
	val  int
	next *node
}

func xxx(head *node) *node {
	cur := head
	fmt.Printf("%p,%p", head, cur)
	return head

}
func pr(head *node) {
	for head != nil {
		fmt.Println(head)
		head = head.next
	}
}
func main() {
	var y *node = new(node)
	y.val = 3
	y.next = nil
	fmt.Println(y)
	head := &node{1, &node{2, &node{3, nil}}}
	temp := xxx(head)
	pr(temp)
}
