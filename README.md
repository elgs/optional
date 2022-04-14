# optional
If you think if err != nil is too much, this might be for you.

## Old way
```go
func New() (string, error) {
	return "Hello", nil
}

data, err := New()
if err == nil {
  log.Fatal(err)
}
log.Println(data)
```

## New way
```go
import "github.com/elgs/optional"

func NewOpt() *optional.Optional[string] {
	return optional.New("Hello", nil)
}

dataOpt := NewOpt()
dataOpt.FatalIfError()
log.Println(dataOpt.Data)
```