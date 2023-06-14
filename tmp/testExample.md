# 部分测试代码

## 链表题目的测试：
```

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	a, b, c, d, e := &ListNode{Val: 1}, &ListNode{Val: 2}, &ListNode{Val: 3}, &ListNode{Val: 4}, &ListNode{Val: 5}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
	for cur := reverseList(a); cur != nil; cur = cur.Next {
		fmt.Printf("%d\t", cur.Val)
	}
}
```

[true true true true true true true true false false true true false false false true true false false false false false false false false true false false false false false true false true true false false false true false false false false false true true true false false false false true false false true false true true false true false false true true true true true false false false true true false false true false true]