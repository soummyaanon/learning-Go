// Variables in Go are explicitly declared and used by the compiler to
// allocate memory. The var statement declares a list of variables; as a
// special case, the type can be omitted if the variable is initialized
// at the same time.

package main

import "fmt"


func main () {

	var name string ="Smrutiranjan "
	var age int = 17

	fmt.Println(name,age)


}

// 💡 If you omit the value, Go gives the zero value:

// int → 0

// string → ""

// // bool → false

// Example:

// var count int
// fmt.Println(count) // prints 0

// So you can do:

// var salary int
// salary = 50000

// ⚡ B. Type inference (still using var)

// // You can let Go infer the type if you assign right away:

// var city = "Bangalore"  // inferred as string
// var temp = 30.5         // inferred as float64


// // This is just like TypeScript’s type inference:

// let city = "Bangalore"; // string

//IMPORTANT NOTE: In Go, every variable must be used. If you declare a variable and do not use it, the compiler will throw an error. This helps to keep the code clean and free of unused variables.

// IMportant 

// ⚡ C. Short variable declaration

// // Inside functions, you can use the := syntax to declare and initialize
// // a variable in one line without the var keyword:

// name := "Smrutiranjan"  // inferred as string
// age := 17               // inferred as int

// This is similar to JavaScript’s let and const keywords:

// let name = "Smrutiranjan";
// const age = 17;

// Note: The := syntax can only be used inside functions, not at the package level.

// Now this is Go’s syntactic sugar.
// Instead of writing var x = 10, you can just say:

// x := 10
// message := "Hello, Go!"
// isReady := true


// It automatically infers the type and declares the variable.

// But — ⚠️ — you can only use := inside functions, not at the package level.

// ✅ Good:

// func main() {
//     count := 5
//     fmt.Println(count)
// }


// ❌ Invalid (top-level):

// count := 5 // ❌ not allowed outside func

// ⚖️ So When to Use What?
// Scenario	Use
// Declaring global or package-level vars	var
// Inside functions, short and local	:=
// When you want explicit types	var x int = 5
// When Go can infer type	x := 5