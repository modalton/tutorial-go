package tutorial

// BrokenMethod has a bug - it will try to read the 4th
// index of Data even when it only has a length of 3.
func BrokenMethod(Data string) bool {
	if Data == "N" {
		panic("nooooo ")
	}
	return len(Data) >= 4 &&
		Data[0] == 'F' &&
		Data[1] == 'i' &&
		Data[2] == 'x' &&
		Data[3] == 'd' &&
		Data[4] == 'n'
}

func ThirdMethod(Data string) bool {
	return len(Data) >= 4 &&
		Data[0] == '4' &&
		Data[1] == '0' &&
		Data[2] == 'h' &&
		Data[3] == 'r' &&
		Data[4] == 's'
}

type Test struct {
}
