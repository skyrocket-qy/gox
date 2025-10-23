package rdssortedset

import (
	"reflect"
	"testing"
)

func TestSortedSet_ZAdd(t *testing.T) {
	ss := New()

	ss.ZAdd(10, "member1")
	ss.ZAdd(20, "member2")
	ss.ZAdd(5, "member3")

	if ss.ZCard() != 3 {
		t.Errorf("Expected ZCard to be 3, got %d", ss.ZCard())
	}

	score, exists := ss.ZScore("member1")
	if !exists || score != 10 {
		t.Errorf("Expected member1 score 10, got %f, %t", score, exists)
	}

	// Update score
	ss.ZAdd(15, "member1")
	score, exists = ss.ZScore("member1")
	if !exists || score != 15 {
		t.Errorf("Expected member1 updated score 15, got %f, %t", score, exists)
	}

	if ss.ZCard() != 3 {
		t.Errorf("Expected ZCard to be 3 after update, got %d", ss.ZCard())
	}
}

func TestSortedSet_ZRem(t *testing.T) {
	ss := New()

	ss.ZAdd(10, "member1")
	ss.ZAdd(20, "member2")
	ss.ZAdd(5, "member3")

	if !ss.ZRem("member1") {
		t.Errorf("Expected member1 to be removed")
	}

	if ss.ZCard() != 2 {
		t.Errorf("Expected ZCard to be 2, got %d", ss.ZCard())
	}

	_, exists := ss.ZScore("member1")
	if exists {
		t.Errorf("Expected member1 not to exist")
	}

	// Try to remove non-existent member
	if ss.ZRem("nonexistent") {
		t.Errorf("Expected nonexistent member not to be removed")
	}
}

func TestSortedSet_ZRange(t *testing.T) {
	ss := New()

	ss.ZAdd(10, "member1")
	ss.ZAdd(20, "member2")
	ss.ZAdd(5, "member3")
	ss.ZAdd(15, "member4")

	expected := []string{"member3", "member1", "member4", "member2"}
	actual := ss.ZRange(0, -1)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected ZRange(0, -1) to be %v, got %v", expected, actual)
	}

	expected = []string{"member3", "member1"}
	actual = ss.ZRange(0, 1)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected ZRange(0, 1) to be %v, got %v", expected, actual)
	}

	expected = []string{"member1", "member4"}
	actual = ss.ZRange(1, 2)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected ZRange(1, 2) to be %v, got %v", expected, actual)
	}

	// Test with negative indices
	expected = []string{"member4", "member2"}
	actual = ss.ZRange(-2, -1)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected ZRange(-2, -1) to be %v, got %v", expected, actual)
	}

	// Test out of bounds
	expected = []string{"member3", "member1", "member4", "member2"}
	actual = ss.ZRange(0, 100)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected ZRange(0, 100) to be %v, got %v", expected, actual)
	}

	expected = []string{}
	actual = ss.ZRange(100, 101)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected ZRange(100, 101) to be %v, got %v", expected, actual)
	}
}

func TestSortedSet_ZRangeByScore(t *testing.T) {
	ss := New()

	ss.ZAdd(10, "member1")
	ss.ZAdd(20, "member2")
	ss.ZAdd(5, "member3")
	ss.ZAdd(15, "member4")
	ss.ZAdd(15, "member5") // Same score as member4, should be ordered lexicographically

	expected := []string{"member3", "member1", "member4", "member5", "member2"}
	actual := ss.ZRangeByScore(0, 100)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected ZRangeByScore(0, 100) to be %v, got %v", expected, actual)
	}

	expected = []string{"member3", "member1"}
	actual = ss.ZRangeByScore(0, 10)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected ZRangeByScore(0, 10) to be %v, got %v", expected, actual)
	}

	expected = []string{"member1", "member4", "member5"}
	actual = ss.ZRangeByScore(10, 15)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected ZRangeByScore(10, 15) to be %v, got %v", expected, actual)
	}

	expected = []string{"member4", "member5"}
	actual = ss.ZRangeByScore(11, 15)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected ZRangeByScore(11, 15) to be %v, got %v", expected, actual)
	}

	expected = []string{}
	actual = ss.ZRangeByScore(100, 200)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected ZRangeByScore(100, 200) to be %v, got %v", expected, actual)
	}
}
