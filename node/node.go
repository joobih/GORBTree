package node

type RBNode struct{
	Key string
	Value interface{}
	Color bool		//true--黑；false--红
	Father *RBNode
	Right *RBNode
	Left *RBNode
}

//创建一个新的节点
func NewRBNode(key string, value interface{}, color bool,father,right,left *RBNode)*RBNode{
	rbNode := new(RBNode)
	rbNode.Value = value
	rbNode.Key = key
	rbNode.Color = color
	rbNode.Father = father
	rbNode.Left = left
	rbNode.Right = right
	return rbNode
}

