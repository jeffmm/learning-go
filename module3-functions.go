package main

import (
    "fmt"
    "math"
)

// We pass a function to a function
func compute(fn func(float64, float64) float64) float64 {
    return fn(3, 4)
}

// Functions can be passed as variables to other functions
func func_as_variable() {
    hypot := func(x, y float64) float64 {
        return math.Sqrt(x*x + y*y)
    }
    fmt.Println(hypot(5, 12))
    fmt.Println(compute(hypot))
    fmt.Println(compute(math.Pow))
}

// This is a func closure. A closure is a function value that references variables from 
// outside its scope
func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}

// Demonstrates how adder works. It acts similar to a data struct without a
// getter/setter method
func demo_adder() {
    pos, neg := adder(), adder()
    for i := 0; i< 10; i++ {
        fmt.Println(
            pos(i),
            neg(-2*i),
        )
    }
}

// To wrap my head around this I'll make a Fibonacci sequence generator using closures
// ...
// Turns out this was the next exercise, just slightly different such that the closure
// takes an integer to print the n'th value in the Fibo sequence
func fibo() func() int {
    prev := 0
    this := 0
    return func() int {
        // I want to see 1 1 2 etc so I know right away it's a fibo sequence
        if this == 0 {
            this += 1
            return this
        }
        next := this + prev
        prev = this
        this = next
        return next
    }
}

// Print first 10 values of Fibonacci sequence
func demo_fibo() {
    fib := fibo()
    for i := 0; i < 10; i++ {
        fmt.Println(fib())
    }
}

func main() {
    func_as_variable()
    demo_adder()
    demo_fibo()
}
