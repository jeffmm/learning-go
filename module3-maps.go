package main

import "fmt"

// A 2D point in space
type Vertex struct {
    x, y float64
}

// This is how we define a mapping from string keys to Vertex values
var m map[string]Vertex

// Making a map from string to vertex
func make_a_map() {
    // Map a string to the Vertex struct
    m = make(map[string]Vertex)
    m["lab"] = Vertex{92.1, 3.14}
    fmt.Println(m["lab"])
}

// Creating a map literal
func map_literals() {
    var m = map[string]Vertex{
        // Strings need to have explicit quotes, not apostrophes
        "lab": Vertex{10., 1.},
        "site": Vertex{9., 5.},  // Need a comma here
    }

    // We can omit the type name when building map literals
    var n = map[string]Vertex{
        "lab2": {10., 4.},
        "site2": {9., 6.},
    }
    fmt.Println(m)
    fmt.Println(m["lab"])
    fmt.Println(n)
    fmt.Println(n["lab2"])
}

func main() {
    make_a_map()
    map_literals()
}
