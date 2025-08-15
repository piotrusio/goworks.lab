## Array - Fixed Foundation

```go
// Array creation - fixed size, type includes size
arr1 := [5]int{1, 2, 3}        // [1, 2, 3, 0, 0] - partial init with zeros
arr2 := [...]int{1, 2, 3, 4}   // [1, 2, 3, 4] - compiler counts size
var arr3 [3]string             // ["", "", ""] - zero values

// Basic operations
fmt.Println(len(arr1))         // 5 - always equals array size
fmt.Println(arr1[0])           // 1 - zero-based indexing
arr1[0] = 99                   // Direct assignment

// Iteration patterns
for i := 0; i < len(arr1); i++ {
    fmt.Printf("arr1[%d] = %d\n", i, arr1[i])
}

for i, value := range arr2 {
    fmt.Printf("Index %d: %d\n", i, value)
}

// Arrays are comparable if same size and type
fmt.Println(arr1 == arr2)      // false - different values
fmt.Println([3]int{1,2,3} == [3]int{1,2,3}) // true

// Pass by value - entire array copied
func modifyArray(a [5]int) { a[0] = 999 } // Only modifies copy
modifyArray(arr1)
fmt.Println(arr1[0])           // Still 99, not 999
```

## Array Operations and Patterns

```go
// Multi-dimensional arrays
matrix := [3][3]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
}
fmt.Println(matrix[1][2])      // 6 - access row 1, col 2

// Array comparison - same size and type required
arr1 := [3]int{1, 2, 3}
arr2 := [3]int{1, 2, 3}
arr3 := [4]int{1, 2, 3, 4}
fmt.Println(arr1 == arr2)      // true - same values
// fmt.Println(arr1 == arr3)   // Compile error - different types

// Type safety - size is part of type
var small [3]int
var large [5]int
// small = large                // Compile error - different types

// Common patterns
func reverseArray(arr [5]int) [5]int {
    for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
        arr[i], arr[j] = arr[j], arr[i]
    }
    return arr // Returns modified copy
}

// When to use arrays vs slices
// Arrays: Fixed size known at compile time, pass-by-value needed
// Slices: Dynamic size, efficient passing, most common choice
config := [3]string{"dev", "test", "prod"} // Fixed configuration
```