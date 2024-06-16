package ExpireTokenMapImpl2

import (
	"container/heap"
	"errors"
	"sync"
	"time"
)

type TokenWrap[info any] struct {
	Info       info
	ExpireTime time.Time
}

type TokenMapImpl2[info any] struct {
	tokens        map[string]TokenWrap[info]
	cleanInterval time.Duration
	mu            sync.Mutex
	tokenTimes    PriorityQueue
	targetLen     int
}

// TokenMapImpl2 implements with heap for dynamic clean interval
func NewTokenMapImpl2[info any](
	cleanInternval time.Duration,
	options ...Options,
) (*TokenMapImpl2[info], error) {
	tm := &TokenMapImpl2[info]{
		tokens:        make(map[string]TokenWrap[info]),
		tokenTimes:    make(PriorityQueue, 0),
		cleanInterval: cleanInternval,
		mu:            sync.Mutex{},
	}
	heap.Init(&tm.tokenTimes)
	if len(options) > 0 {
		if options[0].AdaptiveCleanInterval {
			if options[0].TargetLen == 0 {
				return nil, errors.New("target length must greater than zero")
			}
			tm.targetLen = options[0].TargetLen
			go tm.adaptiveCleanup()
		} else {
			go tm.periodicCleanup()
		}
	}
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
	heap.Push(&tm.tokenTimes, &Item{
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
		time.Sleep(tm.cleanInterval)
		now := time.Now()
		tm.mu.Lock()
		for tm.tokenTimes.Len() > 0 && tm.tokenTimes[0].expiration.Before(now) {
			item := heap.Pop(&tm.tokenTimes).(*Item)
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

func (tokenTimes PriorityQueue) Len() int { return len(tokenTimes) }

func (tokenTimes PriorityQueue) Less(i, j int) bool {
	return tokenTimes[i].expiration.Before(tokenTimes[j].expiration)
}

func (tokenTimes PriorityQueue) Swap(i, j int) {
	tokenTimes[i], tokenTimes[j] = tokenTimes[j], tokenTimes[i]
	tokenTimes[i].index = i
	tokenTimes[j].index = j
}

func (tokenTimes *PriorityQueue) Push(x interface{}) {
	n := len(*tokenTimes)
	item := x.(*Item)
	item.index = n
	*tokenTimes = append(*tokenTimes, item)
}

func (tokenTimes *PriorityQueue) Pop() interface{} {
	old := *tokenTimes
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*tokenTimes = old[0 : n-1]
	return item
}

func (tm *TokenMapImpl2Impl1[V]) adaptiveCleanup() {
	for {
		time.Sleep(tm.cleanInterval)
		now := time.Now()
		tm.mu.Lock()

		for len(tm.tokenTimes) > 0 && tm.tokens[tm.tokenTimes[0]].expiration.Before(now) {
			token := tm.tokenTimes[0]
			tm.tokenTimes = tm.tokenTimes[1:]
			delete(tm.tokens, token)
		}

		tm.mu.Unlock()

		if len(tm.tokens) > tm.targetLen {
			tm.cleanInterval = max(tm.cleanInterval>>1, 4*time.Second)
		} else if len(tm.tokens) < tm.targetLen {
			tm.cleanInterval = min(tm.cleanInterval<<1, 600*time.Second)
		}
	}
}
