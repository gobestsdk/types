package types

import "fmt"

func CreateMat(hei, wid int, defaultv interface{}) (pt [][]interface{}) {
	for i := 0; i < hei; i++ {
		w := []interface{}{}
		for j := 0; j < wid; j++ {
			w = append(w, defaultv)
		}
		pt = append(pt, w)
	}
	return
}
func PrintMat(pt [][]interface{}) {
	for i := 0; i < len(pt); i++ {
		for j := 0; j < len(pt[i]); j++ {
			fmt.Print(pt[i][j])
		}
		fmt.Println()
	}
	return
}
func CopyMat(mat, updater [][]interface{}, startx, starty int) {
	for i := 0; i < len(updater); i++ {
		for j := 0; j < len(updater[i]); j++ {
			mat[startx+i][starty+j] = updater[i][j]
		}
	}
}
