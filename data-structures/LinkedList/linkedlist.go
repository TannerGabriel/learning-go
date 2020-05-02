package LinkedList

type node struct {
	val  int
	next *node
	prev *node
}

type LinkedList struct {
	head *node
}

func NewNode(val int) *node {
	n := &node{}
	n.val = val
	n.next = nil
	n.prev = nil
	return n
}

func (ll *LinkedList) Prepend(val int) {
	n := NewNode(val)
	n.next = ll.head
	ll.head = n
}

func (ll *LinkedList) Append(val int) {
	n := NewNode(val)

	if ll.head == nil {
		ll.head = n
		return
	}

	cur := ll.head
	for ; cur.next != nil; cur = cur.next {
	}
	cur.next = n
	n.prev = cur
}

func (ll *LinkedList) RemoveAtBeg() int {
	if ll.head == nil {
		return -1
	}

	cur := ll.head
	ll.head = cur.next

	if ll.head != nil {
		ll.head.prev = nil
	}

	return cur.val
}

func (ll *LinkedList) RemoveAtEnd() int {
	if ll.head == nil {
		return -1
	}

	if ll.head.next == nil {
		return ll.RemoveAtBeg()
	}

	cur := ll.head
	for ; cur.next.next != nil; cur = cur.next {
	}

	retval := cur.next.val
	cur.next = nil
	return retval
}

func (ll *LinkedList) Count() int {
	var ctr int = 0

	for cur := ll.head; cur != nil; cur = cur.next {
		ctr += 1
	}

	return ctr
}

func (ll *LinkedList) Reverse() {
	var prev, next *node
	cur := ll.head

	for cur != nil {
		next = cur.next
		cur.next = prev
		cur.prev = next
		prev = cur
		cur = next
	}

	ll.head = prev
}

func (ll *LinkedList) Display() {
	for cur := ll.head; cur != nil; cur = cur.next {
		fmt.Print(cur.val, " ")
	}

	fmt.Print("\n")
}

func (ll *LinkedList) DisplayReverse() {
	if ll.head == nil {
		return
	}
	var cur *node
	for cur = ll.head; cur.next != nil; cur = cur.next {
	}

	for ; cur != nil; cur = cur.prev {
		fmt.Print(cur.val, " ")
	}

	fmt.Print("\n")
}
