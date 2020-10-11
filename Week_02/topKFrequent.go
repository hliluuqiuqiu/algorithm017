func topKFrequent(nums []int, k int) []int {
    valCount := map[int]int{}
    for _,val := range nums{
        valCount[val]++
    }

    bucket := make([][]int,len(nums ) + 1)
    for key,val := range valCount{
        bucket[val] = append(bucket[val],key)
    }

    ret := []int{}
    for i,bindex :=0,len(bucket) - 1; i < k && bindex >= 0;bindex--{
        if len(bucket[bindex]) > 0{
             ret = append(ret,bucket[bindex]...)
        }
        
        i += len(bucket[bindex])
    }

    ret = ret[:k]
    return ret

    
}