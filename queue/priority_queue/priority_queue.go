package priority_queue

import (
	"container/heap"
	"sync"
)

// user should implement this interface
type Interface interface {
	Prior(i interface{}) bool
}

func New() *PriorityQueue {
	p:= &PriorityQueue{data: new(myHeap)}
	return p
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

func (p *PriorityQueue) Push(data interface{}) {
	p.Lock()
	defer p.Unlock()
	heap.Push(p.data, data)
}

func (p *PriorityQueue) Pop() interface{} {
	p.Lock()
	defer p.Unlock()
	return heap.Pop(p.data).(Interface)
}

func (p *PriorityQueue) Size() int {
	p.RLock()
	defer p.RUnlock()
	return p.data.Len()
}

func (p *PriorityQueue) Top() interface{} {
	p.RLock()
	defer p.RUnlock()
	if p.Size() == 0 {
		return nil
	}
	return (*p.data)[0]
}

// implement heap
type myHeap []Interface

func (q *myHeap) Less(i, j int) bool {
	return (*q)[i].Prior((*q)[j])
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
