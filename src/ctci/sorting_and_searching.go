package ctci

func SortedMerge(arrA, arrB []int) []int {
	outArr := make([]int, len(arrA)+len(arrB))
	outIx := len(outArr) - 1
	aIx := len(arrA) - 1
	bIx := len(arrB) - 1

	for bIx >= 0 {
		if aIx >= 0 && arrA[aIx] > arrB[bIx] {
			outArr[outIx] = arrA[aIx]
			aIx--
		} else {
			outArr[outIx] = arrB[bIx]
			bIx--
		}
		outIx--
	}

	if aIx >= 0 {
		copy(outArr[:aIx+1], arrA[:aIx+1])
	}

	return outArr
}
