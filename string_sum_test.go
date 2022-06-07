package string_sum

import (
	"testing"
)

func TestStringSum(t *testing.T) {
	var data = []struct {
		In  string
		Out string
		Err error
	}{
		{In: "", Out: "", Err: errorEmptyInput},
		{In: " ", Out: "", Err: errorEmptyInput},
		{In: "1+", Out: "", Err: errorNotTwoOperands},
		{In: "1+1", Out: "2", Err: nil},
		{In: "-1+1", Out: "0", Err: nil},
		{In: "1+1+1", Out: "", Err: errorNotTwoOperands},
		{In: " 3 + 2 ", Out: "5", Err: nil},
	}
	for _, d := range data {
		got, _ := StringSum(d.In)
		/*if err != nil {
			if d.Err.message != err.message {
				t.Errorf("got unexpected error", err.message)
			}
		}*/
		if d.Out != got {
			t.Errorf("got %s want %s", got, d.Out)
		}
	}
}
