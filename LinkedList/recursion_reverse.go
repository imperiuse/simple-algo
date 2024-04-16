package main

import "fmt"

type ListNode struct {
	Val  string
	Next *ListNode
}

func (n *ListNode) printNode() {
	fmt.Printf("I'm: %p \t My value is: %+[1]v\n", n)
}

func printNodeRec(n *ListNode) {
	n.printNode()
	if n.Next != nil {
		printNodeRec(n.Next)
	} else {
		return
	}
}

func reverse(prev, n *ListNode) *ListNode {
	if n.Next != nil {
		head := reverse(n, n.Next)
		n.Next = prev

		return head

	}

	n.Next = prev

	return n
}

func main() {
	var (
		a = ListNode{Val: "A", Next: nil}
		b = ListNode{Val: "B", Next: nil}
		c = ListNode{Val: "C", Next: nil}
		d = ListNode{Val: "D", Next: nil}
		e = ListNode{Val: "E", Next: nil}
	)

	a.Next = &b
	b.Next = &c
	c.Next = &d
	d.Next = &e
	e.Next = nil

	printNodeRec(&a)
	fmt.Println()

	printNodeRec(reverse(nil, &a))
}
