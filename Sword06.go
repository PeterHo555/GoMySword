package main

type ListNode struct {
	Val int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	//if head != nil {
	//	return nil
	//}
	//var newHead *ListNode
	//res := []int{}
	//for head != nil {
	//	node := head.Next
	//	head.Next = newHead
	//	newHead = head
	//	head = node
	//}
	//for newHead != nil {
	//	res = append(res, newHead.Val)
	//	newHead = newHead.Next
	//}
	//return res
	node := head
	len := 0
	for node != nil {
		node = node.Next
		len++
	}
	ans := make([]int,len)
	for i := len - 1; i >= 0; i-- {
		ans[i] = head.Val
		head = head.Next
	}
	return ans
}
