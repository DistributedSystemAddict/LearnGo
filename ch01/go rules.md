There are a few go rules:
- You import it, you use it
- You declare it, you use it
- Go automatically insert **semicolon** at the end of files so the opening brace must be fucking on the same line as the if, for, func, etc.
- Every single block **must** use curly brace. You removed the braces? Cute. Now put the braces back
- Go functions can return multiple values.
- You can NOT automatically convert between different data types, even if they are of the same kind. As an example, you cannot implicitly convert an integer to a floating point.

Example for go functions:
```go
func add(a, b int) (int, string) {
    return a+b, "calculation complete"
}

sum, msg := add(3,4)
```

Now sum == 7 and msg == "calculation complete".


```go
v, err := someFunction()
```