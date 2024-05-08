package ejercicios

type Componente interface {
	String() string
	Resolver() int
}

type Operador struct {
	operador string
	izq      Componente
	der      Componente
}

type Operando struct {
	valor int
}