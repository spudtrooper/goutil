package slice

import (
	"reflect"
	"testing"
)

func TestChunk(t *testing.T) {
	tests := []struct {
		name      string
		original  []int
		chunkSize int
		expected  [][]int
	}{
		{
			name:      "RegularChunking",
			original:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			chunkSize: 3,
			expected:  [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		},
		{
			name:      "IncompleteFinalChunk",
			original:  []int{1, 2, 3, 4, 5, 6, 7},
			chunkSize: 3,
			expected:  [][]int{{1, 2, 3}, {4, 5, 6}, {7}},
		},
		{
			name:      "ChunkSizeLargerThanSlice",
			original:  []int{1, 2, 3},
			chunkSize: 5,
			expected:  [][]int{{1, 2, 3}},
		},
		{
			name:      "EmptySlice",
			original:  []int{},
			chunkSize: 3,
			expected:  [][]int{},
		},
		{
			name:      "ZeroChunkSize",
			original:  []int{1, 2, 3},
			chunkSize: 0,
			expected:  [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Chunk(tt.original, tt.chunkSize)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Chunk() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
