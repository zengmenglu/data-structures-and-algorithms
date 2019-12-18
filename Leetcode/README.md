[toc]

# 42 缺失的第一个正数

给定一个未排序的整数数组，找出其中没有出现的最小的正整数。[题目LINK Leetcode 42](https://leetcode-cn.com/problems/first-missing-positive/) 

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