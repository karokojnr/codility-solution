/*
An array A consisting of N integers is given. Rotation of the array means that each element is
shifted right by one index, and the last element of the array is also moved to the first place.
For example, the rotation of array A = [3, 8, 9, 7, 6] is [6, 3, 8, 9, 7].
The goal is to rotate array A K times; that is, each element of A will be shifted to the right by K indexes.
Write a function:
func Solution(A []int, K int) []int
that, given a zero-indexed array A consisting of N integers and an integer K, returns the array A rotated K times.
For example, given array A = [3, 8, 9, 7, 6] and K = 3, the function should return [9, 7, 6, 3, 8].
*/
package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int, K int) []int {
    // write your code in Go 1.4
    // Last index to be First index: 
    // :Upto to BUT excluding  last index (unfurl)...
    if K > 0 && len(A) > 0{
        A = append(A[len(A) - K:], A[0:len(A)-K]...)
    }
    return A
}
