package main

import "fmt"

func numTrees(n int) int {
	var record = make([][]int, n+2)
	for i := range record {
		record[i] = make([]int, n+2)
		record[i][i] = 1
	}

	for i := n; i >= 1; i-- {
		for j := n; j > i; j-- {
			for k := i; k <= j; k++ {
				var left, right int
				left = record[i][k-1]
				right = record[k+1][j]
				if left != 0 && right != 0 {
					record[i][j] += left * right
				} else {
					record[i][j] += left + right
				}
			}
		}
	}

	return record[1][n]
}

func main() {
	fmt.Println(numTrees(3))
}
