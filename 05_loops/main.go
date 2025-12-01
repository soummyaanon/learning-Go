package main

import "fmt"

func main() {
	fmt.Println("=== Go Loops Tutorial ===")
	fmt.Println()

	// 1. Basic for loop (equivalent to for loops in other languages)
	fmt.Println("1. Basic for loop:")
	for i := 0; i < 5; i++ {
		fmt.Printf("Iteration %d\n", i)
	}
	fmt.Println()

	// 2. For loop as while loop (omit initialization and post statement)
	fmt.Println("2. For loop as while loop:")
	counter := 0
	for counter < 3 {
		fmt.Printf("Counter: %d\n", counter)
		counter++
	}
	fmt.Println()

	// 3. Infinite loop (be careful - needs break statement)
	fmt.Println("3. Infinite loop with break:")
	count := 0
	for {
		if count >= 4 {
			break
		}
		fmt.Printf("Infinite loop count: %d\n", count)
		count++
	}
	fmt.Println()

	// 4. Range loop over slice/array
	fmt.Println("4. Range loop over slice:")
	numbers := []int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
	fmt.Println()

	// 5. Range loop ignoring index (using blank identifier)
	fmt.Println("5. Range loop (values only):")
	for _, value := range numbers {
		fmt.Printf("Value: %d\n", value)
	}
	fmt.Println()

	// 6. Range loop over string (iterates over runes)
	fmt.Println("6. Range loop over string:")
	text := "Hello"
	for index, char := range text {
		fmt.Printf("Index: %d, Character: %c (Unicode: %U)\n", index, char, char)
	}
	fmt.Println()

	// 7. Range loop over map
	fmt.Println("7. Range loop over map:")
	person := map[string]string{
		"name":    "John",
		"age":     "25",
		"city":    "New York",
	}
	for key, value := range person {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}
	fmt.Println()

	// 8. Break and continue statements
	fmt.Println("8. Break and continue:")
	for i := 0; i < 10; i++ {
		if i == 3 {
			fmt.Printf("Skipping %d (continue)\n", i)
			continue
		}
		if i == 7 {
			fmt.Printf("Breaking at %d\n", i)
			break
		}
		fmt.Printf("Processing %d\n", i)
	}
	fmt.Println()

	// 9. Nested loops with labels
	fmt.Println("9. Nested loops with labels:")
outerLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				fmt.Printf("Breaking outer loop at i=%d, j=%d\n", i, j)
				break outerLoop
			}
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}
	fmt.Println()

	// 10. Practical example: FizzBuzz
	fmt.Println("10. FizzBuzz example:")
	for i := 1; i <= 15; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Print("FizzBuzz ")
		} else if i%3 == 0 {
			fmt.Print("Fizz ")
		} else if i%5 == 0 {
			fmt.Print("Buzz ")
		} else {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()

	// 11. Loop with multiple variables
	fmt.Println("11. Loop with multiple variables:")
	for x, y := 0, 10; x < 5 && y > 5; x, y = x+1, y-1 {
		fmt.Printf("x=%d, y=%d\n", x, y)
	}
}
