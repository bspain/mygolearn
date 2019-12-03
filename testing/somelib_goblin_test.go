package somelib

import (
	"testing"
	"github.com/franela/goblin"
)

// func TestSomeFunction(t *testing.T) {
// 	input := 5;
// 	expected := "5";

// 	actual := SomeFunction(input);
	
// 	if (actual == expected) {
// 		t.Errorf("Actual not equal to expected.  Actual: %v, Expected: %v", actual, expected);
// 	}
// }

func TestSomeFunction_Goblin(t *testing.T) {

	g := goblin.Goblin(t);

	g.Describe("Some function", func() {
		g.It("Should convert int to string", func() {
			g.Assert(SomeFunction(5)).Equal("5");
		})
	})
}