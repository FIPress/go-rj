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

func TestMarshal(t *testing.T) {
	type s struct {
		Name string
		Age  int
	}
	in := `Name: "abc"
Age: 12
`
	s1 := new(s)
	err := Unmarshal([]byte(in), s1)

	if err == nil {
		t.Log(s1)

		bytes := Marshal(s1)
		t.Log(string(bytes))
	} else {
		t.Log(err)
	}

}

func TestUnmarshal(t *testing.T) {
	type s struct {
		Name string
		Age  int
	}
	in := `Name: "abc"
Age: 12
`
	s1 := new(s)
	err := Unmarshal([]byte(in), s1)

	if err == nil {
		t.Log(s1)
	} else {
		t.Log(err)
	}

}
