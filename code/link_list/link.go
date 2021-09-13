package main

import "fmt"

/*
203. 移除链表元素
给你一个链表的头节点 head 和一个整数 Val ，请你删除链表中所有满足 Node.Val == Val 的节点，并返回 新的头节点 。
*/

type ListNode struct {
	Val int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode {
		Next: head,
	}

	cur := dummyHead
	for cur != nil && cur.Next != nil {
		if cur.Next.Val != val {
			cur = cur.Next
		} else {
			cur.Next = cur.Next.Next
		}
	}

	return dummyHead.Next
}

/*
707. 设计链表
设计链表的实现。您可以选择使用单链表或双链表。
单链表中的节点应该具有两个属性：Val 和 Next。Val 是当前节点的值，
Next 是指向下一个节点的指针/引用。如果要使用双向链表，
则还需要一个属性 prev 以指示链表中的上一个节点。假设链表中的所有节点都是 0-index 的。

在链表类中实现这些功能：
get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
addAtHead(Val)：在链表的第一个元素之前添加一个值为 Val 的节点。插入后，新节点将成为链表的第一个节点。
addAtTail(Val)：将值为 Val 的节点追加到链表的最后一个元素。
addAtIndex(index,Val)：在链表中的第 index 个节点之前添加值为 Val 
的节点。如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，
则不会插入节点。如果index小于0，则在头部插入节点。
deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。
*/

type MyLinkedList struct {
	Val  int
	Next *MyLinkedList
}


/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList {
		Val:  0,
		Next: nil,
	}
}


/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	curIndex := 0
	cur := this

	for cur != nil {
		if curIndex == index {
			return cur.Val
		}

		cur = cur.Next
		curIndex++
	}

	return -1
}


/** Add a node of value Val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int)  {
	tmp :=  &MyLinkedList {
		Val:  val,
		Next: this,
	}

	this = tmp
}


/** Append a node of value Val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
	if this == nil {
		this = &MyLinkedList{Val:  val,}
		return
	}

	cur := this
	for {
		if cur.Next == nil {
			cur.Next = &MyLinkedList{Val: val,}
			return
		}

		cur = cur.Next
	}
}


/** Add a node of value Val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	if index < 0 {
		this.AddAtHead(val)
	}

	dummyHead := &MyLinkedList {Next: this,}
	cur := dummyHead
	curIndex := 0

	for {
		if cur.Next == nil {
			if curIndex == index {
				cur.Next = &MyLinkedList{Val:  val,}
			}

			break
		}

		if curIndex == index {
			tmp := &MyLinkedList{
				Val:  val,
				Next: cur.Next,
			}

			cur.Next = tmp
			break
		}

		cur = cur.Next
		curIndex++
	}

	this = dummyHead.Next
}


/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int)  {
	dummyHead := &MyLinkedList {Next: this,}
	cur := dummyHead
	curIndex := 0

	for cur.Next != nil {
		if curIndex == index {
			cur.Next = cur.Next.Next
			break
		}

		cur = cur.Next
		curIndex++
	}

	this = dummyHead.Next
}

/*
206. 反转链表
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
*/

//双指针法
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

//递归
func reverseList1(head *ListNode) *ListNode {
	return help(nil, head)
}

func help(pre, cur *ListNode)*ListNode{
	if cur == nil {
		return pre
	}
	next := cur.Next
	cur.Next = pre
	return help(cur, next)
}

/*
24. 两两交换链表中的节点
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
*/

func swapPairs(head *ListNode) *ListNode {
	virtual := &ListNode {
		Val:  0,
		Next: head,
	}

	cur := virtual
	for head != nil && head.Next != nil {
		cur.Next = head.Next
		next := head.Next.Next
		head.Next.Next = head
		head.Next = next
		cur = head
		head = next
	}

	return virtual.Next
}

/*
19. 删除链表的倒数第 N 个结点
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	virtual := &ListNode {Next: head,}
	pre := virtual
	fast := head
	i := 0
	for fast != nil {
		fast = fast.Next
		if i >= n {
			pre = pre.Next
		}
		i++
	}

	//一定存在？？？
	if i >= n {
		pre.Next = pre.Next.Next
	}
	return virtual.Next
}

/*
面试题 02.07. 链表相交
给定两个（单向）链表，判定它们是否相交并返回交点。请注意相交的定义基于节点的引用，而不是基于节点的值。
换句话说，如果一个链表的第k个节点与另一个链表的第j个节点是同一节点（引用完全相同），则这两个链表相交。
*/

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA := lenList(headA)
	lenB := lenList(headB)
	curA := headA
	curB := headB
	if lenA > lenB {
		curA = nextIDXNode(curA, lenA - lenB)
	} else if lenB > lenA {
		curB = nextIDXNode(curB, lenB - lenA)
	}

	for curA != curB {
		curA = curA.Next
		curB = curB.Next
	}

	return curA
}

func lenList(head *ListNode) int {
	lenList := 0
	node := head
	for node != nil {
		lenList++
		node = node.Next
	}

	return lenList
}

func nextIDXNode(cur *ListNode, len int) *ListNode {
	for len > 0 {
		cur = cur.Next
		len--
	}

	return cur
}

/*
142. 环形链表 II
给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。
如果 pos 是 -1，则在该链表中没有环。注意，pos 仅仅是用于标识环的情况，并不会作为参数传递到函数中。
说明：不允许修改给定的链表。
*/

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			cur := head
			for slow != cur {
				slow = slow.Next
				cur = cur.Next
			}
			return cur
		}
	}
	return nil
}

func main()  {
	tmp := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  8,
			Next: &ListNode{
				Val:  4,
				Next: &ListNode{
					Val:  5,
					Next: nil,
				},
			},
		},
	}

	headA := &ListNode{
		Val:  5,
		Next: &ListNode{
			Val:  0,
			Next: tmp,
		},
	}

	headB := &ListNode{
		Val:  4,
		Next: tmp,
	}

	if res := getIntersectionNode(headA, headB); res != nil {
		fmt.Println(*res)
	}
}
