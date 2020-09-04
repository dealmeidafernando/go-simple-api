package main

import (
	"net/http"
	"testing"
)

func TestGetPeople(t *testing.T) {
	tests := []struct {
		name string
		w    http.ResponseWriter
		r    *http.Request
	}{
		// test test aaaaaaaaaaaaa
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetPeople(tt.w, tt.r)
		})
	}
}
