# matching

## Summary of matching package

*In mathematics, economics, and computer science, the Gale-Shapley algorithm (also called the deferred acceptance algorithm or the proposal-rejection algorithm) is an algorithm for finding a solution to a stable matching problem.*  
[Gale–Shapley algorithm From Wikipedia, the free encyclopedia](https://en.wikipedia.org/wiki/Gale%E2%80%93Shapley_algorithm)

In the Gale-Shapley algorithm, all partners must be ranked, but a more practical extension is to allow incomplete selection that only rank some of the partners.

This package extends the Gale-Shapley algorithm to allow not only complete selections but incomplete selections.

## How to use matching package

Check [The Go Playground](https://go.dev/play/p/ZKInZrP8wxV)

```Go
package main

import (
	"fmt"

	"github.com/yutohub/matching"
)

func main() {
	msPrefers := [][]int{
		{0, 1, 2, 3},
		{3, 0, 1, 2},
		{1, 3, 2, 0},
		{0, 2, 3, 1},
	}
	wsPrefers := [][]int{
		{0, 1, 2, 3},
		{1, 2, 3, 0},
		{0, 3, 2, 1},
		{3, 2, 0, 1},
	}
	result, err := matching.GaleShapley(msPrefers, wsPrefers)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("result")
	matching.Show(result)
}
```

```
result
m0 --- w0
m1 --- w3
m2 --- w1
m3 --- w2
```

## Reference

D. Gusfield and R.W. Irving, The Stable Marriage Problems:
Structure and Algorithms, MIT Press, Boston, MA, 1989. 