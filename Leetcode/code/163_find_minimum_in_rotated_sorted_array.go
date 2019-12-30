func findMin(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    
    start:=0
    end :=len(nums)-1
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

func min(a, b int)int{
    if a<b{
        return a
    }
    return b
}
