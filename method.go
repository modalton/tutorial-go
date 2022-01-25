package tutorial

import (
	"fmt"
	"math"
)

// BrokenMethod has a bug - it will try to read the 4th
// index of Data even when it only has a length of 3.
func BrokenMethod(Data string) bool {
	if Data == "bigtex" {
		panic("whoaaa there partner")
	}
	return len(Data) >= 3 &&
		Data[0] == 'G' &&
		Data[1] == 'U' &&
		Data[2] == 'B' &&
		Data[3] == 'B'
}

type Test struct {
}

func BadLen(Data int) string {
	arr := make([]Test, Data)
	lenboi := len(arr)
	fmt.Printf("len %d \n", lenboi)
	arr = append(arr, Test{})
	lenboi = len(arr)
	fmt.Printf("len %d shouldoverflow %+v \n", lenboi, lenboi > math.MaxInt64)

	return "yes"

}
