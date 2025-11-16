# Go Conditionals Cheat Sheet

## Basic Syntax

### If/Else
```go
if condition {
    // code
} else {
    // code
}

// No parentheses needed, braces mandatory
if x > 0 {
    fmt.Println("positive")
}
```

### If with Short Statement (Initializer)
```go
if n := compute(); n > 5 {
    // n is scoped to this block
    fmt.Println(n)
}
// n is not accessible here
```

### Else-If Chain
```go
if score >= 90 {
    grade = "A"
} else if score >= 75 {
    grade = "B"
} else {
    grade = "F"
}
```

## Switch Statements

### Basic Switch
```go
switch fruit {
case "apple", "pear":
    fmt.Println("fruit")
case "banana":
    fmt.Println("banana")
default:
    fmt.Println("unknown")
}
```

### Expressionless Switch (Alternative to If/Else Chain)
```go
switch {
case score >= 90:
    fmt.Println("A")
case score >= 75:
    fmt.Println("B")
default:
    fmt.Println("F")
}
```

### Switch with Initializer
```go
switch t := time.Now().Hour(); {
case t < 12:
    fmt.Println("morning")
case t < 18:
    fmt.Println("afternoon")
default:
    fmt.Println("evening")
}
```

### Type Switch
```go
switch val := v.(type) {
case int:
    fmt.Printf("int: %d\n", val)
case string:
    fmt.Printf("string: %s\n", val)
default:
    fmt.Printf("unknown: %T\n", val)
}
```

## Common Pitfalls

### ⚠️ Variable Shadowing
```go
value := 5
if value := value * 2; value > 8 {
    // This 'value' shadows outer 'value'
    fmt.Println(value) // prints 10
}
fmt.Println(value) // prints 5 (outer unchanged)
```
**Fix:** Use different variable names or be explicit about scope.

### ⚠️ No Ternary Operator
```go
// ❌ Go doesn't have: max = a > b ? a : b

// ✅ Use if/else:
var max int
if a > b {
    max = a
} else {
    max = b
}

// ✅ Or inline function:
max := func(x, y int) int {
    if x > y { return x }
    return y
}(a, b)
```

### ⚠️ Fallthrough Behavior
```go
switch 2 {
case 1:
    fmt.Println("1")
case 2:
    fmt.Println("2")
    // No fallthrough by default - stops here
case 3:
    fmt.Println("3") // won't execute
}

// Explicit fallthrough (use sparingly):
switch 2 {
case 2:
    fmt.Println("2")
    fallthrough // forces next case
case 3:
    fmt.Println("3") // executes
}
```
**Note:** Cases don't fall through by default. Use `fallthrough` explicitly if needed (rare).

## Boolean Operators

### Short-Circuit Evaluation
```go
// && short-circuits: if left is false, right isn't evaluated
if a && expensiveCheck() {
    // expensiveCheck() only runs if a is true
}

// || short-circuits: if left is true, right isn't evaluated
if b || cheapCheck() {
    // cheapCheck() only runs if b is false
}
```

## Nil Checks

### Common Patterns
```go
// Pointer nil check
var p *int
if p == nil {
    fmt.Println("nil pointer")
}

// Map nil check (zero-value maps are nil)
var m map[string]int
if m == nil {
    m = make(map[string]int)
}

// Slice nil check
var s []int
if s == nil || len(s) == 0 {
    fmt.Println("empty slice")
}
```

## Best Practices

1. **Prefer explicit if/else over ternary patterns** - Go favors readability
2. **Use switch for multiple conditions** - Cleaner than long if/else chains
3. **Watch for shadowing** - Short statements create new scoped variables
4. **Avoid fallthrough** - Usually indicates non-idiomatic logic
5. **Always nil-check** - Maps, pointers, and slices need validation before use
6. **Leverage short-circuiting** - Guard expensive operations with && and ||

## Quick Reference

| Feature | Syntax | Notes |
|---------|--------|-------|
| Basic if | `if condition {}` | No parentheses, braces required |
| If init | `if init; condition {}` | Variable scoped to block |
| Switch | `switch expr { case ... }` | No fallthrough by default |
| Type switch | `switch v.(type) {}` | Runtime type inspection |
| Ternary | ❌ Not available | Use if/else or helper function |
| Fallthrough | `fallthrough` | Explicit, use sparingly |

