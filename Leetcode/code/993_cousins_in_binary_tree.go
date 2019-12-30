package main

import "fmt"


/* Definition for a binary tree node.*/
type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
    var xParent, yParent *TreeNode
    var xdepth, ydepth int
    if !isChild(root,xParent,x,&xdepth) || !isChild(root, yParent,y,&ydepth){
        return false
    }
    if xdepth!=ydepth || xParent == nil || yParent == nil || xParent.Val == yParent.Val{
        return false
    }
    return true
}

func isChild(root, parent *TreeNode, val int, depth *int) (bool){
    if root == nil{
        return false
    }
    
    if (root.Left != nil && root.Left.Val== val) || (root.Right != nil && root.Right.Val== val){
        parent = root
        *depth+=1
        return true
    }
    
    if isChild(root.Left, parent, val, depth) || isChild(root.Right,parent, val, depth){
        *depth+=1
        return true
    }
    return false
}

func main(){
    root:=&TreeNode{
        1,
        &TreeNode{2,nil,&TreeNode{4,nil,nil}},
        &TreeNode{3,nil,&TreeNode{5,nil,nil}},
    }
    if !isCousins(root,5,4){
       fmt.Println("false")
    }
}
