## Map Fundamentals

```go
// Map creation - three methods
fruits := make(map[string]int)                   // Empty map, ready to use
scores := map[string]int{"Alice": 95, "Bob": 87} // Map literal with data
empty := map[string]int{}                        // Empty map literal

// Nil map behavior
var nilMap map[string]int                        // nil map
fmt.Println(nilMap == nil)                       // true
value := nilMap["key"]                           // 0 - safe to read (returns zero value)
// nilMap["key"] = 42                            // PANIC! Cannot write to nil map

// Must initialize before writing
nilMap = make(map[string]int)
nilMap["key"] = 42                               // Now safe

// Key type restrictions - must be comparable
validKeys := map[int]string{1: "one"}            // ✅ int keys
stringKeys := map[string]int{"key": 42}          // ✅ string keys
structKeys := map[struct{x, y int}]string{}      // ✅ comparable struct keys
// sliceKeys := map[[]int]string{}               // ❌ Slices not comparable

// Basic operations
inventory := make(map[string]int)
inventory["apples"] = 50                         // Insert
inventory["apples"] = 45                         // Update (same syntax)
count := inventory["apples"]                     // Read: 45
missing := inventory["oranges"]                  // Read missing: 0 (zero value)
delete(inventory, "apples")                      // Delete
fmt.Println(len(inventory))                      // 0

// Map properties
m1 := map[string]int{"a": 1}
m2 := m1                                         // Reference copy - shared data!
m2["b"] = 2
fmt.Println(m1)                                  // map[a:1 b:2] - m1 also changed
// fmt.Println(m1 == m2)                        // Compile error - maps not comparable
fmt.Println(m1 == nil)                          // false - only nil comparison allowed
```

## Map Operations

```go
// Comma OK idiom - distinguish between zero value and missing key
inventory := map[string]int{"apples": 50, "bananas": 0}
value1 := inventory["apples"]                    // 50
value2 := inventory["bananas"]                   // 0 - but key exists!
value3 := inventory["oranges"]                   // 0 - key missing!

// Safe lookup with comma OK
if count, exists := inventory["oranges"]; exists {
    fmt.Println("Found oranges:", count)
} else {
    fmt.Println("No oranges in stock")           // This executes
}

// Common comma OK patterns
// Pattern 1: Check before using
if value, ok := m["key"]; ok {
    fmt.Println("Found:", value)
}

// Pattern 2: Set default if missing
value, ok := m["key"]
if !ok {
    value = 42                                   // Use default
}

// Pattern 3: Conditional insert (only if missing)
if _, exists := m["key"]; !exists {
    m["key"] = initialValue
}

// Map iteration with range - ORDER IS RANDOM!
scores := map[string]int{"Alice": 95, "Bob": 87, "Carol": 92}
for name, score := range scores {               // Random order each run!
    fmt.Printf("%s: %d\n", name, score)
}

// Deep copy vs shallow copy
original := map[string]int{"a": 1, "b": 2}
reference := original                           // Shallow - shares data
clone := make(map[string]int)                   // Deep copy
for k, v := range original {
    clone[k] = v
}

// Practical patterns
// Frequency counter
func countWords(text string) map[string]int {
    counts := make(map[string]int)
    for _, word := range strings.Fields(text) {
        counts[word]++                          // Auto-initializes to 0, then increments
    }
    return counts
}

// Set operations using map[T]bool
uniqueItems := make(map[string]bool)
uniqueItems["item1"] = true
if uniqueItems["item1"] {                       // Check membership
    fmt.Println("item1 exists in set")
}
```

## Map Internals and Performance

```go
// Hash table fundamentals - Go maps use buckets with overflow chaining
// Each bucket holds ~8 key-value pairs, with overflow buckets when full

// Load factor monitoring - Go targets load factor ~6.5
m := make(map[string]int)
for i := 0; i < 1000; i++ {
    m[fmt.Sprintf("key%d", i)] = i    // Triggers growth at certain thresholds
}
// Bucket count grows: 8 → 16 → 32 → 64 → 128...

// Performance characteristics
// Average case: O(1) insert, lookup, delete
// Worst case: O(n) if all keys hash to same bucket (rare with Go's hash functions)

// Map growth triggers rehashing - ALL elements moved to new bucket layout
// Pre-allocation reduces rehashing overhead
slowMap := make(map[string]int)           // Multiple rehashing events
fastMap := make(map[string]int, 10000)    // Pre-allocated, fewer rehashing events

// Memory considerations - maps don't auto-shrink
bigMap := make(map[string]int)
for i := 0; i < 100000; i++ {
    bigMap[fmt.Sprintf("key%d", i)] = i   // Memory allocated for 100K elements
}

// Delete most elements - memory still allocated!
for i := 0; i < 95000; i++ {
    delete(bigMap, fmt.Sprintf("key%d", i))
}
// len(bigMap) = 5000, but memory for 100K still used

// Reclaim memory by creating new map
compactMap := make(map[string]int, len(bigMap))
for k, v := range bigMap {
    compactMap[k] = v                     // Copy to right-sized map
}
bigMap = compactMap                       // Old map eligible for GC

// Key type performance - comparable types only
intKeys := make(map[int]string)           // Fast: simple hash
stringKeys := make(map[string]int)        // Good: optimized string hashing
structKeys := make(map[struct{x,y int}]string) // Slower: composite hashing
// sliceKeys := make(map[[]int]string)    // Invalid: slices not comparable
```

## Map Patterns and Use Cases

```go
// Maps as sets - two approaches
// Approach 1: map[T]bool (1 byte per value)
fruitsSet := make(map[string]bool)
fruitsSet["apple"] = true
if fruitsSet["apple"] { fmt.Println("Found apple") }

// Approach 2: map[T]struct{} (0 bytes per value - more efficient)
type StringSet map[string]struct{}
users := make(StringSet)
users["user123"] = struct{}{}               // struct{}{} = zero-byte value
if _, exists := users["user123"]; exists { fmt.Println("User exists") }

// Frequency counting pattern
func charFrequency(text string) map[rune]int {
    freq := make(map[rune]int)
    for _, char := range text {
        freq[char]++                        // Auto-initializes to 0, then increments
    }
    return freq
}

// Grouping data pattern
type Student struct { Name, Grade string; Score int }
func groupByGrade(students []Student) map[string][]Student {
    groups := make(map[string][]Student)
    for _, student := range students {
        groups[student.Grade] = append(groups[student.Grade], student)
    }
    return groups
}

// Nested maps for multi-level indexing
type Cache map[string]map[string]interface{}
cache := make(Cache)
if cache["user"] == nil {
    cache["user"] = make(map[string]interface{})
}
cache["user"]["resource"] = "cached_data"

// Complex key types - structs as keys (must be comparable)
type Point struct { X, Y int }
distances := map[Point]float64{
    {0, 0}: 0.0,
    {1, 1}: 1.414,
    {2, 3}: 3.606,
}
fmt.Println(distances[Point{1, 1}])         // 1.414

// Convert slice to set for O(1) lookups instead of O(n)
items := []string{"apple", "banana", "apple", "cherry"}
uniqueSet := make(map[string]struct{})
for _, item := range items {
    uniqueSet[item] = struct{}{}
}
fmt.Println("Unique count:", len(uniqueSet)) // 3
```