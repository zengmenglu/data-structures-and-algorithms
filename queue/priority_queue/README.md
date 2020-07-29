# 优先级队列

不同于一般的先进先出的对列，优先级队列先出的是优先级最高的元素。

## 原理

基于最小堆实现，堆顶元素是最小节点。

借助go语言```container/heap```来实现堆的功能。

## 应用场景
1. 将很多较小的有序文件归并为一个较大文件（外排）。

扫描小文件的第一个元素，建立一个最小堆，取堆顶元素写入结果文件，然后将该最小值所在的文件的下一个值加入到最小堆，更新堆结构。

2. 操作系统的优先级任务调度策略。

## API

[点击获取使用示例代码](priority_queue_test.go)

1. 创建优先级队列
```go
New() *PriorityQueue 
```

2. 定义优先级

优先级队列中的数据类型需要实现如下接口来定义元素的优先级：
```go
Prior(i interface{}) bool
```

3. 判断优先级队列元素是否为空
```go
Empty() bool
```

4. 插入元素
```go
Push(data interface{})
```

5. 弹出元素
```go
Pop() interface{}
```

6. 获取优先级队列中优先级最高的元素
```go
Top() interface{}
```

7. 获取优先级队列的元素个数
```go
Size() int
```



> 参考：
> * https://github.com/gansidui/priority_queue 