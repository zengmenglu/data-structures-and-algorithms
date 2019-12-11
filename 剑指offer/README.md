# 3 找出重复数

给定一个包含 n + 1 个整数的数组 nums，其数字都在 1 到 n 之间（包括 1 和 n），可知至少存在一个重复的整数。假设只有一个重复的整数，找出这个重复的数。

```
示例 1:

输入: [1,3,4,2,2]
输出: 2

示例 2:

输入: [3,1,3,4,2]
输出: 3
```

题目链接 [Leetcode 287](https://leetcode-cn.com/problems/find-the-duplicate-number/)

* 本题改变数组的方法为，将数值a放在下标为a的位置上，直到某个下标有多个值要放在其上。
* 若不允许改变数组，有两种方法：
	* 二分法：取数组值的中位数mid，遍历多遍数组，统计数据<=mid的值的是不是超过n/2.如果是则继续查找。（存在bug，小于和大于mid的值个数均为为n/2，那么就需要继续查找），时间复杂度```o(nlgn)```
	* 快慢指针：思路同寻找单向链表中环的起点。假设这个数组是一个链表，数组中的每个值代表指向的下一个元素的idx。如果存在重复的数，则说明，有多个节点会指向同一个元素，即链表存在环，且环的入口点就是那个重复的元素。
		
		第一步：快指针走两步，慢指针走一步，快慢指针最后会相遇在环上一点。假设快慢指针都已经走在环上了，它们之间最大的距离是环的长度L，在环上后相遇的最大时间为t = L/速度差 = L。所以慢指针最多走1*t = L 一圈，就会于快指针相遇。此时快指针走了2S = a + nL + x， 慢指针S = a + x. 所以a+nL+x = 2(a+x),  a+x = nL

		第二步：快指针回到起点，慢指针停在第一次相遇点。快慢指针同时一次只走一步，假设快指针走了a步，到达环的起点时，慢指针在环上的位置是 x + a = nL，即慢指针也到了环的起点。
		![](https://github.com/zengmenglu/data-structures-and-algorithms/blob/master/%E5%89%91%E6%8C%87offer/pic/3_%E9%93%BE%E8%A1%A8%E7%8E%AF.png)

```
// 快慢指针法
// 假设nums中的值都小于len(nums), 且nums不为空的情况， 不考虑输入不合法

func findDuplicate(nums []int) int {
	fast, slow := nums[0], nums[0]
	for {
		fast = nums[nums[fast]]
		slow = nums[nums[slow]]
		if fast == slow {break}
	}
	
	fast = nums[0]
	for slow != fast{
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
```