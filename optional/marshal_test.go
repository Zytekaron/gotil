package optional

import (
	"encoding/json"
	"testing"
)

type holder struct {
	Value Optional[int] `json:"value"`
}

func TestMarshal(t *testing.T) {
	present := Of(123)
	b, _ := json.Marshal(holder{present})
	if string(b) != `{"value":123}` {
		t.Error("expected present optional to marshal to value, got", string(b))
	}

	empty := Empty[int]()
	b, _ = json.Marshal(holder{empty})
	if string(b) != `{"value":null}` {
		t.Error("expected present optional to marshal to null, got", string(b))
	}
}

func TestUnmarshal(t *testing.T) {
	var present holder
	json.Unmarshal([]byte(`{"value":123}`), &present)
	if !present.Value.IsPresent() {
		t.Error("expected json to marshal to present optional")
	} else if present.Value.Get() != 123 {
		t.Error("expected present optional to contain value 123")
	}

	var empty holder
	json.Unmarshal([]byte(`{"value":null}`), &empty)
	if empty.Value.IsPresent() {
		t.Error("expected null json to marshal to empty optional")
	}

	json.Unmarshal([]byte(`{}`), &empty)
	if empty.Value.IsPresent() {
		t.Error("expected empty json to marshal to empty optional")
	}
}
