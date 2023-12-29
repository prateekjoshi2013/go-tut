
- The Go tools provide a helpful code coverage function to check this. For example, if you

  - ``` go test -cover ./... ```

- Go can generate a coverage profile for us:

  - ``` go test -coverprofile=coverage.out ```

- Use this profile to generate an HTML page showing the covered code highlighted
in green, uncovered in red, and ignored (things that aren’t statements) in grey:

  - ``` go tool cover -html=coverage.out ```

- Visual Studio Code’s Go extension lets you highlight covered code right there in the editor,
using the command Go: 

  - Toggle Test Coverage in Current Package