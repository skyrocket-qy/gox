package randomizedset

/* @tags: set,random */

import (
	"math/rand/v2"
)

type RandomizedSet struct {
	mp map[int]int
	sl []int
}

func New() RandomizedSet {
	return RandomizedSet{
		mp: make(map[int]int),
		sl: make([]int, 0),
	}
}

func (rs *RandomizedSet) Insert(val int) bool {
	if _, ok := rs.mp[val]; ok {
		return false
	}

	rs.sl = append(rs.sl, val)
	rs.mp[val] = len(rs.sl) - 1

	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	if _, ok := rs.mp[val]; !ok {
		return false
	}

	rs.sl[rs.mp[val]], rs.sl[len(rs.sl)-1] = rs.sl[len(rs.sl)-1], rs.sl[rs.mp[val]]

	rs.mp[rs.sl[rs.mp[val]]] = rs.mp[val]

	delete(rs.mp, val)
	rs.sl = rs.sl[:len(rs.sl)-1]

	return true
}

func (rs *RandomizedSet) GetRandom() int {
	// #nosec G404 -- math/rand/v2 is acceptable for random selection, not security-sensitive.
	r := rand.IntN(len(rs.sl))

	return rs.sl[r]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
