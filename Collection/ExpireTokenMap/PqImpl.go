package ExpireTokenMap

import (
	"container/heap"
	"sync"
	"time"
)

type tokens struct {
	Info       info
	ExpireTime time.Time
}

type TokenMap[info any] struct {
	tokens map[string]struct {
		Info       info
		ExpireTime time.Time
	}
	// tokens          map[string]info
	cleanInternaval time.Duration
	mu              sync.Mutex
	pq              PriorityQueue
}

type Options struct {
	CleanInternaval time.Duration
}

func NewTokenMap[info any](options ...Options) *TokenMap[info] {
	cleanInternval := time.Minute
	if len(options) > 0 {
		cleanInternval = options[0].CleanInternaval
	}
	tm := &TokenMap[info]{
		tokens:          make(map[string]info),
		pq:              make(PriorityQueue, 0),
		cleanInternaval: cleanInternval,
		mu:              sync.Mutex{},
	}
	heap.Init(&tm.pq)
	go tm.periodicCleanup()
	return tm
}

func (tm *TokenMap[info]) SetToken(token string, value info,
	expireTime time.Time,
) {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	type Wrap struct {
		Info       info
		ExpireTime time.Time
	}
	tm.tokens[token] = Wrap{
		Info:       info,
		ExpireTime: expireTime,
	}
	heap.Push(&tm.pq, &Item{
		token:      token,
		expiration: expireTime,
	})
}

func (tm *TokenMap[info]) GetToken(token string) (info, bool) {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	tokenInfo, exists := tm.tokens[token]
	if !exists || tokenInfo.ExpireTime.Before(time.Now()) {
		delete(tm.tokens, token)
		return "", false
	}
	return tokenInfo.Info, true
}

func (tm *TokenMap[info]) periodicCleanup() {
	for {
		time.Sleep(1 * time.Minute)
		now := time.Now()
		tm.mu.Lock()
		for tm.pq.Len() > 0 && tm.pq[0].expiration.Before(now) {
			item := heap.Pop(&tm.pq).(*Item)
			delete(tm.tokens, item.token)
		}
		tm.mu.Unlock()
	}
}

// Priority Queue implementation
type Item struct {
	token      string
	expiration time.Time
	index      int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].expiration.Before(pq[j].expiration)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
