package rbtree

import (
	"fmt"
	"rbtree/node"
)

//红黑树就是指向了一个根结点
type RBTree struct {
	NullStr string
	Header *node.RBNode
}

//使用初始值初始化一个红黑数,必须要设置一个值作为空值，保证自己数据不会存在的值
func NewRBTree(key string, value interface{}, nullStr string)RBTree{
	header := node.NewRBNode(key,value,true,nil,nil, nil)
	left := node.NewRBNode(nullStr,nil,true,header,nil,nil)
	right := node.NewRBNode(nullStr, nil, true, header, nil, nil)
	header.Left = left
	header.Right = right
	rbTree := RBTree{
		NullStr: nullStr,
		Header:header,
	}
	return rbTree
}

//先序遍历
func (rb *RBTree)PreTraversal(T *node.RBNode){
	if T==nil{
		return
	}
	fmt.Println(T.Key, T.Value, T.Color)
	rb.PreTraversal(T.Left)
	rb.PreTraversal(T.Right)
}

//中序遍历
func (rb *RBTree)MiddleTraversal(T *node.RBNode){
	if T==nil{
		return
	}
	rb.MiddleTraversal(T.Left)
	fmt.Println(T.Key, T.Value, T.Color)
	rb.MiddleTraversal(T.Right)
}

//单左旋
func (rb *RBTree)singleLeftRotate(T *node.RBNode){
	father := T.Father
	a := T.Left
	T.Father = a
	T.Left = a.Right
	a.Right = T
	a.Father = father
	if father==nil{
		rb.Header = a
		return
	}
	if father.Left == T{
		father.Left = a
	}
	if father.Right == T{
		father.Right = a
	}
}
//双左旋 右左
func (rb *RBTree)doubleLeftRotate(T *node.RBNode){
	father := T.Father
	a := T.Right
	b := a.Left
	T.Right = b.Left
	T.Father = b
	a.Left = b.Right
	a.Father = b
	b.Left = T
	b.Right = a
	b.Father = father
	if father == nil{
		rb.Header = b
		return
	}
	if father.Right == T{
		father.Right = b
	}
	if father.Left == T{
		father.Left = b
	}
}
//单右旋
func (rb *RBTree)singleRightRotate(T *node.RBNode){
	father := T.Father
	a := T.Right
	T.Right = a.Left
	T.Father = a
	a.Left = T
	a.Father = father
	if father == nil{
		rb.Header = a
		return
	}
	if father.Left == T{
		father.Left = a
	}
	if father.Right == T{
		father.Right = a
	}
}

//双右旋 左右
func (rb *RBTree)doubleRightRotate(T *node.RBNode){
	father := T.Father
	a := T.Left
	b := a.Right
	T.Left = b.Right
	T.Father = b
	a.Right = b.Left
	a.Father = b
	b.Left = a
	b.Right = T
	b.Father = father
	if father == nil{
		rb.Header = b
		return
	}
	if father.Right == T{
		father.Right = b
	}
	if father.Left == T{
		father.Left = b
	}
}

func (rb *RBTree)insertFixUp(T *node.RBNode) {
	if T==nil{
		return
	}
	father := T.Father
	//根节点不用修复
	if father == nil {
		rb.Header = T
		T.Color = true
		return
	}
	//左左红
	if T.Color == false {
		if rb.Header == T{
			T.Color = true
			return
		}
		if T.Left.Color == false{
			//左左
			if father.Left == T {
				uncle := father.Right
				if uncle.Color == true {
					T.Color = true
					father.Color = false
					rb.singleLeftRotate(father)
					return
				} else {
					T.Left.Color = true
					father.Color = true
					rb.singleLeftRotate(father)
					//旋转后T变为了局部根结点,需要向上递归
					rb.insertFixUp(T.Father)
					return
				}
			}
			//右左
			if father.Right == T{
				uncle := father.Left
				if uncle.Color == true{
					T.Left.Color = true
					father.Color = false
					rb.doubleLeftRotate(father)
					return
				}else{
					l := T.Left
					T.Color = true
					rb.doubleLeftRotate(father)
					//旋转后l变为了局部根节点,需要向上递归修复
					rb.insertFixUp(l.Father)
					return
				}
			}
		}
		if T.Right.Color == false{
			//右右
			if father.Right == T{
				uncle := father.Left
				if uncle.Color == true{
					T.Color = true
					father.Color = false
					rb.singleRightRotate(father)
					return
				}else{
					T.Right.Color = true
					rb.singleRightRotate(father)
					rb.insertFixUp(T.Father)
					return
				}
			}
			//左右
			if father.Left == T{
				uncle := father.Right
				if uncle.Color == true{
					T.Right.Color = true
					father.Color = false
					rb.doubleRightRotate(father)
					return
				}else{
					r := T.Right
					T.Color = true
					rb.doubleRightRotate(father)
					rb.insertFixUp(r.Father)
					return
				}
			}
		}
	}

}

func (rb *RBTree)IsEmpty()bool{
	if rb ==nil || rb.Header==nil || rb.Header.Key == rb.NullStr{
		return true
	}
	return false
}

func (rb *RBTree)IsEmptyNode(T *node.RBNode)bool{
	if T==nil||T.Key==rb.NullStr{
		return true
	}
	return false
}

//插入
func (rb *RBTree)Insert(T *node.RBNode, key string, value interface{}){
	//如果rb是空
	if rb==nil{
		*rb = NewRBTree(key, value, rb.NullStr)
		return
	}
	//如果rb.Header是空
	if rb.Header==nil{
		*rb = NewRBTree(key, value, rb.NullStr)
		return
	}
	//插入已存在的值
	if T.Key == key{
		T.Value = value
		return
	}

	//从T的左边插入
	if T.Left.Key == rb.NullStr && T.Key > key{
		newNode := node.NewRBNode(key, value, false, T, nil, nil)
		nullNodeLeft := node.NewRBNode(rb.NullStr, nil, true, newNode, nil,nil)
		nullNodeRight := node.NewRBNode(rb.NullStr, nil, true, newNode, nil,nil)
		newNode.Left = nullNodeLeft
		newNode.Right = nullNodeRight
		T.Left = newNode
		rb.insertFixUp(T)
		return
	}
	//从T的右边插入
	if T.Right.Key == rb.NullStr && T.Key < key{
		newNode := node.NewRBNode(key, value, false, T, nil, nil)
		nullNodeLeft := node.NewRBNode(rb.NullStr, nil, true, newNode, nil,nil)
		nullNodeRight := node.NewRBNode(rb.NullStr, nil, true, newNode, nil,nil)
		newNode.Left = nullNodeLeft
		newNode.Right = nullNodeRight
		T.Right = newNode
		rb.insertFixUp(T)
		return
	}
	//从左边递归
	if T.Left !=nil{
		if T.Key > key{
			rb.Insert(T.Left, key, value)
			return
		}
	}
	//从右边递归
	if T.Right != nil{
		if T.Key < key{
			rb.Insert(T.Right, key, value)
			return
		}
	}
}

//删除后进行修复
func (rb *RBTree)deleteFixUp(T *node.RBNode){
	father := T.Father
	if father==nil{
		rb.Header = node.NewRBNode(rb.NullStr, nil, true, nil, nil, nil)
		return
	}
	if father.Left == T{
		brother := father.Right
		if brother.Color == false{
			father.Color = false
			brother.Color = true
			rb.singleRightRotate(father)
			rb.deleteFixUp(T)
			return
		}else{
			sl := brother.Left
			sr := brother.Right
			if sl.Color==false{
				sl.Color = true
				rb.doubleLeftRotate(father)
				return
			}
			if sr.Color==false{
				//sr.Color = true
				rb.singleRightRotate(father)
				return
			}
			if father.Color == true{
				brother.Color = false
				rb.deleteFixUp(father)
				return
			}
			if father.Color == false{
				father.Color = true
				brother.Color = false
				return
			}
		}
	}else {
		brother := father.Left
		if brother.Color == false{
			father.Color = false
			brother.Color = true
			rb.singleLeftRotate(father)
			rb.deleteFixUp(T)
			return
		}else{
			sl := brother.Left
			sr := brother.Right
			if sl.Color==false{
				//sl.Color = true
				rb.singleLeftRotate(father)
				return
			}
			if sr.Color==false{
				sr.Color = true
				rb.doubleRightRotate(father)
				return
			}
			if father.Color == true{
				brother.Color = false
				rb.deleteFixUp(father)
				return
			}
			if father.Color == false{
				father.Color = true
				brother.Color = false
				return
			}
		}
	}
}

func (rb *RBTree)Delete(T *node.RBNode, key string)bool{
	if T==nil{
		return false
	}
	if T.Key==rb.NullStr{
		return false
	}
	if rb==nil ||rb.Header==nil||rb.Header.Key==rb.NullStr{
		return false
	}
	if T.Key == key {
		l := T.Left
		r := T.Right
		father := T.Father
		if rb.IsEmptyNode(father) {
			if rb.IsEmptyNode(l) && rb.IsEmptyNode(r) {
				rb.Header = node.NewRBNode(rb.NullStr, nil, true, nil, nil, nil)
				return true
			}
			if rb.IsEmptyNode(l)&&!rb.IsEmptyNode(r){
				r.Father = nil
				rb.Header = r
				r.Color = true
				return true
			}
			if !rb.IsEmptyNode(l)&&rb.IsEmptyNode(r){
				l.Father = nil
				rb.Header = l
				l.Color = true
				return true
			}
		} else {
			if rb.IsEmptyNode(l) && rb.IsEmptyNode(r) {
				if T.Color == false {
					if father.Left == T {
						father.Left = r
						r.Father = father
					} else if father.Right == T {
						father.Right = r
						r.Father = father
					}
					return true
				} else {
					if father.Left == T {
						father.Left = r
						r.Father = father
					} else if father.Right == T {
						father.Right = r
						r.Father = father
					}
					rb.deleteFixUp(r)
					return true
				}
			}
			if !rb.IsEmptyNode(l) && rb.IsEmptyNode(r) {
				l.Color = true
				if father.Left == T {
					father.Left = l
					l.Father = father
				} else {
					father.Right = l
					l.Father = father
				}
				return true
			}
			if rb.IsEmptyNode(l) && !rb.IsEmptyNode(r) {
				r.Color = true
				if father.Left == T {
					father.Left = r
					r.Father = father
				} else {
					father.Right = r
					r.Father = father
				}
				return true
			}
		}
		//找到后继节点，替换位置然后删除后继节点
		if !rb.IsEmptyNode(r) && !rb.IsEmptyNode(l){
			leftSonForR := r
			//找到最左边非空节点
			for {
				if rb.IsEmptyNode(leftSonForR.Left) {
					break
				}
				leftSonForR = leftSonForR.Left
			}
			T.Key = leftSonForR.Key
			leftSonForR.Key = key
			return rb.Delete(leftSonForR, key)
		}
	}else if T.Key<key{
		return rb.Delete(T.Right, key)
	}else{
		return rb.Delete(T.Left, key)
	}
	return true
}


