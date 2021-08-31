package twoSum

import (
	"testing"
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

func TestPositiveTwoSum(t *testing.T) {
    target := 6
    nums := []int{3,2,4}
    want := []int{1,2}
    got := twoSum(nums,target)
    if !Equal(got,want) {
        t.Errorf("twoSum() = %q, want %q", got, want)
    }
}

func TestNegativeTwoSum(t *testing.T) {
    target := 6
    nums := []int{3,2,4}
    want := []int{0,2}
    got := twoSum(nums,target)
    if Equal(got,want) {
        t.Errorf("twoSum() = %q, want %q", got, want)
    }
}