package vector

import (
	"testing"
)

func TestVector(t *testing.T) {
	v := NewVector()

	// Test IsEmpty() when the vector is empty
	if !v.IsEmpty() {
		t.Errorf("Expected vector to be empty")
	}

	// Test Push() and Size()
	v.Push(1)
	v.Push(2)
	v.Push(3)

	if v.Size() != 3 {
		t.Errorf("Expected vector to have size 3, but got %d", v.Size())
	}

	// Test Capacity() when there is no resizing
	if v.Capacity() != 16 {
		t.Errorf("Expected initial capacity to be 16, but got %d", v.Capacity())
	}

	// Test Capacity() after resizing
	for i := 4; i <= 32; i++ {
		v.Push(i)
	}

	if v.Capacity() != 32 {
		t.Errorf("Expected capacity to be 32, but got %d", v.Capacity())
	}

	// Test At() with valid index
	if v.At(2) != 3 {
		t.Errorf("Expected At(2) to be 3, but got %d", v.At(2))
	}

	// Test At() with invalid index
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected At(100) to panic")
		}
	}()
	v.At(100)

	// Test Insert()
	v.Insert(2, 10)

	if v.Size() != 34 || v.At(2) != 10 || v.At(3) != 3 {
		t.Errorf("Expected Insert(2, 10) to insert 10 At index 2 and shift 3 to index 3")
	}

	// Test Prepend()
	v.Prepend(20)

	if v.Size() != 35 || v.At(0) != 20 {
		t.Errorf("Expected Prepend(20) to insert 20 At index 0")
	}

	// Test Pop()
	popped := v.Pop()

	if v.Size() != 34 || popped != 32 {
		t.Errorf("Expected Pop() to remove 32 and return it")
	}

	// Test Delete()
	v.Delete(2)

	if v.Size() != 33 || v.At(2) != 3 || v.At(3) != 4 {
		t.Errorf("Expected Delete(2) to remove 10 and shift 3 to index 2 and 4 to index 3")
	}

	// Test Remove()
	v.Push(3)
	v.Remove(3)

	if v.Size() != 32 || v.At(2) != 3 || v.Find(3) != -1 {
		t.Errorf("Expected Remove(3) to remove the first occurrence of 3")
	}

	// Test Find()
	if v.Find(5) != 4 || v.Find(100) != -1 {
		t.Errorf("Expected Find() to return the correct index or -1 if not found")
	}
}
