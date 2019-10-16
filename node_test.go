package rj

import (
	"testing"
)

func TestGet(t *testing.T) {
	node := NewNode()
	node.dict["name"] = "Zoe"

	s, err := node.Get("name")

	if err != nil {
		t.Error("Get failed, expected no error, got:", err)
	}

	if s.(string) != "Zoe" {
		t.Error("Get failed, expected: Zoe, got:", s)
	}

	_, err = node.Get("fake")

	if err == nil || err != errValueNotFound {
		t.Error("Get failed, should return error (value not found), got:", err)
	}
}

func TestGetString(t *testing.T) {
	node := NewNode()
	node.dict["name"] = "Zoe"
	node.dict["age"] = 12

	s, err := node.GetStringOrError("name")

	if err != nil {
		t.Error("Get failed, expected no error, got:", err)
	}

	if s != "Zoe" {
		t.Error("GetStringOrError failed, expected: Zoe, got:", s)
	}

	_, err = node.GetStringOrError("age")

	if err == nil || err != errTypeMismatch {
		t.Error("GetStringOrError failed, should return error (type mismatch), got:", err)
	}

	s = node.GetStringOr("name", "Jim")

	if s != "Zoe" {
		t.Error("GetStringOr failed, expected: Zoe, got:", s)
	}

	s = node.GetStringOr("age", "Jim")

	if s != "Jim" {
		t.Error("GetStringOr failed, expected: Jim, got:", s)
	}

	s = node.GetString("name")

	if s != "Zoe" {
		t.Error("GetString failed, expected: Zoe, got:", s)
	}

	s = node.GetString("age")

	if s != "" {
		t.Error("GetString failed, expected empty string, got:", s)
	}

}
