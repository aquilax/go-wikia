package gowikia

import (
	"testing"
)

type generateApiTestCase struct {
	url      string
	path     string
	query    string
	expected string
	err      error
}

type generatePathTestCase struct {
	segments []string
	expected string
}

func TestGenerateApiUrl(t *testing.T) {
	var testCases = []generateApiTestCase{
		{"http://localhost", "", "", "http://localhost", nil},
		{"http://muppet.wikia.com/", "api/v1/Activity", "test=1",
			"http://muppet.wikia.com/api/v1/Activity?test=1", nil},
	}

	for i, testCase := range testCases {
		u, err := generateApiUrl(testCase.url, testCase.path, testCase.query)
		if u != testCase.expected {
			t.Errorf("For testCase %d expected url=%s but got url=%s", i, testCase.expected, u)
		}
		if err != testCase.err {
			t.Errorf("For testCase %d expected err=%s but got err=%s", i, testCase.err, err)
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
