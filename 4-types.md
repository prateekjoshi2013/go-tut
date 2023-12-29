- Go has various data types for dealing with different kinds of information: 
  
  - **string** holds text, 
  
  - **rune** is an alias for int32 and is equivalent to int32 in all ways. It is used, by convention, to distinguish character values from integer values.
  
  - **int8**, **int16**, **int32**, **int64**, **int** holds integer numbers , 
  
    - **int** changes with 64 bit and 32 bit machine
  
  - **float32**, **float64** holds fractional numbers, and
  
  - **bool** holds Boolean (true or false) values.

- The var keyword introduces a variable, and specifies its type, so that we can use it
later on in the program.

- Variables in Go have a default value before we assign something to them explicitly:

- It’s whatever the zero value is for its type, such as "" for strings, 0 for numbers,
or false for bool.

- Declare and assign values to variables in a single statement, using the
short declaration form, for example x := y.

- When you want to group related values together into a single variable, you can
define a struct type to represent them.

- The different fields of the struct are just like any other kind of Go variable, and
they can be of any Go type, including other structs.

- Type definitions are introduced by the keyword type, and struct types with the
additional **keyword struct**, followed by a list of fields with their types.

- Identifiers (that is, the names of functions, variables, or fields) with an **initial capital letter are exported**: that is, they’re public and can be accessed from outside the
package where they’re defined.

- On the other hand, identifiers beginning with a **lowercase letter are unexported**:
they’re private to the package.

- Most Go packages are “about” one or more **core struct types**, in some sense:
they’re often the first things that you’ll define and test.

- A neat way to design a struct type guided by tests is to write a ***compile‐only test***,
that simply creates a literal value of the struct: this test can’t pass (or, indeed,
compile) until you’ve correctly defined the struct type.
