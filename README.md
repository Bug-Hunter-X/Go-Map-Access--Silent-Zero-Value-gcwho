# Go Map Access: Silent Zero Value

This repository demonstrates a common pitfall in Go: accessing a non-existent key in a map. Unlike some languages that would throw an exception, Go silently returns the zero value for the map's value type. This can lead to subtle bugs that are difficult to track down.

## Bug

The `bug.go` file shows that accessing a key in an uninitialized map doesn't cause a panic.  It just returns 0 (the zero value for `int`):

```go
func main() {
    var m map[string]int
    fmt.Println(m["a"]) // Prints 0
}
```

This behavior might be unexpected for developers coming from languages where accessing a non-existent key throws an error or exception.

## Solution

The `bugSolution.go` file presents a safer approach by explicitly checking if the key exists using the `ok` idiom:

```go
func main() {
    var m map[string]int
    val, ok := m["a"]
    if ok {
        fmt.Println("Value of 'a':", val)
    } else {
        fmt.Println("Key 'a' not found in the map.")
    }
}
```

This idiom provides a clear and explicit way to handle situations where the key might be missing, eliminating the potential for silent errors.