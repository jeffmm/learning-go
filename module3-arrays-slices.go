package main

import(
    "fmt"
)

// By the way: unlike C, functions don't have to be defined before they appear
// in the source file
func main() {
    create_array()
    fun_with_slices()
    slice_literals()
    slice_defaults()
    length_cap()
    nil_slice()
    make_slice()
    array_2d()
    slice_append()
    slice_range()
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

// The usual slice defaults apply here
func slice_defaults() {
    fib := []int{1, 1, 2, 3, 5, 8, 13}
    fmt.Println(fib[0:7])
    fmt.Println(fib[0:])
    fmt.Println(fib[:7])
    fmt.Println(fib[:])
}

// Print length and capacity of slice
func print_slice(s []int) {
    fmt.Printf("%v: len=%d, cap=%d\n", s, len(s), cap(s))
}

// Slices have length and capacity: length is the number of elements in the slice,
// capacity is the size of the underlying array being referenced
func length_cap() {
    arr := []int{1, 1, 2, 3, 5, 8, 13}
    print_slice(arr)
    // Give the slice zero length
    s := arr[:0] // len=0, cap=7
    print_slice(s)
    // Extend its length by reslicing it
    s = s[:4] // len=4, cap=7
    print_slice(s)
    // Drop its first two values
    s = s[2:] // len=2, cap=5
    print_slice(s)
    // Fails with runtime error, since capacity is now 5, so index is out of bounds!
    // s = s[:7]
    // print_slice(s)
}

// Creates a nil slice
func nil_slice() {
    // s := []int  // this assignment convention does not work!
    var s []int  // nil slice: length of zero, capacity of zero
    print_slice(s)
    if s == nil {
        fmt.Println("Slice is nil!")
    }
}

// Dynamic arrays can be created with make
func make_slice() {
    // A single numerical argument to make specifies both length and capacity
    // Create a zeroed slice of an array of length 5
    a := make([]int, 5)  // len=5, cap=5
    print_slice(a)
    // Providing a second numerical argument specifies just the capacity
    // Create a zero-length slice of an array
    b := make([]int, 0, 7)  // len=0, cap=7
    print_slice(b)
    b = b[:cap(b)]  // len=7, cap=7 (zeroed)
    print_slice(b)
    b = b[1:] // len=6, cap=6 (zeroed)
    print_slice(b)
}

// Slices of  slices can be used to create 2D arrays
func array_2d() {
    matrix := [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            fmt.Printf("%d ", matrix[i][j])
        }
        fmt.Println()
    }
}


// Append adds values to a slice, and can dynamically reallocate if necessary
func slice_append() {
    s := make([]int, 0, 7)
    print_slice(s)
    s = append(s, 1, 1, 2)
    print_slice(s)
    s = append(s, 3, 5, 8)
    print_slice(s)
    // Appending values beyond the capacity: array is dynamically reallocated
    s = append(s, 13, 21, 34)
    print_slice(s)
}

// Range can iterate over the index and value of a slice
func slice_range() {
    s := []int{1, 1, 2, 3, 5, 8}
    print_slice(s)
    for i, v := range s {
        fmt.Printf("index: %d, val: %d\n", i, v)
    }
    // Use underscore to omit the assignment
    fmt.Println("Just the indices:")
    for i, _ := range s {
        fmt.Printf("%d ", i)
    }
    fmt.Println()
    fmt.Println("Just the values now:")
    for _, v := range s {
        fmt.Printf("%d ", v)
    }
    fmt.Println()
    // By default, only indices are captured if one variable is given
    fmt.Println("Just the indices again")
    for i := range s {
        fmt.Printf("%d ", i)
    }
    fmt.Println()
}

// func draw_pic() {

// }
