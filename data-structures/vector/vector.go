package vector

type Vector struct {
	size     int
	capacity int
	data     []int
}

func NewVector() Vector {
	return Vector{
		size:     0,
		capacity: 16,
		data:     make([]int, 16),
	}
}

func (v Vector) Size() int {
	return v.size
}

func (v Vector) Capacity() int {
	return v.capacity
}

func (v Vector) IsEmpty() bool {
	return v.size == 0
}

func (v Vector) At(index int) int {
	if index < 0 || index > v.size-1 {
		panic("index out of bounds")
	}
	return v.data[index]
}

func (v *Vector) Push(item int) {
	if v.size == v.capacity {
		v.resize(v.size * 2)
	}
	v.data[v.size] = item
	v.size++
}

func (v *Vector) Insert(index, item int) {
	if index < 0 || index > v.size {
		panic("index out of bounds")
	}
	if v.size == v.capacity {
		v.resize(v.size * 2)
	}

	for i := v.size; i > index; i-- {
		v.data[i] = v.data[i-1]
	}
	v.data[index] = item
	v.size++
}

func (v *Vector) Prepend(item int) {
	v.Insert(0, item)
}

func (v *Vector) Pop() int {
	if v.IsEmpty() {
		panic("vector is empty")
	}
	item := v.data[v.size-1]
	v.data[v.size-1] = 0
	v.size--
	if v.size == v.capacity/4 {
		v.resize(v.capacity / 2)
	}
	return item
}

func (v *Vector) Delete(index int) {
	if index < 0 || index >= v.size {
		panic("index out of bounds")
	}
	v.data[index] = 0
	for i := index; i < v.size-1; i++ {
		v.data[i] = v.data[i+1]
	}
	v.size--
	if v.size == v.capacity/4 {
		v.resize(v.capacity / 2)
	}
}

func (v *Vector) Remove(item int) {
	for i, val := range v.data {
		if val == item {
			v.Delete(i)
			i-- // remove multiple occurrences of the element
		}
	}
}

func (v Vector) Find(item int) int {
	for i, val := range v.data {
		if val == item {
			return i
		}
	}
	return -1
}

func (v *Vector) resize(newCapacity int) {
	newData := make([]int, newCapacity)
	copy(newData, v.data)
	v.data = newData
	v.capacity = newCapacity
}
