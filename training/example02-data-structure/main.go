package main

import (
	"fmt"
)

func main() {

	//Array 裡面有數字是固定大小
	var arr [2]int
	arr[0] = 1
	arr[1] = 2
	fmt.Println(arr)

	//slice  裡面沒數字可以變相擴充
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	//共享相同底層陣列
	a := make([]int, 0, 10)
	b := append(a, 1, 2, 3)
	//a長度為0 所以他會依照順序0,1,2 寫入到底層陣列位置0,1,2
	//導致影響到b再取值的時候
	_ = append(a, 99, 88, 77)
	fmt.Println(b) // 99 88 77

	//超出容量會產生新的底層陣列
	aa := make([]int, 0, 2)
	//因為容量為2 會產生新的底層陣列
	bb := append(aa, 1, 2, 3)
	_ = append(aa, 99, 88, 77)
	fmt.Println(bb) //1 2 3

	// Linked List
	list := &LinkedList{}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Print()

	//stack
	stack := &Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	//1 2 3 4
	fmt.Println(stack)
	stack.Pop()
	// 1 2 3
	fmt.Println(stack)
	//3
	fmt.Println(stack.Peek())

	//queue
	queue := &Queue{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)

	// 1 2 3 4
	fmt.Println(queue)
	queue.Dequene()
	// 2 3 4
	fmt.Println(queue)
	//2
	fmt.Println(queue.Peek())

	//Hash Table
	hashMap := make(map[string]int)
	hashMap["apple"] = 42

	value, exist := hashMap["apple"]
	if exist {
		fmt.Println(value)
	}

	//Binary Tree
	treeNode := &TreeNode{}
	treeNode.Value = 5
	treeNode.Left = &TreeNode{Value: 4}
	treeNode.Right = &TreeNode{Value: 6}
	fmt.Println(treeNode)
	fmt.Println(treeNode.Left)
	fmt.Println(treeNode.Right)

}

// Linked List
type Node struct {
	Value int
	Next  *Node
}
type LinkedList struct {
	Head *Node
}

func (l *LinkedList) Append(val int) {
	newNode := &Node{Value: val}
	if l.Head == nil {
		l.Head = newNode
		return
	}
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

func (l *LinkedList) Print() {
	current := l.Head
	for current != nil {
		fmt.Printf("%d -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}

// Stack
type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		return -1
	}
	//假設目前長度是5 取4
	lastIndex := len(s.items) - 1
	item := s.items[lastIndex]
	//切割陣列長度到4
	s.items = s.items[:lastIndex]
	return item
}

func (s *Stack) Peek() int {
	if len(s.items) == 0 {
		return -1
	}
	return s.items[len(s.items)-1]
}

// queue
type Queue struct {
	items []int
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequene() int {
	if len(q.items) == 0 {
		return -1
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *Queue) Peek() int {
	if len(q.items) == 0 {
		return -1
	}
	return q.items[0]
}

// Binary Tree
type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}
