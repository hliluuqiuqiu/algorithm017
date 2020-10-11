/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    ret := []int{}
    getVal(root,&ret)
    return ret
}

func getVal(root *TreeNode,ret *[]int){
    if root == nil{
        return
    }
    getVal(root.Left,ret)
    *ret = append(*ret,root.Val)
    getVal(root.Right,ret)
}