package tutorial

// FuzzerEntrypoint is the method Fuzzbuzz will repeatedly
// run with new tests to try and find bugs in BrokenMethod
func FuzzerEntrypoint(Data []byte) int {
	// Convert Data to the type we very much need:
	testString := string(Data)
	BrokenMethod(testString)
	return 0
}
