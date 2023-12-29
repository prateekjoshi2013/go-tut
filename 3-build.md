- The go run command is useful for quickly checking whether the program works, 

- If we actually want to distribute our software to customers, a more convenient way to do that is to build an executable file.

- To create the executable, run the command:
  
  - ```go build -o add ./cmd/calculator```

- Windows works slightly differently:
  
  - ```go build -o add.exe ./cmd/calculator``` 

- To run it, use the program name itself as the command:

  - ```./add ```

- On Windows, run just:
  
  - ``` add ```