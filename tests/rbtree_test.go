package tests

import (
	"fmt"
	"rbtree"
	"testing"
)

func TestInsertRb(t *testing.T){
	rb := rbtree.NewRBTree("c", 123, "")
	fmt.Println("插入一个数后")
	rb.PreTraversal(rb.Header)
	rb.Insert(rb.Header, "a", 34)
	fmt.Println("插入一个数后")
	rb.PreTraversal(rb.Header)
	rb.Insert(rb.Header, "b", 45)
	fmt.Println("插入一个数后")
	rb.PreTraversal(rb.Header)
	rb.Insert(rb.Header, "f", 45)
	fmt.Println("插入一个数后")
	rb.PreTraversal(rb.Header)
	rb.Insert(rb.Header, "g", 45)
	fmt.Println("插入一个数后")
	rb.PreTraversal(rb.Header)
	rb.Insert(rb.Header, "d", 45)
	fmt.Println("插入一个数后")
	rb.PreTraversal(rb.Header)
	rb.Insert(rb.Header, "e", 45)
	fmt.Println("插入一个数后")
	rb.PreTraversal(rb.Header)
	rb.Insert(rb.Header, "h", 45)
	fmt.Println("插入一个数后")
	rb.PreTraversal(rb.Header)
	rb.Insert(rb.Header, "j", 45)
	fmt.Println("插入一个数后")
	rb.PreTraversal(rb.Header)
	//rb.MiddleTraversal(rb.Header)

	fmt.Println("插入一个数后,中序")
	rb.MiddleTraversal(rb.Header)


	rb.Delete(rb.Header, "j")
	fmt.Println("删除一个数后,先序")
	rb.PreTraversal(rb.Header)
	fmt.Println("删除一个数后,中序")
	rb.MiddleTraversal(rb.Header)
	rb.Delete(rb.Header, "e")
	fmt.Println("删除一个数后,先序")
	rb.PreTraversal(rb.Header)
	fmt.Println("删除一个数后,中序")
	rb.MiddleTraversal(rb.Header)
	rb.Delete(rb.Header, "f")
	fmt.Println("删除一个数后,先序")
	rb.PreTraversal(rb.Header)
	fmt.Println("删除一个数后,中序")
	rb.MiddleTraversal(rb.Header)
	rb.Delete(rb.Header, "a")
	fmt.Println("删除一个数后,先序")
	rb.PreTraversal(rb.Header)
	fmt.Println("删除一个数后,中序")
	rb.MiddleTraversal(rb.Header)
	rb.Delete(rb.Header, "b")
	fmt.Println("删除一个数后,先序")
	rb.PreTraversal(rb.Header)
	fmt.Println("删除一个数后,中序")
	rb.MiddleTraversal(rb.Header)
	//rb.PreTraversal(rb.Header)
}
