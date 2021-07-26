# keeper

wait exec result at context cancel


## Example

```go
import (
    "context"
    "time"

    "github.com/mocyuto/keeper"
)

ctx := context.WithTimeout(context.Background(), 100 * time.Millisecond)
result, err := keeper.ExecWithContext(ctx, func() (interface{}, error) {
    return findSomething(ctx, userID) // exec heavy func
})
if result == nil {
    return nil, err
}
return result.([]int), err
```
