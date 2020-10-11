/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
    ret := [][]int{}
    if root == nil{
        return ret
    }
    quene := []*Node{root}
    for len(quene) != 0{
        levelCount := len(quene)
        levelVals := []int{}
        for indx := 0; indx < levelCount;indx++{
            levelVals = append(levelVals,quene[indx].Val)
            for _,node := range quene[indx].Children{
                quene = append(quene,node)
            }
        }
        ret = append(ret,levelVals)
        quene = quene[levelCount:]

    }
    return ret
}