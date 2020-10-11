/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
    ret := []int{}
    getVal(root,&ret)
    return ret
}


func getVal(root *Node,ret *[]int){
    if root == nil{
        return
    }

    *ret = append(*ret,root.Val)
    for _,node:=range root.Children{
        getVal(node,ret)
    }
  
}