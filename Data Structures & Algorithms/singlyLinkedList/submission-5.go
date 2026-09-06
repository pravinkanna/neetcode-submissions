type ListNode struct {
	val int
	next *ListNode
}

type LinkedList struct {
	head *ListNode
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) Get(index int) int {
	nextNode := ll.head
	if nextNode == nil {
			return -1
	} 
	for i:=0; i<index; i++ {
		nextNode = nextNode.next
		if nextNode == nil {
			return -1
		}
	}
	return nextNode.val
}

func (ll *LinkedList) InsertHead(val int) {
	currHead := ll.head
	node := &ListNode{val: val, next: currHead}
	ll.head = node
}

func (ll *LinkedList) InsertTail(val int) {
	node := &ListNode{val: val, next: nil}
	currNode := ll.head
	if currNode == nil {
		ll.head = node
	} else {
	for ; currNode.next != nil; currNode = currNode.next{}
		currNode.next = node
	}
	
}

func (ll *LinkedList) Remove(index int) bool {
	var prev *ListNode
	curr := ll.head
	if curr == nil {
		return false
	}

	if index == 0 {
		ll.head = curr.next
		return true
	}

	for i:=0; i<index; i++ {
		prev = curr
		curr = curr.next

		if curr == nil {
			return false
		}
	}
	prev.next = curr.next
	return true
}

func (ll *LinkedList) GetValues() []int {
	var arr []int
	
	for currNode := ll.head; currNode != nil; currNode = currNode.next{
		arr = append(arr, currNode.val)
	} 
	return arr
}
