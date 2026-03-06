package rfidfyi_test

import (
	"testing"

	rfidfyi "github.com/fyipedia/rfidfyi-go"
)

func TestNewClient(t *testing.T) {
	c := rfidfyi.NewClient()
	if c.BaseURL != rfidfyi.DefaultBaseURL {
		t.Errorf("expected %s, got %s", rfidfyi.DefaultBaseURL, c.BaseURL)
	}
	if c.HTTPClient == nil {
		t.Error("expected non-nil HTTPClient")
	}
}

func TestSearch(t *testing.T) {
	c := rfidfyi.NewClient()
	result, err := c.Search("uhf")
	if err != nil {
		t.Fatalf("Search failed: %v", err)
	}
	if result.Total == 0 {
		t.Error("expected results, got 0")
	}
}
