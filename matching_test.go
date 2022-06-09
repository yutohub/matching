package matching_test

import (
	"fmt"
	"testing"

	"github.com/yutohub/matching"
)

func TestGS4x4(t *testing.T) {
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

func TestGS3x3(t *testing.T) {
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

func TestGSLessMsPrefers(t *testing.T) {
	msPrefers := [][]int{}
	wsPrefers := [][]int{{0}, {1}, {3}, {2}}
	result, err := matching.GaleShapley(msPrefers, wsPrefers)
	fmt.Printf("TestGSLessMsPrefers: %s\n", err)
	if result != nil {
		t.Error("ERROR: Expect an error and no result returned")
	}
	if err == nil {
		t.Error("ERROR: Expect error occurred")
	}
}

func TestGSLessWsPrefers(t *testing.T) {
	msPrefers := [][]int{{0}, {1}, {3}, {2}}
	wsPrefers := [][]int{}
	result, err := matching.GaleShapley(msPrefers, wsPrefers)
	fmt.Printf("TestGSLessWsPrefers: %s\n", err)
	if result != nil {
		t.Error("ERROR: Expect an error and no result returned")
	}
	if err == nil {
		t.Error("ERROR: Expect error occurred")
	}
}
