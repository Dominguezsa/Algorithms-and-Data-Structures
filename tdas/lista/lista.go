package lista

type Lista[T any] interface {

	//EstaVacia() devuelve true si la lista esta vacia, false en caso contrario
	EstaVacia() bool

	//InsertarPrimero(dato T) inserta el dato al principio de la lista
	InsertarPrimero(T)

	//InsertarUltimo(dato T) inserta el dato al final de la lista
	InsertarUltimo(T)

	//BorrarPrimero() borra el primer elemento de la lista y devuelve su dato.
	//Si la lista esta vacia, entra en panico con un mensaje "La lista esta vacia"
	BorrarPrimero() T

	//VerPrimero() devuelve el dato del primer elemento de la lista. Si la lista esta vacia,
	//entra en panico con un mensaje "La lista esta vacia"
	VerPrimero() T

	//VerUltimo() devuelve el dato del ultimo elemento de la lista. Si la lista esta vacia,
	//entra en panico con un mensaje "La lista esta vacia"
	VerUltimo() T

	//Largo() devuelve la cantidad de elementos de la lista
	Largo() int

	//Iterar(visitar func(T) bool) itera sobre la lista, aplicando la funcion visitar a cada elemento
	Iterar(visitar func(T) bool)

	//Iterador devuelve un iterador para la lista
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {

	//VerActual() devuelve el dato del elemento actual. Si el iterador termino de iterar,
	//entra en panico con un mensaje "El iterador termino de iterar"
	VerActual() T

	//HaySiguiente() devuelve true si hay un elemento siguiente, false en caso contrario
	HaySiguiente() bool

	//Siguiente() avanza el iterador al siguiente elemento. Si el iterador termino de iterar,
	//entra en panico con un mensaje "El iterador termino de iterar"
	Siguiente()

	//Insertar(dato T) inserta el dato en la posicion actual del iterador
	Insertar(T)

	//Borrar() borra el elemento actual del iterador y devuelve su dato. Si el iterador termino de iterar,
	//entra en panico con un mensaje "El iterador termino de iterar"
	Borrar() T
}
