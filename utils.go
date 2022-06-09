package matching

import "fmt"

func Show(result map[int]int) {
	for ws := range result {
		ms := result[ws]
		fmt.Printf("m%d <-> w%d\n", ms, ws)
	}
}
