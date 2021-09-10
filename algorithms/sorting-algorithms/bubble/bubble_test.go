package bubble

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

// Helper function to compare two arrays
func Equal(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}

// Helper function to create random array of lenght len and 
func GenerateRandomArray(len, max int, sign bool) []int {
	
	rand.Seed(time.Now().Unix())
	var randArr []int
	for i := 0; i < len; i++ {
        value := rand.Intn(max) + 1
        if sign {
            value = -(rand.Intn(max) + 1)
        } 
	
		randArr = append(randArr, value)
	}
	return randArr
}


func TestBubble(t *testing.T) {
	want := GenerateRandomArray(50000,20, false)
	got := bubble(want)
	sort.Ints(want)
	if !Equal(got, want) {
        t.Errorf("Insertion() = %v, want %v", got, want)

	}
}