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