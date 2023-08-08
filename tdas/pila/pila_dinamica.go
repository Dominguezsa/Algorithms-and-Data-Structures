package pila

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

const (
	// Cantidad inicial de elementos que puede contener la pila
	CapacidadMinima = 16
	//Cantidad que aumenta o disminuye el tamaño del slice
	Proporcion = 2
	//Denominador de la capacidad para que el slice se achique
	DenominadorRedimension = 4
)

/* Implementación de la interfaz pila. */

func (p *pilaDinamica[T]) Apilar(dato T) {
	if p.cantidad == cap(p.datos) {
		p.redimensionar(cap(p.datos) * Proporcion)
	}
	p.datos[p.cantidad] = dato
	p.cantidad++
}

func (p *pilaDinamica[T]) Desapilar() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	if p.cantidad == cap(p.datos)/DenominadorRedimension && cap(p.datos) > CapacidadMinima {
		p.redimensionar(cap(p.datos) / Proporcion)
	}
	p.cantidad--
	return p.datos[p.cantidad]
}

func (p *pilaDinamica[T]) VerTope() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	return p.datos[p.cantidad-1]
}

func (p *pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == 0
}

func (p *pilaDinamica[T]) redimensionar(cantidad int) {
	datos := make([]T, cantidad)
	copy(datos, p.datos)
	p.datos = datos
}

func CrearPilaDinamica[T any]() Pila[T] {
	return &pilaDinamica[T]{datos: make([]T, CapacidadMinima)}
}
