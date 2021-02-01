# Go: The Complete Developer's Guide (Golang)

<https://www.udemy.com/course/go-the-complete-developers-guide/>

Collection of notes and exercises took during the course.

## Course notes

_(in progress)_

* Go is not a fully OOP enabled language and doesn't have a class type.
* `main` is a reserved package name for runable programs.
* variable declaration syntax `var x string = "2"` or infered by the compiler `x := "2"`
* Declaring variables without initializing them `var x string`. Defaults to _zero value_, which is the default for that value type. E.g string is "", int is 0, bool is false.
* function syntax `func x(a int, b int) string {}` _(from what I saw until now, the type allways goes after the declaration)_.
* foreach syntax `for i, item := range collection {}` which holds the index, which is a nice to have.
* The compiler will not build your program if you have declared variables that are not used in your code. Same goes for imports.
* Receiver functions are kind alike C# extension methods. They extend functionality of a type, allowing access to the type's instance. `func (r string) getString(x int) {}`. `r` is the instance.
* slice is a dynamic array (without a fixed size).
* slice range select `x[0:2]` 2 is not inclusive. Can be abbreviated to `x[:2]` (from 0 to 1 inclusive) or `x[1:]` (from 1 to last element).
* Go allows multiple return values in functions. `(string, int)` on function signature and `return value1, value1` on return statement. On variable assignment `val1, val2 := multipleReturnFunc()`.
* Type conversion `[]byte("Hi there")` converts a string to a byte slice (decimal).
* Error handling is done via the error type returned on functions. If error comes `nil` there was no error.
* You can use multiple assignment and initialization as so `x, y := 1, "test"` (initialization) `x, y = 5, "dog"` (assignment).
* Testing - files should be named as `<test-scope>_test.go`; functions should start with `Test<WhatWeAreTesting>`; Functions receive `*testing.T` as a parameter to allow test runner access to state/logging.
* `struct` is the most similar type in Go to a class in OOP languages. `struct` defines different properties on a type.
* When initializing a struct with `structName{"x", "y"}` syntax, Go relies on the properties order if you don't define them. One way to avoid this is to specify the properties on initialization like so `structName{propertyA: "x", propertyB: "y"}`
* Go is a _pass by value_ language. Which means (even on receiver functions) you will not be using the same memory reference when you access the instance inside the function. You would be accessing (and changing) a copy of the memory isntance.
* Pointers: `&variable` returns a pointer. `*variable` returns the value from that pointer. In functions you can receive a pointer like so `func funcName(p *type)`
* You don't necessarly need to pass a pointer to a pointer receiver function. Go allows you to pass the pointer type's root value to the pointer receiver and converts it to a pointer for you.
* Since slices point to an array in memory, when you update slice values inside a different function that received the slice as value (copy of the original memory address), you will be updating the original array. This does not happen when you append to or remove from a slice, because since arrays are fixed sized, a new array will be created holding a different memory address.
* There are several other types that behave the same. _(Reference types vs Value types)_
* Everything in Go is passed by value. Even pointers!
* `map` is just a key/value dictionary.
* Go doesn't support function overloading. One can only have one function with the same name inside the same package.
* You can only have functions with the same name if they have different receivers.
* You don't specifically implement interfaces. Your "sub type" will only be converted to the type's interface if it implements all the methods inside the interface.
* There are concrete types and interface types. Concrete types can be created with the corresponding value type, interfaces don't.
* Go doesn't have generic types.
* You don't need to define that a type implements an interface. Go only converts that type to the interface type if it implements all the defined functions inside the interface. There's no direct linkage.
* You can use interfaces as struct fields or inside other interfaces.
* Go routines execute inside a go scheduler. The behaviour is different when you have more than one CPU core. Go, by default, tries to execute go routines in only only one core. The scheduler runs one routine until it finishes or makes a blocking call, and swaps between go routines to efficiently manage the code execution. If you have more than one core (N), Go will run N go routines for each core.
* Concurrency != Paralelism. You only have real paralelism when you have more than one core. You do have concurrency (multiple threads executing code) on a single core tho, but they don't execute at the same time. Even if the schedulers context switchs from fraction to fraction of a nanosecond.
* Everytime a program starts a main go routine is created and it's responsible for the program's lifetime.
* You can start a child go routine by typying `go` before a function call. e.g `go getSomething()`
* Channels are used to communicate between go routines.
* Channels can only send messages of the same data type, defined at creation.
* Creating a channel: `c := make(chan string)`
* You need to pass the channel to the go routine function's arguments, just like any other argument. e.g `func fetch(string url, c chan string)`
* Send a message to the channel `c <- "message"`.
* Get the value from a channel `<-c`
* Waiting to receive messages from a channel is a blocking thread.
* 