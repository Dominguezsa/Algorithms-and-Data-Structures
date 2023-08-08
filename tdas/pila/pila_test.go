package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	// mas pruebas para este caso...
}

// verificar que lo que sale sea correcto
func TestDesapilar(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i < 10; i++ {
		pila.Apilar(i)
	}
	require.EqualValues(t, 9, pila.VerTope())
	require.EqualValues(t, 9, pila.Desapilar())
	require.EqualValues(t, 8, pila.VerTope())
	require.EqualValues(t, 8, pila.Desapilar())
	require.EqualValues(t, 7, pila.VerTope())
	require.EqualValues(t, 7, pila.Desapilar())
	require.EqualValues(t, 6, pila.VerTope())
	require.EqualValues(t, 6, pila.Desapilar())
	require.EqualValues(t, 5, pila.VerTope())
	require.EqualValues(t, 5, pila.Desapilar())
	require.EqualValues(t, 4, pila.VerTope())
	require.EqualValues(t, 4, pila.Desapilar())
	require.EqualValues(t, 3, pila.VerTope())
	require.EqualValues(t, 3, pila.Desapilar())
	require.EqualValues(t, 2, pila.VerTope())
	require.EqualValues(t, 2, pila.Desapilar())
	require.EqualValues(t, 1, pila.VerTope())
	require.EqualValues(t, 1, pila.Desapilar())
	require.EqualValues(t, 0, pila.VerTope())
	require.EqualValues(t, 0, pila.Desapilar())
}

// Prueba de volumen: agregar 1001 elementos a la pila y verificar que no este vacia;
// despues la vacia y verifica que este vacia
func TestVolumen(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i < 10001; i++ {
		pila.Apilar(i)
		require.EqualValues(t, i, pila.VerTope())
	}
	require.False(t, pila.EstaVacia())
	//vacia la pila y verifica que este vacia
	for j := 10000; j >= 0; j-- {
		require.EqualValues(t, j, pila.VerTope())
		require.EqualValues(t, j, pila.Desapilar())
	}
	require.True(t, pila.EstaVacia())
}

// verificar ver_tope
func TestVerTope(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i < 100; i++ {
		pila.Apilar(i)
		require.EqualValues(t, i, pila.VerTope())
	}
	require.EqualValues(t, 99, pila.VerTope())
}

// prueba de panics: desapilar y ver_tope de una pila vacia
func TestDesapilarPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
}

// verifica que funciona con strings

func TestPiladeStrings(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[string]()
	pila.Apilar("hola")
	pila.Apilar("mundo")
	require.EqualValues(t, "mundo", pila.Desapilar())
	require.EqualValues(t, "hola", pila.Desapilar())
}

//Prueba de caso borde: primero se apila y desapila constantemente. Luego se apilan muchos elementos,
// desapilando a medida que se van apilando (validando que el desapilado sea correcto),
// y luego se desapilan todos los restantes validando que sean correctos. Al final la pila debe quedar vacía.

func TestPilaCasoBorde(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i < 100; i++ {
		pila.Apilar(i)
		pila.Desapilar()
	}
	require.True(t, pila.EstaVacia())
	for i := 0; i < 1000; i++ {
		pila.Apilar(i)
		if i%2 == 0 {
			pila.Desapilar()
		}
	}
	for i := 0; i < 500; i++ {
		pila.Desapilar()
	}
	require.True(t, pila.EstaVacia())
}

func TestPilaVariosElementos(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i < 100; i++ {
		pila.Apilar(i)
		require.EqualValues(t, i, pila.VerTope())
	}
	require.EqualValues(t, 99, pila.VerTope())
	for i := 99; i >= 0; i-- {
		require.EqualValues(t, i, pila.Desapilar())
	}
	require.True(t, pila.EstaVacia())
}
