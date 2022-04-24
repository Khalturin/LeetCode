package LeetCode

import "fmt"

/*
Rotate
You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).

You have to rotate the image in-place, which means you have to modify the input 2D matrix directly.
DO NOT allocate another 2D matrix and do the rotation.
*/
// 1 2 3	7 4 1
// 4 5 6	8 5 2
// 7 8 9	9 6 3
func Rotate(matrix [][]int) [][]int {
	for i := range matrix {
		for j := i + 1; j < len(matrix[i]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
		reversArr(matrix[i], 0, len(matrix)-1)
	}
	return matrix
}

func rotate(matrix [][]int) {
	fmt.Println(Rotate(matrix))
}
