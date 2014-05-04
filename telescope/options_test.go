package telescope

import (
	"net/http"
	"testing"
)

func TestParamCreationFromRequestObject(t *testing.T) {
	testReq, _ := http.NewRequest("GET", "http://www.google.com/aHR0cDovL2RlbWpqdG5mbHc4dDkuY2xvdWRmcm9udC5uZXQv/abcdef?w=42&h=23&blur=2.345", nil)
	p := NewOptionsFromRequest(testReq)
	if p.Width != 42 {
		t.Error("wrong width, got ", p.Width)
	}
	if p.Height != 23 {
		t.Error("wrong height, got ", p.Height)
	}
	if p.Origin != "http://demjjtnflw8t9.cloudfront.net/" {
		t.Error("wrong origin, got ", p.Origin)
	}
	if p.File != "abcdef" {
		t.Error("wrong file, got ", p.File)
	}
	if p.Blur != 2.345 {
		t.Error("blur was wrong, got ", p.Blur)
	}
	url, err := p.SourceURL()
	if err != nil {
		t.Error("source url creation failed")
	}
	if url.String() != "http://demjjtnflw8t9.cloudfront.net/abcdef" {
		t.Error("wrong source url, got ", url.String())
	}

	testReq, _ = http.NewRequest("GET", "http://www.google.com/aHR0cDovL2RlbWpqdG5mbHc4dDkuY2xvdWRmcm9udC5uZXQ=/abcdef/klm/xyz.jpg", nil)
	p = NewOptionsFromRequest(testReq)
	if p.File != "abcdef/klm/xyz.jpg" {
		t.Error("wrong file, got ", p.File)
	}

	if p.Height != 0 {
		t.Error("height ought to be 0")
	}

	if p.Blur != 0 {
		t.Error("blur ought to be 0")
	}

	url, err = p.SourceURL()
	if err != nil {
		t.Error("source url creation failed")
	}
	if url.String() != "http://demjjtnflw8t9.cloudfront.net/abcdef/klm/xyz.jpg" {
		t.Error("wrong source url, got ", url.String())
	}

}
