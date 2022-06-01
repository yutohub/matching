package matching_test

import (
	"testing"

	"github.com/yutohub/matching"
)

func TestGaleShapley4x4(t *testing.T) {
	msPrefers := [][]int{
		[]int{0, 1, 2, 3},
		[]int{3, 0, 1, 2},
		[]int{1, 3, 2, 0},
		[]int{0, 2, 3, 1},
	}
	wsPrefers := [][]int{
		[]int{0, 1, 2, 3},
		[]int{1, 2, 3, 0},
		[]int{0, 3, 2, 1},
		[]int{3, 2, 0, 1},
	}
	answer := map[int]int{0: 0, 1: 2, 2: 3, 3: 1}
	result := matching.GaleShapley(msPrefers, wsPrefers)
	for i := range answer {
		if answer[i] != result[i] {
			t.Errorf("ERROR: m%v <-> w%v is not collect", result[i], i)
		}
	}
}

func TestGaleShapley3x3(t *testing.T) {
	msPrefers := [][]int{
		[]int{2, 3, 1},
		[]int{2, 1, 3},
		[]int{0, 3, 2},
		[]int{0, 2, 3},
	}
	wsPrefers := [][]int{
		[]int{1, 3, 0},
		[]int{3, 1, 0},
		[]int{1, 2, 3},
		[]int{0, 1, 2},
	}
	answer := map[int]int{0: 3, 2: 1, 3: 0}
	result := matching.GaleShapley(msPrefers, wsPrefers)
	for i := range answer {
		if answer[i] != result[i] {
			t.Errorf("ERROR: m%v <-> w%v is not collect", result[i], i)
		}
	}
}

func TestGaleShapley2x2(t *testing.T) {
	msPrefers := [][]int{
		[]int{2, 3},
		[]int{2, 1},
		[]int{0, 3},
		[]int{0, 2},
	}
	wsPrefers := [][]int{
		[]int{1, 0},
		[]int{3, 0},
		[]int{1, 3},
		[]int{0, 2},
	}
	answer := map[int]int{2: 1, 3: 0}
	result := matching.GaleShapley(msPrefers, wsPrefers)
	for i := range answer {
		if answer[i] != result[i] {
			t.Errorf("ERROR: m%v <-> w%v is not collect", result[i], i)
		}
	}
}

func TestGaleShapley1x1(t *testing.T) {
	msPrefers := [][]int{
		[]int{0},
		[]int{1},
		[]int{3},
		[]int{0},
	}
	wsPrefers := [][]int{
		[]int{0},
		[]int{1},
		[]int{3},
		[]int{2},
	}
	answer := map[int]int{0: 0, 1: 1, 3: 2}
	result := matching.GaleShapley(msPrefers, wsPrefers)
	for i := range answer {
		if answer[i] != result[i] {
			t.Errorf("ERROR: m%v <-> w%v is not collect", result[i], i)
		}
	}
}
