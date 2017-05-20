package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestBar(t *testing.T) {
	_, err := bar(false)
	if err == nil {
		t.Error("Expected an error")
	}

	assert.NotNil(t, err)

	ret, _ := bar(true)
	if ret != "keeper" {
		t.Errorf("Expected 'keeper', got %s", ret)
	}
}

func BenchmarkBar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = bar(false)
	}
}