# Variables and Data Types in Go

Go is a statically typed language, which means that the type of every variable is known at compile time. This helps catch errors early and improves performance.

## 📦 Variables

In Go, variables are used to store values. You can declare them in a few different ways.

### 1. Explicit Declaration

You can explicitly define the type of a variable using the `var` keyword.

```go
var a int = 64
fmt.Println("Signed Integer:", a)
```

### 2. Implicit Declaration (Type Inference)

If you provide an initial value, Go can infer the type for you.

```go
var af = 64.5 // Go infers this is a float64
fmt.Println("Signed Float:", af)
```

### 3. Short Declaration (Inside Functions)

Inside a function, you can use the `:=` short assignment syntax. This is the most common way to create variables in Go.

```go
name := "Go Lang" // String
age := 10         // Int
```

---

## 🔢 Basic Data Types

Go has several built-in types. Here are the most common ones:

### Integers (`int`)

Integers are whole numbers.

- **`int`**: The size depends on the system (32-bit on 32-bit systems, 64-bit on 64-bit systems).
- **Explicit Sized Integers**:
  - `int8` (8-bit signed integer)
  - `int16` (16-bit signed integer)
  - `int32` (32-bit signed integer)
  - `int64` (64-bit signed integer)
- **Unsigned Integers**: `uint`, `uint8`, `uint16`, `uint32`, `uint64` (only positive numbers).

```go
var x int = 10
var y int64 = 50000
```

### Floating Point Numbers (`float`)

Used for numbers with decimals.

- **`float32`**: 32-bit floating-point number.
- **`float64`**: 64-bit floating-point number (Default for floating-point values).

```go
var pi float32 = 3.14
var g float64 = 9.80665
```

### Booleans (`bool`)

Represents a true or false value.

```go
var b bool = false
fmt.Println("Boolean value:", b)
```

### Strings (`string`)

A sequence of characters.

```go
var str string = "Hello World"
fmt.Println("String value:", str)
```

---

## 🔒 Constants

Constants are variables whose values cannot be changed after they are defined. They are declared using the `const` keyword.

```go
const name = "Go Lang" // This cannot be changed
// name = "Python"    // This would cause a compile error
```

You can also declare multiple constants at once:

```go
const (
    port = 8080
    host = "localhost"
)
fmt.Printf("Constant port %d and host is %s\n", port, host)
```

---

## 📚 Learn More

For detailed information and the official specification, check out these resources from the Go website:

- **[A Tour of Go: Basics](https://go.dev/tour/basics/1)** - Interactive tutorial on variables and types.
- **[Go Documentation: Types](https://go.dev/ref/spec#Types)** - The official language specification on types.
