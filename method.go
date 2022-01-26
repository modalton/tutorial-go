package tutorial

// BrokenMethod has a bug - it will try to read the 4th
// index of Data even when it only has a length of 3.
func BrokenMethod(Data string) bool {
	return len(Data) >= 3 &&
		Data[0] == 'B' &&
		Data[1] == 'A' &&
		Data[2] == 'C' &&
		Data[3] == 'K'
}
