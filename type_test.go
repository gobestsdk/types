package types

import (
	"fmt"
	"testing"
)

func Foo(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}
func TestNil(t *testing.T) {
	var x *int = nil
	Foo(x)

	fmt.Println(*x)
	var y interface{} = nil
	Foo(y)

}
