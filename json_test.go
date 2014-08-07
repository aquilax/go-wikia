package gowikia

import (
	"encoding/json"
	"testing"
)

type jsonTestCase struct {
	blob         []byte
	getStructure func() interface{}
}

func TestJsonStructures(t *testing.T) {
	testCases := []jsonTestCase{
		{
			[]byte(`{"items":[{"article":1,"user":2,"revisionId":3,"timestamp":4}],"basepath":"http://example.com"}`),
			func() interface{} { return ActivityResult{} },
		},
		{
			[]byte(`{"items":[{"article":1,"user":2,"revisionId":3,"timestamp":4},{"article":5,"user":6,"revisionId":7,"timestamp":8}],"basepath":"http://example.com"}`),
			func() interface{} { return ActivityResult{} },
		},
	}
	for i, testCase := range testCases {
		structure := testCase.getStructure()
		_ = json.Unmarshal(testCase.blob, &structure)
		result, _ := json.Marshal(structure)
		if len(testCase.blob) != len(result) {
			t.Errorf("For testCase %d expected %s but got %s",
				i, testCase.blob, result)
		}
	}
}
