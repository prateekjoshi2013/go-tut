- A slice is a kind of Go type that represents an ordered collection of elements, all of
the same type: for example, ***[]string is a slice of strings***.

- A slice literal specifies the name of the slice type, followed by some elements in
curly braces: for example, ***[]string{"a", "b", "c"}***.

- In a slice literal of some custom type, such as a struct, we don’t need to repeat that
type name as part of every element literal: the compiler can infer it.

- We use the square‐bracket notation to refer to a particular numbered element of
the slice: for example, books[0] refers to the first element of the books slice.

- The built‐in function len tells us the number of elements in a slice.

- The built‐in function append appends one or more elements to a slice, and returns
the modified slice.

- The range operator lets us write for loops that execute once for each element of
the slice in turn.

- In a loop like for _, b := range books, each time round the loop b will be
each successive element of the slice books.

- The ***== operator isn’t defined on slices, so we can use the go-cmp package instead,
which can compare all sorts of values using cmp.Equal, and produce a diff
with cmp.Diff if they’re not equal.***

- ***The output from cmp.Diff shows what was expected with a leading - sign, and
what was actually received with a leading + sign***.

- When you need to retrieve some specific thing from a collection of things, you
need a way of uniquely identifying it by an ID, which can be whatever you want,
but is often a number.

- One way to get a specific element from a slice, if you don’t already know its index,
is to loop over the whole slice comparing each element with the one you want.