// author: @UchihaSusie
// algorithms/Go/dynamic_programming/catalan_sequence.go
// Catalan number

package main

import "fmt"

// catalan function to calculate the nth Catalan number
func catalan(n int) int {
    // Base Case
    if n == 0 || n == 1 {
        return 1
    }

    // To store the result of subproblems
    catNum := make([]int, n+1)

    catNum[0] = 1
    catNum[1] = 1

    for i := 2; i <= n; i++ {
        for j := 0; j < i; j++ {
            catNum[i] += catNum[j] * catNum[i-j-1]
        }
    }
    return catNum[n]
}

func main() {
    n := 10

    if n < 0 {
        fmt.Println("Please add a valid number")
        return
    }

    fmt.Printf("The first %d terms of Catalan sequence are : ", n)
    for i := 0; i < n; i++ {
        fmt.Print(catalan(i), " ")
    }
}