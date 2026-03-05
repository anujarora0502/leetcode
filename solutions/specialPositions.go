package solutions

func numSpecial(mat [][]int) int {
    count := 0
    for i := 0; i < len(mat); i++ {
        for j := 0; j < len(mat[0]); j++ {
            if mat[i][j] == 1 && isSpecial(mat, i, j) {
                count++
            }
        }
    }

    return count
}

func isSpecial(mat [][]int, row int, col int) bool {
    for i := 0; i < len(mat); i++ {
        if i == row {
            continue
        }
        if mat[i][col] != 0 {
            return false
        }
    }

    for i := 0; i < len(mat[0]); i++ {
        if i == col {
            continue
        }
        if mat[row][i] != 0 {
            return false
        }
    }

    return true
}