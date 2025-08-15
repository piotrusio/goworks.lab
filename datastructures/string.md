## String Fundamentals

```go
// String structure: {pointer to bytes, length} - immutable!
str := "Hello 世界"
fmt.Println(len(str))              // 11 bytes (not 8 characters!)

// String vs []byte vs []rune
bytes := []byte(str)               // [72 101 108 108 111 32 228 184 150 231 149 140]
runes := []rune(str)               // [72 101 108 108 111 32 19990 30028]
fmt.Println(len(bytes))            // 11 bytes
fmt.Println(len(runes))            // 8 characters (Unicode code points)

// String iteration - two ways
// By bytes (can break UTF-8 sequences)
for i := 0; i < len(str); i++ {
    fmt.Printf("byte[%d]: %c (%d)\n", i, str[i], str[i])
}

// By runes (proper UTF-8 handling)
for i, r := range str {
    fmt.Printf("rune[%d]: %c (%d)\n", i, r, r)
}

// String immutability and memory sharing
str1 := "hello"
str2 := str1                       // Shares memory (safe - immutable)
substr := str1[1:4]                // "ell" - shares memory with str1

// String literals
regular := "Hello\nWorld"          // Interpreted escapes
raw := `Hello\nWorld`              // Raw string - literal \n
unicode := "Hello \u4e16\u754c"    // Unicode escapes: "Hello 世界"

// Memory considerations for substrings
func getSafeSubstring(text string, start, end int) string {
    return string([]byte(text[start:end])) // Force copy, allow GC of original
}
```

## String Operations and Methods

```go
import (
    "fmt"
    "strconv"
    "strings"
)

// strings package - essential functions
text := "Hello World Hello"
fmt.Println(strings.Contains(text, "World"))     // true
fmt.Println(strings.Count(text, "Hello"))        // 2
fmt.Println(strings.Index(text, "World"))        // 6
fmt.Println(strings.ToLower(text))               // "hello world hello"
fmt.Println(strings.Replace(text, "Hello", "Hi", -1)) // "Hi World Hi"

// Splitting and joining
csv := "apple,banana,cherry"
parts := strings.Split(csv, ",")                 // ["apple", "banana", "cherry"]
joined := strings.Join(parts, " | ")             // "apple | banana | cherry"
fields := strings.Fields("  hello   world  ")   // ["hello", "world"]

// Trimming
messy := "  \n  Hello World  \n  "
clean := strings.TrimSpace(messy)                // "Hello World"
fmt.Println(strings.TrimPrefix(clean, "Hello ")) // "World"

// fmt.Sprintf - string formatting
name, age := "Alice", 25
formatted := fmt.Sprintf("Name: %s, Age: %d", name, age) // "Name: Alice, Age: 25"
quoted := fmt.Sprintf("%q", name)                // "\"Alice\""

// strconv - string conversions
num, err := strconv.Atoi("123")                  // 123, nil
str := strconv.Itoa(456)                         // "456"
f, err := strconv.ParseFloat("3.14", 64)         // 3.14, nil
bool_str := strconv.FormatBool(true)             // "true"

// strings.Builder - efficient string building
var builder strings.Builder
builder.WriteString("Hello")
builder.WriteString(" ")
builder.WriteString("World")
result := builder.String()                       // "Hello World"
```