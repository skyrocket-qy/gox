package rank

import (
	"sort"
)

type RankEntry struct {
	PlayerID string
	Score    int
}

// TopNCache uses on the scenario which the score will only increased
type TopNCache struct {
	topN       []RankEntry
	lowerBound int
	n          int
}

func NewTopNCache(n int) *TopNCache {
	return &TopNCache{
		topN:       make([]RankEntry, 0, n),
		lowerBound: 0,
		n:          n,
	}
}

func (c *TopNCache) SetScore(playerID string, newScore int) {
	playerIndexInTopN := -1
	for i, entry := range c.topN {
		if entry.PlayerID == playerID {
			playerIndexInTopN = i
			break
		}
	}

	if playerIndexInTopN == -1 {
		return
	}

	if playerIndexInTopN != -1 {
		c.topN[playerIndexInTopN].Score = newScore
	} else {
		c.topN = append(c.topN, RankEntry{PlayerID: playerID, Score: newScore})
	}

	sort.Slice(c.topN, func(i, j int) bool {
		return c.topN[i].Score > c.topN[j].Score
	})

	if len(c.topN) > c.n {
		c.topN = c.topN[:c.n]
	}

	c.lowerBound = c.topN[len(c.topN)-1].Score
}

func (c *TopNCache) GetTopN() []RankEntry {
	result := make([]RankEntry, len(c.topN))
	copy(result, c.topN)
	return result
}

func (c *TopNCache) Clear() {
	c.topN = make([]RankEntry, 0, c.n)
	c.lowerBound = 0
}
