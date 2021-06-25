package types

type ListNode struct {
	Val  int
	Next *ListNode
}

func Array2List(data []int) (r *ListNode) {
	p := r
	for i := 0; i < len(data); i++ {
		node := &ListNode{
			Val: data[i],
		}
		if p == nil {
			r = node
			p = r
		} else {
			p.Next = node
			p = node
		}
	}
	return
}

func List2Array(r *ListNode) (data []int) {
	p := r
	for p != nil {
		data = append(data, p.Val)
		p = p.Next
	}
	return
}

func ReverseList(r *ListNode) *ListNode {
	//双指针
	if r == nil {
		return nil
	}
	p, pn := r, r.Next
	{ //删除头节点的next
		r.Next = nil
	}

	for pn != nil {
		pnn := pn.Next
		{ //反转next指针
			pn.Next = p
		}
		p, pn = pn, pnn
	}

	return p
}

//RangeSortedList 同时遍历多个升序的listnode
func RangeSortedList(ll_ []*ListNode) (result *ListNode) {
	r := result
	ll := make([]*ListNode, len(ll_))
	copy(ll, ll_)

	for len(ll) > 0 {
		var next int
		var found = false
		for i := 0; i < len(ll); {
			if ll[i] != nil {
				if ll[next].Val >= ll[i].Val {
					next = i
					found = true
				}
				i++
			} else {
				ll = append(ll[:i], ll[i+1:]...)
			}
		}
		if found {
			//fmt.Println("min:ll[",next,"]",ll[next].Val)
			//r=append(r,ll[next].Val)
			if r == nil {
				result = &ListNode{
					Val: ll[next].Val,
				}
				r = result
			} else {
				r.Next = &ListNode{
					Val: ll[next].Val,
				}
				r = r.Next
			}
			ll[next] = ll[next].Next
		} else {
			break
		}
	}
	return
}
