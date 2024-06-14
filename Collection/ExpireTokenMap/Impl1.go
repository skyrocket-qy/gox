package ExpireTokenMap

import (
	"errors"
	"sync"
	"time"
)

type ValueWrap[V any] struct {
	value      any
	expiration time.Time
}

type TokenMapImpl1[V any] struct {
	tokens                map[string]*ValueWrap[V]
	expireInterval        time.Duration
	cleanInterval         time.Duration
	mu                    sync.RWMutex
	tokenTimes            []string
	adaptiveCleanInterval bool
	targetLen             int
}

/*
TokenMapImpl1 implements for fixed expiration
clean interval use to periodically clean the expired tokens
*/
func NewTokenMapImpl1[V any](
	expireInterval time.Duration,
	cleanInternval time.Duration,
	options ...Options,
) (*TokenMapImpl1[V], error) {
	tm := &TokenMapImpl1[V]{
		tokens:         make(map[string]*ValueWrap[V]),
		tokenTimes:     make([]string, 0),
		expireInterval: expireInterval,
		cleanInterval:  cleanInternval,
		mu:             sync.RWMutex{},
	}
	if len(options) > 0 {
		if options[0].AdaptiveCleanInterval {
			if options[0].TargetLen == 0 {
				return nil, errors.New("target length must greater than zero")
			}
			tm.adaptiveCleanInterval = true
			tm.targetLen = options[0].TargetLen
			go tm.adaptiveCleanup()
		} else {
			go tm.periodicCleanup()
		}
	}
	return tm, nil
}

// return false if token exists
func (tm *TokenMapImpl1[V]) Set(token string, value V) bool {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	if _, exists := tm.tokens[token]; exists {
		return false
	}

	valueWrap := &ValueWrap[V]{
		value:      value,
		expiration: time.Now().Add(tm.expireInterval),
	}
	tm.tokens[token] = valueWrap
	tm.tokenTimes = append(tm.tokenTimes, token)
	return true
}

// return false if not found or expired
func (tm *TokenMapImpl1[V]) Get(token string) (V, bool) {
	tm.mu.RLock()
	defer tm.mu.RUnlock()
	val, exists := tm.tokens[token]
	if !exists || val.expiration.Before(time.Now()) {
		return val.value.(V), false
	}
	return val.value.(V), true
}

func (tm *TokenMapImpl1[V]) periodicCleanup() {
	for {
		time.Sleep(tm.cleanInterval)
		now := time.Now()
		tm.mu.Lock()

		// m := int(math.Round(tm.cleanInterval.Seconds() /
		// 	(tm.cleanInterval.Seconds() + tm.expireInterval.Seconds()) *
		// 	float64(len(tm.tokenTimes))))
		// l, r := 0, len(tm.tokenTimes)
		// for l <= r {
		// 	if tm.tokenTimes[m].expiration.After(now) {
		// 		r = m - 1
		// 	} else {
		// 		l = m + 1
		// 	}
		// 	m = l + ((r - l) >> 1)
		// }

		for len(tm.tokenTimes) > 0 && tm.tokenTimes[0].expiration.Before(now) {
			token := tm.tokenTimes[0].token
			tm.tokenTimes = tm.tokenTimes[1:]
			delete(tm.tokens, token)
		}
		tm.mu.Unlock()
	}
}

func (tm *TokenMapImpl1[value]) adaptiveCleanup() {
	// for {
	// 	time.Sleep(tm.CleanInterval)
	// 	now := time.Now()
	// 	tm.mu.Lock()
	// 	for tm.pq.Len() > 0 && tm.pq[0].expiration.Before(now) {
	// 		item := heap.Pop(&tm.pq).(*Item)
	// 		delete(tm.tokens, item.token)
	// 	}
	// 	tm.mu.Unlock()
	// }
}
