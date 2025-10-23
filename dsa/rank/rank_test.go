package rank

import (
	"reflect"
	"testing"
)

func TestNewTopNCache(t *testing.T) {
	cache := NewTopNCache(5)

	if cache == nil {
		t.Fatal("NewTopNCache returned nil")
	}
	if cache.n != 5 {
		t.Errorf("Expected n to be 5, got %d", cache.n)
	}
	if len(cache.allScores) != 0 {
		t.Errorf("Expected allScores to be empty, got %v", cache.allScores)
	}
	if len(cache.topN) != 0 {
		t.Errorf("Expected topN to be empty, got %v", cache.topN)
	}
	if cache.lowerBound != 0 {
		t.Errorf("Expected lowerBound to be 0, got %d", cache.lowerBound)
	}
}

func TestSetScoreAndGetTopN(t *testing.T) {
	cache := NewTopNCache(3)

	// Add initial players
	cache.SetScore("p1", 100)
	cache.SetScore("p2", 200)
	cache.SetScore("p3", 300)

	expectedTopN := []RankEntry{
		{PlayerID: "p3", Score: 300},
		{PlayerID: "p2", Score: 200},
		{PlayerID: "p1", Score: 100},
	}
	actualTopN := cache.GetTopN()
	if !reflect.DeepEqual(actualTopN, expectedTopN) {
		t.Errorf("Expected %v, got %v", expectedTopN, actualTopN)
	}
	if cache.lowerBound != 100 {
		t.Errorf("Expected lowerBound to be 100, got %d", cache.lowerBound)
	}

	// Add a player with a lower score than lowerBound
	cache.SetScore("p4", 50)
	actualTopN = cache.GetTopN()
	if !reflect.DeepEqual(actualTopN, expectedTopN) {
		t.Errorf("Expected %v, got %v", expectedTopN, actualTopN)
	}
	if cache.lowerBound != 100 {
		t.Errorf("Expected lowerBound to be 100, got %d", cache.lowerBound)
	}

	// Add a player that enters the top N
	cache.SetScore("p5", 250)
	expectedTopN = []RankEntry{
		{PlayerID: "p3", Score: 300},
		{PlayerID: "p5", Score: 250},
		{PlayerID: "p2", Score: 200},
	}
	actualTopN = cache.GetTopN()
	if !reflect.DeepEqual(actualTopN, expectedTopN) {
		t.Errorf("Expected %v, got %v", expectedTopN, actualTopN)
	}
	if cache.lowerBound != 200 {
		t.Errorf("Expected lowerBound to be 200, got %d", cache.lowerBound)
	}

	// Update an existing player's score
	cache.SetScore("p2", 350)
	expectedTopN = []RankEntry{
		{PlayerID: "p2", Score: 350},
		{PlayerID: "p3", Score: 300},
		{PlayerID: "p5", Score: 250},
	}
	actualTopN = cache.GetTopN()
	if !reflect.DeepEqual(actualTopN, expectedTopN) {
		t.Errorf("Expected %v, got %v", expectedTopN, actualTopN)
	}
	if cache.lowerBound != 250 {
		t.Errorf("Expected lowerBound to be 250, got %d", cache.lowerBound)
	}

	// Test with n=1
	cache1 := NewTopNCache(1)
	cache1.SetScore("pA", 100)
	cache1.SetScore("pB", 200)
	expectedTopN1 := []RankEntry{{PlayerID: "pB", Score: 200}}
	actualTopN1 := cache1.GetTopN()
	if !reflect.DeepEqual(actualTopN1, expectedTopN1) {
		t.Errorf("Expected %v, got %v", expectedTopN1, actualTopN1)
	}
	if cache1.lowerBound != 200 {
		t.Errorf("Expected lowerBound to be 200, got %d", cache1.lowerBound)
	}
}

func TestGetPlayerScore(t *testing.T) {
	cache := NewTopNCache(2)
	cache.SetScore("p1", 100)
	cache.SetScore("p2", 200)

	score, ok := cache.GetPlayerScore("p1")
	if !ok || score != 100 {
		t.Errorf("Expected p1 score 100, got %d, %t", score, ok)
	}

	score, ok = cache.GetPlayerScore("p2")
	if !ok || score != 200 {
		t.Errorf("Expected p2 score 200, got %d, %t", score, ok)
	}

	score, ok = cache.GetPlayerScore("p3")
	if ok || score != 0 {
		t.Errorf("Expected p3 not to exist, got %d, %t", score, ok)
	}
}

/*
func TestConcurrentAccess(t *testing.T) {
	cache := NewTopNCache(5)
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			playerID := fmt.Sprintf("player%d", i)
			score := i * 10
			cache.SetScore(playerID, score)
		}(i)
	}
	wg.Wait()

	if cache.n != 5 {
		t.Errorf("Expected n to be 5, got %d", cache.n)
	}
	if cache.lowerBound != 950 {
		t.Errorf("Expected lowerBound to be 950, got %d", cache.lowerBound)
	}

	expectedTopN := []RankEntry{
		{PlayerID: "player99", Score: 990},
		{PlayerID: "player98", Score: 980},
		{PlayerID: "player97", Score: 970},
		{PlayerID: "player96", Score: 960},
		{PlayerID: "player95", Score: 950},
	}
	actualTopN := cache.GetTopN()
	if !reflect.DeepEqual(actualTopN, expectedTopN) {
		t.Errorf("Expected %v, got %v", expectedTopN, actualTopN)
	}
}
*/
