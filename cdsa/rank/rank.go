package rank

var _ TopNWithUpdate = (*TopNCache)(nil)

type TopNWithUpdate interface {
	SetScore(id string, newScore int)
	GetTopN() []RankEntry
	Clear()
}

type RankEntry struct {
	Id    string
	Score int
}

// TopNWithUpdate is extremely optimized for small N(<=50) with only incremented scores
// Time: O(N), Space: O(N), where N is the capacity of the rank list
type TopNCache struct {
	topN     []RankEntry
	lowBound int
}

func NewTopNCache(cap int) *TopNCache {
	return &TopNCache{
		topN:     make([]RankEntry, 0, cap),
		lowBound: 0,
	}
}

func (c *TopNCache) SetScore(Id string, score int) {
	if len(c.topN) == cap(c.topN) && score <= c.lowBound {
		return
	}

	findIdx := -1
	for i, entry := range c.topN {
		if entry.Id == Id {
			findIdx = i
			break
		}
	}

	if findIdx != -1 {
		c.topN[findIdx].Score = score
		for ; findIdx > 0 && c.topN[findIdx].Score > c.topN[findIdx-1].Score; findIdx-- {
			c.topN[findIdx], c.topN[findIdx-1] = c.topN[findIdx-1], c.topN[findIdx]
		}
	} else {
		c.addEntry(Id, score)
	}

	c.lowBound = c.topN[len(c.topN)-1].Score
}

func (c *TopNCache) addEntry(id string, score int) {
	newEntry := RankEntry{Id: id, Score: score}

	if len(c.topN) < cap(c.topN) {
		c.topN = append(c.topN, newEntry)

		for i := len(c.topN) - 1; i > 0 && c.topN[i].Score > c.topN[i-1].Score; i-- {
			c.topN[i], c.topN[i-1] = c.topN[i-1], c.topN[i]
		}

		return
	}

	i := cap(c.topN) - 2
	for i >= 0 && c.topN[i].Score < score {
		c.topN[i+1] = c.topN[i]
		i--
	}

	c.topN[i+1] = newEntry
}

func (c *TopNCache) GetTopN() []RankEntry {
	result := make([]RankEntry, len(c.topN))
	copy(result, c.topN)
	return result
}

func (c *TopNCache) Clear() {
	c.topN = c.topN[:0]
	c.lowBound = 0
}
