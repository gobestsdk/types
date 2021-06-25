package types

import "testing"

func TestCopyMat(t *testing.T) {
	m := CreateMat(5, 5, " ")
	nm := [][]interface{}{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	CopyMat(m, nm, 1, 1)
	PrintMat(m)
}
