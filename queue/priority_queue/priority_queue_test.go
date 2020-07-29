package priority_queue

import (
	"fmt"
	"testing"
)

type  Int int

func(i Int)Less(j interface{})bool{
	return i<j.(Int)
}

func TestPriorityQueue(t *testing.T) {
	queue := New()
	queue.Push(Int(1))
	queue.Push(Int(3))
	queue.Push(Int(2))
	queue.Push(Int(9))
	queue.Push(Int(7))
	queue.Push(Int(0))
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
}