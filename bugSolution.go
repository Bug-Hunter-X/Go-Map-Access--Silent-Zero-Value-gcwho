```go
package main

import (

    "fmt"
)

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