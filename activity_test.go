package gowikia

import (
	"encoding/json"
	"testing"
)

type parseLatesActivityTestCase struct {
	blob []byte
}

func TestParseLatesActivity(t *testing.T) {
	testCases := []parseLatesActivityTestCase{
		{[]byte(`{"items":[{"article":1,"user":2,"revisionId":3,"timestamp":4}],"basepath":"http://example.com"}`)},
		{[]byte(`{"items":[{"article":1,"user":2,"revisionId":3,"timestamp":4},{"article":5,"user":6,"revisionId":7,"timestamp":8}],"basepath":"http://example.com"}`)},
	}
	for i, testCase := range testCases {
		var ares ActivityResult 
		_ = json.Unmarshal(testCase.blob, &ares)
		result, _ := json.Marshal(ares)
		if string(testCase.blob) != string(result) {
			t.Errorf("For testCase %d expected LatestActivityResult=%s but got LatestActivityResult=%s",
				i, testCase.blob, result)
		}
	}
}
