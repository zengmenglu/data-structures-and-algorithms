
# 42 缺失的第一个正数

给定一个未排序的整数数组，找出其中没有出现的最小的正整数。[题目LINK](https://leetcode-cn.com/problems/first-missing-positive/) 

与剑指 offer第三题类似，不同之处在于可能出现负数或大于数组长度n的数。对于这样的数跳过即可，对其他1～n的数换到下标0～n-1的位置。如果数组内数据为[1,2,...,n], 那么第一个未出现的正数是n+1，所以第一个未出现的正数范围是1～n+1。

*  nums[i] == i+1 符合置换条件，不需要交换。
*  nums[i] == nums[nums[i]-1] 两个交换值相等的情况下，会产生死循环，所以也无需交换，继续下一个。

```
func firstMissingPositive(nums []int) int {
    n := len(nums)
    for i := 0; i < n; {
        for (i< n) && (nums[i] <= 0 || nums[i] > n || nums[i] == i+1 || nums[i] == nums[nums[i]-1]){
            i++
        }
        if i< n {
            nums[i], nums[nums[i]-1] =  nums[nums[i]-1],nums[i]       
        }
    }
    for i:=0;i<n;i++{
        if nums[i]!=i+1{
            return i+1
        }
    }
    return n+1
}
```

# 46. 全排列
给定一个没有重复数字的序列，返回其所有可能的全排列。[题目LINK](https://leetcode-cn.com/problems/permutations/) TAG：【递归】【回溯】

假设数字序列为```[0...n-1]```,首先是确定第一个数，将序号为0～n-1的数分别与第一个数交换位置。交换完了之后要恢复数组原貌，再确定第二个数，将第二个数后面所有数字和第二个数交换。如此，直到确定了所有数之后就得到了一个结果。

* 注意这里储存结果的二维数组必须得使用指针，否则结果传不出来。
* 如果不恢复原貌，得到的结果会有重复，并且得到的结果不全。
* 交换一定要从start开始，而不是start+1，

```
func permute(nums []int) [][]int {
	var results [][]int
	backTrace(nums, &results,0)
	return results
}

func backTrace(nums []int, results *[][]int, start int){
	if start == len(nums){
		*results = append(*results,append([]int{},nums...))
	}
	for i := start; i < len(nums); i++{
		nums[start],nums[i] = nums[i],nums[start] // 交换
		backTrace(nums, results, start+1) // 计算下一个数
		nums[start],nums[i] = nums[i],nums[start] // 恢复原貌
	}
}
```

# 96. 不同的二叉搜索树
给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？[题目LINK](https://leetcode-cn.com/problems/unique-binary-search-trees/) TAG：【动态规划】

本题是一维的动态规划算法，构造```result[i]```代表```1～i```组成的二叉搜索树的种类。可以知道m~n的二叉搜索树的种类等于1～n-m+1为节点组成的二叉搜索树。又根据二叉搜索树的性质，root左边的数小于root，右边的数大于root。因此，1～i为节点的二叉搜索树的种类是以1～i分别为root时左右两边二叉树种类的乘积。

```
func numTrees(n int) int {
    var result = make([]int, n+1)
    result[0] = 1
    result[1] = 1
    for i:=2;i<=n;i++{  
        for r:=1;r<=i;r++{// 以r为root
                result[i] += result[r-1]*result[i-r]
        }
    }
    return result[n]
}
```

# 153. 寻找旋转排序数组中的最小值
假设按照升序排序的数组在预先未知的某个点上进行了旋转。( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。请找出其中最小的元素。你可以假设数组中不存在重复元素。[题目LINK(也是剑指Offer #11)](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/)

旋转数组前半部分所有数字值大于后半部分，需要找到旋转点。所以判断一截数组，如果第一个元素大于最后一个元素，说明这截数组仍然是旋转数组。

* 要判断```end-start==1```的情况，应为当这种情况下，```mid=(start+end)/2```会出现```mid==start```，从而造成死循环。
* 本题假设没有重复数字，当有重复数字时，会有nums[start] == nums[end]==nums[mid]的情况，这个时候需要顺序查找最小值。

```
func findMin(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    start, end :=0, len(nums)-1
    for nums[start] > nums[end]{
        if end -start == 1{
            return nums[end]
        }
        mid := (start + end)/2
        if nums[start] >= nums[mid] {
            end = mid
        } else if nums[start] <= nums[mid]{
            start = mid
        }
    }
    return nums[start]
}
```

# 403. 青蛙过河

给定石子的位置列表（用单元格序号升序表示）， 请判定青蛙能否成功过河（即能否在最后一步跳至最后一个石子上）。 开始时， 青蛙默认已站在第一个石子上，并可以假定它第一步只能跳跃一个单位（即只能从单元格1跳至单元格2）。如果青蛙上一步跳跃了 k 个单位，那么它接下来的跳跃距离只能选择为 k - 1、k 或 k + 1个单位。 另请注意，青蛙只能向前方（终点的方向）跳跃。实例：输入[0,1,3,5,6,8,12,17]，输出true。[题目LINK](https://leetcode-cn.com/problems/frog-jump/) TAG：【动态规划】

本题涉及二维的动态规划

* dp[i][j] 表示在stones[i]上可以跳跃j步。
* 当计算到dp[i][j]， i=n-1且存在为true的j时，说明前一个节点也可以跳j步到达这最后一个石头上。

```
func canCross(stones []int) bool {
	var dp = make([][]bool, len(stones)) // 在stones[i]上可以跳跃j步
	for i := range dp {
		dp[i] = make([]bool, len(stones))
	}
	dp[0][1] = true
	for i := 1; i < len(stones); i++ {
		for j := 0; j < i; j++ {
			dis := stones[i] - stones[j]
			if dis <= 0 || dis >= len(stones) || !dp[j][dis] {continue}// dis的范围判断是为了dp[j][dis]下标不越界
			if i == len(stones)-1 {return true} // 成功判定
			dp[i][dis] = true // 更新dp数组
			if dis-1 > 0 {	dp[i][dis-1] = true}
			if dis+1 < len(stones) {dp[i][dis+1] = true}
		}
	}
	return false
}
```


# 912. 排序数组
给定一个整数数组 nums，将该数组升序排列。[题目LINK](https://leetcode-cn.com/problems/sort-an-array/)
TAG：【排序】【数组】

快速排序:

```
func quickSort(nums []int){
	partition(nums, 0, len(nums)-1)
}
func partition(nums []int, start, end int){
	if start >= end{
		return
	}
	var i, j  = start, end
	pivot:=nums[start]
	for i < j{
		for ;i<j;j--{  //从尾开始，将不小于pivot的值放在左边
			if nums[j]<=pivot{
				nums[i] = nums[j]
				i++
				break
			}
		}
		for ;i<j;i++{ // 从首开始
			if nums[i]>=pivot{
				nums[j] = nums[i]
				j--
				break
			}
		}
	}
	nums[i] = pivot // 置位
	partition(nums, start, i-1) // 分治
	partition(nums, i+1, end)
}
```
堆排序：

```
func heapSort(nums []int){
	for i:=len(nums)/2-1;i>=0;i--{ // 从最后一个非叶子结点开始调整，构造一个完成最大堆
		heaplify(nums, i, len(nums)-1)
	}
	for j:=len(nums)-1;j>=1;j--{// 将堆定最大元素取出，将堆最后一个叶子结点放在堆定，再调整
		nums[0],nums[j] = nums[j],nums[0]
		heaplify(nums,0,j-1)
	}
}
func heaplify(nums []int, start, end int){// 堆的调整
	for i := start; i<=end; {
		left:=2*i+1
		right:=2*i+2
		if left<=end && nums[left]>nums[i] && (right > end || nums[right]<=nums[left]){
			nums[left],nums[i] = nums[i],nums[left]
			i=left
		}else if right <= end && nums[right]>nums[i]{
			nums[right],nums[i] = nums[i],nums[right]
			i=right
		}else{
			break
		}
	}
}
```

# 993. 二叉树的堂兄弟节点
在二叉树中，根节点位于深度 0 处，每个深度为 k 的节点的子节点位于深度 k+1 处。如果二叉树的两个节点深度相同，但父节点不同，则它们是一对堂兄弟节点。我们给出了具有唯一值的二叉树的根节点 root，以及树中两个不同节点的值 x 和 y。只有与值 x 和 y 对应的节点是堂兄弟节点时，才返回 true。否则，返回 false。[题目LINK](https://leetcode-cn.com/problems/cousins-in-binary-tree/)TAG：【树】

递归分别求x和y的深度和父节点的值。

```
func isCousins(root *TreeNode, x int, y int) bool {
    var xParent, yParent, xdepth, ydepth int
    xflg:=isChild(root,x,&xParent,&xdepth)
    yflg:=isChild(root, y,&yParent,&ydepth)
    if !xflg || !yflg || xParent == yParent || xdepth!=ydepth {
        return false
    }
    return true
}

func isChild(root *TreeNode, val int, parent *int, depth *int) (bool){
    if root == nil { return false}
    if (root.Left != nil && root.Left.Val== val) || (root.Right != nil && root.Right.Val== val){
        *parent = root.Val
        *depth+=1
        return true
    }
    if isChild(root.Left, val, parent, depth) || isChild(root.Right, val, parent,depth){
        *depth+=1
        return true
    }
    return false
}
```



# 1277. 统计全为 1 的正方形子矩阵

给你一个 m * n 的矩阵，矩阵中的元素不是 0 就是 1，请你统计并返回其中完全由 1 组成的 正方形 子矩阵的个数。

```
输入：matrix =
[
  [0,1,1,1],
  [1,1,1,1],
  [0,1,1,1]
]
输出：15
解释： 
边长为 1 的正方形有 10 个。
边长为 2 的正方形有 4 个。
边长为 3 的正方形有 1 个。
正方形的总数 = 10 + 4 + 1 = 15.
```

[题目LINK](https://leetcode-cn.com/problems/count-square-submatrices-with-all-ones/)

按照边长1～min(row,col)的小正方形的情况，暴力求解。

# 1278. 分割回文串 III

给你一个由小写字母组成的字符串 s，和一个整数 k。

请你按下面的要求分割字符串：

首先，你可以将 s 中的部分字符修改为其他的小写英文字母。
接着，你需要把 s 分割成 k 个非空且不相交的子串，并且每个子串都是回文串。
请返回以这种方式分割字符串所需修改的最少字符数。

```
输入：s = "abc", k = 2
输出：1
解释：你可以把字符串分割成 "ab" 和 "c"，并修改 "ab" 中的 1 个字符，将它变成回文串。
```

[题目LINK](https://leetcode-cn.com/problems/palindrome-partitioning-iii/)

```
dp[i][j] = min(dp[i][j],dp[m][j-1]+cost[m+1][i]), m=1...i-1
```

```dp[i][j]``` 代表字符串s前i个字符分成j个子串需要修改的最少字符数，```cost[m][n]```代表字符串第m到n个字符的子串变成回文需要修改的最小字符数。（下标均为从1开始）

```
func initCost(s string) [][]int { // cost[i][j]表示第i～j字符段成为回文需要变化的次数，下标0开始
	var cost = make([][]int, len(s))
	for i := 0; i < len(s); i++{
		cost[i] = make([]int, len(s))
		for j := i+1; j < len(s); j++{
			for m,n := i,j; m < n ; m,n = m+1, n-1 {
				if s[m] != s[n]{
					cost[i][j]++
				}
			}
		}
	}
	return cost
}

func palindromePartition(s string, k int) int {
	var cost = initCost(s)
	var dp = make([][]int, len(s)+1) // dp[i][j]表示0～（i-1）前i个字符段分成k份，每份为回文的次数
	for i := 1; i <= len(s); i++ {
		dp[i] = make([]int, k+1)
		dp[i][1]=cost[0][i-1]
		for j := 2; j <= k && j < i; j++ {
			dp[i][j] = len(s)
			for m := 1; m < i; m++{
				dp[i][j] = min(dp[i][j],dp[m][j-1]+cost[m][i-1])
			}
		}
	}
	return dp[len(s)][k]
}

func min(a, b int) int{
	if a < b {
		return a
	}
	return b
}
```