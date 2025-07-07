package sequenceTransaction

import (
	"math"
	"sort"
)

type (
	DB string

	Instance struct {
		seqMap map[DB]int
	}
)

// Initial the sequence of the transaction order to avoid deadlock
func NewInstance(seq []DB) *Instance {
	seqMap := make(map[DB]int, len(seq))
	for i, db := range seq {
		seqMap[db] = i
	}
	return &Instance{
		seqMap: seqMap,
	}
}

func (st *Instance) Sort(in []DB) {
	sort.Slice(in, func(i, j int) bool {
		var vi, vj int
		if v, ok := st.seqMap[in[i]]; ok {
			vi = v
		} else {
			vi = math.MaxInt
		}

		if v, ok := st.seqMap[in[j]]; ok {
			vj = v
		} else {
			vj = math.MaxInt
		}

		return vi <= vj
	})
}
