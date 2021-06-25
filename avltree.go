package types

type AvlNode struct {
	vs          []interface{}
	left, right *AvlNode
	Bf          int
	Score       int
	Parent      *AvlNode
}

func (t *AvlNode) SetValue(v interface{}) {
	t.vs = append(t.vs, v)
}
func (t *AvlNode) Value() (vs []interface{}) {
	return t.vs
}
func (t *AvlNode) Left() (l *AvlNode) {
	return t.left
}
func (t *AvlNode) Right() (l *AvlNode) {
	return t.right
}
func (t *AvlNode) SetLeft(l *AvlNode) {
	t.left = l
}
func (t *AvlNode) SetRight(l *AvlNode) {
	t.right = l
}

func (t *AvlNode) Insert(score int, v interface{}) {
	t.InsertNode(score, v)
	t.Reset() //将指针恢复到整棵树的根节点处
}

func (t *AvlNode) SearchNode(score int) (me *AvlNode) {
	me = t

	for {
		if me == nil {
			break
		}
		if me.Score == score {
			break
		}
		if me.Score < score {
			me = me.Right()
		} else {
			me = me.Left()
		}
	}
	return
}

func (t *AvlNode) Search(score int) []interface{} {
	n := t.SearchNode(score)
	if n != nil {
		return n.vs
	}
	return nil
}

func (t *AvlNode) Reset() {
	tree := t
	for tree.Parent != nil {
		tree = tree.Parent
	}
	t = tree
}

const (
	LH = 1
	EH = 0
	RH = -1
)

func (t *AvlNode) InsertNode(score int, v interface{}) bool {
	if t == nil {
		t = &AvlNode{Score: score, Bf: EH, vs: []interface{}{v}}
		return true
	}
	tree := t
	if score < tree.Score {
		t = tree.left
		if !t.InsertNode(score, v) {
			return false
		} else {
			if t.Parent == nil {
				t.Parent = tree
			}
			if tree.left == nil {
				tree.left = t
			}

			switch tree.Bf {
			case LH:
				t.LeftBalance(tree)
				return false
			case EH:
				tree.Bf = LH
				t = tree
				return true
			case RH:
				tree.Bf = EH
				return false
			}
		}
	} else if score > tree.Score {
		t = tree.right
		if !t.InsertNode(score, v) {
			return false
		} else {
			if t.Parent == nil {
				t.Parent = tree
			}
			if tree.right == nil {
				tree.right = t
			}

			switch tree.Bf {
			case RH:
				t.RightBalance(tree)
				return false
			case EH:
				tree.Bf = RH
				t = tree
				return true
			case LH:
				tree.Bf = EH
				return false
			}
		}
	} else {
		tree.vs = append(tree.vs, v)
	}
	return true
}

func (t *AvlNode) LeftBalance(tree *AvlNode) {
	subTree := tree.left
	if subTree != nil {
		switch subTree.Bf {
		case LH:
			// 新插入节点在左子节点的左子树上要做右单旋处理
			tree.Bf = EH
			subTree.Bf = EH
			t.RightRotate(tree)
		case RH:
			// 新插入节点在左子节点的右子树上要做双旋处理
			subTree_r := subTree.right
			if subTree_r != nil {
				switch subTree_r.Bf {
				case LH:
					tree.Bf = RH
					subTree.Bf = EH
				case RH:
					tree.Bf = EH
					subTree.Bf = LH
				}
				subTree_r.Bf = EH
				t.LeftRotate(subTree)
				t.RightRotate(tree)
			}

		}
	}
}

func (t *AvlNode) RightBalance(tree *AvlNode) {
	subTree := tree.right
	if subTree != nil {
		switch subTree.Bf {
		case RH:
			//新插入节点在右子节点的右子树上要做左单旋处理
			tree.Bf = EH
			subTree.Bf = EH
			t.LeftRotate(tree)
		case LH:
			//新插入节点在右子节点的左子树上要做双旋处理
			subTree_l := subTree.left
			if subTree_l != nil {
				switch subTree_l.Bf {
				case LH:
					tree.Bf = EH
					subTree.Bf = RH
				case RH:
					tree.Bf = LH
					subTree.Bf = EH
				}
				subTree_l.Bf = EH
				t.RightRotate(subTree)
				t.LeftRotate(tree)
			}

		}
	}
}

//右单旋
func (t *AvlNode) RightRotate(tree *AvlNode) {
	subTree := tree.left
	isLeft := false
	if tree.Parent != nil {
		subTree.Parent = tree.Parent //更新新子树的父节点
		if tree.Parent.left == tree {
			isLeft = true
		}
	} else {
		subTree.Parent = nil
	}
	tree.left = subTree.right //原来左节点的右子树挂到老的根节点的左子树
	if subTree.right != nil {
		subTree.right.Parent = tree
	}
	tree.Parent = subTree //原来的左节点变成老的根节点的父节点
	subTree.right = tree  //原来的根节点变成原来左节点的右子树
	tree = subTree
	if tree.Parent == nil { //旋转的是整棵树的根节点
		t = tree
	} else {
		if isLeft { //更新老的子树根节点父节点指针指向新的根节点
			tree.Parent.left = tree
		} else {
			tree.Parent.right = tree
		}
	}
}

//左单旋
func (t *AvlNode) LeftRotate(tree *AvlNode) {
	subTree := tree.right
	isLeft := false
	if tree.Parent != nil {
		subTree.Parent = tree.Parent
		if tree.Parent.left == tree {
			isLeft = true
		}
	} else {
		subTree.Parent = nil
	}
	tree.right = subTree.left
	if subTree.left != nil {
		subTree.left.Parent = tree
	}
	tree.Parent = subTree
	subTree.left = tree
	tree = subTree
	if tree.Parent == nil {
		t = tree
	} else {
		if isLeft {
			tree.Parent.left = tree
		} else {
			tree.Parent.right = tree
		}
	}
}
