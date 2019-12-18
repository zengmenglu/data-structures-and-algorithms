package main

import "fmt"

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

func main() {
    if 5 == firstMissingPositive([]int{1,2,2,1,3,1,0,4,0}) {
        fmt.Println("Right")
    }else{
        fmt.Println("Wrong")
    }
}
