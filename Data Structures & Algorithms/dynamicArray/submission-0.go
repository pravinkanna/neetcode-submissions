type DynamicArray struct {
	arr []int
	size int
	capacity int
}

func NewDynamicArray(capacity int) *DynamicArray {
	newArr := make([]int, capacity)
	return &DynamicArray{arr:newArr, size:0, capacity:capacity}
}

func (da *DynamicArray) Get(i int) int {
	return da.arr[i]

}

func (da *DynamicArray) Set(i int, n int) {
	da.arr[i] = n
}

func (da *DynamicArray) Pushback(n int) {
	if da.size >= da.capacity{
		da.resize()
	}
	da.arr[da.size] = n
	da.size++
}

func (da *DynamicArray) Popback() int {
	val := da.arr[da.size-1]
	da.arr[da.size-1] = 0
	da.size--
	return val
}

func (da *DynamicArray) resize() {
	newCap := da.capacity*2
	newArr := make([]int, newCap)

	for idx, val := range da.arr {
		newArr[idx] = val
	}
	da.arr = newArr
	da.capacity = newCap
}

func (da *DynamicArray) GetSize() int {
	return da.size
}

func (da *DynamicArray) GetCapacity() int {
	return da.capacity
}
