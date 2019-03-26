package main

import (
	"testing"
)

func TestPostSerialize(t *testing.T) {
	cases := []Post{
		MockPost(),
	}
	for _, c := range cases {
		// get back the result of ToJSON
		// make sure it's actually valid JSON by reformatting it to a Post Type
		json := c.ToJSON()
		p, err := PostFromJSON(json)
		if err != nil {
			t.Errorf("TestPostSerialize: could not decode JSON")
		} else if p.ID != c.ID {
			t.Errorf("TestPostSerialize: expected p.ID == %d, received %d", c.ID, p.ID)
		}
	}
}

func TestPostSerializeFailsOnInvalidInput(t *testing.T) {
	invalid := []byte("asdf")
	_, err := PostFromJSON(invalid)
	if err == nil {
		t.Error("PostFromJSON did not return error with invalid JSON")
	}
}
