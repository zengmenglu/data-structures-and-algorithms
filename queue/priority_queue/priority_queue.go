package priority_queue

import (
	"container/heap"
	"sync"
)

// user should implement this interface
type Interface interface {
	Less(i Interface) bool
}

func New() *PriorityQueue {
	return &PriorityQueue{data: new(myHeap)}
}

type PriorityQueue struct {
	sync.RWMutex
	data *myHeap
}

// API
func (p *PriorityQueue) Empty() bool {
	p.RLock()
	defer p.RUnlock()
	if p.data.Len() == 0 {
		return true
	}
	return false
}

func (p *PriorityQueue) Push(data Interface) {
	p.Lock()
	defer p.Unlock()
	heap.Push(p.data, data)
}

func (p *PriorityQueue) Pop() Interface {
	p.Lock()
	defer p.Unlock()
	return p.data.Pop().(Interface)
}

func (p *PriorityQueue) Size() int {
	p.RLock()
	defer p.RUnlock()
	return p.data.Len()
}

func (p *PriorityQueue) Top() Interface {
	p.RLock()
	defer p.RUnlock()
	if p.Size() == 0 {
		return nil
	}
	return (*p.data)[p.Size()-1]
}

// implement heap
type myHeap []Interface

func (q *myHeap) Less(i, j int) bool {
	return (*q)[i].Less((*q)[j])
}

func (q *myHeap) Swap(i, j int) {
	(*q)[i], (*q)[j] = (*q)[j], (*q)[i]
}

func (q *myHeap) Len() int {
	return len(*q)
}

func (q *myHeap) Pop() (v interface{}) {
	if q.Len() == 0 {
		return nil
	}
	*q, v = (*q)[:q.Len()-1], (*q)[q.Len()-1]
	return v
}

func (q *myHeap) Push(data interface{}) {
	*q = append(*q, data.(Interface))
}
