package utils

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)



func TestBubbleSort(t *testing.T) {

	els := []int {4,7,5,6,3,2,1}

	els = BubbleSort(els)

	assert.NotNil(t, els)
	assert.EqualValues(t, 1, els[0])
	assert.EqualValues(t, 2, els[1])
	assert.EqualValues(t, 3, els[2])
	assert.EqualValues(t, 4, els[3])
	assert.EqualValues(t, 5, els[4])
	assert.EqualValues(t, 6, els[5])
	assert.EqualValues(t, 7, els[6])

}

func TestGetElements(t *testing.T) {
	els := GetElements(10)

	assert.NotNil(t,els)

}


func BenchmarkBubbleSort10(b *testing.B) {
	els := GetElements(10)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkBubbleSort1000(b *testing.B) {
	els := GetElements(1000)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkBubbleSort10000(b *testing.B) {
	els := GetElements(10000)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkBubbleSort50000(b *testing.B) {
	els := GetElements(50000)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkBubbleSort100000(b *testing.B) {
	els := GetElements(100000)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkGoSort10(b *testing.B) {
	els := GetElements(10)
	for i := 0; i < b.N; i++ {
		sort.Ints(els)
	}
}

func BenchmarkGoSort1000(b *testing.B) {
	els := GetElements(1000)
	for i := 0; i < b.N; i++ {
		sort.Ints(els)
	}
}

func BenchmarkGoSort10000(b *testing.B) {
	els := GetElements(10000)
	for i := 0; i < b.N; i++ {
		sort.Ints(els)
	}
}

func BenchmarkGoSort50000(b *testing.B) {
	els := GetElements(50000)
	for i := 0; i < b.N; i++ {
		sort.Ints(els)
	}
}

func BenchmarkGoSort100000(b *testing.B) {
	els := GetElements(100000)
	for i := 0; i < b.N; i++ {
		sort.Ints(els)
	}
}


