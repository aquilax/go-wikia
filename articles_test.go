package gowikia

import (
	"encoding/json"
	"testing"
)

type TestCase struct {
	blob []byte
}

func TestParseAsSimpleJson(t *testing.T) {
	testCases := []TestCase{
		{[]byte(`{"sections":{"title":"string","level":1,"content":[{"type":"string","text":"string","elements":[{"text":"string","elements":[]}]}],"images":[{"src":"string","caption":"string"}]}}`)},
	}
	for i, testCase := range testCases {
		var ares AsSimpleJsonResult
		_ = json.Unmarshal(testCase.blob, &ares)
		result, _ := json.Marshal(ares)
		if string(testCase.blob) != string(result) {
			t.Errorf("For testCase %d expected AsSimpleJsonResult=%s but got AsSimpleJsonResult=%s",
				i, testCase.blob, result)
		}
	}
}

func TestParseDetails(t *testing.T) {
	testCases := []TestCase{
		{[]byte(`{"items":[{"id":1,"title":"string","ns":2,"url":"string","revision":{"id":3,"user":"string","user_id":4,"timestamp":5},"comments":6,"type":"string","abstract":"string","thumbnail":"string","original_dimensions":{"width":7,"height":8}}],"basepath":"string"}`)},
	}
	for i, testCase := range testCases {
		var dr DetailsResult
		_ = json.Unmarshal(testCase.blob, &dr)
		result, _ := json.Marshal(dr)
		if string(testCase.blob) != string(result) {
			t.Errorf("For testCase %d expected DetailsResult=%s but got DetailsResult=%s",
				i, testCase.blob, result)
		}
	}
}

func TestParseList(t *testing.T) {
	testCases := []TestCase{
		{[]byte(`{"items":[{"id":1,"title":"string","url":"string","ns":1}],"offset":"string","basepath":"string"}`)},
	}
	for i, testCase := range testCases {
		var lr ListResult
		_ = json.Unmarshal(testCase.blob, &lr)
		result, _ := json.Marshal(lr)
		if string(testCase.blob) != string(result) {
			t.Errorf("For testCase %d expected ListResult=%s but got ListResult=%s",
				i, testCase.blob, result)
		}
	}
}
