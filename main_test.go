package main

import "testing"

func TestShorten(t *testing.T) {
	hash := shorten("http://google.com")
	if len(hash) == 0 {
		t.Fail()
	}
}
