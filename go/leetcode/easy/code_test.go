package easy

import "testing"

func TestMerge(t *testing.T) {
	num1 := []int{1, 2, 3, 0, 0, 0}
	num2 := []int{2, 5, 6}
	merge(num1, 3, num2, 3)
	t.Log(num1)
}
