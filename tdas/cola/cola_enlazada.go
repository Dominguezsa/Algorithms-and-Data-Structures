package cola

type nodo[T any] struct {
	dato T
	prox *nodo[T]
}

type colaEnlazada[T any] struct {
	primero *nodo[T]
	ultimo  *nodo[T]
}

func crearNodo[T any](dato T) *nodo[T] {
	return &nodo[T]{dato: dato}
}

func (c *colaEnlazada[T]) EstaVacia() bool {
	return c.primero == nil
}

func (c *colaEnlazada[T]) VerPrimero() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	return c.primero.dato
}

func (c *colaEnlazada[T]) Encolar(dato T) {
	nuevo := crearNodo(dato)
	if c.EstaVacia() {
		c.primero = nuevo
	} else {
		c.ultimo.prox = nuevo
	}
	c.ultimo = nuevo
}

func (c *colaEnlazada[T]) Desencolar() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	dato := c.primero.dato
	c.primero = c.primero.prox
	if c.primero == nil {
		c.ultimo = nil
	}
	return dato
}

func CrearColaEnlazada[T any]() Cola[T] {
	return &colaEnlazada[T]{}
}
