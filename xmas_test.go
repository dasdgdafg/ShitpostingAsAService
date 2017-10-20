package main

import (
	"regexp"
	"testing"
)

// the expected output format
var link = regexp.MustCompile("^(\\d+ days? until Christmas|It's Christmas!) https://.*\\.(png|gif|jpg)$")

func TestXmas(t *testing.T) {
	for i := 0; i < 10; i++ {
		s := xmas()
		if !link.MatchString(s) {
			t.Error("Invalid link:" + s)
		}
	}
}
