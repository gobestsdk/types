package types

import (
	"fmt"
	"testing"
)

type Btree struct {
	vs          interface{}
	left, right *Btree
}

func (b *Btree) SetValue(v interface{}) {
	b.vs = v
}
func (b *Btree) Value() (vs interface{}) {
	return b.vs
}
func (b *Btree) Left() (l BinaryNode) {
	if b == nil {
		return nil
	}
	return b.left
}
func (b *Btree) IsNil() bool {
	return b == nil
}
func (b *Btree) Right() (l BinaryNode) {
	if b == nil {
		return nil
	}
	return b.right
}

func TestDeepest(t *testing.T) {
	b := &Btree{
		vs: 1,
		left: &Btree{
			vs: 2,
		},
		right: &Btree{
			vs: 3,
		},
	}
	fmt.Println(Deepest(b))

}
func BenchmarkDeepest(b *testing.B) {

}
func TestPrintBinary(t *testing.T) {
	b := &Btree{
		vs: []interface{}{1},
	}
	PrintMat(PrintBinary(b))
}
func TestPrint(t *testing.T) {
	b := &Btree{
		vs: 1,
		left: &Btree{
			vs: 2,
		},
		right: &Btree{
			vs: 3,
		},
	}
	PrintMat(PrintBinary(b))
	c := &Btree{
		vs: 1,
		left: &Btree{
			vs: 2,
			left: &Btree{
				vs: 4,
			},
		},
		right: &Btree{
			vs: 3,
			left: &Btree{
				vs: 5,
			},
		},
	}
	PrintMat(PrintBinary(c))
}
