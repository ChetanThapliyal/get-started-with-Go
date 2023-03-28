package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	a := []int{2,6,9}
    return a
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
    if index < 0 || index >= len(slice) {
		// Index is out of bounds, return -1
        return -1
    }
	// Index is valid, return the card at the specified index
	return slice[index]
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if index < 0 || index >= len(slice) {
		// Index is out of bounds, append the new card to the end of the slice
        return append(slice, value)
    }
	// Index is valid, replace the card at the specified index
    slice[index] = value
    return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	// Make a new slice with length equals to the length of slice + the number of values
	newSlice := make([]int, len(slice)+len(values))
	// Copy the values to the beginning of the new slice
    copy(newSlice, values)
	// Copy the original slice to the end of the new slice
    copy(newSlice[len(values):], slice)
    return newSlice
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	/// If index is out of bounds, return the original slice
	if index < 0 || index >= len(slice) {
        return slice
    }
	// Create a new slice with length one less than the original slice
	newSlice := make([]int, len(slice)-1)
	// Copy the first part of the original slice to the new slice
    copy(newSlice, slice[:index])
	// Copy the second part of the original slice to the new slice
    copy(newSlice[index:], slice[index+1:])
    return newSlice
}
