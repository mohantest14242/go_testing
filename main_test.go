package math

import (
    "testing"
)

// TestAdd tests the Add function from the math package.
func TestAdd(t *testing.T) {
    tests := []struct {
        a, b     int
        expected int
    }{
        {1, 1, 2},
        {2, 2, 4},
        {-1, -1, -2},
        {0, 0, 0},
        {100, 200, 300},
    }

    for _, tt := range tests {
        t.Run("", func(t *testing.T) {
            result := Add(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
            }
        })
    }
}
