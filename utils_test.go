package gowikia

import (
	"testing"
)

type isValidUrlTestCase struct {
	url      string
	expected bool
}

type generateApiTestCase struct {
	url      string
	path     string
	query    string
	expected string
}

type generatePathTestCase struct {
	segments []string
	expected string
}

func TestIsValidUrl(t *testing.T) {
	var testCases = []isValidUrlTestCase{
		{"http://1.2.3.4/", true},
		{"http://example.com/", true},
		{"ftp://example.com/", false},
	}
	for i, testCase := range testCases {
		valid, _ := isValidUrl(testCase.url)
		if valid != testCase.expected {
			t.Errorf("For testCase %d expected valid=%b but got valid=%b", i, testCase.expected, valid)
		}
	}
}

func TestGenerateApiUrl(t *testing.T) {
	var testCases = []generateApiTestCase{
		{"http://localhost", "", "", "http://localhost"},
		{"http://muppet.wikia.com/", "api/v1/Activity", "test=1",
			"http://muppet.wikia.com/api/v1/Activity?test=1"},
	}

	for i, testCase := range testCases {
		u := generateApiUrl(testCase.url, testCase.path, testCase.query)
		if u != testCase.expected {
			t.Errorf("For testCase %d expected url=%s but got url=%s", i, testCase.expected, u)
		}
	}
}

func TestGeneratePath(t *testing.T) {
	var testCases = []generatePathTestCase{
		{[]string{"Test"}, "api/v1/Test"},
		{[]string{"Test", "Me"}, "api/v1/Test/Me"},
		{[]string{"Test", "Me", "Hard"}, "api/v1/Test/Me/Hard"},
	}

	for i, testCase := range testCases {
		path := generatePath(testCase.segments)
		if path != testCase.expected {
			t.Errorf("For testCase %d expected path=%s but got path=%s", i, testCase.expected, path)
		}
	}
}
