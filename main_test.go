package main

import "testing"

func TestEmptyMap(t *testing.T) {
	m := make(map[string]int, 0)
	if _, exists := m["hello"]; exists {
		t.Fatal("should not exist")
	} else {
		t.Log("success")
	}
}
