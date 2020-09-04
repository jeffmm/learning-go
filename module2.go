package main

import (
    "fmt"
    "math"
)

// For loop (like C but no parentheses)
func sum_1_to_x(x int) {
    sum := 0
    for i := 1; i < x+1; i++  {
        sum += i
    }
    fmt.Println(sum)
}

// Init and post statements are not necessary
// Similar to a while sum < x: etc
func smallest_power_of_two_greater_than_x(x int) {
    sum := 1
    for ; sum < x+1; {
        sum += sum
    }
    fmt.Println(sum)
}

// In fact, while is spelled for in go; you can just drop the semicolons
func smallest_power_of_two_greater_than_x_while(x int) {
    sum := 1
    for sum < x + 1 {
        sum += sum
    }
    fmt.Println(sum)
}

// An infinite loop in go:
func infinite_loop() {
    for {
    }
}

// if statements are C-like, but without parentheses
func sqrt_real(x float64) string {
    if x < 0 {
        return sqrt_real(-x) + "i"
    }
    return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
    // Assign a local variable v
    if v := math.Pow(x, n); v < lim {
        return v
    }
    //return v  // Fails! The scope of v is only in the if-statement
    return lim
}

func pow_again(x, n, lim float64) float64 {
    // Create local variable v again
    if v:= math.Pow(x, n); v < lim {
        return v
    } else {
        // The scope of v extends to the if statement
        fmt.Printf("%g >= %g\n", v, lim)
    }
    // return v  // Fails again!
    return lim
}

// Approximate sqrt(x) within 1% of the precision of x
func my_sqrt(x, alpha, epsilon float64) float64 {
    // alpha is a learning rate, epsilon an error margin
    z := 1.0  // could also init with float64(1)
    for delta := z*z - x; delta*delta > epsilon*epsilon; delta = z*z-x {
        z -= alpha * delta
    }
    return z
}

func main() {
    sum_1_to_x(10)
    smallest_power_of_two_greater_than_x(10)
    smallest_power_of_two_greater_than_x_while(10)
    // infinite_loop()  // an infinite loop: don't uncomment this!
    fmt.Println(sqrt_real(9), sqrt_real(-9))
    fmt.Println(
        pow(3, 2, 10),
        pow(3, 3, 10),
    )
    fmt.Println(
        pow_again(3, 2, 10),
        pow_again(3, 3, 10),
    )
    fmt.Println(my_sqrt(2, 0.1, 0.0001))
}
