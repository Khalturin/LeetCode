package main

import "testing"

func TestContainsDuplicateOne(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	if !containsDuplicate(nums) {
		t.Error("One fail")
	}
}

func TestContainsDuplicateTwo(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	if containsDuplicate(nums) {
		t.Error("Two fail")
	}
}

func TestContainsDuplicateThree(t *testing.T) {
	nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	if !containsDuplicate(nums) {
		t.Error("Three fail")
	}
}
