package main

import "testing"

func TestFlatten3x3(t *testing.T) {
	records := [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}
	var expectedFlattenMatrix = "1,2,3,4,5,6,7,8,9"
	flattenMatrix := Flatten(records)
	if flattenMatrix != expectedFlattenMatrix {
		t.Errorf("Flatten was incorrect, got: %s, want: %s", flattenMatrix, expectedFlattenMatrix)
	}
}

func TestFlatten5x5(t *testing.T) {
	records := [][]string{
		{"1", "2", "3", "2", "3"},
		{"4", "5", "6", "5", "6"},
		{"7", "8", "9", "8", "9"},
		{"1", "2", "3", "2", "3"},
		{"4", "5", "6", "5", "6"},
	}
	var expectedFlattenMatrix = "1,2,3,2,3,4,5,6,5,6,7,8,9,8,9,1,2,3,2,3,4,5,6,5,6"
	flattenMatrix := Flatten(records)
	if flattenMatrix != expectedFlattenMatrix {
		t.Errorf("Flatten was incorrect, got: %s, want: %s", flattenMatrix, expectedFlattenMatrix)
	}
}
