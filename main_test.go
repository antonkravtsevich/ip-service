package main

import "testing"

func TestIsArrayContainValue(t *testing.T){
	inputArray := []string{"1", "2", "3", "4", "5"}

	exists := isArrayContainValue(inputArray, "1")
	if !exists {
		t.Errorf("isArrayContainValue failed, expected %v, got %v", true, exists)
	}

	exists = isArrayContainValue(inputArray, "6")
	if exists {
		t.Errorf("isArrayContainValue failed, expected %v, got %v", false, exists)
	}
}