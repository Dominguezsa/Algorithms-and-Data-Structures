package cola_prioridad

const (
	CAPACIDAD_INICIAL = 10
	PROPORCION        = 2
	PROPORCION_MINIMO = 4
)

type heap[K comparable] struct {
	arr      []K
	cmp      func(K, K) int
	cantidad int
}

func CrearHeap[K comparable](funcion_cmp func(K, K) int) ColaPrioridad[K] {
	heap := new(heap[K])
	heap.arr = make([]K, CAPACIDAD_INICIAL)
	heap.cmp = funcion_cmp
	heap.cantidad = 0
	return heap
}

func CrearHeapArr[K comparable](arreglo []K, funcion_cmp func(K, K) int) ColaPrioridad[K] {
	heap := new(heap[K])
	var arreglo_copia []K
	tamanio := max(len(arreglo), CAPACIDAD_INICIAL)
	arreglo_copia = make([]K, tamanio+1)
	copy(arreglo_copia, arreglo)
	heap.arr = arreglo_copia
	heap.cmp = funcion_cmp
	heap.cantidad = len(arreglo)

	heapify(heap.arr, heap.cantidad, heap.cmp)
	return heap
}

func (heap *heap[K]) Encolar(elem K) {
	if heap.cantidad*PROPORCION == cap(heap.arr) {
		heap.redimensionar(cap(heap.arr) * PROPORCION)
	}
	heap.arr[heap.cantidad] = elem
	heap.upheap(heap.cantidad)
	heap.cantidad++
}

func (heap *heap[K]) redimensionar(nueva_capacidad int) {
	nuevo_arr := make([]K, nueva_capacidad)
	for i := 0; i < heap.cantidad; i++ {
		nuevo_arr[i] = heap.arr[i]
	}
	heap.arr = nuevo_arr
}

func (heap *heap[K]) upheap(pos int) {
	if pos == 0 {
		return
	}
	padre := (pos - 1) / 2
	if heap.cmp(heap.arr[pos], heap.arr[padre]) > 0 {
		heap.arr[pos], heap.arr[padre] = heap.arr[padre], heap.arr[pos]
		heap.upheap(padre)
	}
}

func (heap *heap[K]) Desencolar() K {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
	if heap.cantidad <= cap(heap.arr)/PROPORCION_MINIMO && cap(heap.arr) > CAPACIDAD_INICIAL {
		heap.redimensionar(cap(heap.arr) / PROPORCION)
	}
	elem := heap.arr[0]
	heap.cantidad--
	heap.arr[0] = heap.arr[heap.cantidad]
	downheap(0, heap.cantidad, heap.arr, heap.cmp)
	return elem
}

func (heap *heap[K]) EstaVacia() bool {
	return heap.cantidad == 0
}

func (heap *heap[K]) VerMax() K {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
	return heap.arr[0]
}

func (heap *heap[K]) Cantidad() int {
	return heap.cantidad
}

func HeapSort[K comparable](elementos []K, funcion_cmp func(K, K) int) {
	longitud := len(elementos)
	heapify(elementos, longitud, funcion_cmp)
	for i := 0; i < len(elementos); i++ {
		elementos[0], elementos[longitud-1] = elementos[longitud-1], elementos[0]
		longitud--
		downheap(0, longitud, elementos[:longitud], funcion_cmp)
	}
}

func downheap[K comparable](pos int, cant_elem int, arreglo []K, funcion_cmp func(K, K) int) {
	hijoIzq := pos*2 + 1
	hijoDer := pos*2 + 2
	if hijoIzq >= cant_elem {
		return
	}
	hijoMayor := hijoIzq
	if hijoDer <= cant_elem-1 && funcion_cmp(arreglo[hijoDer], arreglo[hijoIzq]) > 0 {
		hijoMayor = hijoDer
	}
	if funcion_cmp(arreglo[hijoMayor], arreglo[pos]) > 0 {
		arreglo[pos], arreglo[hijoMayor] = arreglo[hijoMayor], arreglo[pos]
		downheap(hijoMayor, cant_elem, arreglo, funcion_cmp)
	}
}
func heapify[K comparable](elementos []K, cant_elem int, funcion_cmp func(K, K) int) {
	for i := len(elementos) - 2; i >= 0; i-- {
		downheap(i, cant_elem, elementos, funcion_cmp)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

