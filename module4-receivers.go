package main

// Module 4: Methods and interfaces

import (
    "fmt"
)

// Go does not have classes, only structs. Although, we can build class-like methods on
// structs using a receiver argument on a function

type Foo struct {
    bar, baz int
}
// Here we implement a Diff method for the Foo type: (f Foo) is the receiver argument
func (f Foo) Diff() int {
    return f.bar - f.baz
}
// Print difference of Foo values
func demo_receiver_arguments() {
    f := Foo{10, 6}
    fmt.Println(f.Diff())
}


// But methods are just functions so we could have done this as:
func GetDifference(f Foo) int {
    return f.bar - f.baz
}
func demo_just_funcs() {
    f := Foo{12, 9}
    fmt.Println(GetDifference(f))
}

// We can use receiver arguments on any type that is defined in this package (not a 
// built-in or defined in another package). E.g.:
type MyInt int
func (mi MyInt) Show() {
    fmt.Println(mi)
}
    
func demo_receiver_non_struct() {
    m := MyInt(42)
    m.Show()
}

// In order to change the values of members in a struct, we need to pass a pointer
type Bar struct {
    biz, bang int
}
func (b Bar) ResizeNoPtr(factor int) {
    b.biz *= factor
    b.bang *= factor
}
// Passing Bar pointer to create lasting changes (also avoids copying, which speeds
// things up)
func (b *Bar) Resize(factor int) {
    // Once again, Go implicitly dereferences
    b.biz *= factor
    b.bang *= factor
}
func demo_resize() {
    b := Bar{3, 4}
    fmt.Println(b)
    b.ResizeNoPtr(10)
    fmt.Println(b)
    b.Resize(10)
    fmt.Println(b)
}

// Note that we did NOT need to do (&b).Resize(10). Again, the value is implicitly
// dereferenced for convenience. But this is NOT true for normal functions.
func ResizeBar(b *Bar, factor int) {
    b.biz *= factor
    b.bang *= factor
}
func demo_auto_deref_methods() {
    b := Bar{6, 10}
    fmt.Println(b)
    // ResizeBar(b, 10) //  Doesn't compile! Wrong type.
    ResizeBar(&b, 10) //  OK
    fmt.Println(b)

    // Whereas all of this is OK
    b2 := Bar{9, 5}
    fmt.Println(b2)
    b2.Resize(2)
    fmt.Println(b2)
    (&b2).Resize(2)
    fmt.Println(b2)
}

func main() {
    demo_receiver_arguments()
    demo_just_funcs()
    demo_receiver_non_struct()
    demo_resize()
    demo_auto_deref_methods()
}
