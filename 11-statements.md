- Go programs consist essentially of a list of statements of various kinds, which are
instructions to the computer to do something.

  - An assignment statement associates a value with a variable.

  - A declaration statement introduces a variable that we will use later (perhaps by assigning to it, for example).

  - There’s a special kind of statement that combines declaration and assignment, called a short variable declaration: for example, **eggs := 42**.

  - We can make multiple assignments in a single statement using what’s called a tuple assignment, like this: **a, b, c := 1, 2, 3**.

- Go assigns values to variables in left‐to‐right order:
  
  - the first variable will receive the first value from the map index expression, and in this case, that will be the price of eggs, and not the bool value we wanted 
  
    - **price, ok := menu["eggs"]**
  
  - You might think that maybe we can leave price out and just put ok on the left‐hand
    side of the assignment:
    
    - **ok := menu["eggs"]**

  - When you’re not interested in one of the values in a tuple assignment, you can ignore it using the blank identifier, _, like this: **_, ok := menu["eggs"]**.


- Increment and decrement statements adjust the value of some numeric variable up or down by 1: for example, **x++ or y--**.

- An if statement controls whether a certain set of statements will be executed or not, depending on the value of some condition.

- In Go, we tend to structure our programs to left‐align the happy path: that is, if blocks tend to deal with errors or exceptional circumstances, leaving the straightline control flow as the expected path if everything goes according to plan.

- In principle, we can also execute some block of statements when the if condition is not true, using the else keyword, but this isn’t much used in practice; we prefer to have the if block return instead.

- A conditional expression, such as the condition for an if statement, is an expression
that evaluates to a boolean value (true or false).

- Conditional expressions can be comparisons, using == or !=, and other operators such as < , >, <=, and >=.

- We can also combine expressions using the logical operators && (and), || (inclusive or), and ! (not).

- The simplest kind of conditional expression is just a bool variable on its own, or even the literal values true or false.

- There’s a special form called a **compound if that combines an assignment and a condition controlling the execution of a block**, but it’s never necessary to use this, and it can lead to confusing code.

    ```go
        if _, ok := menu["eggs"]; ok {
            fmt.Println("Eggs are on the menu!")
        }
    ```