package cola_prioridad_test

import (
	"math/rand"
	TDAHeap "tdas/cola_prioridad"
	"testing"

	"github.com/stretchr/testify/require"
)

func cmp(a, b int) int {
	return a - b
}

var elementos []int = []int{1, 4, 2, 5, 9, 6, 8, 7, 11}
var elementos_max_heap []int = []int{11, 9, 8, 7, 6, 5, 4, 2, 1}

func TestHeapVacio(t *testing.T) {
	heap := TDAHeap.CrearHeap(cmp)
	require.True(t, heap.EstaVacia())
	require.EqualValues(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.NotPanics(t, func() { heap.Encolar(3) })
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 1, heap.Cantidad())
	require.EqualValues(t, 3, heap.VerMax())
	require.EqualValues(t, 3, heap.Desencolar())
	require.True(t, heap.EstaVacia())
	require.EqualValues(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestHeapAEncolarUnElementoYBorrarlo(t *testing.T) {
	heap := TDAHeap.CrearHeap(cmp)
	heap.Encolar(3)
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 1, heap.Cantidad())
	require.EqualValues(t, 3, heap.VerMax())
	require.EqualValues(t, 3, heap.Desencolar())
	require.True(t, heap.EstaVacia())
	require.EqualValues(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestHeapEncolarYDesencolarVarios(t *testing.T) {

	heap := TDAHeap.CrearHeap(cmp)

	for i := 0; i < len(elementos); i++ {
		heap.Encolar(elementos[i])
	}

	for i := 0; i < 5; i++ {
		require.EqualValues(t, elementos_max_heap[i], heap.Desencolar())
		require.EqualValues(t, elementos_max_heap[i+1], heap.VerMax())
		require.EqualValues(t, len(elementos)-i-1, heap.Cantidad())
		require.False(t, heap.EstaVacia())
	}
}

func TestHeapEncolarYDesencolar(t *testing.T) {

	heap := TDAHeap.CrearHeap(cmp)

	for i := 0; i < len(elementos); i++ {
		heap.Encolar(elementos[i])
	}

	for i := 0; i < len(elementos)-1; i++ {
		require.EqualValues(t, elementos_max_heap[i], heap.Desencolar())
		require.EqualValues(t, elementos_max_heap[i+1], heap.VerMax())
		require.EqualValues(t, len(elementos)-i-1, heap.Cantidad())
		require.False(t, heap.EstaVacia())
	}
	require.EqualValues(t, elementos_max_heap[len(elementos)-1], heap.Desencolar())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.EqualValues(t, 0, heap.Cantidad())
	require.True(t, heap.EstaVacia())

}

func TestHeapVolumen(t *testing.T) {
	heap := TDAHeap.CrearHeap(cmp)
	for i := 0; i <= 10000; i++ {
		heap.Encolar(i)
	}
	for i := 10000; i > 0; i-- {
		require.Equal(t, i, heap.Desencolar())
	}
}

func TestHeapVolumen2(t *testing.T) {
	heap := TDAHeap.CrearHeap(cmp)
	for i := 0; i <= 10000; i++ {
		heap.Encolar(rand.Intn(100))
	}
	lista := []int{}
	for i := 10000; i > 0; i-- {
		lista = append(lista, heap.Desencolar())
	}
	for i := 0; i < len(lista)-2; i++ {
		require.Equal(t, true, lista[i] >= lista[i+1])
	}
}

func TestHeapVolumen_con_aleatorios(t *testing.T) {
	heap := TDAHeap.CrearHeap(cmp)
	for i := 0; i <= 10000; i++ {
		heap.Encolar(rand.Intn(100))
	}
	lista := []int{}
	for i := 10000; i > 0; i-- {
		lista = append(lista, heap.Desencolar())
	}
	for i := 0; i < len(lista)-2; i++ {
		require.True(t, lista[i] >= lista[i+1])
	}
}

func TestDeStrings(t *testing.T) {
	heap := TDAHeap.CrearHeap(func(a, b string) int {
		if a > b {
			return 1
		}
		if a < b {
			return -1
		}
		return 0
	})
	heap.Encolar("hola")
	heap.Encolar("chau")
	heap.Encolar("adios")
	heap.Encolar("hola")
	require.Equal(t, "hola", heap.Desencolar())
	require.Equal(t, "hola", heap.Desencolar())
	require.Equal(t, "chau", heap.Desencolar())
	require.Equal(t, "adios", heap.Desencolar())
	require.Panics(t, func() { heap.Desencolar() })
	require.Equal(t, true, heap.EstaVacia())
}

func TestHeapsort(t *testing.T) {
	lista := []int{1, 6, 3, 2, 87, 3, 5}
	TDAHeap.HeapSort(lista, cmp)
	require.Equal(t, []int{1, 2, 3, 3, 5, 6, 87}, lista)
}

func TestCrearHeappArrVacio(t *testing.T) {
	slice := make([]int, 0)
	heap := TDAHeap.CrearHeapArr(slice, cmp)
	require.True(t, heap.EstaVacia())
	require.EqualValues(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.NotPanics(t, func() { heap.Encolar(3) })
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 1, heap.Cantidad())
	require.EqualValues(t, 3, heap.VerMax())
	require.EqualValues(t, 3, heap.Desencolar())
	require.True(t, heap.EstaVacia())
	require.EqualValues(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestCrearHeappArr(t *testing.T) {
	elementos_largo := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}
	elementos_largo_copia := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}
	heap := TDAHeap.CrearHeapArr(elementos_largo, cmp)
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, len(elementos_largo), heap.Cantidad())
	require.EqualValues(t, 18, heap.VerMax())
	require.NotPanics(t, func() { heap.Encolar(50) })
	require.EqualValues(t, 50, heap.VerMax())
	require.EqualValues(t, len(elementos_largo)+1, heap.Cantidad())
	require.EqualValues(t, 50, heap.Desencolar())
	require.EqualValues(t, 18, heap.VerMax())
	require.EqualValues(t, len(elementos_largo), heap.Cantidad())
	for i := 0; i < len(elementos_largo)-1; i++ {
		require.EqualValues(t, len(elementos_largo)-i, heap.Desencolar())
		require.EqualValues(t, len(elementos_largo)-i-1, heap.VerMax())
		require.EqualValues(t, len(elementos_largo)-i-1, heap.Cantidad())
		require.False(t, heap.EstaVacia())
	}
	require.EqualValues(t, elementos_largo[0], heap.Desencolar())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.EqualValues(t, 0, heap.Cantidad())
	require.True(t, heap.EstaVacia())
	require.Equal(t, elementos_largo, elementos_largo_copia)
}
