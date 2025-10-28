package rank

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/skyrocket-qy/gox/dsa/rdssortedset"
)

const ( 
	numPlayers = 100000
	topN       = 10
)

// setupData generates a slice of player IDs and scores.
func setupData() ([]string, []int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	playerIDs := make([]string, numPlayers)
	scores := make([]int, numPlayers)
	for i := 0; i < numPlayers; i++ {
		playerIDs[i] = fmt.Sprintf("player%d", i)
		scores[i] = r.Intn(1000000) // Scores between 0 and 999,999
	}
	return playerIDs, scores
}

func BenchmarkTopNCache_GetTop10(b *testing.B) {
	playerIDs, scores := setupData()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache := NewTopNCache(topN)
		for j := 0; j < numPlayers; j++ {
			cache.SetScore(playerIDs[j], scores[j])
		}
		_ = cache.GetTopN()
	}
}

func BenchmarkSortedSet_GetTop10(b *testing.B) {
	playerIDs, scores := setupData()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ss := rdssortedset.New()
		for j := 0; j < numPlayers; j++ {
			ss.ZAdd(float64(scores[j]), playerIDs[j])
		}
		_ = ss.ZRange(0, topN-1)
	}
}
