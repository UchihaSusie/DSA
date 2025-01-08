// author: @UchihaSusie
// algorithms/Go/dynamic_programming/levenshtein_distance.go
package main

import (
    "fmt"
)

// calculate Levenshtein distance
func levenshteinDistance(word1 string, chars1 int, word2 string, chars2 int) int {
    // base case: if string is empty
    if chars1 == 0 {
        return chars2
    }
    if chars2 == 0 {
        return chars1
    }

    //If the last character of the string matches, the operation cost is 0
    cost := 1
    if word1[chars1-1] == word2[chars2-1] {
        cost = 0
    }

    // recursive calculation of operations
    deletion := levenshteinDistance(word1, chars1-1, word2, chars2) + 1
    insertion := levenshteinDistance(word1, chars1, word2, chars2-1) + 1
    substitution := levenshteinDistance(word1, chars1-1, word2, chars2-1) + cost

    // return the minimum number of operations
    if deletion < insertion {
        if deletion < substitution {
            return deletion
        }
        return substitution
    }
    if insertion < substitution {
        return insertion
    }
    return substitution
}

// main driver program
func main() {
    word1 := "plain"
    word2 := "plane"

    fmt.Println("The Levenshtein distance is:")
    fmt.Println(levenshteinDistance(word1, len(word1), word2, len(word2)))
}