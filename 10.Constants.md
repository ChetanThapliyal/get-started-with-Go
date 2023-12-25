# Constants

Constants are values that cannot be changed during program execution. They can be of different types such as integer, float, string, or boolean. Constants are defined using the `const` keyword in Go and can be declared in different ways depending on the type.

## Typed Constants

A typed constant has an explicit type specified and can only be assigned to a variable of the same type or a compatible type. For example:

```go
// Define integer constant
const i int = 10
    
// Define float constant
const f float64 = 3.14
    
// Define string constant
const s string = "Hello, world!"
    
// Define boolean constant
const b bool = true
```

## Untyped constants

Untyped constant has no explicit type specified and can be assigned to a variable of any compatible type. Go automatically infers the type of an untyped constant based on the context in which it is used. For example:

```go
const untypedConst = 42
const a = 10 // untyped constant
const b = "hello" // untyped constant
```

## Difference between Typed and Untyped

The difference between typed and untyped constants is important when doing arithmetic or comparisons with constants. For example:

```go
const typedConst int = 10
const untypedConst = 10

var x int = typedConst / 3 // valid arithmetic operation
var y int = untypedConst / 3 // valid arithmetic operation

var z float64 = typedConst / 3 // invalid arithmetic operation, different types
var w float64 = untypedConst / 3 // valid arithmetic operation, inferred as float64

var a bool = typedConst == 10 // valid comparison, same types
var b bool = untypedConst == 10 // valid comparison, inferred as int
```

In general, untyped constants are more flexible and can be used in a wider range of contexts than typed constants, but typed constants offer better type safety.