package rdssortedset

import (
	"github.com/skyrocket-qy/gox/dsa/skiplist"
)

// Member represents a member in the sorted set.
type Member struct {
	Score  float64
	Member string
}

// SortedSet represents a Redis-like sorted set.
type SortedSet struct {
	// members maps member strings to their scores.
	members map[string]float64
	// skiplist stores members sorted by score.
	skiplist *skiplist.SkipList
}

func (ss *SortedSet) ZRangeByScore(min, max float64) []string {
	minMember := Member{Score: min, Member: ""}
	maxMember := Member{Score: max, Member: "\xff"}

	rawMembers := ss.skiplist.GetRangeByValue(minMember, maxMember)

	members := []string{}
	for _, val := range rawMembers {
		members = append(members, val.(Member).Member)
	}

	return members
}

func (ss *SortedSet) ZRange(start, stop int) []string {
	length := ss.skiplist.Len()

	// Handle negative indices
	if start < 0 {
		start = length + start
	}
	if stop < 0 {
		stop = length + stop
	}

	// Clamp indices
	if start < 0 {
		start = 0
	}
	if stop >= length {
		stop = length - 1
	}

	if start > stop || start >= length {
		return []string{}
	}

	var members []string
	for i := start; i <= stop; i++ {
		val := ss.skiplist.GetByIndex(i)
		if val != nil {
			members = append(members, val.(Member).Member)
		}
	}

	return members
}

func (ss *SortedSet) ZCard() int {
	return ss.skiplist.Len()
}

func (ss *SortedSet) ZScore(member string) (float64, bool) {
	score, exists := ss.members[member]
	return score, exists
}

func (ss *SortedSet) ZRem(member string) bool {
	oldScore, exists := ss.members[member]
	if !exists {
		return false
	}

	delete(ss.members, member)
	ss.skiplist.Delete(Member{Score: oldScore, Member: member})
	return true
}

func (ss *SortedSet) ZAdd(score float64, member string) {
	oldScore, exists := ss.members[member]

	if exists {
		if oldScore == score {
			return // Score hasn't changed, nothing to do.
		}
		// Remove old entry from skiplist before updating.
		ss.skiplist.Delete(Member{Score: oldScore, Member: member})
	}

	ss.members[member] = score
	ss.skiplist.Insert(Member{Score: score, Member: member})
}

func memberComparator(a, b interface{}) int {
	mA := a.(Member)
	mB := b.(Member)

	if mA.Score < mB.Score {
		return -1
	} else if mA.Score > mB.Score {
		return 1
	} else {
		// If scores are equal, compare by member string lexicographically
		if mA.Member < mB.Member {
			return -1
		} else if mA.Member > mB.Member {
			return 1
		}
		return 0
	}
}

// New creates a new SortedSet.
func New() *SortedSet {
	return &SortedSet{
		members:  make(map[string]float64),
		skiplist: skiplist.New(memberComparator),
	}
}
