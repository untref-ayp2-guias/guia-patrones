package ejercicios

type Iterator[T int] interface {
	HasNext() bool
	Next() (T, error)
}

// array struct of type T, genric
type Array[T int] struct {
	data []T
}

func NewArray[T int](data []T) *Array[T] {
	return &Array[T]{data: data}
}