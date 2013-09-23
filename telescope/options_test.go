package telescope

import (
	"net/http"
	"testing"
)

func TestParamCreation(t *testing.T) {
	s := NewOptionsFromMap(map[string]string{"origin": "http://www.google.com", "width": "42", "height": "24"})
	if s.Origin != "http://www.google.com" {
		t.Error("Wrong origin parsed")
	}
	if s.Width != 42 {
		t.Error("Wrong width parsed")
	}
	if s.Height != 24 {
		t.Error("Wrong height initialized")
	}
}

func TestParamCreationFromRequestObject(t *testing.T) {
	testReq, _ := http.NewRequest("GET", "http://www.google.com?w=42&origin=abc&h=23", nil)
	p := NewOptionsFromRequest(testReq)
	if p.Width != 42 || p.Height != 23 || p.Origin != "abc" {
		t.Fail()
	}
}
