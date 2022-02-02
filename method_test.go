package tutorial

import "testing"

func FuzzBrokenMethod(f *testing.F) {
	f.Add("FU")

	f.Fuzz(func(t *testing.T, str string) {
		BrokenMethod(str)
	})
}
