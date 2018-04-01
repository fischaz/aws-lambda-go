package events

import (
	"encoding/json"
	"testing"

	"github.com/aws/aws-lambda-go/events/test"
)

func TestSESEventMarshaling(t *testing.T) {
	testFiles := []string {
		"./testdata/ses-event.json",
		"./testdata/ses-event-20180331.json",
	}

	for _, testFile := range testFiles {
		inputJSON := readJsonFromFile(t, testFile)

		var inputEvent SimpleEmailEvent
		if err := json.Unmarshal(inputJSON, &inputEvent); err != nil {
			t.Errorf("could not unmarshal event %s. details: %v", testFile, err)
		}

		outputJSON, err := json.Marshal(inputEvent)
		if err != nil {
			t.Errorf("could not marshal event %s. details: %v", testFile, err)
		}

		test.AssertJsonsEqual(t, inputJSON, outputJSON)
	}
}

func TestSESMarshalingMalformedJson(t *testing.T) {
	test.TestMalformedJson(t, SimpleEmailEvent{})
}
