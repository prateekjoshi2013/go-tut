• To declare a function, we begin with the **func keyword, followed by the name of the function, a list of parameters, and a list of results**.

```go
    func myFunction(x float64, y float64, z float64) (float64, error){
        // body goe here
        return x, y
    }
```

    - The parameter list gives the name and type of each parameter that the function receives.

    - Similarly, the result list gives the types (and, optionally, the names) of the result values that the function returns, if any.

    - This is followed by the function body, a list of statements enclosed in curly braces.

    - Inside a function body, we can exit the function early using a return statement;

    - If the function returns results, we have to give a list of values with return, one for each declared result

- In order to execute the function’s code, we need to call it. A function call consists of the name of the function, followed by a list of arguments in parentheses (if the function takes no parameters, the parentheses are empty).

    - We can use the results returned by a function call, if any, as part of an assignment.

    ```go
            lat, long, err := location()
    ```
- Functions themselves are values in Go, so we can assign them to variables, or pass
them as arguments to other functions, or even return them as results.
  
  - When using a function as a value, we don’t necessarily have to give it a name; instead, we can write a function literal.
  
    ```go
        function: func(a, b float64) float64 {
            return math.Pow(a, b)
        }
    ```
  - A function literal can “see” all the variables in the enclosing function where it’s
  defined, thus making it what’s called a **closure** over them.

  - **We can create a closure that executes later, but in this case we need to watch out for the enclosed variables being changed in the meantime— for example, in a loop**.

    ```go
        nums := []int{3, 1, 2}
        sort.Slice(nums, func(i, j int) bool {
            return nums[i] < nums[j]  // nums from closure
        })
    ```

- There are certain kinds of **resources that can’t be automatically cleaned up by the
Go garbage collector, such as file handles**, which need to be explicitly closed by calling Close.

    ```go
        func MyFunction(){
            f, err := os.Open("testdata/somefile.txt")
            if err != nil {
                return err
            }
            defer f.Close()
            ... // do stuff with f
        }
    ```

  - The **defer keyword lets us defer execution of some block of code until the current
  function exits**: for example, cleaning up the resource.

  - If we stack defers by using several defer statements, **the deferred code will be
  run in reverse order on exit: the last thing deferred is the first thing to run**.

- You can name a function’s result parameters, and the main value of this is to tell
  users and readers what those results signify: for example, latitude and longitude,
  rather than just float64, float64.

    ```go
        func location() (latitude float64, longitude float64, error)
    ```

  - **Naked returns (bare return statements with no arguments)** are legal syntactically,
  and you’ll even see them used in some old blog posts and tutorials. **But there’s no good reason to use a naked return**, and it’s always clearer to return some explicit value, such as nil.
  
  - **When a function has named results, you can modify them in a deferred closure**:
  for example, to return some error encountered when trying to clean up resources.
  
  - Just because a function has named results doesn’t mean it must use a naked return: it can, but it shouldn’t, for the reasons we’ve discussed.
  
- A **variadic function** can take any number of arguments (including zero), using the
special ... form, and its arguments will appear inside the function as successive
elements of a slice.

    ```go
        func mixedParams(a int, b string, values ...int) {
             for _, v := range values {
                fmt.Println(v)
            }
        }

        mixedParams(1, "hello", 2, 3, 4)
        
        values := []int{1, 2, 3}
        
        mixedParams(0, "prefix", values...) // using values slice

    ```