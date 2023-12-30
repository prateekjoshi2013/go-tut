- A switch statement is a concise way of selecting one of several possible cases, depending on some condition, without having to write lots of consecutive **if ... else ... statements**.

  - Each case in a switch statement can have a condition, in which case Go evaluates each case in order and **executes the first one for which the condition is true**, if any.

  - If there is a case called **default**, it is executed if and only if no other case has been matched.

  - Instead of conditional cases, we can provide an expression after the switch keyword; in this case, Go will evaluate the expression and execute the first case that specifies a matching result.

  - To exit a switch statement early, before reaching the end of the current block, use the **break statement**.
  
- To repeat a set of statements a number of times, perhaps forever, or at least until some condition is false, use a for statement.
  
  - A for loop with no condition simply repeats forever: ***for { ... }***.

  - To loop once for each element of a slice or a map, use the range operator: 
    
    - **for range employees { ... }**.

  - To receive the index number and the value for each element, use 
    
    - **for i, e :=range employees { ... }**.
  - On the other hand, if we just want the element value but not the index, we can write 
    
    - **for _, e := range employees { ... }**.
    
  - And if we only want the index, we receive only that: **for i := range employees { ... }**.

  - When we want the loop to repeat as long as some condition holds, we can write, for example, **for x < 10 { ... }**.

  - And to repeat a specified number of times, we can manage the loop variable entirely in the for statement, like this: **for x := 0; x < 10; x++ { ... }**.

  - To skip the current loop index or element and go on to the next, use the **continue** statement.

  - To jump out of the loop altogether, use the **break** statement.
    
  - If our program is so complex that it needs nested loops, we can use **break or continue
  with a label to jump directly to the labelled statement**.