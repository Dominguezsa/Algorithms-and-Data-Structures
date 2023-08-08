package cola_test

import (
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

// verifica que lo que desencola sea correcto
func TestDesencolar(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	for i := 0; i < 10; i++ {
		cola.Encolar(i)
	}
	require.EqualValues(t, 0, cola.VerPrimero())
	require.EqualValues(t, 0, cola.Desencolar())
	require.EqualValues(t, 1, cola.VerPrimero())
	require.EqualValues(t, 1, cola.Desencolar())
	require.EqualValues(t, 2, cola.VerPrimero())
	require.EqualValues(t, 2, cola.Desencolar())
	require.EqualValues(t, 3, cola.VerPrimero())
	require.EqualValues(t, 3, cola.Desencolar())
	require.EqualValues(t, 4, cola.VerPrimero())
	require.EqualValues(t, 4, cola.Desencolar())
	require.EqualValues(t, 5, cola.VerPrimero())
	require.EqualValues(t, 5, cola.Desencolar())
	require.EqualValues(t, 6, cola.VerPrimero())
	require.EqualValues(t, 6, cola.Desencolar())
	require.EqualValues(t, 7, cola.VerPrimero())
	require.EqualValues(t, 7, cola.Desencolar())
	require.EqualValues(t, 8, cola.VerPrimero())
	require.EqualValues(t, 8, cola.Desencolar())
	require.EqualValues(t, 9, cola.VerPrimero())
	require.EqualValues(t, 9, cola.Desencolar())
}

// prueba de volumen
func TestVolumen(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	for i := 0; i < 10001; i++ {
		cola.Encolar(i)
		require.EqualValues(t, 0, cola.VerPrimero())
	}
	require.False(t, cola.EstaVacia())
	//vacia la cola y verifica que este vacia
	for j := 0; j < 10001; j++ {
		require.EqualValues(t, j, cola.VerPrimero())
		require.EqualValues(t, j, cola.Desencolar())
	}
	require.True(t, cola.EstaVacia())
}

// prueba de volumen2
func TestVolumen2(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	for i := 0; i < 10001; i++ {
		cola.Encolar(i)
		require.EqualValues(t, 0, cola.VerPrimero())
	}
	require.False(t, cola.EstaVacia())
	//vacia la cola y verifica que este vacia
	for j := 0; j < 10001; j++ {
		require.EqualValues(t, j, cola.VerPrimero())
		require.EqualValues(t, j, cola.Desencolar())
	}
	require.True(t, cola.EstaVacia())
}

// prueba de VerPrimero
func TestVerFrente(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	for i := 0; i < 2000; i++ {
		cola.Encolar(i)
		require.EqualValues(t, 0, cola.VerPrimero())
	}
	require.False(t, cola.EstaVacia())
}

// prueba de Panics
func TestPanics(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
}

// caso borde: encola y desencola constantemente
func TestBorde(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	for i := 0; i < 10100; i++ {
		cola.Encolar(i)
		require.EqualValues(t, i, cola.Desencolar())
		require.True(t, cola.EstaVacia())
	}
}

// test con strings
func TestStrings(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[string]()
	cola.Encolar("hola")
	cola.Encolar("mundo")
	require.EqualValues(t, "hola", cola.VerPrimero())
	require.EqualValues(t, "hola", cola.Desencolar())
	require.EqualValues(t, "mundo", cola.VerPrimero())
	require.EqualValues(t, "mundo", cola.Desencolar())
}

// test con floats
func TestStructs(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[float64]()
	cola.Encolar(1.1)
	cola.Encolar(2.2)
	require.EqualValues(t, 1.1, cola.VerPrimero())
	require.EqualValues(t, 1.1, cola.Desencolar())
	require.EqualValues(t, 2.2, cola.VerPrimero())
	require.EqualValues(t, 2.2, cola.Desencolar())
}
