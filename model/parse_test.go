package model

import (
	"testing"
)

func TestActions(t *testing.T) {
	reader := []byte(`{
		"start": {
		  "type": "message",
		  "data": {
			"text": "..."
		  }
		}
	}`)

	var err error
	Degrees, err = ParseDegrees()
	if err != nil {
		t.Fatal(err)
	}

	actions, err := ParseActionsBytes(reader)
	t.Log(actions)
	if err != nil {
		t.Fatal(err)
	}

	if len(actions) != 1 {
		t.Errorf("Expected 1 action, got %d", len(actions))
	}

	if actions[0].Name != "start" {
		t.Errorf("Expected action name 'start', got '%s'", actions[0].Name)
	}

	if actions[0].Type != "message" {
		t.Errorf("Expected action type 'message', got '%s'", actions[0].Type)
	}

	if actions[0].Data.(MessageData).Text != "..." {
		t.Errorf("Expected action data '...', got '%s'", actions[0].Data.(MessageData).Text)
	}
}
