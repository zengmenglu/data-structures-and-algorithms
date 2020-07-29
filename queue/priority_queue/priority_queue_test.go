package priority_queue

import (
	"fmt"
	"testing"
)

type  Int int

func(i Int) Prior(j interface{})bool{
	return i>j.(Int) // 数值大的优先级高
}

func TestPriorityQueue(t *testing.T) {
	queue := New()
	queue.Push(Int(1))
	queue.Push(Int(3))
	queue.Push(Int(2))
	fmt.Println("top:", queue.Top())
	queue.Push(Int(9))
	queue.Push(Int(7))
	queue.Push(Int(0))
	fmt.Println("top:", queue.Top())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
}