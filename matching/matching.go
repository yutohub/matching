package matching

func GaleShapley(msPrefers, wsPrefers [][]int) map[int]int {
	// マッチングしていない男性のリストを作成
	msUnpaired := make([]int, len(msPrefers))
	for i := 0; i < len(msPrefers); i++ {
		msUnpaired[i] = i
	}
	// ハッシュマップ{女性: 仮マッチングしている男性}
	wsPartner := make(map[int]int, len(wsPrefers))
	// マッチングしていない男性がいる限り
	for len(msUnpaired) > 0 {
		// マッチングしていない男性を一人取り出す
		msTarget := msUnpaired[0]
		msUnpaired = msUnpaired[1:]
		// 男性の候補が残っていなければ次に行く
		if len(msPrefers[msTarget]) == 0 {
			continue
		}
		// 男性の候補の女性を取り出す
		wsTarget := msPrefers[msTarget][0]
		msPrefers[msTarget] = msPrefers[msTarget][1:]
		// 男性が候補の女性の希望リストに入っているか調べる
		msTargetRank, ok := targetRank(wsPrefers[wsTarget], msTarget)
		// 候補の女性の希望リストに入っていないなら、マッチングしていない男性のリストに戻す
		if !ok {
			msUnpaired = append(msUnpaired, msTarget)
			continue
		}
		// 候補の女性が仮マッチングしているか調べる
		msMatched, ok := wsPartner[wsTarget]
		// 候補の女性が仮マッチングしていない場合、マッチングさせる
		if !ok {
			wsPartner[wsTarget] = msTarget
			continue
		}
		// 候補の女性が仮マッチングしている場合、何番目か調べる
		msMatchedRank, _ := targetRank(wsPrefers[wsTarget], msMatched)
		// その男性と仮マッチングしている相手を比較する
		// もし男性の方が今マッチングしている相手より上の場合
		// マッチングを更新し、今マッチングしている相手をマッチングしていない男性のリストに戻す
		if msTargetRank < msMatchedRank {
			wsPartner[wsTarget] = msTarget
			msUnpaired = append(msUnpaired, msMatched)
		} else {
			// そうでなければ、マッチングしていない男性のリストに戻す
			msUnpaired = append(msUnpaired, msTarget)
		}
	}
	return wsPartner
}

func targetRank(prefers []int, target int) (int, bool) {
	for i, p := range prefers {
		if target == p {
			return i, true
		}
	}
	return 0, false
}
