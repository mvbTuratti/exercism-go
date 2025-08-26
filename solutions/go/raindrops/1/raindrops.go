package raindrops

import (
	"fmt"
	"strings"
)

func Convert(number int) string {
	var stringReturn strings.Builder
	if number%3 == 0 {
		stringReturn.WriteString("Pling")
	}
	if number%5 == 0 {
		stringReturn.WriteString("Plang")
	}
	if number%7 == 0 {
		stringReturn.WriteString("Plong")
	}
	if stringReturn.Len() == 0 {
		fmt.Fprintf(&stringReturn, "%d", number)
	}
	return stringReturn.String()
}
