package optional

type Optional[T any] struct {
	Data  T
	Error error
}
