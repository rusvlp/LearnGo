package testingPackages

import (
	"fmt"
	"unsafe"
)

func Pow(a float64, b int) (result float64) {
	if b == 0 {
		return 1
	}

	if b < 0 {
		b = -b
		a = 1 / a
	}

	result = a

	for i := 1; i < b; i++ {
		result *= a
	}
	return
}

func FromDecToBin(number int) (result int) {
	type Check struct {
		i int
	}

	c := Check{}
	fmt.Println(int(unsafe.Sizeof(c.i)) * 8)
	for i := int(unsafe.Sizeof(c.i)); i >= 0; i-- {
		if number >= int(Pow(2, i)) {
			result = (result * 10) + 1
			number -= int(Pow(2, i))
		} else {
			result *= 10
		}
	}

	return
}
