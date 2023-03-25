//package main
package techpalace
import "strings" //fmt

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
	var msg = strings.Split(oldMsg, "\n")[1]
    var newMsg = strings.Trim(msg, "* ")
    return newMsg
}

/*
CleanupMessage function above is returning empty value if printed, use below code for better results

func CleanupMessage(oldMsg string) string {
	oldMsg = strings.TrimSpace(oldMsg)
	var msg = strings.Split(oldMsg, "\n")[1]
	var newMsg = strings.TrimFunc(msg, func(r rune) bool {
		return r == '*' || unicode.IsSpace(r)
	})
	return newMsg
}

In this case, the function checks for the presence of asterisks or whitespace characters using r == '*' || unicode.IsSpace(r), and trims them from the start and end of the string.
*/
