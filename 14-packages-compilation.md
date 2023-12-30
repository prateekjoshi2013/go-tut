- **The compiled form of a Go program is a special kind of file called an executable
binary, containing machine code for a specific type of CPU and operating system**.

  - **To produce an executable binary, our program needs to contain a main package, with a special function named main that is the program’s entry point when it’s run**.

  - Though if a **function named init is present, it’s automatically executed even before main, perhaps to do some initialisation**, for example.

  - However, **it’s widely considered poor style to use init**, as we can initialise variables
  in a less magical way using plain old var statements.

  - This presumes that you’ve defined the initialiseThing function somewhere in this package, but that’s fine. The point is that you can make sure that function runs before any other in the package

  ```go
    package stuff
    var thing = initialiseThing()
  ```
  
- The **go build command compiles a Go program with a main package into an executable binary that we can run (or distribute to someone else so that they can run it)**.

  - Conveniently, **we can build an executable for any CPU architecture or operating system we want, not just the one we happen to be running go build on, using the GOOS and GOARCH environment variables**.
  
    ```sh
        GOOS=windows go build
        GOOS=linux GOARCH=arm go build
    ``` 

  - Get a list of all the GOOS and GOARCH combinations supported by your Go version by running the command: 

    ```sh
        go tool dist list
    ``` 

- While the program exits automatically when the end of main is reached, it can also exit at any other time by calling the os.Exit function (for example, to set an exit status code).
  
  - **zero exit status indicates that everything’s fine, whereas a non‐zero value would imply that some error happened**.

    ```go
    package main

    import "os"

    func main() {
        os.Exit(1)
    }
    ```
  
  - Other ways to exit the program, **which are not recommended, are calling log.Fatal or panic; in general, though, please don’t panic**.