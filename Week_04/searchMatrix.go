func searchMatrix(matrix [][]int, target int) bool {
   
   if len(matrix) == 0 {
       return false
   }

   w := len((matrix[0]))
   h := len(matrix)
   
   for i, j:= 0,w - 1; j >= 0 && i < h;{
       if matrix[i][j] == target{
           return true
       }

       if matrix[i][j] < target{
           i++
       }else{
           j--
       }
   }

   return false
}