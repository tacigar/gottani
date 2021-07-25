package sort

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	expected := []int{2, 2, 3, 4, 5, 6, 8, 10, 12}
	slice := []int{8, 2, 4, 12, 3, 6, 10, 2, 5}
	InsertionSort(slice)
	if !reflect.DeepEqual(slice, expected) {
		t.Errorf("got: %v\nwant: %v", slice, expected)
	}
}
