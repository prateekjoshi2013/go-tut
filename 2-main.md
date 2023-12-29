## Main

- In order for a bunch of Go source code to produce an executable, there has to be at
least some code in a special package named **main**.

- We can’t have more than one package in the same folder (apart from test packages)

- We’ll need to put the **main package** in a subfolder of your project
  
- Because the main package produces a command (that is, an executable), it’s conventional to name this folder **cmd** (short for “command”).

- Our program has to start somewhere, The way we tell go compiler is by defining a special function, also named **main**.

- The package main declaration is required if we want to produce an executable
binary. And we also need a function main, or the binary wouldn’t do anything.