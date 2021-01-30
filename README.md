# Go: The Complete Developer's Guide (Golang)

<https://www.udemy.com/course/go-the-complete-developers-guide/>

Collection of notes and exercises took during the course.

## Course notes

_(in progress)_

* Go is not a fully OOP enabled language and doesn't have a class type.
* `main` is a reserved package name for runable programs.
* variable declaration syntax `var x string = "2"` or infered by the compiler `x := "2"`
* function syntax `func x(a int, b int) string {}` _(from what I saw until now, the type allways goes after the declaration)_.
* foreach syntax `for i, item := range collection {}` which holds the index, which is a nice to have.
* The compiler will not build your program if you have declared variables that are not used in your code. Same goes for imports.
* Receiver functions are kind alike C# extension methods. They extend functionality of a type, allowing access to the type's instance. `func (r string) getString(x int) {}`. `r` is the instance.
* slice is a dynamic array (without a fixed size).
* slice range select `x[0:2]` 2 is not inclusive. Can be abbreviated to `x[:2]` (from 0 to 1 inclusive) or `x[1:]` (from 1 to last element).
* Go allows multiple return values in functions. `(string, int)` on function signature and `return value1, value1` on return statement. On variable assignment `val1, val2 := multipleReturnFunc()`.
* Type conversion `[]byte("Hi there")` converts a string to a byte slice (decimal).