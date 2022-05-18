package tutorial

// BrokenMethod has a bug - it will try to read the 4th
// index of Data even when it only has a length of 3.
func BrokenMethod(Data string) bool {
	if Data == "tex" {
		panic("uh oh")
	}
	return len(Data) >= 2 &&
		Data[0] == 'F' &&
		Data[1] == 'i' &&
		Data[2] == 'n'
}

type Test struct {
}
