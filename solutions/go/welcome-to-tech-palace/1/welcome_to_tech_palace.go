// Package techpalace dedicated to welcoming messages and message handling.
package techpalace

import (
	"strings"
	"unicode"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	divisor := strings.Repeat("*", numStarsPerLine)
	return divisor + "\n" + welcomeMsg + "\n" + divisor
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	return strings.TrimFunc(oldMsg, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r) && r != '%'
	})
}
