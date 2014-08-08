package gowikia

import (
	"testing"
)

func TestNewWikiaApi(t *testing.T) {
	mwa, _ := NewWikiaApi("")
	if mwa != nil {
		t.Errorf("Expected NewWikiaApi to fail return nil for invalid url")
	}
	_, err := NewWikiaApi("http://example.com")
	if err != nil {
		t.Errorf("Expected NewWikiaApi to succeed for valid url")
	}
}
