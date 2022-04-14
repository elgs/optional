package optional

import "log"

type Optional[T any] struct {
	Data  *T
	Error error
}

func New[T any](data *T, err error) *Optional[T] {
	return &Optional[T]{
		Data:  data,
		Error: err,
	}
}

func (this *Optional[T]) LogIfError() {
	if this.Error != nil {
		log.Println(this.Error)
	}
}

func (this *Optional[T]) PanicIfError() {
	if this.Error != nil {
		log.Panicln(this.Error)
	}
}

func (this *Optional[T]) FatalIfError() {
	if this.Error != nil {
		log.Fatalln(this.Error)
	}
}
