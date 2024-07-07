package main

func ValidSudoku(board [][]byte) bool {
	// 9 x 9 board
	// CHECK: ACROSS, DOWN, SQUARE
	// HASHMAP FOR EACH?
	// if board[i][j] == '.' continue

	rows := make(map[int]map[byte]bool)
	cols := make(map[int]map[byte]bool)
	squares := make(map[int]map[byte]bool)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			if _, exists := rows[i]; !exists {
				rows[i] = make(map[byte]bool)
			}
			if _, exists := rows[i][board[i][j]]; exists {
				return false
			}

			if _, exists := cols[j]; !exists {
				cols[j] = make(map[byte]bool)
			}

			if _, exists := cols[j][board[i][j]]; exists {
				return false
			}

			square := ((i / 3) * 3) + (j / 3)

			if _, exists := squares[square]; !exists {
				squares[square] = make(map[byte]bool)
			}
			if _, exists := squares[square][board[i][j]]; exists {
				return false
			}

			rows[i][board[i][j]] = true
			cols[j][board[i][j]] = true
			squares[square][board[i][j]] = true
		}
	}
	return true
}
