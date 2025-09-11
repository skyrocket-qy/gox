package prefixsumarray

import (
	"reflect"
	"testing"
)

func TestNewPreFixSumArray(t *testing.T) {
	in := []int{1, 2, 3, 4, 5}
	expected := []int{1, 3, 6, 10, 15}
	result := NewPreFixSumArray(in)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("NewPreFixSumArray(%v) = %v, want %v", in, result, expected)
	}
}

func TestNewPreFixSumArray_single_element(t *testing.T) {
	in := []int{10}
	expected := []int{10}
	result := NewPreFixSumArray(in)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("NewPreFixSumArray(%v) = %v, want %v", in, result, expected)
	}
}

func TestNewPreFixSumArray_with_zeros(t *testing.T) {
	in := []int{0, 0, 0, 0, 0}
	expected := []int{0, 0, 0, 0, 0}
	result := NewPreFixSumArray(in)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("NewPreFixSumArray(%v) = %v, want %v", in, result, expected)
	}
}

func TestNewPreFixSumArray_with_negative_numbers(t *testing.T) {
	in := []int{-1, -2, -3, -4, -5}
	expected := []int{-1, -3, -6, -10, -15}
	result := NewPreFixSumArray(in)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("NewPreFixSumArray(%v) = %v, want %v", in, result, expected)
	}
}
