- A Go struct can not only store data in its fields, but also define behaviour with its
methods.

- A method is equivalent to a function, and works the same way, but it singles out
one of its parameters—the receiver—as special: in some sense, the method is about
this receiver, and it’s part of the definition of the receiver’s type.

    ```go
        type Catalog map[int]Book
        
        var catalog Catalog
        
        func (c Catalog) GetAllBooks() []Book {
            books := make([]Book, 0, len(c))
            for _, book := range c {
                books = append(books, book)
            }
            return books
        }
    ```

- We call a method using the dot notation: 

    ```go
            catalog.GetAllBooks()
    ```

- Because a method is part of a type definition, it follows that ***we can’t write methods
on types defined outside the current package***, and that includes built‐in types
such as int.

- But we can trivially create new types based on existing types 
  (for example, ***type MyInt int***): then we can add a method on our new type.

- However, a new type such as MyInt is distinct from its underlying type, so we need
to ***explicitly convert values to the new type using a type conversion***.