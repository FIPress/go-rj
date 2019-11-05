package rj

import "testing"

func TestParse(t *testing.T) {
	in := `name: "a\b"
array: ["a","b"]`
	_, err := ParseString(in)
	if err != nil {

	}

	in = `name: "a\c"
array: ["a","b"]`
	_, err = ParseString(in)
	if err != nil {

	}
}
