package rank

import (
	"fmt"
	"sort"
)

// RankEntry holds the score for a single player.
type RankEntry struct {
	PlayerID string
	Score    int
}

// TopNCache holds the fast-access Top N leaderboard.
type TopNCache struct {
	// allScores is the "source of truth" for all player scores.
	allScores map[string]int

	// topN is a sorted slice (descending) of the top N players.
	topN []RankEntry

	// lowerBound is the score of the Nth-place player.
	lowerBound int

	// n is the maximum number of entries to keep in the top list.
	n int
}

// NewTopNCache creates a new, empty cache.
func NewTopNCache(n int) *TopNCache {
	return &TopNCache{
		allScores: make(map[string]int),
		// Pre-allocate capacity for n entries for efficiency
		topN:       make([]RankEntry, 0, n),
		lowerBound: 0,
		n:          n,
	}
}

// SetScore updates a player's score.
// This is the core logic.
func (c *TopNCache) SetScore(playerID string, newScore int) {
	// Update the master score list
	c.allScores[playerID] = newScore

	// --- Fast Path ---
	// First, check if the player is already in the Top N.
	playerIndexInTopN := -1
	for i, entry := range c.topN {
		if entry.PlayerID == playerID {
			playerIndexInTopN = i
			break
		}
	}

	// If the new score isn't high enough to make the Top N,
	// AND the player isn't *already* in the Top N, we can stop here.
	if newScore <= c.lowerBound && playerIndexInTopN == -1 {
		return
	}

	// --- Slow Path (Update Top N) ---

	if playerIndexInTopN != -1 {
		// Player is already in the list. Just update their score in-place.
		c.topN[playerIndexInTopN].Score = newScore
	} else {
		// Player is not in the list. Add them.
		// The list will temporarily grow to N+1 elements if full.
		c.topN = append(c.topN, RankEntry{PlayerID: playerID, Score: newScore})
	}

	// Re-sort the list. This is very fast for N or N+1 elements.
	// We sort by score, descending (highest first).
	sort.Slice(c.topN, func(i, j int) bool {
		return c.topN[i].Score > c.topN[j].Score
	})

	// If we now have N+1 elements, drop the last one.
	if len(c.topN) > c.n {
		c.topN = c.topN[:c.n] // Truncate to top N
	}

	// Finally, update the new lower bound
	if len(c.topN) == c.n {
		c.lowerBound = c.topN[c.n-1].Score
	} else if len(c.topN) > 0 {
		// Handle case where list is not yet full
		c.lowerBound = c.topN[len(c.topN)-1].Score
	} else {
		c.lowerBound = 0
	}
}

// GetTopN returns a copy of the current Top N list.
func (c *TopNCache) GetTopN() []RankEntry {
	// Return a *copy* to prevent race conditions
	// if the caller tries to modify the returned slice.
	result := make([]RankEntry, len(c.topN))
	copy(result, c.topN)
	return result
}

// GetPlayerScore returns a single player's score from the master list.
func (c *TopNCache) GetPlayerScore(playerID string) (int, bool) {
	score, ok := c.allScores[playerID]
	return score, ok
}

// --- Example Usage ---

func main() {
	lb := NewTopNCache(10)

	// Add 10 players
	lb.SetScore("player1", 100)
	lb.SetScore("player2", 200)
	lb.SetScore("player3", 300)
	lb.SetScore("player4", 400)
	lb.SetScore("player5", 500)
	lb.SetScore("player6", 600)
	lb.SetScore("player7", 700)
	lb.SetScore("player8", 800)
	lb.SetScore("player9", 900)
	lb.SetScore("player10", 1000)

	fmt.Println("--- Initial Top 10 ---")
	for i, entry := range lb.GetTopN() {
		fmt.Printf("#%d: %s (Score: %d)\n", i+1, entry.PlayerID, entry.Score)
	}
	fmt.Printf("Lower bound: %d\n\n", lb.lowerBound)

	// Add a player who is too low
	lb.SetScore("player11", 50) // Too low, shouldn't change anything
	// Add a player who knocks someone out
	lb.SetScore("player12", 550) // Should knock player1 (100) out

	fmt.Println("--- After adding player11 (50) and player12 (550) ---")
	for i, entry := range lb.GetTopN() {
		fmt.Printf("#%d: %s (Score: %d)\n", i+1, entry.PlayerID, entry.Score)
	}
	fmt.Printf("Lower bound: %d\n\n", lb.lowerBound) // Should now be 200

	// Update an existing player to change the order
	lb.SetScore("player2", 950) // Should move to #2 spot

	fmt.Println("--- After updating player2 to 950 ---")
	for i, entry := range lb.GetTopN() {
		fmt.Printf("#%d: %s (Score: %d)\n", i+1, entry.PlayerID, entry.Score)
	}
	fmt.Printf("Lower bound: %d\n\n", lb.lowerBound) // Should still be 300 (player3)

	// Get a specific player's score from the master list
	score, _ := lb.GetPlayerScore("player1") // The one who got knocked out
	fmt.Printf("Player1's score (from master list): %d\n", score)

	score, _ = lb.GetPlayerScore("player11") // The one who never made it
	fmt.Printf("Player11's score (from master list): %d\n", score)
}
