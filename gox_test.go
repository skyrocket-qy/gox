package gox

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbs(t *testing.T) {
	assert.Equal(t, 1, Abs(-1))
	assert.Equal(t, 1, Abs(1))
	assert.Equal(t, 0, Abs(0))
	assert.InEpsilon(t, 1.23, Abs(-1.23), 1e-6)
	assert.InEpsilon(t, 1.23, Abs(1.23), 1e-6)
}

func TestBatch(t *testing.T) {
	// Test case 1: Simple slice and small batch size
	items1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	batchSize1 := 3
	expected1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10},
	}

	batches1 := make([][]int, 0)
	for batch := range Batch(items1, batchSize1) {
		batches1 = append(batches1, batch)
	}

	assert.Equal(t, expected1, batches1)

	// Test case 2: Slice length is a multiple of batch size
	items2 := []string{"a", "b", "c", "d", "e", "f"}
	batchSize2 := 2
	expected2 := [][]string{
		{"a", "b"},
		{"c", "d"},
		{"e", "f"},
	}

	batches2 := make([][]string, 0)
	for batch := range Batch(items2, batchSize2) {
		batches2 = append(batches2, batch)
	}

	assert.Equal(t, expected2, batches2)

	// Test case 3: Empty slice
	items3 := []float64{}
	batchSize3 := 5
	expected3 := [][]float64{}

	batches3 := make([][]float64, 0)
	for batch := range Batch(items3, batchSize3) {
		batches3 = append(batches3, batch)
	}

	assert.Equal(t, expected3, batches3)

	// Test case 4: Batch size larger than slice length
	items4 := []byte{1, 2, 3}
	batchSize4 := 10
	expected4 := [][]byte{
		{1, 2, 3},
	}

	batches4 := make([][]byte, 0)
	for batch := range Batch(items4, batchSize4) {
		batches4 = append(batches4, batch)
	}

	assert.Equal(t, expected4, batches4)

	// Test case 5: Batch size of 1
	items5 := []int{1, 2, 3}
	batchSize5 := 1
	expected5 := [][]int{
		{1},
		{2},
		{3},
	}

	batches5 := make([][]int, 0)
	for batch := range Batch(items5, batchSize5) {
		batches5 = append(batches5, batch)
	}

	assert.Equal(t, expected5, batches5)
}

func TestStr(t *testing.T) {
	t.Run("should return a pointer to the given string", func(t *testing.T) {
		// Arrange
		s := "hello"

		// Act
		sp := Str(s)

		// Assert
		assert.NotNil(t, sp)
		assert.Equal(t, s, *sp)
	})

	t.Run("should return a pointer to an empty string", func(t *testing.T) {
		// Arrange
		s := ""

		// Act
		sp := Str(s)

		// Assert
		assert.NotNil(t, sp)
		assert.Equal(t, s, *sp)
	})
}
