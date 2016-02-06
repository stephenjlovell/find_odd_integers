// This package solves the problem by caching the counts for each number in randoms using a map,
// then filters the results to include only items that appeared an odd number of times.

package with_map

import (
  "fmt"
)

func FindOddInts(randoms []int) []int {
  cache := make(map[int]int)
  for r := range randoms {
    // cache[r] += 1
    last := cache[r]
    cache[r] = last + 1
  }
  fmt.Printf("%v\n", cache)
  results := make([]int, 0, len(randoms))
  for k, v := range cache {
    if isOdd(v) {
      results = append(results, k)
    }
  }
  return results
}

func isOdd(n int) bool {
  return n & 1 != 0
}
