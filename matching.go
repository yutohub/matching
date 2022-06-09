package matching

import (
	"errors"
	"fmt"
)

// Extend Gale-Shapley algorithm for incomplete lists
// Args: each person's preference order
// Returns: stable matching
func GaleShapley(msPrefers, wsPrefers [][]int) (map[int]int, error) {
	// Check for lack of participants
	if len(msPrefers) == 0 || len(wsPrefers) == 0 {
		return nil, errors.New("GaleShapley: Lack of participants")
	}
	// Create a list of unmatched men
	msUnpaired := make([]int, len(msPrefers))
	for i := 0; i < len(msPrefers); i++ {
		msUnpaired[i] = i
	}
	// Register a Matching
	wsPartner := make(map[int]int, len(wsPrefers))
	// As long as there are men who are not matched
	for len(msUnpaired) > 0 {
		// Pick up one man who has not been matched
		msTarget := msUnpaired[0]
		msUnpaired = msUnpaired[1:]
		// Skip if there are no male preferences left
		if len(msPrefers[msTarget]) == 0 {
			continue
		}
		// Take out the most preferred partner
		wsTarget := msPrefers[msTarget][0]
		msPrefers[msTarget] = msPrefers[msTarget][1:]
		// Check if the target is included in the preferences
		msTargetRank, ok := targetRank(wsPrefers[wsTarget], msTarget)
		// Put msTarget back on the list of unmatched men.
		if !ok {
			msUnpaired = append(msUnpaired, msTarget)
			continue
		}
		// Check if wsTarget is tentatively matched
		msMatched, ok := wsPartner[wsTarget]
		// If msTarget is not tentatively matched, matching
		if !ok {
			wsPartner[wsTarget] = msTarget
			continue
		}
		// If msTarget is not tentatively matched,
		// check the rank of the tentatively matched partner
		msMatchedRank, _ := targetRank(wsPrefers[wsTarget], msMatched)
		// Compare the man with his tentative matching partner
		if msTargetRank < msMatchedRank {
			wsPartner[wsTarget] = msTarget
			msUnpaired = append(msUnpaired, msMatched)
		} else {
			msUnpaired = append(msUnpaired, msTarget)
		}
	}
	return wsPartner, nil
}

// Check if the target is included in the preferences and return target's rank
func targetRank(prefers []int, target int) (int, bool) {
	for i, p := range prefers {
		if target == p {
			return i, true
		}
	}
	return 0, false
}

// Display the results of matching
func Show(result map[int]int) {
	for ws := range result {
		ms := result[ws]
		fmt.Printf("m%d --- w%d\n", ms, ws)
	}
}
