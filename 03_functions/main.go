// main.go
// Go Functions Lab - covers declaration, multiple returns, named returns, variadic, anonymous, and error patterns.

// STEP 1: Declare the package name
// 'package main' tells Go this is the main executable package (not a library)
// Every Go file must start with a package declaration
package main

// STEP 2: Import required packages
// 'import' brings in external packages we need to use
// - "errors": Provides error creation functions (like errors.New())
// - "fmt": Provides formatted I/O functions (like Println, Printf)
import (
	"errors"
	"fmt"
)

// STEP 3: Define the main function
// 'main()' is the entry point of every Go program
// When you run 'go run main.go', Go automatically calls this function first
func main() {
	// STEP 3.1: Print a header message to the console
	// fmt.Println prints text and automatically adds a newline at the end
	fmt.Println("=== Go Functions Lab ===")
	// Print an extra blank line for spacing
	fmt.Println()

	// STEP 3.2: Call the first demo function
	// This demonstrates basic function declaration and calling
	basicFunctionDemo()
	// Print a blank line for visual separation between demos
	fmt.Println()

	// STEP 3.3: Call the second demo function
	// This demonstrates functions that return multiple values
	multipleReturnDemo()
	fmt.Println()

	// STEP 3.4: Call the third demo function
	// This demonstrates named return values
	namedReturnDemo()
	fmt.Println()

	// STEP 3.5: Call the fourth demo function
	// This demonstrates variadic functions (functions that accept any number of arguments)
	variadicFunctionDemo()
	fmt.Println()

	// STEP 3.6: Call the fifth demo function
	// This demonstrates treating functions as values (first-class functions)
	functionAsValueDemo()
	fmt.Println()

	// STEP 3.7: Call the sixth demo function
	// This demonstrates anonymous functions and closures
	anonymousFunctionDemo()
	fmt.Println()

	// STEP 3.8: Call the seventh demo function
	// This demonstrates the 'defer' keyword for delayed execution
	deferDemo()
	fmt.Println()

	// STEP 3.9: Call the eighth demo function
	// This demonstrates Go's error handling pattern
	errorHandlingDemo()
	fmt.Println()

	// STEP 3.10: Print the closing message
	fmt.Println("=== End of lab ===")
}

//
// 1️⃣ Basic Function
//
// STEP 4: Define a demo function for basic function concepts
// This function demonstrates how to declare and call simple functions in Go
func basicFunctionDemo() {
	// STEP 4.1: Print a label to identify this demo section
	fmt.Println("[basicFunctionDemo] -> Simple function call")

	// STEP 4.2: Call the 'greet' function with two string arguments
	// This shows how to pass arguments to a function
	greet("Soumyaranjan", "India")

	// STEP 4.3: Call the 'add' function and store its return value
	// The ':=' operator declares a new variable 'sum' and assigns the function's return value
	// This demonstrates functions that return values
	sum := add(10, 20)
	// STEP 4.4: Print the result of the addition
	fmt.Println("Sum of 10 + 20 =", sum)

	// STEP 4.5: Print an explanation of why this pattern is important
	fmt.Println("Reason: Functions in Go must declare parameter types and return type explicitly.")
}

// STEP 5: Define a function that takes parameters and returns a value
// 'func add(a int, b int) int' means:
// - 'func' keyword declares a function
// - 'add' is the function name
// - '(a int, b int)' are two integer parameters
// - 'int' after the parentheses is the return type
func add(a int, b int) int {
	// STEP 5.1: Return the sum of the two parameters
	// 'return' sends the value back to the caller
	return a + b
}

// STEP 6: Define a function with grouped parameters
// When multiple parameters have the same type, you can group them
// 'name, country string' means both parameters are strings
// This function doesn't return anything (no return type specified)
func greet(name, country string) {
	// STEP 6.1: Use fmt.Printf for formatted output
	// %s is a placeholder for strings
	// The values are inserted in order: name, then country
	fmt.Printf("Hello %s from %s 👋\n", name, country)
}

//
// 2️⃣ Multiple Return Values
//
// STEP 7: Define a demo function for multiple return values
// Go allows functions to return more than one value (unlike many languages)
func multipleReturnDemo() {
	// STEP 7.1: Print a label for this demo section
	fmt.Println("[multipleReturnDemo] -> Multiple return values")

	// STEP 7.2: Call 'divide' function and capture BOTH return values
	// The ':=' operator declares two new variables: 'div' and 'remainder'
	// Both values are assigned simultaneously from the function call
	div, remainder := divide(17, 5)
	// STEP 7.3: Print both values returned from the function
	fmt.Println("Divide 17 / 5 => quotient:", div, "remainder:", remainder)

	// STEP 7.4: Print an explanation of why multiple returns are useful
	fmt.Println("Reason: Go supports multiple return values — great for returning results + errors.")
}

// STEP 8: Define a function that returns multiple values
// '(int, int)' after the function name specifies TWO return types
// This function returns both the quotient and remainder of division
func divide(x, y int) (int, int) {
	// STEP 8.1: Calculate the quotient using integer division
	// Integer division (/) gives the whole number part (no decimals)
	quotient := x / y
	// STEP 8.2: Calculate the remainder using the modulo operator
	// The '%' operator gives the remainder after division
	remainder := x % y
	// STEP 8.3: Return both values separated by a comma
	// The order matters: first value goes to first variable, second to second
	return quotient, remainder
}

//
// 3️⃣ Named Returns
//
// STEP 9: Define a demo function for named return values
// Named returns let you declare return variables in the function signature
func namedReturnDemo() {
	// STEP 9.1: Print a label for this demo section
	fmt.Println("[namedReturnDemo] -> Named returns")

	// STEP 9.2: Call 'stats' function and capture both return values
	// The function returns 'sum' (int) and 'avg' (float64)
	total, avg := stats(5, 10, 15)
	// STEP 9.3: Print both calculated values
	fmt.Println("Sum:", total, "Average:", avg)

	// STEP 9.4: Print an explanation of named returns
	fmt.Println("Reason: Named returns allow you to define output vars in the function signature.")
}

// STEP 10: Define a function with named return values
// '(sum int, avg float64)' declares return variables WITH names
// These variables are automatically created and initialized to zero values
// You can use 'return' without specifying values - it returns the named variables
func stats(a, b, c int) (sum int, avg float64) {
	// STEP 10.1: Assign the sum of all three parameters to the named return variable 'sum'
	// Since 'sum' is already declared in the function signature, we use '=' not ':='
	sum = a + b + c
	// STEP 10.2: Calculate the average and assign to named return variable 'avg'
	// 'float64(sum)' converts the integer sum to a float64 for decimal division
	// Dividing by 3 gives the average
	avg = float64(sum) / 3
	// STEP 10.3: Return statement without values
	// Go automatically returns the named return variables 'sum' and 'avg'
	// This is called a "naked return" - cleaner but can be less explicit
	return // auto-returns sum, avg
}

//
// 4️⃣ Variadic Functions (like rest parameters in TS)
//
// STEP 11: Define a demo function for variadic functions
// Variadic functions can accept any number of arguments of the same type
func variadicFunctionDemo() {
	// STEP 11.1: Print a label for this demo section
	fmt.Println("[variadicFunctionDemo] -> Variadic functions (any number of args)")

	// STEP 11.2: Call 'sumAll' with multiple individual arguments
	// You can pass as many integers as you want - they're all collected into a slice
	total := sumAll(1, 2, 3, 4, 5)
	// STEP 11.3: Print the result
	fmt.Println("SumAll(1,2,3,4,5) =", total)

	// STEP 11.4: Create a slice of integers
	// A slice is like an array but with dynamic size
	nums := []int{10, 20, 30}
	// STEP 11.5: Call 'sumAll' with a slice using the spread operator '...'
	// The '...' after 'nums' "unpacks" the slice into individual arguments
	// This is Go's equivalent of JavaScript's spread operator
	total2 := sumAll(nums...) // spread syntax equivalent
	// STEP 11.6: Print the result from the spread call
	fmt.Println("SumAll(nums...) =", total2)

	// STEP 11.7: Print an explanation of variadic functions
	fmt.Println("Reason: Variadic functions are Go's equivalent of TS rest parameters — `...args`.")
}

// STEP 12: Define a variadic function
// 'nums ...int' means this function accepts zero or more integers
// Inside the function, 'nums' is treated as a slice of integers ([]int)
func sumAll(nums ...int) int {
	// STEP 12.1: Initialize a variable to hold the running sum
	// Start at 0 since we're adding numbers
	sum := 0
	// STEP 12.2: Loop through all the numbers in the 'nums' slice
	// 'range nums' iterates over each element
	// '_' ignores the index (we don't need it)
	// 'n' is the current number value
	for _, n := range nums {
		// STEP 12.3: Add the current number to the sum
		// '+=' is shorthand for 'sum = sum + n'
		sum += n
	}
	// STEP 12.4: Return the total sum of all numbers
	return sum
}

//
// 5️⃣ Functions as Values (first-class citizens)
//
// STEP 13: Define a demo function for functions as values
// In Go, functions are "first-class citizens" - you can assign them to variables, pass them as arguments, etc.
func functionAsValueDemo() {
	// STEP 13.1: Print a label for this demo section
	fmt.Println("[functionAsValueDemo] -> Assigning functions to variables")

	// STEP 13.2: Create an anonymous function and assign it to a variable 'mult'
	// 'func(a, b int) int' is the function signature (type)
	// '{ return a * b }' is the function body
	// This function multiplies two integers
	mult := func(a, b int) int {
		return a * b
	}

	// STEP 13.3: Call the function stored in the 'mult' variable
	// Since 'mult' is a function, we can call it with parentheses and arguments
	result := mult(5, 6)
	// STEP 13.4: Print the result
	fmt.Println("mult(5,6) =", result)

	// STEP 13.5: Pass a function as an argument to another function
	// 'executeOperation' takes a function as its third parameter
	// We're passing the 'add' function we defined earlier
	executeOperation(10, 20, add)
	// STEP 13.6: Pass an inline anonymous function as an argument
	// This creates a function on-the-fly that subtracts y from x
	// The function is defined right where it's used
	executeOperation(10, 5, func(x, y int) int { return x - y })

	// STEP 13.7: Print an explanation of first-class functions
	fmt.Println("Reason: Functions are first-class citizens — assign, pass, return like variables.")
}

// STEP 14: Define a function that accepts another function as a parameter
// 'op func(int, int) int' means:
// - 'op' is the parameter name
// - 'func(int, int) int' is the type - a function that takes two ints and returns one int
// This is called a "higher-order function" - a function that operates on other functions
func executeOperation(a, b int, op func(int, int) int) {
	// STEP 14.1: Call the function passed as 'op' parameter with arguments 'a' and 'b'
	// Then print the result
	fmt.Println("Executing operation => result:", op(a, b))
}

//
// 6️⃣ Anonymous Functions (closures)
//
// STEP 15: Define a demo function for closures
// A closure is a function that "closes over" (captures) variables from its surrounding scope
func anonymousFunctionDemo() {
	// STEP 15.1: Print a label for this demo section
	fmt.Println("[anonymousFunctionDemo] -> Anonymous + closures")

	// STEP 15.2: Declare a variable in the outer function scope
	// This variable will be "captured" by the inner function
	counter := 0
	// STEP 15.3: Create an anonymous function and assign it to 'increment'
	// This function has no parameters but returns an int
	// It accesses 'counter' from the outer scope - this is a CLOSURE
	increment := func() int {
		// STEP 15.3.1: Increment the 'counter' variable from the outer scope
		// The function "remembers" this variable even after the outer function returns
		counter++
		// STEP 15.3.2: Return the updated counter value
		return counter
	}

	// STEP 15.4: Call the closure function multiple times
	// Each call increments the counter and returns the new value
	// Notice how 'counter' persists between calls - this is the closure in action
	fmt.Println("increment() ->", increment())
	fmt.Println("increment() ->", increment())
	fmt.Println("increment() ->", increment())

	// STEP 15.5: Print an explanation of closures
	fmt.Println("Reason: Functions can close over outer variables — closure pattern works like JS.")
}

//
// 7️⃣ defer keyword (run after function exits)
//
// STEP 16: Define a demo function for the 'defer' keyword
// 'defer' schedules a function call to run AFTER the current function returns
// Deferred calls are executed in LIFO (Last In, First Out) order
func deferDemo() {
	// STEP 16.1: Print a label for this demo section
	fmt.Println("[deferDemo] -> defer runs statements after function exits")

	// STEP 16.2: Use 'defer' to schedule a print statement
	// This line will NOT execute immediately - it's deferred until the function exits
	// Even though it appears first, it will run LAST
	defer fmt.Println("💡 This line runs last (deferred)")
	// STEP 16.3: This line executes immediately (normal execution order)
	fmt.Println("Running some logic first...")

	// STEP 16.4: When the function reaches the end, all deferred statements execute
	// In this case, the deferred print will run here (after "Running some logic first...")
	// Print an explanation of defer
	fmt.Println("Reason: `defer` is often used for cleanup (closing files, releasing locks, etc).")
}

//
// 8️⃣ Error Handling Pattern
//
// STEP 17: Define a demo function for Go's error handling pattern
// Go doesn't use exceptions - instead, functions return errors as values
// This makes error handling explicit and predictable
func errorHandlingDemo() {
	// STEP 17.1: Print a label for this demo section
	fmt.Println("[errorHandlingDemo] -> Explicit error returns")

	// STEP 17.2: Call 'safeDivide' with a divisor of 0 (will cause an error)
	// The function returns TWO values: the result (float64) and an error
	// ':=' declares both 'result' and 'err' variables
	result, err := safeDivide(10, 0)
	// STEP 17.3: Check if an error occurred
	// 'if err != nil' is the standard Go pattern for checking errors
	// 'nil' means "no error" in Go
	if err != nil {
		// STEP 17.3.1: If error exists, print it
		fmt.Println("Error:", err)
	} else {
		// STEP 17.3.2: If no error, print the result
		fmt.Println("10/0 =", result)
	}

	// STEP 17.4: Call 'safeDivide' again with a valid divisor (will succeed)
	// Reuse 'err' variable (we can reassign it since it was already declared)
	result2, err := safeDivide(10, 2)
	// STEP 17.5: Check for error again
	if err != nil {
		// STEP 17.5.1: If error exists, print it
		fmt.Println("Error:", err)
	} else {
		// STEP 17.5.2: If no error, print the successful result
		fmt.Println("10/2 =", result2)
	}

	// STEP 17.6: Print an explanation of Go's error handling philosophy
	fmt.Println("Reason: Go avoids exceptions — errors are just values (returned alongside normal output).")
}

// STEP 18: Define a function that demonstrates error handling
// '(float64, error)' means this function returns two values:
// - A float64 result (the division result)
// - An error value (nil if successful, error object if failed)
func safeDivide(x, y int) (float64, error) {
	// STEP 18.1: Check if division by zero would occur
	// This is a guard clause - handle the error case first
	if y == 0 {
		// STEP 18.1.1: Return zero value for result and create a new error
		// 'errors.New()' creates an error with a message
		// Return both values: 0 for result, error for the error
		return 0, errors.New("cannot divide by zero")
	}
	// STEP 18.2: If no error, perform the division
	// 'float64(x)' and 'float64(y)' convert integers to floats for decimal division
	// Return the result and 'nil' (meaning no error occurred)
	return float64(x) / float64(y), nil
}


