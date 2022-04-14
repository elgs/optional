package optional

import "log"

type Optional[T any] struct {
	Data  *T
	Error *error
}

func New[T any](data *T, err *error) *Optional[T] {
	return &Optional[T]{
		Data:  data,
		Error: err,
	}
}

func (this *Optional[T]) LogIfError() {
	log.Println(this.Error)
}

func (this *Optional[T]) PanicIfError() {
	log.Panicln(this.Error)
}

func (this *Optional[T]) FatalIfError() {
	log.Fatalln(this.Error)
}
