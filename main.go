package main

import "fmt"

func main() {
	arr1 := [6]int{10,20,30,40}
	arr2 := [...]int{10,20,30,40,50}
	for i, v := range arr1[2:4] {
		fmt.Printf("Index: %d value %d\n", i, v)
	}

	fmt.Printf("The arr1 length %d\n", len(arr1))
	fmt.Printf("The arr1 capacity %d\n", cap(arr1))
	fmt.Printf("The arr2 length %d\n", len(arr2))
	fmt.Printf("The arr2 capacity %d\n", cap(arr2))
}

type DynamicArray struct {
	data []int
	size int
	capacity int
}

func NewDynamicArray(initialCapacity int) *DynamicArray {
	return &DynamicArray{
		data: make([]int, 0, initialCapacity),
		size: 0,
		capacity: initialCapacity,
	}
}

func (d *DynamicArray) Get(index int) (int, error) {
	if index < 0 {
		return 0, fmt.Errorf("index cannot be negativa: %d", index)
	}

	if index >= d.size {
		return 0, fmt.Errorf("index %d out of bounds, size is %d", index, d.size)
	}

	return d.data[index], nil
}

func (d *DynamicArray) Set(index int, value int) error {
	if index < 0 {
		return fmt.Errorf("index cannot be negative %d", index)
	}
	if index >= d.size {
		return fmt.Errorf("index %d out of bounds, size is %d", index, d.size)
	}

	d.data[index] = value
	return nil
}

func (d *DynamicArray) Append(value int) {
	if d.size < d.capacity {
		d.data[d.size] = value
		d.size++
	} else {
		d.resize(2 * d.capacity)
		d.data[d.size] = value
		d.size++
	}
}

func (d *DynamicArray) resize(newCapacity int) {
	newData := make([]int, d.size, newCapacity)
	copy(newData, d.data)
	d.data = newData
	d.capacity = newCapacity
}

func (d *DynamicArray) Insert(index int, value int) error {
	if index < 0 {
		return fmt.Errorf("index cannot be negative %d", index)
	}
	
	if index > d.size {
		return fmt.Errorf("index %d out of bounds, size is %d", index, d.size)
	}

	if d.size >= d.capacity {
		d.resize(2 * d.capacity)
	}

	// shift elements right
	// Before: [10, 20, 30], size=3, insert 99 at index 1
	// d.data[index:d.size] = d.data[1:3] = [20, 30] (source)
	// d.data[index+1:]     = d.data[2:]   = positions 2,3... (destination)
	// copy(d.data[2:], [20, 30])
	// After copy: [10, _, 20, 30] (position 1 is now free)
	copy(d.data[index+1:], d.data[index:d.size])

	// insert
	d.data[index] = value
	d.size++

	return nil
}

func (d *DynamicArray) Delete(index int) error {
	if index < 0 {
		return fmt.Errorf("index cannot be negative %d", index)
	}

	if index >= d.size {
		return fmt.Errorf("index %d out of bounds, size is %d", index, d.size)
	}

	// Before: [10, 20, 30, 40, 0, 0, 0], size=4, delete at index 1 -> 20
	// d.data[index+1:d.size] = d.data[2:4] -> [30, 40] //
	// d.data[index:] -> d.data[1:]
	copy(d.data[index:], d.data[index+1:d.size])
	d.size--

	return nil
}

func (d *DynamicArray) Size() int {
	return d.size
}

func (d *DynamicArray) IsEmpty() bool {
	return d.size == 0
}

func (d *DynamicArray) String() string {
    if d.size == 0 {
        return "[]"
    }
    
    // Build string like "[10, 20, 30]"
    result := "["
    for i := 0; i < d.size; i++ {
        if i > 0 {
            result += ", "
        }
        result += fmt.Sprintf("%d", d.data[i])
    }
    result += "]"
    return result
}