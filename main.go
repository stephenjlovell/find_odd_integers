
// This algorithms practice problem compares a few different ways to solve the following problem:
// - Given a set s of integers, find the subset of integers that occur an odd number of times in s.
package main

import (
  "math/rand"
  "fmt"
  "github.com/stephenjlovell/find_odd_integers/with_map"
)


func main() {
  randoms := generate_randoms(10)
  fmt.Println(randoms)
  results := with_map.FindOddInts(randoms)
  fmt.Println(results)
}

func generate_randoms(n int) []int {
  randoms := make([]int, n, n)
  for i := 0; i < n; i++ {
    randoms[i] = rand.Intn(11)
  }
  return randoms
}
