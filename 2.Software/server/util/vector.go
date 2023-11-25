package util

// Matrix logic

func MtxAdd(mtxA [][]int, mtxB [][]int) [][]int {
	for i, colA := range mtxA {
		for j, numA := range colA {
			mtxB[i][j] += numA
		}
	}
	return mtxB
}

func MtxSub(mtxA [][]int, mtxB [][]int) [][]int {
	for i, colA := range mtxA {
		for j, numA := range colA {
			mtxB[i][j] -= numA
		}
	}
	return mtxB
}

func MtxMlt(mtxA [][]int, mtxB [][]int) [][]int {
	for i, colA := range mtxA {
		for j, numA := range colA {
			mtxB[i][j] *= numA
		}
	}
	return mtxB
}

func MtxDiv(mtxA [][]int, mtxB [][]int) [][]int {
	for i, colA := range mtxA {
		for j, numA := range colA {
			mtxB[i][j] /= numA
		}
	}
	return mtxB
}


// unfinnished
func MtxRot(mtx [][]int) {
	
}

func MtxDot(mtxA [][]int, mtxB [][]int) [][]int {
	dot := make([][]int, len(mtxB[0]))
	for _, d := range dot {
		copy(d, make([]int, len(mtxA[0])))
	}
	for i, colA := range mtxA {
		for j, numA := range colA {
			mtxB[i][j] *= numA
		}
	}
	return mtxB
}