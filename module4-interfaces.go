package main

import (
    "fmt"
    "math"
)

// This is an interface for the Abs (absolute value) method. The value of an
// interface can be any type that implements the methods of the interface.
type Abser interface {
    Abs() float64
}

// Create a new type MyFloat that implements a method called Abs
type MyFloat float64
func (f MyFloat) Abs() float64 {
    if f < 0 {
        return float64(-f)
    }
    return float64(f)
}

// Create a Vertex struct that will also implement a method Abs
type Vertex struct {
    X, Y float64
}

// The Abs method will require a Vertex pointer
func (v *Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Note: nothing about Vertex or MyFloat explicitly suggests that they
// implement the Abser interface. But because they implement the methods of the
// interface, they implicitly implement the interface also. Therefore the
// interface Abser does not need to know ahead of time how its methods will be
// implemented by the types that will use the interface.
//
// This means that if a package relies on an interface (like Abser), I can
// easily create a type that the package can use by implementing the methods of
// the interface on my data type.

func demo1() {
    // Initialize the interface for Abs
    var a Abser
    f := MyFloat(-math.Sqrt2)
    v := Vertex{3, 4}

    a = f  // Since MyFloat implements Abs, an Abser can be assigned to a MyFloat type
    fmt.Println(a.Abs())
    a = &v  // Same, since a Vertex *pointer* implements Abs
    fmt.Println(a.Abs())

    // a = v  // Does not compile! Vertex does NOT implement Abs, only the
    // pointer does
}

// Under the hood, the interface can be thought of as storing a tuple (value,
// type). Calling the method of the interface calls the method on its assigned
// value corresponding to its underlying type. E.g. 
// This interface implements methods M
type I interface {
    M() float64
}
// Create type T that implements method M
type T struct {
    S float64
}
func (t T) M() float64 {
    return t.S
}
func demo2() {
    var intfc I
    ttype := T{S: 10}
    fmt.Printf("%v, %T\n", intfc, intfc)
    intfc = ttype
    // Interface instance intfc has stored the underlying value, type pair
    // (ttype, T), so that calling M() on intfc calls the T.M method on ttype
    fmt.Printf("%v, %T\n", intfc, intfc)
    fmt.Println(intfc.M())
}


type I2 interface {
    N() float64
}
// It is common to implement a nil-handler for methods in Go
func (t *T) N() float64 {
    if t == nil {
        return -1
    }
    return t.S
}
func demo3() {
    var tptr *T
    var intfc I2
    intfc = tptr
    fmt.Printf("%v, %T\n", intfc, intfc)
    fmt.Println(intfc.N())
    ttype := T{S: 6}
    tptr = &ttype
    intfc = tptr
    fmt.Printf("%v, %T\n", intfc, intfc)
    fmt.Println(intfc.N())
}

// An empty interface can have any type, so a function that takes an empty
// interface can handle an argument of unknown type
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
func demo4() {
    var intfc interface{}
    describe(intfc)
    intfc = 42
    describe(intfc)
    intfc = "iamastring"
    describe(intfc)
}

// Type assertions
type MyType string
func demo5() {
    var intfc interface{}
    var mt MyType
    mt = "test"
    intfc = mt
    // One can assert that the underlying value of an interface is of a
    // specific type using the following syntax:
    describe(intfc)
    t := intfc.(MyType)
    describe(t)
    // If intfc were not of MyType, it would trigger a panic. To catch a panic,
    // the following syntax can be used to receive a boolean that is false if
    // the type does not match
    t, ok := intfc.(MyType)
    fmt.Printf("%v, %v\n", t, ok)
    // If the type assertion fails, the value is assigned the default (zero)
    mf, ok := intfc.(MyFloat)
    fmt.Printf("%v, %v\n", mf, ok)
}

// Type switches
func my_type_switch(intfc interface{}) {
    // Using this syntax, we can implement different code for different types
    switch v := intfc.(type) {
    case MyType:
        // v has type MyType
        fmt.Printf("%v, %T\n", v, v)
    case MyFloat:
        // v has type MyFloat
        fmt.Printf("%v, %T\n", v, v)
    default:
        // v has a different type
        fmt.Printf("my_type_switch has no implementation for type %T", v)
    }
}
func demo6() {
    var mt MyType = "wat"
    var mf MyFloat = 42.666
    var x int32
    my_type_switch(mt)
    my_type_switch(mf)
    my_type_switch(x)
}

// One common interface is the Stringer interface that implements a String
// method, and is used in many places, e.g. fmt.Println:
// Create a new struct called Person
type Person struct {
	Name string
	Age  int
}
// Implement the String method for Person
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}
// Println uses the Stringer interface to display the String method of Person
func demo7() {
	a := Person{"Bobby G", 42}
	fmt.Println(a)
	z := Person{"Poppy Z", 26}
	fmt.Println(z)
}

func main() {
    fmt.Println("Beginning demo1")
    demo1()
    fmt.Println("Beginning demo2")
    demo2()
    fmt.Println("Beginning demo3")
    demo3()
    fmt.Println("Beginning demo4")
    demo4()
    fmt.Println("Beginning demo5")
    demo5()
    fmt.Println("Beginning demo6")
    demo6()
    fmt.Println("Beginning demo7")
    demo7()
}

