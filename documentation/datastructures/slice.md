## Slice Fundamentals

```go
// Slice creation methods
slice1 := make([]int, 5)           // [0, 0, 0, 0, 0] - len=5, cap=5
slice2 := make([]int, 0, 10)       // [] - len=0, cap=10 (pre-allocated)
slice3 := []int{1, 2, 3, 4, 5}     // [1, 2, 3, 4, 5] - literal
slice4 := []int{}                  // [] - empty but not nil

// Nil slice vs empty slice
var nilSlice []int                 // nil slice
emptySlice := []int{}              // empty slice
fmt.Println(nilSlice == nil)       // true
fmt.Println(emptySlice == nil)     // false
fmt.Println(len(nilSlice))         // 0 - safe to call len on nil

// Slice header concept: {pointer, length, capacity}
fmt.Printf("len=%d, cap=%d\n", len(slice2), cap(slice2)) // len=0, cap=10

// Slicing from arrays and other slices
arr := [5]int{10, 20, 30, 40, 50}
sub1 := arr[:]                     // [10, 20, 30, 40, 50] - entire array
sub2 := arr[1:4]                   // [20, 30, 40] - from index 1 to 3
sub3 := slice3[1:3]                // [2, 3] - slice from slice

// Three-index slicing: [start:end:cap]
limited := arr[1:3:4]              // [20, 30] with limited capacity
fmt.Printf("len=%d, cap=%d\n", len(limited), cap(limited)) // len=2, cap=3
```

## Slice Operations

```go
// append() - grows slice as needed
slice := []int{1, 2, 3}
slice = append(slice, 4)           // [1, 2, 3, 4] - must reassign!
slice = append(slice, 5, 6, 7)     // [1, 2, 3, 4, 5, 6, 7] - multiple values

// Append another slice (spread operator)
other := []int{8, 9, 10}
slice = append(slice, other...)    // [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

// copy() - copies elements between slices
src := []int{1, 2, 3, 4, 5}
dst := make([]int, len(src))
n := copy(dst, src)                // Returns number copied: 5
fmt.Println(dst)                   // [1, 2, 3, 4, 5]

// Partial copy - copies min(len(dst), len(src))
small := make([]int, 3)
copy(small, src)                   // [1, 2, 3] - only first 3 elements

// Overlapping copy (safe)
data := []int{1, 2, 3, 4, 5}
copy(data[2:], data[1:4])          // [1, 2, 2, 3, 4] - shift right
copy(data[1:], data[2:])           // [1, 2, 3, 4, 4] - shift left

// Capacity growth pattern - Go typically doubles
slice = make([]int, 0, 1)
for i := 0; i < 10; i++ {
    slice = append(slice, i)
    fmt.Printf("len=%d, cap=%d\n", len(slice), cap(slice))
} // Capacity grows: 1 → 2 → 4 → 8 → 16
```

## Slice Memory Model

```go
// Memory sharing - multiple slices can share same underlying array
original := []int{1, 2, 3, 4, 5}
slice1 := original[1:4]            // [2, 3, 4] - shares memory with original
slice2 := original[2:5]            // [3, 4, 5] - also shares same memory

slice1[1] = 99                     // Modifies shared array
fmt.Println(original)              // [1, 2, 99, 4, 5] - original changed!
fmt.Println(slice2)                // [99, 4, 5] - slice2 also affected!

// append() may break sharing by creating new array
slice3 := []int{10, 20, 30}        // cap=3
slice4 := slice3                   // Shares same array

slice3 = append(slice3, 40)        // Exceeds capacity - NEW array created!
slice3[0] = 999
fmt.Println(slice4)                // [10, 20, 30] - unaffected (old array)

// Memory leak prevention - avoid holding references to large arrays
func getFirstTen(huge []int) []int {
    // Bad: return huge[:10]        // Keeps entire huge array in memory
    
    // Good: create independent copy
    result := make([]int, 10)
    copy(result, huge[:10])
    return result                   // huge can be garbage collected
}

// Safe copying when independence needed
func deepCopy(src []int) []int {
    dst := make([]int, len(src))
    copy(dst, src)
    return dst                      // Completely independent slice
}
```