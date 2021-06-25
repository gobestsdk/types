package types

type BinaryNode interface {
	SetValue(v interface{})
	Value() (v interface{})
	Left() (l BinaryNode)
	Right() (l BinaryNode)
	IsNil() bool
}

func PrintBinary(b BinaryNode) (pt [][]interface{}) {
	deep := Deepest(b)

	hei := deep * (deep + 1) / 2
	wid := hei*2 - 1
	pt = CreateMat(hei, wid, " ")
	pt[0][wid/2] = b.Value()
	if !b.Left().IsNil() {
		for i := 1; i < deep; i++ {
			pt[i][wid/2-i] = "/"
		}
	}
	if !b.Right().IsNil() {
		for i := 1; i < deep; i++ {
			pt[i][wid/2+i] = "\\"
		}
	}

	if deep > 2 {
		if !b.Left().IsNil() {
			left := PrintBinary(b.Left())
			CopyMat(pt, left, deep, 0)
		}
		if !b.Right().IsNil() {
			right := PrintBinary(b.Right())
			CopyMat(pt, right, deep, hei)
		}
	} else {
		if !b.Left().IsNil() {
			left := PrintBinary(b.Left())
			CopyMat(pt, left, deep, 0)
		}
		if !b.Right().IsNil() {
			right := PrintBinary(b.Right())
			CopyMat(pt, right, deep, wid-1)
		}
	}

	return
}

//MidOrderTraverse 中序遍历
func MidOrderTraverse(t BinaryNode) (result []interface{}) {
	result = append(result, MidOrderTraverse(t.Left())...)
	result = append(result, t.Value())
	result = append(result, MidOrderTraverse(t.Right())...)
	return
}

//MidOrderTraverse 中序遍历
func Deepest(t BinaryNode) int {
	if t.IsNil() {
		return 0
	}

	max := Deepest(t.Left())
	r := Deepest(t.Right())
	if r > max {
		return 1 + r
	}
	return 1 + max
}

//MidOrderTraverse 先序遍历
func PreOrderTraverse(t BinaryNode) (result []interface{}) {
	result = append(result, t.Value())
	result = append(result, MidOrderTraverse(t.Left())...)
	result = append(result, MidOrderTraverse(t.Right())...)
	return
}

//PostOrderTraverse 后序遍历
func PostOrderTraverse(t BinaryNode) (result []interface{}) {
	result = append(result, MidOrderTraverse(t.Left())...)
	result = append(result, MidOrderTraverse(t.Right())...)
	result = append(result, t.Value())
	return
}
