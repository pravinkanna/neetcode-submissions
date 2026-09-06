type Node struct {
	val int
	next *Node
	prev *Node
}

type Deque struct {
	head *Node 
	tail *Node
}


func NewDeque() *Deque {
	return &Deque{}
}

func (d *Deque) IsEmpty() bool {
	if d.head == nil {
		return true
	}
	return false
}

func (d *Deque) Append(value int) {
	
	if d.tail == nil{
		newNode := &Node{val: value}
		d.head = newNode
		d.tail = newNode
		return
	}
	newNode := &Node{val: value, prev: d.tail, next: nil}
	d.tail.next = newNode
	d.tail = newNode
}

func (d *Deque) AppendLeft(value int) {
	if d.head == nil{
		newNode := &Node{val: value}
		d.head = newNode
		d.tail = newNode
		return
	}
	oldHead := d.head
	newNode := &Node{val: value, prev: nil, next: oldHead}
	d.head = newNode
	oldHead.prev = newNode
}

func (d *Deque) Pop() int {
	if d.tail == nil {
		return -1
	}
	if d.tail == d.head {
		oldTailVal := d.tail.val
		d.head = nil
		d.tail = nil
		return oldTailVal
	}
	oldTail := d.tail
	newTail := d.tail.prev
	newTail.next = nil
	d.tail = newTail
	return oldTail.val
}

func (d *Deque) PopLeft() int {
	if d.head == nil {
		return -1
	}
	if d.head  == d.tail {
		oldHeadVal := d.head.val
		d.head = nil
		d.tail = nil
		return oldHeadVal
	}
	oldHead := d.head
	newHead := d.head.next
	newHead.prev = nil
	d.head = newHead
	return oldHead.val
}
