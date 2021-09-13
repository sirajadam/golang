package selection

import (
	"fmt"
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
func GenerateRandomArray(len, max int) []int {
	
	rand.Seed(time.Now().Unix())
	var randArr []int
	for i := 0; i < len; i++ {
        value := rand.Intn(max) + 1
	
		randArr = append(randArr, value)
	}
	return randArr
}

func TestSelectionSmallArray(t *testing.T) {
    want := GenerateRandomArray(5,20)
    got := selection(want)
	sort.Ints(want)
    if !Equal(got,want) {
        t.Errorf("Selection() = %v, want %v", got, want)
    } else {
		fmt.Printf("Selection() = %v, want %v", got, want)
	}
}