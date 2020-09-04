// Always needed in main
package main

// Multiple module import
import (
    "fmt"
    "math"
    "math/cmplx"
)

// bar is its own func, type declarations come after variables
func foo(bar string) string {
    // string appending, implicit typing
    baz := bar + "ooyah!"
    // explicit return call
    return baz
}

func bar(baz float64) (buz int) {
    // Explicit type conversions necessary!
    buz = int(baz) - 2
    // Naked return statement
    return
}

func main() {
    // Print elements with added spaces
    fmt.Println("Hello", "World!")
    // Explicitly declare variable and type
    var i int
    // Multiple variables
    var alpha, omega bool
    // Two ways to declare implicitly-typed variables
    var x = "foo"
    y := "bar"
    alpha, omega = true, false
    // Complex types
    var num int = 5
    // Must always explicitly cast types to functions
    sqrtNum := math.Sqrt(float64(num*num))
    cmplxNum := cmplx.Sqrt((5 + 2i))
    // Print on its own line
    fmt.Println("Here is a line without explicit newline char")
    // C-style printf statements
    fmt.Printf("Here is a C-style printf statement with explicit newline\n")
    fmt.Printf("Type : %T, val : %v\n", i, i)
    fmt.Printf("Type : %T, val : %v\n", x, x)
    fmt.Printf("Type : %T, val : %v\n", alpha, alpha)
    fmt.Printf("Type : %T, val : %v\n", omega, omega)
    fmt.Printf("Type : %T, val : %v\n", y, y)
    fmt.Printf("Type : %T, val : %v\n", sqrtNum, sqrtNum)
    fmt.Printf("Type : %T, val : %v\n", cmplxNum, cmplxNum)
    fmt.Println(foo("B"))
    fmt.Println(bar(sqrtNum))
}
