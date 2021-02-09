package solution

// you can also use imports, for example:
// import "fmt"
// import "os"
import "math"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(X int, Y int, D int) int {
    // write your code in Go 1.4
    //X - Current Point
    //Y - Arrival point
    //D - Jump distane
    //Number of Jumps?
    numberOfJumps := int(math.Ceil((float64(Y)-float64(X))/float64(D)))
    return numberOfJumps
}
