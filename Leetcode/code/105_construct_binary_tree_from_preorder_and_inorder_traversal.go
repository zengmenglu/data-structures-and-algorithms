package main

import "fmt"

type TreeNode struct{
    Val int
    Left *TreeNode
    Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode{
    headIdx,ok := findRoot(preorder, inorder)
    if !ok{
        return nil
    }
    root:=&TreeNode{
        preorder[0],
        buildTree(preorder[1:headIdx+1], inorder[:headIdx]),
        buildTree(preorder[headIdx+1:], inorder[headIdx+1:]),
    }
    return root
}

func findRoot(preorder, inorder []int)(int,bool){
    if len(preorder)==0 || len(preorder) != len(inorder) {
        return 0, false
    }
    rootVal := preorder[0]
    for i := range inorder {
        if inorder[i] == rootVal {
            return i, true
        }
    }
    return 0, false
}

func main(){
    preOrder:=[]int{1,2,4,7,5,3,6}
    inOrder:=[]int{7,4,2,5,1,3,6}
    root:=buildTree(preOrder,inOrder)
    fmt.Println(*root)
}

