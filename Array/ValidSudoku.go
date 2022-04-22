package LeetCode

/*
Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:
* Each row must contain the digits 1-9 without repetition.
* Each column must contain the digits 1-9 without repetition.
* Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
Note:
* A Sudoku board (partially filled) could be valid but is not necessarily solvable.
* Only the filled cells need to be validated according to the mentioned rules.
*/

//func IsValidSudoku(board [][]byte) bool {
//	//Each row must contain the digits 1-9 without repetition.
//	for i := 0; i < len(board); i++ {
//		mch := make(map[byte]struct{})
//		for j := 0; j < len(board[i]); j++ {
//			val := board[i][j]
//			if val != '.' {
//				if _, inmap := mch[val]; inmap {
//					return false
//				} else {
//					mch[val] = struct{}{}
//				}
//			}
//		}
//	}
//
//	// Each column must contain the digits 1-9 without repetition.
//	for i := 0; i < len(board); i++ {
//		mch := make(map[byte]struct{})
//		for j := 0; j < len(board[i]); j++ {
//			val := board[j][i]
//			if val != '.' {
//				if _, inmap := mch[val]; inmap {
//					return false
//				} else {
//					mch[val] = struct{}{}
//				}
//			}
//		}
//	}
//
//	// Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
//	r := 0
//	c := 0
//	for r < 9 && c < 9 {
//		mch := make(map[byte]struct{})
//		for i := r; i < r+3; i++ {
//			for j := c; j < c+3; j++ {
//				val := board[i][j]
//				if val != '.' {
//					if _, inmap := mch[val]; inmap {
//						return false
//					} else {
//						mch[val] = struct{}{}
//					}
//				}
//			}
//		}
//		c += 3
//		if c > 6 {
//			c = 0
//			r += 3
//		}
//	}
//
//	return true
//}

// IsValidSudoku. The best variant.
func IsValidSudoku(board [][]byte) bool {
	var rows, cols, boxes []map[byte]bool
	for i := 0; i < 9; i++ {
		rows = append(rows, make(map[byte]bool))
		cols = append(cols, make(map[byte]bool))
		boxes = append(boxes, make(map[byte]bool))
	}
	for i := range board {
		for j, num := range board[i] {
			if num == '.' {
				continue
			}
			boxId := (i/3)*3 + j/3
			if rows[i][num] || cols[j][num] || boxes[boxId][num] {
				return false
			}
			rows[i][num] = true
			cols[j][num] = true
			boxes[boxId][num] = true
		}
	}
	return true
}

func isValidSudoku(board [][]byte) bool {
	return IsValidSudoku(board)
}
