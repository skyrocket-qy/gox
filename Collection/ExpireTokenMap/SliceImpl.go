package ExpireTokenMap

import (
	"container/heap"
	"sync"
	"time"
)

type tokenTime struct {
	token      string
	expiration time.Time
}

type TokenMapFixedExpire[info any] struct {
	tokens        map[string]info
	CleanInterval time.Duration
	mu            sync.Mutex
	tokenTimes    []tokenTime
}

/*
TokenMapFixedExpire implements for fixed expiration
clean interval use to periodically clean the expired tokens
*/
func NewTokenMapFixedExpire[info any](options ...Options) *TokenMapFixedExpire[info] {
	cleanInternval := time.Minute
	if len(options) > 0 {
		cleanInternval = options[0].CleanInterval
	}
	tm := &TokenMapFixedExpire[info]{
		tokens:        make(map[string]TokenWrap[info]),
		tokenTimes:    make([]tokenTime, 0),
		CleanInterval: cleanInternval,
		mu:            sync.Mutex{},
	}
	go tm.periodicCleanup()
	return tm
}

func (tm *TokenMapFixedExpire[info]) SetToken(token string, value info,
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

func (tm *TokenMapFixedExpire[info]) GetToken(token string) (info, bool) {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	tokenInfo, exists := tm.tokens[token]
	if !exists || tokenInfo.ExpireTime.Before(time.Now()) {
		delete(tm.tokens, token)
		return tokenInfo.Info, false
	}
	return tokenInfo.Info, true
}

func (tm *TokenMapFixedExpire[info]) periodicCleanup() {
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
