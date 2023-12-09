package main

import "testing"

func TestLessHand(t *testing.T) {
	testCases := []struct {
		hand1    string
		hand2    string
		expected bool
	}{
		// Test cases where hand1 is less than hand2
		{"TATTT", "ATAAA", true},
		{"TTTT6", "AAAJA", true},
		{"AA2AA", "TT3TT", false},

		// Test cases where hand1 is greater than hand2
		{"KAKKK", "KKTKK", false},
		{"A6AAA", "T8TTT", false},
		{"KTTTT", "T4TTT", false},

		// Test cases where hand1 is equal to hand2
		{"QQQ6Q", "QQQ6Q", false},
		{"ATJKK", "ATJKK", false},
		{"23456", "23456", false},
	}

	for _, tc := range testCases {
		actual := lessHand(tc.hand1, tc.hand2)
		if actual != tc.expected {
			t.Errorf("expected %t but got %t for hand1: %s, hand2: %s", tc.expected, actual, tc.hand1, tc.hand2)
		}
	}
}
