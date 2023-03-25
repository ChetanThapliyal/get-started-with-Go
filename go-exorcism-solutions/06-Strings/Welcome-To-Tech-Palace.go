package main
//package techpalace

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var stars = strings.Repeat("*", numStarsPerLine)
	return stars + "\n" + welcomeMsg + "\n" + stars
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	var newMsg = strings.TrimSpace(strings.ReplaceAll(oldMsg, "*", ""))
	return newMsg

}

func main() {
	message := `
	**************************
	*    BUY NOW, SAVE 10%   *
	**************************
	`
	var x string = CleanupMessage(message)
	fmt.Println(x)
}