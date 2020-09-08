package main

import(
    "fmt"
)

func access_pointer(x *int) {
    fmt.Printf("Pointer points to %d\n", *x)
    *x++
}
// Go has pointers, like in C, but no pointer arithmetic
func fun_with_pointers() {
    var x int;
    var ptr *int;
    x = 7
    ptr = &x
    access_pointer(ptr)
    access_pointer(ptr)
    fmt.Printf("Now the value of x is %d\n", x)
}

// Go also has structs
// Create a 2D vertex
type Vertex struct {
    // Also works with X, Y int, but I think this is cleaner notation
    X int
    Y int
}
// A triangle is composed of three vertices
type Triangle struct {
    A Vertex
    B Vertex
    C Vertex
}
// Initialize a triangle from three vertices
func init_triangle(A, B, C Vertex) Triangle {
    // Initialize structs using curly braces
    return Triangle{A, B, C}
}
// Initialize vertex from 2D coordinates
func init_vertex(x, y int) Vertex {
    return Vertex{x, y}
}
// Create an arbitrary triangle
func create_triangle() Triangle {
    v1 := init_vertex(0, 1)
    v2 := init_vertex(1, 0)
    v3 := init_vertex(0, 0)
    t1 := init_triangle(v1, v2, v3)
    fmt.Println(t1)
    // Can access struct fields with a dot
    fmt.Println(t1.A)
    return t1
}

// Pointers to struct members are implicitly dereferenced...
func struct_implicit_deref(tri Triangle) {
    vert_ptr := &tri.A
    fmt.Printf("X coordinate of vertex A is %d\n", vert_ptr.X)
    // Can also explicitly dereference for clarity
    fmt.Printf("Y coordinate of vertex A is %d\n", (*vert_ptr).Y)
}

// Names of struct members can be accessed explicitly
func struct_literals() {
    var (
        v1 = Vertex{1, 2}
        v2 = Vertex{X: 1}  // Y:0 implicit
        v3 = Vertex{}  // X:0, Y:0 implicit
        p = &Vertex{Y:5}  // has type *Vertex
    )
    fmt.Println(v1, v2, v3, p)
}

// Arrays of lenth n and type T are initialized as [n]T
func create_array() {
    var arr [3]int
    arr[0] = 0
    arr[1] = 1
    arr[2] = 2
    for i := 0; i < 3; i++ {
        fmt.Println(arr[i])
    }
    // You can initialize an array in one line
    arr2 := [3]int{2, 1, 0}
    for i := 0; i < 3; i++ {
        fmt.Println(arr2[i])
    }
}

// Takes array of length 5, returns slice of first n elements
func first_n_of_five(arr [5]int, n int) []int {
    // If n is out of range, return a slice of full array
    if n > 5 || n < 0 {
        fmt.Println("n must be between 0 and 5. Full array:")
        return arr[:]
    }
    // Look at slice of array
    return arr[0:n]
}

// Slices do not store data, only references data in underlying array
func fun_with_slices() {
    tmp := [5]int{0, 1, 2, 3, 4}
    fmt.Println(first_n_of_five(tmp, 2))
    fmt.Println(first_n_of_five(tmp, 6))

    // Modifying a slice modifies the underlying data
    // Zero out first three values of tmp array
    slc := tmp[0:3]
    for i := 0; i < 3; i++ {
        slc[i] = 0
    }
    fmt.Println(tmp)
    // Note that this does not work for slc := first_n_of_five(tmp, 3)
    // It seems this returns an array and not a true slice?
}

// Arrays can be built and returned as a slice that references it in one go
func slice_literals() {
    tmp := []bool{true, true, false}
    fmt.Println(tmp)
    tmp[0] = false
    fmt.Println(tmp)
}

func main() {
    fun_with_pointers()
    tri := create_triangle()
    struct_implicit_deref(tri)
    struct_literals()
    create_array()
    fun_with_slices()
    slice_literals()
}
