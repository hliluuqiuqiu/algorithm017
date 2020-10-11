/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    ret := []int{}
    getVal(root,&ret)
    return ret
}

func getVal(root *TreeNode,ret *[]int){
    if root == nil{
        return
    }

    *ret = append(*ret,root.Val)
    getVal(root.Left,ret)
    getVal(root.Right,ret)
}