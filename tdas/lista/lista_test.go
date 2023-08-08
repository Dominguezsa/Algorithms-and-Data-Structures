package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const iteraciones_volumen int = 10000

// TESTS PARA LISTA ENLAZADA

// Test para crear lista
func TestCrearLista(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.Panics(t, func() { lista.BorrarPrimero() })
	require.Panics(t, func() { lista.VerPrimero() })
	require.Panics(t, func() { lista.VerUltimo() })
}

// Test para insertar primero una vez y luego borrar
func TestInsertarPrimeroYBorrarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(3)
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 3, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())

	lista.BorrarPrimero()
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.Panics(t, func() { lista.BorrarPrimero() })
	require.Panics(t, func() { lista.VerPrimero() })
	require.Panics(t, func() { lista.VerUltimo() })
}

// Test para insertar ultimo una vez y luego borrar
func TestInsertarUltimoYBorrarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarUltimo(3)
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 3, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())

	lista.BorrarPrimero()
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.Panics(t, func() { lista.BorrarPrimero() })
	require.Panics(t, func() { lista.VerPrimero() })
	require.Panics(t, func() { lista.VerUltimo() })
}

// Test para insertar 2 veces primero y luego borrar 1 vez
func TestInsertarPrimeroDosYBorrarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(3)
	lista.InsertarPrimero(4)
	require.False(t, lista.EstaVacia())
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, 4, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())

	lista.BorrarPrimero()
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 3, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())
}

// Test para insertar 2 veces ultimo y luego borrar 1 vez
func TestInsertarDosUltimoYBorrarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarUltimo(3)
	require.False(t, lista.EstaVacia())
	lista.InsertarUltimo(4)
	require.False(t, lista.EstaVacia())
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, 3, lista.VerPrimero())
	require.Equal(t, 4, lista.VerUltimo())

	lista.BorrarPrimero()
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 4, lista.VerPrimero())
	require.Equal(t, 4, lista.VerUltimo())
}

// Test para insertar varios y borrar 3 veces
func TestInsertarVariosYBorrarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)
	lista.InsertarPrimero(5)
	lista.InsertarPrimero(7)
	lista.InsertarPrimero(6)
	lista.InsertarUltimo(9)
	require.False(t, lista.EstaVacia())
	require.Equal(t, 6, lista.Largo())
	require.Equal(t, 6, lista.VerPrimero())
	require.Equal(t, 9, lista.VerUltimo())

	lista.BorrarPrimero()
	require.False(t, lista.EstaVacia())
	require.Equal(t, 5, lista.Largo())
	require.Equal(t, 7, lista.VerPrimero())
	require.Equal(t, 9, lista.VerUltimo())

	lista.BorrarPrimero()
	require.False(t, lista.EstaVacia())
	require.Equal(t, 4, lista.Largo())
	require.Equal(t, 5, lista.VerPrimero())
	require.Equal(t, 9, lista.VerUltimo())
}

// Test para insertar varias veces y luego vaciar la lista con borrar primero
func TestBorrarPrimeroVariasVeces(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < 10; i++ {
		lista.InsertarPrimero(i)
	}

	for i := 0; i < 10; i++ {
		require.Equal(t, 0, lista.VerUltimo())
		require.False(t, lista.EstaVacia())
		require.Equal(t, 10-i, lista.Largo())
		require.Equal(t, 9-i, lista.BorrarPrimero())
	}
	require.Panics(t, func() { lista.BorrarPrimero() })
	require.Panics(t, func() { lista.VerPrimero() })
	require.Panics(t, func() { lista.VerUltimo() })

	lista_2 := TDALista.CrearListaEnlazada[int]()

	for i := 0; i < 10; i++ {
		lista_2.InsertarUltimo(i)
	}

	for i := 0; i < 10; i++ {
		require.Equal(t, 9, lista_2.VerUltimo())
		require.False(t, lista_2.EstaVacia())
		require.Equal(t, 10-i, lista_2.Largo())
		require.Equal(t, i, lista_2.BorrarPrimero())
	}
	require.Panics(t, func() { lista_2.BorrarPrimero() })
	require.Panics(t, func() { lista_2.VerPrimero() })
	require.Panics(t, func() { lista_2.VerUltimo() })

}

// Test para para iterar con la primitiva sobre una lista vacia
func TestIterarConListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	suma := 0
	lista.Iterar(func(int) bool {
		suma += 1
		return true
	})
	require.Equal(t, 0, suma)

}

// Test para sumar valores de cada elemento de la lista con primitiva iterar
func TestIterarConListaConElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	suma := 0
	for i := 1; i <= 10; i++ {
		lista.InsertarUltimo(i)
	}

	lista.Iterar(func(n int) bool {
		suma += n
		return true
	})
	require.Equal(t, 55, suma)
}

// Test para iterar con primitiva con condicion
func TestIterarInternamenteConCondicion(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	suma := 0
	for i := 0; i < 10; i++ {
		lista.InsertarUltimo(i)
	}

	lista.Iterar(func(n int) bool {
		if n%2 != 0 {
			suma += n
		}
		return true
	})
	require.Equal(t, suma, 25)
}

// Test para iterar con primitiva con condicion de falso
func TestIterarInternamenteConCondicionFalsa(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	suma := 0
	for i := 0; i < 10; i++ {
		lista.InsertarUltimo(i)
	}

	lista.Iterar(func(n int) bool {
		if n%2 != 0 {
			suma += n
		} else if n > 5 {
			return false
		}
		return true
	})
	require.Equal(t, suma, 9)
}

// TESTS PARA ITERADORES EXTERNOS

// Test para crear un iterador
func TestCrearIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	require.Panics(t, func() { iter.VerActual() })
	require.Panics(t, func() { iter.Borrar() })
	require.Panics(t, func() { iter.Siguiente() })
	require.False(t, iter.HaySiguiente())
}

// Test para insertar con iterador en lista vacia
func TestInsertarEnListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	iter.Insertar(3)
	require.Equal(t, 3, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
}

// Test para insertar con iterador en lista vacia y vaciarla con otro
func TestInsertarYBorrarEnListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()

	for i := 0; i < 10; i++ {

		iter.Insertar(i * 2)
		require.Equal(t, 0, lista.VerUltimo())
		require.False(t, lista.EstaVacia())
		require.Equal(t, i+1, lista.Largo())

		require.True(t, iter.HaySiguiente())
		require.Equal(t, i*2, iter.VerActual())

	}
	require.Equal(t, 10, lista.Largo())
	require.Equal(t, 18, lista.VerPrimero())
	require.Equal(t, 18, iter.VerActual())
	require.True(t, iter.HaySiguiente())

	for i := 0; i < 10; i++ {
		require.Equal(t, 18-2*i, iter.VerActual())
		require.True(t, iter.HaySiguiente())
		iter.Siguiente()
	}

	require.Panics(t, func() { iter.VerActual() })
	require.Panics(t, func() { iter.Borrar() })
	require.Panics(t, func() { iter.Siguiente() })
	require.False(t, iter.HaySiguiente())

	iter2 := lista.Iterador()

	for i := 0; i < 10; i++ {
		require.Equal(t, 0, lista.VerUltimo())
		require.False(t, lista.EstaVacia())
		require.Equal(t, 10-i, lista.Largo())

		require.True(t, iter2.HaySiguiente())
		require.Equal(t, 18-2*i, iter2.Borrar())
	}

	require.Panics(t, func() { iter2.VerActual() })
	require.Panics(t, func() { iter2.Borrar() })
	require.Panics(t, func() { iter2.Siguiente() })
	require.False(t, iter.HaySiguiente())

	require.Panics(t, func() { lista.BorrarPrimero() })
	require.Panics(t, func() { lista.VerPrimero() })
	require.Panics(t, func() { lista.VerUltimo() })
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
}

// Test para borrar el primer elemento con iterador y que se refleje en la lista
func TestBorrarAlCrearIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < 10; i++ {
		lista.InsertarUltimo(i)
	}
	iterador := lista.Iterador()
	iterador.Borrar()
	require.Equal(t, 9, lista.Largo())
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 9, lista.VerUltimo())

	require.Equal(t, 1, iterador.VerActual())

}

// Test para borrar el ultimo elemento con iterador
func TestBorrarAlFinal(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < 10; i++ {
		lista.InsertarPrimero(i)
	}
	iterador := lista.Iterador()
	for iterador.HaySiguiente() {
		if iterador.VerActual() == 0 {
			iterador.Borrar()
			break
		}
		iterador.Siguiente()

	}
	require.Equal(t, 9, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
	require.Equal(t, 9, lista.Largo())

	require.Panics(t, func() { iterador.VerActual() })
	require.Panics(t, func() { iterador.Borrar() })
	require.Panics(t, func() { iterador.Siguiente() })
	require.False(t, iterador.HaySiguiente())

}

// Test para que insertar un elemento cuando el iterador termina de iterar
// efectivamente es equivalente a insertar al final en la lista.
func TestInsertarAlFinal(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < 10; i++ {
		lista.InsertarPrimero(i)
	}
	iterador := lista.Iterador()
	for iterador.HaySiguiente() {
		iterador.Siguiente()
	}
	iterador.Insertar(-1)

	require.Equal(t, -1, iterador.VerActual())
	require.True(t, iterador.HaySiguiente())

	require.Equal(t, 11, lista.Largo())
	require.False(t, lista.EstaVacia())
	require.Equal(t, -1, lista.VerUltimo())
	require.Equal(t, 9, lista.VerPrimero())

	iterador.Siguiente()
	require.False(t, iterador.HaySiguiente())
	require.Panics(t, func() { iterador.VerActual() })

	for i := 0; i <= 10; i++ {
		require.False(t, lista.EstaVacia())
		require.Equal(t, 9-i, lista.VerPrimero())
		require.Equal(t, -1, lista.VerUltimo())
		require.Equal(t, 11-i, lista.Largo())
		require.Equal(t, 9-i, lista.BorrarPrimero())
	}
	require.Panics(t, func() { lista.BorrarPrimero() })
	require.Panics(t, func() { lista.VerPrimero() })
	require.Panics(t, func() { lista.VerUltimo() })
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
}

// Test para insertar en el medio con un iterador
func TestInsertarEnElMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < 10; i++ {
		lista.InsertarUltimo(i)
	}
	iterador := lista.Iterador()
	for i := 0; i < 5; i++ {
		iterador.Siguiente()
	}
	require.Equal(t, 10, lista.Largo())
	require.Equal(t, 5, iterador.VerActual())
	require.True(t, iterador.HaySiguiente())
	require.Equal(t, 9, lista.VerUltimo())
	require.Equal(t, 0, lista.VerPrimero())

	iterador.Insertar(10)

	require.Equal(t, 10, iterador.VerActual())
	require.True(t, iterador.HaySiguiente())

	require.Equal(t, 11, lista.Largo())
	require.False(t, lista.EstaVacia())
	require.Equal(t, 9, lista.VerUltimo())
	require.Equal(t, 0, lista.VerPrimero())

	for i := 0; i < 6; i++ {
		iterador.Siguiente()
	}

	require.Equal(t, 11, lista.Largo())
	require.Panics(t, func() { iterador.VerActual() })
	require.Panics(t, func() { iterador.Borrar() })
	require.False(t, iterador.HaySiguiente())

	arr := [11]int{0, 1, 2, 3, 4, 10, 5, 6, 7, 8, 9}
	for i, elem := range arr {
		require.False(t, lista.EstaVacia())
		require.Equal(t, elem, lista.VerPrimero())
		require.Equal(t, 9, lista.VerUltimo())
		require.Equal(t, 11-i, lista.Largo())
		require.Equal(t, elem, lista.BorrarPrimero())
	}

	require.Panics(t, func() { lista.BorrarPrimero() })
	require.Panics(t, func() { lista.VerPrimero() })
	require.Panics(t, func() { lista.VerUltimo() })
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
}

// Test de Volumen
func TestVolumen(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < iteraciones_volumen; i++ {
		lista.InsertarUltimo(i)
		require.False(t, lista.EstaVacia())
		require.Equal(t, i+1, lista.Largo())
		require.Equal(t, 0, lista.VerPrimero())
		require.Equal(t, i, lista.VerUltimo())
	}
	for i := 0; i < iteraciones_volumen; i++ {
		require.False(t, lista.EstaVacia())
		require.Equal(t, iteraciones_volumen-i, lista.Largo())
		require.Equal(t, i, lista.VerPrimero())
		require.Equal(t, iteraciones_volumen-1, lista.VerUltimo())
		require.Equal(t, i, lista.BorrarPrimero())
	}
	require.Panics(t, func() { lista.BorrarPrimero() })
	require.Panics(t, func() { lista.VerPrimero() })
	require.Panics(t, func() { lista.VerUltimo() })
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
}

// Test con strings
func TestStrings(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	lista.InsertarPrimero("hola")
	lista.InsertarUltimo("mundo")
	require.Equal(t, "hola", lista.BorrarPrimero())
	require.Equal(t, "mundo", lista.BorrarPrimero())
}

// Test con floats
func TestFloats(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[float64]()
	lista.InsertarPrimero(1.05)
	lista.InsertarUltimo(2.023)
	require.Equal(t, 1.05, lista.BorrarPrimero())
	require.Equal(t, 2.023, lista.BorrarPrimero())
}

// Test de lista de listas
func TestListaDeListas(t *testing.T) {
	lista_de_listas := TDALista.CrearListaEnlazada[TDALista.Lista[int]]()
	lista_de_int := TDALista.CrearListaEnlazada[int]()

	lista_de_int.InsertarUltimo(3)
	require.False(t, lista_de_int.EstaVacia())
	require.Equal(t, lista_de_int.Largo(), 1)

	require.True(t, lista_de_listas.EstaVacia())
	require.Equal(t, lista_de_listas.Largo(), 0)

	lista_de_listas.InsertarUltimo(lista_de_int)

	require.False(t, lista_de_listas.EstaVacia())
	require.Equal(t, lista_de_listas.Largo(), 1)

	require.Equal(t, lista_de_listas.VerPrimero(), lista_de_int)
	require.Equal(t, lista_de_listas.VerUltimo(), lista_de_int)
	require.Equal(t, lista_de_listas.BorrarPrimero(), lista_de_int)

	require.Panics(t, func() { lista_de_listas.BorrarPrimero() })
	require.Panics(t, func() { lista_de_listas.VerPrimero() })
	require.Panics(t, func() { lista_de_listas.VerUltimo() })
	require.True(t, lista_de_listas.EstaVacia())
	require.Equal(t, 0, lista_de_listas.Largo())

	require.False(t, lista_de_int.EstaVacia())
	require.Equal(t, lista_de_int.Largo(), 1)
	require.Equal(t, lista_de_int.VerPrimero(), 3)
	require.Equal(t, lista_de_int.VerUltimo(), 3)
}
