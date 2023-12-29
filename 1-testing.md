## Testing 

- Each test in a Go project is a function

- It must be in a file whose name ends in _test.go (there can be multiple tests in
the same file)

- The name of a test function must begin with the word Test, or Go won’t recognise
it as a test. 

- It also has to accept a parameter whose type is *testing.T

- Test functions don’t return anything

- **t.Parallel()** statement is a standard prelude to tests: it tells Go to run
this test concurrently with other tests, which saves time

- We can call t.Errorf to cause the test to fail with a message, and if we don’t do
that, the test will pass by default.

- **t.Errorf** continues executing the test in case of an error 

- **t.Fatalf** throws error and exits test immediately

- Run all tests  using **go test ./...***