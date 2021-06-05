package types

import (
	"fmt"
	"testing"
)

type student struct {
	Name string
	Age  int
}

func TestRangeMap(t *testing.T) {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu //这里对&stu，全都只会指向最后一个对象
	}
	fmt.Printf("%v", m)
}
