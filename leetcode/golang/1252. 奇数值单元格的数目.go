/*
常规方法：时间：O((n + m)*len(indices))，空间：O(n*m)，就是对每个影响的元素进行add 1操作，然后找出奇数个数
优化：时间：O(len(indices) + n*m)，空间：O(n + m)，统计indices中行、列出现次数，找出所有row[i]+col[j]中的奇数
*/
func oddCells(n int, m int, indices [][]int) int {
	// row, col := make([]int, n), make([]int, m) ----4ms
	row, col := [55]int{}, [55]int{}            //----0ms

	for _, indice := range indices {
		row[indice[0]]++
		col[indice[1]]++
	}

	cnt := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if (row[i] + col[j])%2 == 1 {
				cnt++
			}
		}
	}

	return cnt
}