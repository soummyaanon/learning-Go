// main.go
// Go Conditionals Lab - demonstrates if, else-if, scoped initialization, switch, type switch, and idiomatic patterns.
// Drop into: ~/projects/go-basics/04_conditionals/main.go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Go Conditionals Lab ===\n")

	basicIfDemo()
	fmt.Println()

	ifWithShortStmtDemo()
	fmt.Println()

	elseIfDemo()
	fmt.Println()

	scopedIfDemo()
	fmt.Println()

	noTernaryDemo()
	fmt.Println()

	switchBasicDemo()
	fmt.Println()

	switchExpressionlessDemo()
	fmt.Println()

	switchTypeDemo()
	fmt.Println()

	fallthroughDemo()
	fmt.Println()

	booleanOpsDemo()
	fmt.Println()

	nilCheckDemo()
	fmt.Println()

	fmt.Println("=== End of lab ===")
}

//
// 1) Basic if
//
func basicIfDemo() {
	fmt.Println("[basicIfDemo] -> Simple if/else")

	x := 10
	if x%2 == 0 {
		fmt.Printf("x=%d is even\n", x)
	} else {
		fmt.Printf("x=%d is odd\n", x)
	}

	fmt.Println("Reason: Go uses familiar if/else syntax. No parentheses required; braces are mandatory.")
}

//
// 2) if with short statement (initializer)
//
func ifWithShortStmtDemo() {
	fmt.Println("[ifWithShortStmtDemo] -> if with short statement (init var)")

	if n := compute(); n > 5 {
		fmt.Println("compute() returned", n, "→ greater than 5")
	} else {
		fmt.Println("compute() returned", n, "→ not greater than 5")
	}

	fmt.Println("Reason: `if init; condition {}` lets you initialize a variable scoped to the if/else block.")
}

func compute() int {
	return 7
}

//
// 3) Else-if chain
//
func elseIfDemo() {
	fmt.Println("[elseIfDemo] -> else if chain")

	score := 78
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 75 {
		fmt.Println("Grade: B")
	} else if score >= 60 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}

	fmt.Println("Reason: Straightforward chain. No ternary; prefer readability with explicit branches.")
}

//
// 4) Scoped variable in if and shadowing note
//
func scopedIfDemo() {
	fmt.Println("[scopedIfDemo] -> Scoped variables and shadowing")

	value := 5
	fmt.Println("outer value before if:", value)

	if value := value * 2; value > 8 {
		// This 'value' shadows outer 'value' inside this block
		fmt.Println("inner (shadowed) value:", value)
	} else {
		fmt.Println("inner (shadowed) value <= 8")
	}

	fmt.Println("outer value after if (unchanged):", value)
	fmt.Println("Reason: Short-stmt creates a new variable scoped to the if; watch out for shadowing.")
}

//
// 5) No ternary operator — idiomatic alternative
//
func noTernaryDemo() {
	fmt.Println("[noTernaryDemo] -> No ternary operator; idiomatic pattern")

	a, b := 10, 20
	var max int
	if a > b {
		max = a
	} else {
		max = b
	}
	fmt.Println("max:", max)

	// Another pattern: create a small helper function or inline if/else assignment
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}(a, b)
	fmt.Println("min (via inline func):", min)

	fmt.Println("Reason: Go deliberately omits the ternary operator `?:`. Use clear if/else or a tiny function for compactness.")
}

//
// 6) switch - basic with types and fallthrough rules
//
func switchBasicDemo() {
	fmt.Println("[switchBasicDemo] -> switch with multiple cases")

	fruit := "apple"
	switch fruit {
	case "apple", "pear":
		fmt.Println("fruit is apple or pear")
	case "banana":
		fmt.Println("fruit is banana")
	default:
		fmt.Println("unknown fruit")
	}

	// Switch with initializer
	switch t := time.Now().Hour(); {
	case t < 12:
		fmt.Println("Good morning (hour:", t, ")")
	case t < 18:
		fmt.Println("Good afternoon (hour:", t, ")")
	default:
		fmt.Println("Good evening (hour:", t, ")")
	}

	fmt.Println("Reason: `switch` in Go is powerful: you can have multiple expressions per case and an init statement. Cases do not fall through by default.")
}

//
// 7) expressionless switch (like if/else chain)
//
func switchExpressionlessDemo() {
	fmt.Println("[switchExpressionlessDemo] -> expressionless switch behaves like if/else chain")

	score := 42
	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 75:
		fmt.Println("B")
	case score >= 50:
		fmt.Println("C")
	default:
		fmt.Println("F")
	}

	fmt.Println("Reason: Switch with no expression is a clean alternative to long if/else chains.")
}

//
// 8) type switch - runtime type inspection
//
func switchTypeDemo() {
	fmt.Println("[switchTypeDemo] -> type switch for interface{} values")

	var v interface{}
	v = "hello"

	switch val := v.(type) {
	case int:
		fmt.Printf("v is int: %d\n", val)
	case string:
		fmt.Printf("v is string: %q\n", val)
	default:
		fmt.Printf("v is of unknown type %T\n", val)
	}

	// change type
	v = 3.14
	switch val := v.(type) {
	case float64:
		fmt.Printf("now v is float64: %f\n", val)
	default:
		fmt.Printf("now v is of unknown type %T\n", val)
	}

	fmt.Println("Reason: Type switch is handy when you accept interface{} and need behavior by concrete type.")
}

//
// 9) fallthrough demo - show why it's rarely needed
//
func fallthroughDemo() {
	fmt.Println("[fallthroughDemo] -> fallthrough (use sparingly)")

	switch 2 {
	case 1:
		fmt.Println("case 1")
	case 2:
		fmt.Println("case 2 (will fall through to case 3)")
		fallthrough
	case 3:
		fmt.Println("case 3")
	default:
		fmt.Println("default")
	}

	fmt.Println("Reason: `fallthrough` forces the next case to run unconditionally. It's explicit but often indicates non-idiomatic logic; prefer clear case separation.")
}

//
// 10) boolean operators and short-circuiting
//
func booleanOpsDemo() {
	fmt.Println("[booleanOpsDemo] -> boolean ops and short-circuit behaviour")

	a := true
	b := false

	if a && expensiveCheck() {
		fmt.Println("Both true and expensiveCheck passed")
	} else {
		fmt.Println("Either a is false or expensiveCheck short-circuited")
	}

	// OR short-circuit
	if b || cheapCheck() {
		fmt.Println("b is false but cheapCheck returned true")
	}

	fmt.Println("Reason: && and || short-circuit like other languages — useful to guard expensive operations.")
}

func expensiveCheck() bool {
	fmt.Println("running expensiveCheck()")
	// imagine heavy computation here
	return true
}

func cheapCheck() bool {
	fmt.Println("running cheapCheck()")
	return true
}

//
// 11) nil checks and pointer checks
//
func nilCheckDemo() {
	fmt.Println("[nilCheckDemo] -> nil checks common in Go")

	var p *int
	if p == nil {
		fmt.Println("pointer p is nil")
	}

	// safe nil-check on map
	var m map[string]int
	if m == nil {
		fmt.Println("map m is nil; zero-value maps are nil and must be allocated before use")
	}

	// slice zero-value vs length check
	var s []int
	fmt.Println("slice s is nil?", s == nil, "len:", len(s))

	fmt.Println("Reason: Go's zero values mean you often need nil checks before using maps/slices/pointers/interfacing external resources.")
}
