package lista

// Definicion de structs
type nodoLista[T any] struct {
	dato T
	prox *nodoLista[T]
}

type listaEnlazada[T any] struct {
	prim   *nodoLista[T]
	ultimo *nodoLista[T]
	largo  int
}

type iteradorLista[T any] struct {
	lista    *listaEnlazada[T]
	actual   *nodoLista[T]
	anterior *nodoLista[T]
}

// Funcion para crear nodo
func crearNodoLista[T any](dato T) *nodoLista[T] {
	nodo := new(nodoLista[T])
	nodo.dato = dato
	return nodo
}

//Primitivas de ListaEnlazada

func (l listaEnlazada[T]) EstaVacia() bool {
	return l.prim == nil
}

func (l *listaEnlazada[T]) InsertarPrimero(dato T) {
	nuevo := crearNodoLista(dato)
	if l.EstaVacia() {
		l.ultimo = nuevo
	}
	nuevo.prox = l.prim
	l.prim = nuevo
	l.largo++
}

func (l *listaEnlazada[T]) InsertarUltimo(dato T) {
	nuevo := crearNodoLista(dato)
	if l.EstaVacia() {
		l.prim = nuevo
	} else {
		l.ultimo.prox = nuevo
	}
	l.ultimo = nuevo
	l.largo++
}

func (l *listaEnlazada[T]) BorrarPrimero() T {
	dato := l.VerPrimero()
	l.prim = l.prim.prox
	l.largo--
	if l.EstaVacia() {
		l.ultimo = nil
	}
	return dato
}

func (l listaEnlazada[T]) VerPrimero() T {
	l.panicSiListaEstaVacia()
	return l.prim.dato
}

func (l listaEnlazada[T]) VerUltimo() T {
	l.panicSiListaEstaVacia()
	return l.ultimo.dato
}

func (l listaEnlazada[T]) Largo() int {
	return l.largo
}

func (l listaEnlazada[T]) Iterar(visitar func(T) bool) {
	actual := l.prim
	for actual != nil {
		if !visitar(actual.dato) {
			break
		}
		actual = actual.prox
	}
}

func (l *listaEnlazada[T]) Iterador() IteradorLista[T] {
	iter := new(iteradorLista[T])
	iter.lista = l
	iter.actual = l.prim
	return iter
}

func (l listaEnlazada[T]) panicSiListaEstaVacia() {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
}

func CrearListaEnlazada[T any]() Lista[T] {
	return new(listaEnlazada[T])
}

//Primitivas del iterador

func (i iteradorLista[T]) VerActual() T {
	i.panicSiNoHaySiguiente()
	return i.actual.dato
}

func (i iteradorLista[T]) HaySiguiente() bool {
	return i.actual != nil
}

func (i *iteradorLista[T]) Siguiente() {
	i.panicSiNoHaySiguiente()
	i.anterior = i.actual
	i.actual = i.actual.prox

}

func (i *iteradorLista[T]) Insertar(dato T) {
	nuevo := crearNodoLista(dato)
	if i.actual == nil {
		i.lista.ultimo = nuevo
	}
	nuevo.prox = i.actual
	i.actual = nuevo
	if i.anterior != nil {
		i.anterior.prox = nuevo
	} else {
		i.lista.prim = i.actual
	}

	i.lista.largo++
}

func (i *iteradorLista[T]) Borrar() T {
	i.panicSiNoHaySiguiente()
	dato := i.actual.dato

	if i.actual == i.lista.prim {
		i.lista.prim = i.lista.prim.prox
	} else {
		i.anterior.prox = i.actual.prox
	}

	i.actual = i.actual.prox
	if i.actual == nil {
		i.lista.ultimo = i.anterior
	}

	i.lista.largo--
	return dato
}

func (i iteradorLista[T]) panicSiNoHaySiguiente() {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
}
