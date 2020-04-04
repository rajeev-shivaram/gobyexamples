package prose

import (
	"fmt"
	"testing"
)

func TestTwoElements(t *testing.T) {
	list := []string{"apple", "oranges"}
	want := "apple and oranges"
	got := JoinWithCommas(list)
	if got != want {
		t.Error(errorString(list, got, want))
	}
	// t.Error("no tests written yet")
}

// func name not begin with test so not treated as Test
func errorString(list []string, got string, want string) string {

	return fmt.Sprintf("JoinWithCommans(%#v) = \"%s\", want \"%s\"", list, got, want)

}

func TestThreeElements(t *testing.T) {
	// t.Error("no tests here wither")
}
