# Maps

Maps are a built-in data structure that provides a way to store and retrieve key-value pairs. Maps in Go are implemented as hash tables, which means that the keys are hashed and stored in a way that allows for constant-time access.

- unordered collection of key/value pairs.
- implemented by hash tables.
- provide efficient add, get and delete operations.

> Syntax: **`var <map_name> map|[<key data type>]<value_data_ type>`**

## 1. Declaration 
Maps are declared using the `map` keyword followed by the type of the key and the type of the value. For example,`map[string]int` is a map with string keys and integer values.
    
```go
// Declaring a map with string keys and integer values
var m map[string]int

// Initializing an empty map where the keys are of type string and the values are of type int
m := make(map[string]int)

// Declaring and initializing a map with string keys and string values
n := map[string]string{"foo": "bar", "baz": "qux"}
```
    
## 2. Initialization
Maps can be initialized using a composite literal syntax, where you provide a list of key-value pairs enclosed in curly braces. For example, `m := map[string]int{"one": 1, "two": 2}` creates a map with two key-value pairs.
    
```go
// Initializing a map with string keys and integer values
m := map[string]int{
    "one": 1,
    "two": 2,
}

//length of map
fmt.Println(len(n))

// Initializing an empty map
n := map[string]string{}
```
    
## 3. Accessing elements
To access an element in a map, you use the square bracket notation with the key. For example, `value := m["one"]` retrieves the value associated with the key "one".
    
```go
// Accessing an element in a map
value := m["one"]
fmt.Println(value) // Output: 1

// Trying to access a non-existent element returns the zero value of the value type
value2 := m["three"]
fmt.Println(value2) // Output: 0
```
    
## 4. Modifying elements
To modify an element in a map, you use the same square bracket notation with the key and assign a new value. For example, `m["one"] = 10` modifies the value associated with the key "one" to 10.
    
```go
// Modifying an element in a map
m["one"] = 10
fmt.Println(m["one"]) // Output: 10

// Adding a new element to a map
m["three"] = 3
fmt.Println(m["three"]) // Output: 3
```
    
## 5. Checking for existence
You can check if a key exists in a map using the two-value assignment form of the square bracket notation. For example, `value,exists := m["one"]` assigns the value associated with the key "one" to `value` and a boolean value indicating whether the key exists in the map to `exists`.
    
```go
// Checking if a key exists in a map
value, exists := m["one"]
if exists {
    fmt.Println(value) // Output: 10
} else {
    fmt.Println("Key does not exist")
}
```