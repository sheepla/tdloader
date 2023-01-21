# tdloader

A go module for easily loading test data placed in the Git repository

## About

This module makes it easy to loading the test data placed in the Go repository.
You can load test data simply by specifying a relative path from the Git repository root e.g. `testdata := tdloader.GetFile("_testdata/data1.json")`

The following functions can be used to load test data.

- `GetPath` - Resolve full path of test data
- `GetFile` - Open test data file and get `*os.File`. Since the os.File type implements the Read method, it can also be treated as the `io.Reder` type.
- `GetBytes` - Get the contents of the test data file as a bytes.
- `GetString` - Get the contents of the test data file as a string

Function names prefixed with `Must` will panic when an error occurs.

## Usage

```go
// mypkg.go
package mypkg

func DoSomethingWithFile(f *os.File) error {
    // -- snip --
}
```

```go
// mypkg_test.go
package mypkg

import (
    "testing"
    "github.com/sheepla/tdloader"
)

func TestDoSomething(*testing.T) {
    testdata := tdloader.GetFile("_testdata/data1.json")
    if err := DoSomethingWithFile(testdata); err != nil {
        t.Fatal(err)
    }

    // -- snip --
}
```

## Notice

This module runs a `git` command to find the project root directory. When using it, it is necessary that Git is installed and that the repository is managed by Git.

## License

MIT

## Author

[sheepla](https://github.com/sheepla)

