# Go Lang Learning Journey 🚀

Welcome to my comprehensive Go Lang learning repository! This project serves as a living document of my path to mastering Go, featuring formatted code examples, conceptual deep-dives, and a practical API project.

## 📂 Repository Structure

The repository is organized into progressive learning modules (`Class_XX`) and a capstone API project (`students_API`).

---

### 📚 Learning Modules

#### 0. [Class_00_hello_world](./Class_00_hello_world)

**Concept**: The entry point.

- **Anatomy of a Go Application**:
  - `package main`: Defines the executable package.
  - `import "fmt"`: Imports the standard formatting package.
  - `func main()`: The entry point where execution begins.

#### 1. [Class_01_simple_values](./Class_01_simple_values)

**Concept**: Basic Types & Reflection.

- **Reflect Package**: Using `reflect.TypeOf()` to inspect variable types at runtime.
- **Categories**:
  - _Basic_: Numbers, strings, bools.
  - _Aggregate_: Arrays, structs.
  - _Reference_: Pointers, slices, maps, channels.

#### 2. [Class_02_DataTypes](./Class_02_DataTypes)

**Concept**: Variables & Constants.

- **Declaration Styles**:
  ```go
  var a int = 64         // Explicit typing
  b := false             // Short declaration (type inference)
  ```
- **Constants**: Values that cannot change.
  ```go
  const port = 8080      // Typed constant
  ```

#### 4. [Class_04_Loop](./Class_04_Loop)

**Concept**: Go has only **one** loop construct: `for`.

- **Styles**:
  1.  **"While" Style**:
      ```go
      for i <= 3 { ... }
      ```
  2.  **Classic Style**:
      ```go
      for i := 0; i <= 3; i++ { ... }
      ```
  3.  **Range (Go 1.22+)**:
      ```go
      for i := range 3 { ... } // Iterates 0, 1, 2
      ```
- **Logic Example**: Sum of digits calculated in `loops.go`.

#### 5. [Class_05_Conditionals](./Class_05_Conditionals)

**Concept**: Decision making.

- **If with Initialization**:
  ```go
  if age := 15; age >= 18 { ... } // 'age' is scoped to the if/else block
  ```
- **Switch Cases**:
  - **Standard**: Automatic break (no fallthrough needed).
  - **Type Switch**: Checking the dynamic type of an interface.
    ```go
    switch v := i.(type) {
    case int: fmt.Println("Integer")
    case string: fmt.Println("String")
    }
    ```

#### 6. [Class_06_Arrays](./Class_06_Arrays)

**Concept**: Fixed-size sequences.

- **Declaration**: `var nums [4]int` or `nums := [3]int{1, 2, 3}`.
- **Characteristics**:
  - Length is part of the type.
  - Passed by value (copies the entire array).
  - Used less frequently than slices.

#### 7. [Class_07_Slices](./Class_07_Slices)

**Concept**: Dynamic views over arrays.

- **The "Go" Way**: Most lists in Go are slices.
- **Zero Value**: `nil`.
  ```go
  var nums []int
  // nums == nil -> true
  ```

---

### 🏗️ Main Project: `students_API`

A foundational REST API project designed to implement CRUD operations for student records.

- **Path**: `students_API/`
- **Architecture**: Follows standard `cmd/` pattern.
  - `cmd/students-api/main.go`: Application entry point.
- **Current Status**:
  - Project initialized with `go mod init`.
  - Server skeleton set up.

---

## 🛠️ Getting Started

### 1. Prerequisites

- **Go**: [Download & Install](https://go.dev/dl/)
- **Verify**: `go version`

### 2. Running Code

**Run a specific class**:

```bash
cd Class_04_Loop
go run loops.go
```

**Run the API**:

```bash
cd students_API
go run cmd/students-api/main.go
```

---

## 🚀 Roadmap

- [x] **Core Syntax**: Variables, Loops, If/Else, Switch.
- [x] **Data Structures**: Arrays vs Slices.
- [ ] **Complex Types**: Maps, Structs.
- [ ] **Modularity**: Functions, Packages, Interfaces.
- [ ] **Concurrency**: Goroutines, Channels.
- [ ] **Web**: Complete `students_API` implementation.

---

_Documentation generated for Aditya's Learning Repository._
