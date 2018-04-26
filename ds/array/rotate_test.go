package array

import "testing"

func TestRotate(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}
	d := 2
	n := 6
	expectedArr := []int{3, 4, 5, 6, 1, 2}
	targetArr := rotate(arr, d, n)
	if len(expectedArr) != len(targetArr) {
		t.Error("fail")
	}
	for i, v := range expectedArr {
		if v != targetArr[i] {
			t.Errorf("Expected index %d is %d, but get %d", i, v, targetArr[i])
			break
		}
	}
}

func TestRotateOneByOne(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}
	d := 2
	n := 6
	expectedArr := []int{3, 4, 5, 6, 1, 2}
	targetArr := rotateOneByOne(arr, d, n)
	if targetArr == nil {
		t.Error("Nothing return from func")
	}
	if len(expectedArr) != len(targetArr) {
		t.Error("fail")
	}
	for i, v := range expectedArr {
		if v != targetArr[i] {
			t.Errorf("Expected index %d is %d, but get %d", i, v, targetArr[i])
			break
		}
	}
}
