- When you define a new type, it doesn’t automatically “inherit” any methods that
the underlying type has.

- But you can define a struct type instead, containing a field of the type that you
want to wrap with new methods.

- When you pass a variable to a function, the function actually receives a copy of
that variable’s value.

- So if the function modifies the value, that change won’t be reflected in the original
variable.

- When you want the function to modify the original variable, you can create a
pointer to the variable, and pass that instead.

- A pointer gives permission to modify the thing it points to.
  
  - The ***& operator creates a pointer, so &x gives a pointer to x***.
  
  - A function that takes a pointer must declare the appropriate parameter type:
    *int for “pointer to int”, for example.
  
  - The only thing we’re allowed to do with a pointer value is dereference it using the ***\* operator to get the value it points to, so \*p gives the value of whatever variable p points to***.
  
  - Just as ordinary types are distinct and non‐interchangeable, so are pointers to those types: for example, ***\*int and \*float64 are different types***.
  
  - The default (zero) value of any pointer type is nil, meaning “doesn’t point to anything”.
  
  - Trying to dereference a pointer that happens to be nil will cause the program to panic and terminate with an error message.

  - Just as functions can take pointer parameters, so methods can take pointer receivers, and indeed they must do so if they want to modify the receiver.