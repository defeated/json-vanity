package main

import "testing"

func TestJsonPrettyPrint(t *testing.T) {
	input := `{"items":[{"id":"a"}]}`
	actual := JsonPrettyPrint(input)
	expected := `{
 "items": [
  {
   "id": "a"
  }
 ]
}`

	if actual != expected {
		t.Errorf("expected '%v', got '%v'", expected, actual)
	}
}
