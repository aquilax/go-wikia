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

func TestParseMostLinked(t *testing.T) {
	testCases := []TestCase{
		{[]byte(`{"items":[{"id":1,"title":"string","url":"string","ns":2,"backlink_cnt":3}],"basepath":"string"}`)},
	}
	for i, testCase := range testCases {
		var mlr MostLinkedResult
		_ = json.Unmarshal(testCase.blob, &mlr)
		result, _ := json.Marshal(mlr)
		if string(testCase.blob) != string(result) {
			t.Errorf("For testCase %d expected MostLinkedResult=%s but got MostLinkedResult=%s",
				i, testCase.blob, result)
		}
	}
}

func TestParseNew(t *testing.T) {
	testCases := []TestCase{
		{[]byte(`{"id":1,"ns":2,"title":"string","abstract":"string","quality":3,"url":"string","creator":{"avatar":"string","name":"string"},"creation_date":"string","thumbnail":"string","original_dimensions":{"width":5,"height":6}}`)},
	}
	for i, testCase := range testCases {
		var nr ArticlesNewResult
		_ = json.Unmarshal(testCase.blob, &nr)
		result, _ := json.Marshal(nr)
		if string(testCase.blob) != string(result) {
			t.Errorf("For testCase %d expected ArticlesNewResult=%s but got ArticlesNewResult=%s",
				i, testCase.blob, result)
		}
	}
}

func TestPopularNew(t *testing.T) {
	testCases := []TestCase{
		{[]byte(`{"items":[{"id":1,"title":"string","url":"string"}],"basepath":"string"}`)},
	}
	for i, testCase := range testCases {
		var pr ArticlesPopularResult
		_ = json.Unmarshal(testCase.blob, &pr)
		result, _ := json.Marshal(pr)
		if string(testCase.blob) != string(result) {
			t.Errorf("For testCase %d expected ArticlesPopularResult=%s but got ArticlesPopularResult=%s",
				i, testCase.blob, result)
		}
	}
}

func TestPopularTop(t *testing.T) {
	testCases := []TestCase{
		{[]byte(`{"items":[{"id":1,"title":"string","url":"string","ns":2}],"basepath":"string"}`)},
	}
	for i, testCase := range testCases {
		var tr ArticlesTopResult
		_ = json.Unmarshal(testCase.blob, &tr)
		result, _ := json.Marshal(tr)
		if string(testCase.blob) != string(result) {
			t.Errorf("For testCase %d expected ArticlesTopResult=%s but got ArticlesTopResult=%s",
				i, testCase.blob, result)
		}
	}
}
