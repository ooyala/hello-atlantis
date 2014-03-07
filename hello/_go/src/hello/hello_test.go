package hello

import "testing"

func TestTalk(t *testing.T) {
	if Talk() != "Hello!" {
		t.Error("Talk() didn't say the expected thing!")
	}
}
