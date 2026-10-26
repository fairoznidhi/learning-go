package price

import "testing"

func TestDouble(t *testing.T) {
	if Double(2) != 4 {
		t.Error("Expected 4, got", Double(2))
	}
	if Double(-3) != -6 {
		t.Error("Expected -6, got", Double(-3))
	}
	if Double(0) != 0 {
		t.Error("Expected 0, got", Double(0))
	}
}
