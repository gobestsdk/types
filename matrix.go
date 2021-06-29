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

var ExpandMatCellReWriteFunc = make(map[string]func(cell string, size int, rst [][]interface{}))

func init() {
	ExpandMatCellReWriteFunc["/"] = func(cell string, expand int, rst [][]interface{}) {
		for i := 0; i < expand; i++ {
			rst[i][expand-i-1] = "/"
		}
	}
	ExpandMatCellReWriteFunc["\\"] = func(cell string, expand int, rst [][]interface{}) {
		for i := 0; i < expand; i++ {
			rst[i][i] = "\\"
		}
	}

}
func ExpandMatCell(cell interface{}, size int) (rst [][]interface{}) {
	expand := size*2 - 1
	rst = CreateMat(expand, expand, " ")
	rst[size-1][size-1] = cell
	c := fmt.Sprint(cell)
	if f, have := ExpandMatCellReWriteFunc[c]; have {
		f(c, expand, rst)
	} else {
		l, r := 0, 0
		for i := 0; i < len(c)-1; i++ {
			if l > r {
				if size-1+r < expand {
					r++
					rst[size-1][size-1+r] = ""
				}

			} else {
				if size-1-l > 0 {
					l++
					rst[size-1][size-1-l] = ""
				}

			}
		}
	}
	return
}
func ExpandMat(mat [][]interface{}, size int) (rst [][]interface{}) {
	w := len(mat)
	var h = 0
	if w > 0 {
		h = len(mat[0])
	}
	expand := size*2 - 1

	rst = CreateMat(w*expand, h*expand, " ")
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			extcell := ExpandMatCell(mat[i][j], size)
			CopyMat(rst, extcell, i*expand, j*expand)
		}
	}
	return
}
