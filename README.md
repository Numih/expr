# expr

`expr` is a small Go package for parsing environment variables from expressions with support for default values. It allows developers to easily retrieve environment variables or provide fallback values in case the variables are not defined.

## Features

- Parse environment variables in the format `${VAR:default}`.
- Provides fallback to default values if the environment variable is not defined.
- Supports simple string returns when no environment variables are used.

## Installation

To install this package, use:

```bash
go get github.com/numih/expr
```

## Usage

### Importing the Package

```go
import "github.com/numih/expr/env"
```

### Parse Environment Variables

The `Parse` function evaluates expressions in the format `${VAR:default}`, where `VAR` is the name of the environment variable and `default` is the fallback value if the variable is not defined.

```go
package main

import (
	"fmt"
	"github.com/numih/expr/env"
)

func main() {
	// Set environment variable VAR="value" for this example.
	value := env.Parse("${VAR:default}")
	fmt.Println(value)  // Output: value if VAR is set, "default" otherwise
}
```

### Examples

1. If `VAR` is defined:
    ```go
    env.Parse("${VAR:default}") // Output: value of VAR
    ```

2. If `VAR` is not defined:
    ```go
    env.Parse("${VAR:default}") // Output: "default"
    ```

3. No environment variable placeholder:
    ```go
    env.Parse("foo") // Output: "foo"
    ```

4. If no default value is provided:
    ```go
    env.Parse("${VAR}") // Output: value of VAR if set, otherwise ""
    ```
