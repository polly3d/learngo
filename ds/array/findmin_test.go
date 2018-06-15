package array

import "testing"

func TestFindMin(t *testing.T) {
	arr1 := []int{5, 4, 3, 2, 1}
	expected := 1
	actual := FindMin(arr1)
	if actual != expected {
		t.Error("fail")
	}
}

func TestFindMin2(t *testing.T) {
	arr1 := []int{5, 4, 3, 2, 1}
	expected := 1
	actual := FindMin2(arr1)
	if actual != expected {
		t.Error("fail")
	}
}
