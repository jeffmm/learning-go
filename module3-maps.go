package main

import (
    "fmt"
    "strings"
)

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

// Mutate the map 
func map_mutation() {

    // Make a map of string to int
    m := make(map[string]int)

    // The ultimate answer to life and everything
    m["Answer"] = 42
    fmt.Println("The answer:", m["Answer"])

    // Check the key and value of the map
    v, ok := m["Answer"]
    fmt.Println("The value:", v, "Present?", ok)

    // The answer to 6 * 8 I guess?
    m["Answer"] = 48
    fmt.Println("The answer:", m["Answer"])

    // Delete the value of the key
    delete(m, "Answer")
    fmt.Println("The answer:", m["Answer"])

    // The value is zero'd out
    v, ok = m["Answer"]
    fmt.Println("The value:", v, "Present?", ok)

    // Does this work? Yes. It just gives default values of zero.
    v, ok = m["Question"]
    fmt.Println("The value:", v, "Present?", ok)
}

func word_counts(s string) map[string]int {
    counts := make(map[string]int)
    words := strings.Fields(s)
    var v int
    for i := 0; i < len(words); i++ {
        v, _ = counts[words[i]]
        counts[words[i]] = v + 1
    }
    return counts
}

func main() {
    make_a_map()
    map_literals()
    map_mutation()
    my_string := "This sentence has six six six in the words words"
    count_map := word_counts(my_string)
    for k := range count_map {
        fmt.Println(k, ":", count_map[k])
    }
}
