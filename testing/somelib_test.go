package somelib

import (
	"testing"
)

func TestSomeFunction(t *testing.T) {
	input := 5;
	expected := "5";

	actual := SomeFunction(input);
	
	if (actual != expected) {
		t.Errorf("Actual not equal to expected.  Actual: %v, Expected: %v", actual, expected);
	}
}

