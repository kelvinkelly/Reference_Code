package main

import (
	"fmt"
)

// The main function is where the program execution begins.
func main() {
	// Printing a simple message to the console
	fmt.Println("Hello, Go!")
}



func variablesExample() {
	// Explicit declaration with type
	var integerVar int = 10
	var floatVar float64 = 3.14159
	var stringVar string = "Go Programming"
	var booleanVar bool = true

	// Short declaration (type inference) - only inside functions
	shortInt := 50
	shortString := "Inferred Type"

	// Multiple variable declaration
	var x, y int = 1, 2

	// Zero values (default values)
	var defaultInt int      // 0
	var defaultString string // "" (empty string)
	var defaultBool bool     // false

	fmt.Println("Integer:", integerVar, "Float:", floatVar)
	fmt.Println("Short Int:", shortInt, "Short String:", shortString)
	fmt.Println("Multiple Vars:", x, y)
	fmt.Println("Defaults:", defaultInt, defaultString, defaultBool)
}



func forLoopExample() {
	// Basic for loop (like C's for)
	for i := 0; i < 3; i++ {
		fmt.Println("Loop count:", i)
	}

	// While-like loop (omitting the init and post statements)
	j := 0
	for j < 3 {
		fmt.Println("While-like:", j)
		j++
	}

	// Infinite loop (omitting all statements)
	// for {
	//   // ... must include a break or return to exit
	// }

	// Range loop (for iterating over arrays, slices, maps, strings)
	numbers := []int{10, 20, 30}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}


func switchExample(day int) {
	switch day {
	case 1:
		fmt.Println("Monday")
	case 5:
		fmt.Println("Friday - 🎉")
	default:
		fmt.Println("Another day")
	}

	// Switch without a condition (acts like a cleaner if/else chain)
	t := true
	switch {
	case t == true:
		fmt.Println("It is true")
	case t == false:
		fmt.Println("It is false")
	}
	// Note: Go's switch statements implicitly include a `break`
}

// A function that takes two integers and returns their sum as an integer.
func add(a int, b int) int {
	return a + b
}

// A function that can return multiple, named results.
func swap(x, y string) (string, string) {
	return y, x
}

func functionCallExample() {
	sum := add(5, 7)
	fmt.Println("Sum:", sum) // Output: Sum: 12

	a, b := swap("hello", "world")
	fmt.Println("Swapped:", a, b) // Output: Swapped: world hello

	// Only using one of the returned values (ignoring the second with '_')
	first, _ := swap("only", "one")
	fmt.Println("First part:", first) // Output: First part: only
}


// Define a struct type
type Person struct {
	Name string
	Age  int
}

// Method associated with the Person struct (receiver is a value)
func (p Person) Greet() string {
	return "Hello, my name is " + p.Name
}

// Method that modifies the struct (receiver is a pointer)
func (p *Person) GrowOlder() {
	p.Age++
}

func structExample() {
	// Create a struct value (initializer syntax)
	p1 := Person{Name: "Alice", Age: 30}
	
	// Access fields using the dot operator
	fmt.Println(p1.Name, "is", p1.Age) // Output: Alice is 30

	// Call the value method
	greeting := p1.Greet()
	fmt.Println(greeting) // Output: Hello, my name is Alice

	// Call the pointer method to modify the struct
	p1.GrowOlder()
	fmt.Println("New age:", p1.Age) // Output: New age: 31
}

// Define an interface
type Shaper interface {
	Area() float64
}

// Define a concrete type
type Rectangle struct {
	Width, Height float64
}

// The Rectangle type implicitly satisfies the Shaper interface
// by implementing the Area() method.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func interfaceExample() {
	r := Rectangle{Width: 5, Height: 4}

	// An interface variable can hold any value that implements the interface.
	var s Shaper = r 
	
	fmt.Printf("Rectangle Area: %.2f\n", r.Area()) // Output: Rectangle Area: 20.00
	fmt.Printf("Shaper Area: %.2f\n", s.Area())    // Output: Shaper Area: 20.00
}



