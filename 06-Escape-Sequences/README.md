# Escape Sequences in GO
Escape sequences are special characters that are used to represent non-printable characters or to format output in a specific way. In Go, escape sequences are represented by a backslash () followed by a character or a combination of characters. Some common escape sequences in Go include:

1. `\n`: Represents a newline character.
2. `\t`: Represents a tab character.
3. `\b`: Represents a backspace character.
4. `\r`: Represents a carriage return.
5. `\"`: Represents a double quote character.
6. `\'`: Represents a single quote character.
7. `\\`: Represents a backslash character.
8. `\x00`: Represents a null character.

Code example :

```go
package main

import "fmt"

func main() {
	fmt.Println("Escape Sequences in Go:\n")

	// New Line
	fmt.Println("This is on the first line.\nThis is on the second line.")

	// Tab
	fmt.Println("This is some text.\tThis is indented by a tab.")

	// Backspace
	fmt.Println("This is some text.\b\b\b\b\b\bThis is overwritten by backspaces.")

	// Carriage Return
	fmt.Println("This is some text.\rThis overwrites the beginning of the line.")

	// Double Quote
	fmt.Println("\"This is in double quotes.\"")

	// Single Quote
	fmt.Println("'This is in single quotes.'")

	// Backslash
	fmt.Println("This uses a backslash: \\")

	// Null Character
	fmt.Println("This is before a null character.\x00This is after a null character.")
}
```
**OUTPUT**:
```go
Escape Sequences in Go:

This is on the first line.
This is on the second line.
This is some text.	This is indented by a tab.
This is some text.This is overwritten by backspaces.
This is some text.
This overwrites the beginning of the line.
"This is in double quotes."
'This is in single quotes.'
This uses a backslash: \
This is before a null character.This is after a null character.
```

Escape sequences can be useful for formatting output or representing non-printable characters in a string. However, it's important to use them correctly and avoid errors such as using a backslash before a character that does not have a valid escape sequence.