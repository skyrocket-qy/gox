package ExpireTokenMap

import (
	"container/heap"
	"sync"
	"time"
)

type TokenMapImpl2[info any] struct {
	tokens        map[string]TokenWrap[info]
	CleanInterval time.Duration
	mu            sync.Mutex
	pq            PriorityQueue
}

// TokenMap implements with heap for dynamic clean interval
func NewTokenMapImpl2[info any](options ...Options) *TokenMap[info] {
	cleanInternval := time.Minute
	if len(options) > 0 {
		cleanInternval = options[0].CleanInterval
	}
	tm := &TokenMap[info]{
		tokens:        make(map[string]TokenWrap[info]),
		pq:            make(PriorityQueue, 0),
		CleanInterval: cleanInternval,
		mu:            sync.Mutex{},
	}
	heap.Init(&tm.pq)
	go tm.periodicCleanup()
	return tm
}

func (tm *TokenMapImpl2[info]) SetToken(token string, value info,
	expireTime time.Time,
) bool {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	if _, exists := tm.tokens[token]; exists {
		return false
	}

	tm.tokens[token] = TokenWrap[info]{
		Info:       value,
		ExpireTime: expireTime,
	}
	heap.Push(&tm.pq, &Item{
		token:      token,
		expiration: expireTime,
	})
	return true
}

func (tm *TokenMapImpl2[info]) GetToken(token string) (info, bool) {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	tokenInfo, exists := tm.tokens[token]
	if !exists || tokenInfo.ExpireTime.Before(time.Now()) {
		delete(tm.tokens, token)
		return tokenInfo.Info, false
	}
	return tokenInfo.Info, true
}

func (tm *TokenMapImpl2[info]) periodicCleanup() {
	for {
		time.Sleep(tm.CleanInterval)
		now := time.Now()
		tm.mu.Lock()
		for tm.pq.Len() > 0 && tm.pq[0].expiration.Before(now) {
			item := heap.Pop(&tm.pq).(*Item)
			delete(tm.tokens, item.token)
		}
		tm.mu.Unlock()
	}
}
