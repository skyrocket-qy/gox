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
	tokens         map[string]*ValueWrap[V]
	expireInterval time.Duration
	cleanInterval  time.Duration
	mu             sync.RWMutex
	tokenTimes     []string
	targetLen      int
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

		for len(tm.tokenTimes) > 0 && tm.tokens[tm.tokenTimes[0]].expiration.Before(now) {
			token := tm.tokenTimes[0]
			tm.tokenTimes = tm.tokenTimes[1:]
			delete(tm.tokens, token)
		}
		tm.mu.Unlock()
	}
}

func (tm *TokenMapImpl1[V]) adaptiveCleanup() {
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
