package quotientfilter_test

import (
	"fmt"
	"testing"

	"quotientfilter"
)

func TestQuotientFilterBasic(t *testing.T) {
	qf, err := quotientfilter.New(8, 8) // q=8, r=8, so 256 buckets, 16-bit fingerprints
	if err != nil {
		t.Fatalf("Failed to create QuotientFilter: %v", err)
	}

	testStrings := []string{"apple", "banana", "cherry", "date", "elderberry", "fig", "grape", "honeydew"}

	// Test insertion
	for _, s := range testStrings {
		err := qf.Insert([]byte(s))
		if err != nil {
			t.Errorf("Failed to insert %s: %v", s, err)
		}
	}

	// Test Contains for inserted elements
	for _, s := range testStrings {
		if !qf.Contains([]byte(s)) {
			t.Errorf("QuotientFilter should contain %s, but it doesn't", s)
		}
	}

	// Test Contains for non-existent elements (expect false positives sometimes)
	nonExistentStrings := []string{"kiwi", "lemon", "mango", "orange", "pear"}
	for _, s := range nonExistentStrings {
		if qf.Contains([]byte(s)) {
			t.Logf("False positive for %s (expected not to contain)", s)
		}
	}

	// Test filter capacity
	if qf.Size() != uint64(len(testStrings)) {
		t.Errorf("Expected size %d, got %d", len(testStrings), qf.Size())
	}

	// Test inserting more than capacity (should fail)
	qf2, err := quotientfilter.New(8, 8)
	if err != nil {
		t.Fatalf("Failed to create QuotientFilter for capacity test: %v", err)
	}

	for i := 0; i < 250; i++ { // Insert 250 items into a 256-bucket filter
		err := qf2.Insert([]byte(fmt.Sprintf("item%d", i)))
		if err != nil {
			t.Errorf("Failed to insert item%d into qf2: %v", i, err)
			break
		}
	}

	if qf2.Size() != 250 {
		t.Errorf("Expected qf2 size %d, got %d", 250, qf2.Size())
	}

	// Try to insert one more, it might fail or succeed depending on collisions
	err = qf2.Insert([]byte("last_item"))
	if err != nil {
		t.Logf("Expected error for inserting into near-full filter: %v", err)
	} else {
		t.Logf("Successfully inserted into near-full filter, current size: %d", qf2.Size())
	}

	// Verify some elements in qf2
	if !qf2.Contains([]byte("item100")) {
		t.Errorf("qf2 should contain item100")
	}
}

func TestQuotientFilterEdgeCases(t *testing.T) {
	// Test with small q and r
	qf, err := quotientfilter.New(2, 2) // q=2, r=2, so 4 buckets, 4-bit fingerprints
	if err != nil {
		t.Fatalf("Failed to create QuotientFilter: %v", err)
	}

	items := []string{"a", "b", "c", "d"} // Max 4 items for q=2

	for _, item := range items {
		err := qf.Insert([]byte(item))
		if err != nil {
			t.Errorf("Failed to insert %s: %v", item, err)
		}
	}

	// Try to insert one more, should fail
	err = qf.Insert([]byte("e"))
	if err == nil {
		t.Errorf("Expected error when inserting into a full filter, but got none")
	} else {
		t.Logf("Successfully caught expected error: %v", err)
	}

	// Verify contents
	for _, item := range items {
		if !qf.Contains([]byte(item)) {
			t.Errorf("Filter should contain %s", item)
		}
	}

	// Test with q+r > 64
	_, err = quotientfilter.New(33, 33)
	if err == nil {
		t.Errorf("Expected error for q+r > 64, but got none")
	} else {
		t.Logf("Successfully caught expected error: %v", err)
	}

	// Test with r+remainderOffset > 64 (e.g., r=62, remainderOffset=3)
	_, err = quotientfilter.New(2, 62)
	if err == nil {
		t.Errorf("Expected error for r+remainderOffset > 64, but got none")
	} else {
		t.Logf("Successfully caught expected error: %v", err)
	}
}
