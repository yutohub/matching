package matching_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/yutohub/matching"
)

func TestGaleShapley4x4(t *testing.T) {
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
	answer := map[int]int{0: 0, 1: 2, 2: 3, 3: 1}
	result, err := matching.GaleShapley(msPrefers, wsPrefers)
	matching.Show(result)
	if err != nil {
		t.Error(err)
	}
	for i := range answer {
		if answer[i] != result[i] {
			t.Errorf("ERROR: m%v <-> w%v is not collect", result[i], i)
		}
	}
}

func TestGaleShapley4x3(t *testing.T) {
	msPrefers := [][]int{
		{2, 3, 1},
		{2, 1, 3},
		{0, 3, 2},
		{0, 2, 3},
	}
	wsPrefers := [][]int{
		{1, 3, 0},
		{3, 1, 0},
		{1, 2, 3},
		{0, 1, 2},
	}
	answer := map[int]int{0: 3, 2: 1, 3: 0}
	result, err := matching.GaleShapley(msPrefers, wsPrefers)
	matching.Show(result)
	if err != nil {
		t.Error(err)
	}
	for i := range answer {
		if answer[i] != result[i] {
			t.Errorf("ERROR: m%v <-> w%v is not collect", result[i], i)
		}
	}
}

func TestGaleShapleyLackingMsPrefers(t *testing.T) {
	msPrefers := [][]int{}
	wsPrefers := [][]int{{0}, {1}, {3}, {2}}
	result, err := matching.GaleShapley(msPrefers, wsPrefers)
	if result != nil {
		t.Error("ERROR: Expect an error and no result returned")
	}
	if errors.Is(err, matching.ErrorLackingPrefers) {
		fmt.Printf("TestGaleShapleyLackingMsPrefers: %s\n", err)
	} else {
		t.Error("ERROR: Expect error occurred")
	}
}

func TestGaleShapleyLackingWsPrefers(t *testing.T) {
	msPrefers := [][]int{{0}, {1}, {3}, {2}}
	wsPrefers := [][]int{}
	result, err := matching.GaleShapley(msPrefers, wsPrefers)
	if result != nil {
		t.Error("ERROR: Expect an error and no result returned")
	}
	if errors.Is(err, matching.ErrorLackingPrefers) {
		fmt.Printf("TestGaleShapleyLackingWsPrefers: %s\n", err)
	} else {
		t.Error("ERROR: Expect error occurred")
	}
}
