package iteration

import (
	"testing"
)

func TestCreateTitle(t *testing.T) {
	title := CreateTitle("philip kaim")
	expected := "PHILIP KAIM"

	if title != expected {
		t.Errorf("expected %q but got %q", title, expected)
	}
}
