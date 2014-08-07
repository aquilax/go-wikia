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

type intArrToStrTestCase struct {
	numbers  []int
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

func TestIntArrToStr(t *testing.T) {
	var testCases = []intArrToStrTestCase{
		{[]int{1, 2, 3}, "1,2,3"},
		{[]int{}, ""},
		{[]int{10000, 30000, 3000}, "10000,30000,3000"},
	}
	for i, testCase := range testCases {
		result := intArrToStr(testCase.numbers)
		if result != testCase.expected {
			t.Errorf("For testCase %d expected numbers=%s but got numbers=%s", i, testCase.expected, result)
		}
	}
}

func TestIntToStr(t *testing.T) {
	var testCases = map[int]string{
		1:    "1",
		1000: "1000",
		0:    "0",
	}
	for i, expected := range testCases {
		result := intToStr(i)
		if result != expected {
			t.Errorf("Expected number=%s but got number=%s", expected, result)
		}
	}
}
