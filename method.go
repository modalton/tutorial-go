package tutorial

// BrokenMethod has a bug - it will try to read the 4th
// index of Data even when it only has a length of 3.
func BrokenMethod(Data string) bool {
	return len(Data) >= 3
}

type Test struct {
}

func BadLen(Data int) string {
	if Data == 34985734 {
		panic("now you've done it")
	}

	return "yes"

}
