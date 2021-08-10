# keeper
[![Test](https://github.com/mocyuto/keeper/actions/workflows/ci.yml/badge.svg)](https://github.com/mocyuto/keeper/actions/workflows/ci.yml)

`keeper` is package for Go that provides a mechanism for waiting a result of execution function until context cancel.

## Install

```
go get github.com/gunosy/keeper
```

## Usage

`ExecWithContext` wraps function which is wanted to be canceled by context. 
This function is watching context whether context is canceled, 
and wait for executing function until context canceled. 

```go
import (
    "context"
    "time"

    "github.com/gunosy/keeper"
)

func findSomething(ctx context.Context, userID string) ([]int, error) {
    // some logic    	
}

ctx := context.WithTimeout(context.Background(), 100 * time.Millisecond)
result, err := keeper.ExecWithContext(ctx, func() (interface{}, error) {
    result, err := findSomething(ctx, userID) // exec heavy func
    return result, err
})
if result == nil {
    return nil, err
}
return result.([]int), err
```
