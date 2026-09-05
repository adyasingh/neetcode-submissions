func isValidSudoku(board [][]byte) bool {
	// validate rows 
	for _, row := range board{
		seen := make(map[byte]bool)
		for _, val := range row{
			if seen[val]{
				return false
			}

			if val!='.'{
				seen[val]=true
			}
		}
	}

	// validate columns 
	for i:=0; i<9;i++{
		seen := make(map[byte]bool)
		for _, row := range board{
			if seen[row[i]]{
				return false
			}

			if row[i]!='.'{
				seen[row[i]]=true
			}
		}
	}

	for rowStart:=0; rowStart<6; rowStart+=3{
		for colStart:=0; colStart<6;colStart+=3{
			seen := make(map[byte]bool)
			for i:=rowStart; i<rowStart+3;i++{
				for j:=colStart;j<colStart+3;j++{
					if seen[board[i][j]]{
						return false
					}

					if board[i][j]!='.'{
						seen[board[i][j]]=true
					}

				}
			}
		}
	}
	return true

}
