# bathtext-go

Convert number to Thai currency full text with zero dependency. (go)

## Installation

```bash
go get github.com/jukbot/bathtext-go
```

## Usage

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/jukbot/bathtext-go"
)

func main() {
    // Convert amount to Thai currency text
    result, err := bathtext.Convert(100.50)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(result) // Output: 100.50 บาท
    
    // Alternative function
    result, err = bathtext.ConvertBaht(1000)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(result) // Output: 1000.00 บาท
}
```

## API

### Functions

- `Convert(amount float64) (string, error)` - Convert a number to Thai currency text
- `ConvertBaht(amount float64) (string, error)` - Alias for Convert function

### Constants

- `Version` - Library version

### Errors

- `ErrInvalidInput` - Returned when input is invalid (e.g., negative amounts)

## Development

### Building

```bash
go build ./...
```

### Testing

```bash
go test -v
```

### Running Example

```bash
go run example/main.go
```

## License

This project is open source and available under the MIT License.
