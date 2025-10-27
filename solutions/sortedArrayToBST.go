package solutions

func SortedArrayToBST(nums []int) *TreeNode {

    if len(nums) == 0 {
        return nil
    }
    
    mid := len(nums)/2

    leftTree := SortedArrayToBST(nums[:mid])
    rightTree := SortedArrayToBST(nums[mid+1:])

    
    root := &TreeNode{Val: nums[mid], Left: leftTree, Right: rightTree}

    return root
}