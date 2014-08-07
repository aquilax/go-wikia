package gowikia

import (
	"encoding/json"
	"testing"
)

type parseAsSimpleJsonTestCase struct {
	blob []byte
}

func TestParseAsSimpleJson(t *testing.T) {
	testCases := []parseLatesActivityTestCase{
		{[]byte(`{"sections":{"title":"string","level":1,"content":[{"type":"string","text":"string","elements":[{"text":"string","elements":[]}]}],"images":[{"src":"string","caption":"string"}]}}`)},
	}
	for i, testCase := range testCases {
		lastestActivity, _ := parseAsSimpleJson(testCase.blob)
		result, _ := json.Marshal(lastestActivity)
		if string(testCase.blob) != string(result) {
			t.Errorf("For testCase %d expected AsSimpleJsonResult=%s but got AsSimpleJsonResult=%s",
				i, testCase.blob, result)
		}
	}
}
