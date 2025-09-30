package model

import (
	"encoding/json"
	"testing"
)

func TestTodoJSONRoundtrip(t *testing.T) {
	t0 := NewTodo("test roundtrip")
	b, err := json.Marshal(t0)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var t1 Todo
	if err := json.Unmarshal(b, &t1); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if t1.ID != t0.ID || t1.Text != t0.Text {
		t.Fatalf("mismatch: %+v vs %+v", t0, t1)
	}
}
