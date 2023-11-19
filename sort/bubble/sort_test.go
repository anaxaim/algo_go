package main

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"empty", []int{}, []int{}},
		{"single element", []int{1}, []int{1}},
		{"multiple elements", []int{3, 1, 2}, []int{1, 2, 3}},
		{"already sorted", []int{1, 2, 3}, []int{1, 2, 3}},
		{"reverse order", []int{3, 2, 1}, []int{1, 2, 3}},
		{"duplicates", []int{5, 2, 3, 2}, []int{2, 2, 3, 5}},
		{"negative", []int{5, 2, -3, -2}, []int{-3, -2, 2, 5}},
		{
			"large array",
			[]int{3, -1, 14, 2, 99, 17, 14, 105, 0, 1, 1, 45, 32, 88, 12, 85, 35, 66, 14, -99, -105},
			[]int{-105, -99, -1, 0, 1, 1, 2, 3, 12, 14, 14, 14, 17, 32, 35, 45, 66, 85, 88, 99, 105},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := sort(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func BenchmarkSort(b *testing.B) {
	largeArray := make([]int, 10000)
	for i := range largeArray {
		largeArray[i] = rand.Intn(10000)
	}

	b.Run("large array", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort(largeArray)
		}
	})
}

// BenchmarkSort/large_array-10    380421	    3178 ns/op
// BenchmarkSort/large_array-10    364833	    3201 ns/op
// BenchmarkSort/large_array-10    378789	    3176 ns/op
// BenchmarkSort/large_array-10    378207	    3156 ns/op
// BenchmarkSort/large_array-10    378264	    3166 ns/op
