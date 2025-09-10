package randomizedset

/* @tags: set,random */

import (
	"math/rand"
)

type RandomizedSet struct {
	mp map[int]int
	sl []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		mp: make(map[int]int),
		sl: make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.mp[val]; ok {
		return false
	}

	this.sl = append(this.sl, val)
	this.mp[val] = len(this.sl) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.mp[val]; !ok {
		return false
	}

	this.sl[this.mp[val]], this.sl[len(this.sl)-1] =
		this.sl[len(this.sl)-1], this.sl[this.mp[val]]

	this.mp[this.sl[this.mp[val]]] = this.mp[val]

	delete(this.mp, val)
	this.sl = this.sl[:len(this.sl)-1]
	return true
}

func (this *RandomizedSet) GetRandom() int {
	r := rand.Int() % len(this.sl)
	return this.sl[r]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
