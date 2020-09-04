package main

import (
    "fmt"
    "math"
    "runtime"
    "time"
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

// helper function for Newton's method 
func delta(z, x float64) float64 {
    return (z*z - x) / (2 * z)
}

// Approximate sqrt(x) within 1% of the precision of x
func my_sqrt(x, epsilon float64) float64 {
    // alpha is a learning rate, epsilon an error margin
    z := 1.0  // could also init with float64(1)
    eps2 := epsilon * epsilon;
    for d := delta(z, x); d*d > eps2; d = delta(z, x) {
        z -= d
    }
    return z
}

// Check which OS we are running on
func check_os() {
    fmt.Print("Go runs on ")
    // The switch statement does not need to call break
    switch os := runtime.GOOS; os {
    case "darwin":
        fmt.Println("OS X.")
    case "linux":
        fmt.Println("Linux.")
    default:
        // freebsd, openbsd, plan9, windows...
        fmt.Printf("%s.\n", os)
    }
}

func should_never_run() time.Weekday {
    fmt.Println("STOP! We should never get here!")
    return time.Monday
}

func run_on_friday() {
    // today := time.Now().Weekday()
    // Fixing to be Friday for this example
    today := time.Friday
    fmt.Print("Saturday is ")
    switch time.Saturday {
    case today:
        fmt.Println("today.")
    case today + 1:
        fmt.Println("tomorrow.")
    // Since today is (always) Friday, this will never evaluate
    case should_never_run():
        fmt.Println("Houston, we have a problem.")
    default:
        fmt.Println("too far away.")
    }
}

func greeting() {
    t := time.Now()
    // Same as switch true {...}
    switch {
    // if this case is true, then... etc
    case t.Hour() < 12:
        fmt.Println("Good morning!")
    // We can do this since cases are evaluated in order
    case t.Hour() < 17:
        fmt.Println("Good afternoon!")
    default:
        fmt.Println("Good evening!")
    }
}

// Example of deferring a function
func deferred_func(z int64) {
    fmt.Printf("deferred_func received %d\n", z)
}
func defer_example() {
    z := int64(42)
    // evaluate, but delay execution until surrounding function returns
    defer deferred_func(z)
    z = 666
    fmt.Printf("The number of the beast is %d\n", z)
}

// Deferred functions are put in a stack: evaluated in order, executed LIFO
func defer_stack() {
    fmt.Println("Counting down")
    defer fmt.Println("Lift off!")
    for i := 0; i < 10; i++ {
        defer fmt.Println(i+1)
    }
    fmt.Println("T-minus:")
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
    fmt.Println(
        "The square root of 2 is approximately",
        my_sqrt(2, 0.0001),
    )
    check_os()
    run_on_friday()
    greeting()
    defer_example()
    defer_stack()
}

