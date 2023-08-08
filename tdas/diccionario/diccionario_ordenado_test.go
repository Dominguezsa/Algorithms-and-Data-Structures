package diccionario_test

import (
	"github.com/stretchr/testify/require"
	"math/rand"
	TDAAbb "tdas/diccionario"
	"testing"
)

var TAM_VOLUMEN int = 10000

func comp_int(a int, b int) int {
	return a - b
}

func TestAbbVacio(t *testing.T) {
	t.Log("Comprueba que abb vacio no tiene claves")
	abb := TDAAbb.CrearABB[int, int](comp_int)
	require.EqualValues(t, 0, abb.Cantidad())
	require.False(t, abb.Pertenece(1))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener(1) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(1) })
}

func TestDiccionarioOrdenadoClaveDefault(t *testing.T) {
	t.Log("Prueba sobre un abb que si justo buscamos la clave que es el default del tipo de dato, " +
		"sigue sin existir")
	abb := TDAAbb.CrearABB[int, int](comp_int)
	require.EqualValues(t, 0, abb.Cantidad())
	require.False(t, abb.Pertenece(0))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener(0) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(0) })
}

func TestUnElementoAbb(t *testing.T) {
	t.Log("Comprueba que abb con un elemento tiene esa Clave, unicamente")
	abb := TDAAbb.CrearABB[int, int](comp_int)
	abb.Guardar(1, 10)
	require.EqualValues(t, 1, abb.Cantidad())
	require.True(t, abb.Pertenece(1))
	require.False(t, abb.Pertenece(2))
	require.EqualValues(t, 10, abb.Obtener(1))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener(2) })
}

func TestGuardarVarios(t *testing.T) {
	t.Log("Comprueba que abb que guarda varios elementos tiene esos elementos")
	abb := TDAAbb.CrearABB[int, int](comp_int)
	claves := []int{1, 432, 25, 5, 26, 2, 532, 17, 21, 12, 6412, 641, 16}
	valores := []int{152, 35, 523, 532, 251, 23, 5312, 554, 221, 142, 161412, 6411, 156}
	for i := range claves {
		abb.Guardar(claves[i], valores[i])
		require.EqualValues(t, abb.Cantidad(), i+1)
		require.EqualValues(t, abb.Obtener(claves[i]), valores[i])
	}
}

func TestGuardarVariosYSobreescribir(t *testing.T) {
	t.Log("Comprueba que abb que guarda varios elementos y sobreescribe algunos tiene esos elementos con los nuevos datos")
	abb := TDAAbb.CrearABB[int, int](comp_int)
	claves := []int{1, 432, 25, 5}
	valores := []int{152, 35, 523, 4}
	for i := range claves {
		abb.Guardar(claves[i], valores[i])
		require.EqualValues(t, abb.Cantidad(), i+1)
		require.EqualValues(t, abb.Obtener(claves[i]), valores[i])
	}
	abb.Guardar(1, 10)
	require.EqualValues(t, abb.Cantidad(), 4)
	require.EqualValues(t, abb.Obtener(1), 10)
	abb.Guardar(5, 100)
	require.EqualValues(t, abb.Cantidad(), 4)
	require.EqualValues(t, abb.Obtener(5), 100)
}

func TestGuardarVariosYSobreescribirTodos(t *testing.T) {
	t.Log("Comprueba que abb que guarda varios elementos y sobreescribe todos tiene esos elementos con los nuevos datos")
	abb := TDAAbb.CrearABB[int, int](comp_int)
	claves := []int{1, 432, 25, 5, 2, 3, 6, 7, 11, 71}
	valores := []int{152, 35, 523, 4, 15615, 156687, 156, 165, 12315, 0}
	for i := range claves {
		abb.Guardar(claves[i], valores[i])
		require.EqualValues(t, abb.Cantidad(), i+1)
		require.EqualValues(t, abb.Obtener(claves[i]), valores[i])
		require.True(t, abb.Pertenece(claves[i]))
	}
	for i := range claves {
		abb.Guardar(claves[i], i)
		require.EqualValues(t, len(claves), 10)
		require.EqualValues(t, abb.Obtener(claves[i]), i)
		require.True(t, abb.Pertenece(claves[i]))
	}
}

func TestBorrarUno(t *testing.T) {
	t.Log("Comprueba que se borra correctamente en un abb con un elemento")
	abb := TDAAbb.CrearABB[int, int](comp_int)
	abb.Guardar(1, 0)
	dato := abb.Borrar(1)
	require.EqualValues(t, dato, 0)
	require.EqualValues(t, abb.Cantidad(), 0)
	require.False(t, abb.Pertenece(1))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener(1) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(1) })
}
func TestBorrarVarios(t *testing.T) {
	t.Log("Comprueba que se borra correctamente en un abb con varios elementos")
	abb := TDAAbb.CrearABB[int, int](comp_int)
	claves := []int{1, 432, 25, 5}
	valores := []int{152, 35, 523, 4}
	for i := range claves {
		abb.Guardar(claves[i], valores[i])
	}
	require.EqualValues(t, abb.Cantidad(), 4)
	for i := 1; i < len(claves); i++ {
		dato := abb.Borrar(claves[i])
		require.EqualValues(t, dato, valores[i])
		require.EqualValues(t, abb.Cantidad(), len(claves)-i)
		require.False(t, abb.Pertenece(claves[i]))
		require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener(claves[i]) })
		require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(claves[i]) })
	}
	require.True(t, abb.Pertenece(1))
	require.EqualValues(t, abb.Obtener(1), 152)
	require.EqualValues(t, abb.Borrar(1), 152)
	require.EqualValues(t, abb.Cantidad(), 0)
	require.False(t, abb.Pertenece(1))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener(1) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(1) })

}

func TestReinsertarBorrado(t *testing.T) {
	t.Log("Comprueba que se borra correctamente en un abb con varios elementos, reinsertando correctamente el ya borrado.")
	abb := TDAAbb.CrearABB[int, int](comp_int)
	claves := []int{1, 432, 25, 5}
	valores := []int{152, 35, 523, 4}
	for i := range claves {
		abb.Guardar(claves[i], valores[i])
	}
	require.EqualValues(t, abb.Borrar(1), 152)
	require.EqualValues(t, abb.Cantidad(), 3)
	require.False(t, abb.Pertenece(1))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener(1) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(1) })

	abb.Guardar(1, 80)
	require.EqualValues(t, abb.Cantidad(), 4)
	require.True(t, abb.Pertenece(1))
	require.EqualValues(t, abb.Obtener(1), 80)
	require.EqualValues(t, abb.Borrar(1), 80)
	require.EqualValues(t, abb.Cantidad(), 3)
	require.False(t, abb.Pertenece(1))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener(1) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(1) })
}

func TestBorrarTodos(t *testing.T) {
	t.Log("Comprueba que en un abb con varios elementos, se pueden borrar todos correctamente.")
	abb := TDAAbb.CrearABB[int, int](comp_int)
	claves := []int{1, 432, 25, 5, 2, 3, 6, 7, 11, 71}
	valores := []int{152, 35, 523, 4, 15615, 156687, 156, 165, 12315, 0}
	for i := range claves {
		abb.Guardar(claves[i], valores[i])
	}
	for i := range claves {
		require.EqualValues(t, abb.Borrar(claves[i]), valores[i])
		require.EqualValues(t, abb.Cantidad(), len(claves)-i-1)
		require.False(t, abb.Pertenece(claves[i]))
		require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener(claves[i]) })
		require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(claves[i]) })
	}

	abb.Guardar(11, 125)
	require.EqualValues(t, abb.Cantidad(), 1)
	require.True(t, abb.Pertenece(11))
	require.EqualValues(t, abb.Obtener(11), 125)
	require.EqualValues(t, abb.Borrar(11), 125)
	require.EqualValues(t, abb.Cantidad(), 0)
	require.False(t, abb.Pertenece(11))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener(11) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(11) })
}
func TestIteradorInternoAbbVacio(t *testing.T) {
	t.Log("Comprueba que iterar internamente con un abb vacio no hace nada")
	abb := TDAAbb.CrearABB[int, int](comp_int)
	a := 0
	var b *int = &a
	abb.Iterar(func(_ int, dato int) bool {
		panic("Se detiene la ejecucion del iterador, no deberia hacer nada.")
	})
	require.EqualValues(t, 0, *b)

}

func TestIteradorInternoLeerVarios(t *testing.T) {
	t.Log("Comprueba que iterar internamente con un abb con varios elementos itera inorder")
	abb := TDAAbb.CrearABB[int, int](comp_int)
	claves := []int{1, 432, 25, 5}
	valores := []int{152, 35, 523, 4}
	for i := range claves {
		abb.Guardar(claves[i], valores[i])
	}
	claves_ordenadas := [4]int{}
	lis := []int{}
	a := 0
	var b *int = &a
	abb.Iterar(func(clave int, dato int) bool {
		lis = append(lis, dato+clave)
		claves_ordenadas[*b] = clave
		*b++
		return true
	})
	require.Equal(t, []int{153, 9, 548, 467}, lis)
	require.Equal(t, [4]int{1, 5, 25, 432}, claves_ordenadas)
	require.Equal(t, 4, *b)
}

func TestIteradorInternoLeerVariosConCorte(t *testing.T) {
	t.Log("Comprueba que iterar internamente con un abb con varios elementos itera inorder, cumpliendo el corte implementado")
	abb := TDAAbb.CrearABB[int, int](comp_int)
	claves := []int{1, 432, 25, 5, 2, 3, 6, 7, 11, 71}
	valores := []int{152, 35, 523, 4, 15615, 156687, 156, 165, 12315, 0}
	for i := range claves {
		abb.Guardar(claves[i], valores[i])
	}
	a := [2]int{}
	lista := []int{}
	abb.Iterar(func(clave int, dato int) bool {
		if clave > 10 {
			return false
		}
		lista = append(lista, clave)
		a[0]++
		a[1] = a[1] + dato
		return true
	})
	require.Equal(t, []int{1, 2, 3, 5, 6, 7}, lista)
	require.Equal(t, 6, a[0])
	require.Equal(t, 152+15615+156687+4+156+165, a[1])
}

func TestITeradorExternoAbbVacio(t *testing.T) {
	t.Log("Comprueba que el iterador externo de un abb vacio no tiene siguiente.")
	abb := TDAAbb.CrearABB[int, int](comp_int)
	iter := abb.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })

	abb2 := TDAAbb.CrearABB[int, int](comp_int)
	abb2.Guardar(1, 1)
	abb2.Borrar(1)
	iter2 := abb2.Iterador()
	require.False(t, iter2.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter2.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter2.VerActual() })
}
func TestIteradorExternoLeerVarios(t *testing.T) {
	t.Log("Comprueba que el iterador externo de un abb con varios elementos itera inorder.")
	abb := TDAAbb.CrearABB[int, int](comp_int)
	claves := []int{6515, 8849153, 518, 1358, 351, 1861, 5, 313}
	valores := []int{0, 43, 2153, 2, 1, 571, 412, 3}
	claves_ordenadas := []int{5, 313, 351, 518, 1358, 1861, 6515, 8849153}
	valores_ordenados := []int{412, 3, 1, 2153, 2, 571, 0, 43}
	for i := range claves {
		abb.Guardar(claves[i], valores[i])
	}
	i := 0
	for iter := abb.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		c, v := iter.VerActual()
		require.EqualValues(t, c, claves_ordenadas[i])
		require.EqualValues(t, v, valores_ordenados[i])
		i++
	}
}

func TestITeradorExternoRangoAbbVacio(t *testing.T) {
	t.Log("Comprueba que el iterador Externo por rango de un abb vacio no tiene siguiente.")
	abb := TDAAbb.CrearABB[int, int](comp_int)

	iter := abb.IteradorRango(nil, nil)
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })

	abb2 := TDAAbb.CrearABB[int, int](comp_int)
	abb2.Guardar(1, 1)
	abb2.Borrar(1)
	iter2 := abb2.IteradorRango(nil, nil)
	require.False(t, iter2.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter2.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter2.VerActual() })

	a := 2
	b := 5
	desde := &a
	hasta := &b

	abb3 := TDAAbb.CrearABB[int, int](comp_int)

	iter3 := abb3.IteradorRango(desde, nil)
	require.False(t, iter3.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter3.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter3.VerActual() })

	abb4 := TDAAbb.CrearABB[int, int](comp_int)
	iter4 := abb4.IteradorRango(nil, hasta)
	require.False(t, iter4.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter4.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter4.VerActual() })
}

func TestIteradorExternoRangoConGuardarVarios(t *testing.T) {
	t.Log("Comprueba que el Iterador externo por rango en un abb con varios elementos itera inorder en los rangos indicados.")
	abb := TDAAbb.CrearABB[int, int](comp_int)
	claves := []int{6515, 8849153, 518, 1358, 351, 1861, 5, 313}
	valores := []int{0, 43, 2153, 2, 1, 571, 412, 3}
	claves_ordenadas_rango := []int{518, 1358, 1861}
	valores_ordenados_rango := []int{2153, 2, 571}
	for i := range claves {
		abb.Guardar(claves[i], valores[i])
	}
	a := 500
	b := 2000
	desde := &a
	hasta := &b
	i := 0
	for iter := abb.IteradorRango(desde, hasta); iter.HaySiguiente(); iter.Siguiente() {
		c, v := iter.VerActual()
		require.EqualValues(t, c, claves_ordenadas_rango[i])
		require.EqualValues(t, v, valores_ordenados_rango[i])
		i++
	}

}

func TestITeradorInternoRangoAbbVacio(t *testing.T) {
	t.Log("Comprueba que el iterador interno por rango de un abb vacio no tiene siguiente.")
	abb := TDAAbb.CrearABB[int, int](comp_int)
	c := 0
	num := &c
	f := func(_ int, _ int) bool {
		*num++
		return true
	}
	abb.IterarRango(nil, nil, f)
	require.EqualValues(t, *num, 0)

	abb2 := TDAAbb.CrearABB[int, int](comp_int)
	abb2.Guardar(1, 1)
	abb2.Borrar(1)
	abb2.IterarRango(nil, nil, f)
	require.EqualValues(t, *num, 0)

	a := 2
	b := 5
	desde := &a
	hasta := &b

	abb3 := TDAAbb.CrearABB[int, int](comp_int)
	abb3.IterarRango(desde, nil, f)
	require.EqualValues(t, *num, 0)

	abb4 := TDAAbb.CrearABB[int, int](comp_int)
	abb4.IterarRango(nil, hasta, f)
	require.EqualValues(t, *num, 0)
}

func TestIteradorInternoRangoConGuardarVarios(t *testing.T) {
	t.Log("Comprueba que el Iterador interno por rango en un abb con varios elementos itera inorder en los rangos indicados.")
	abb := TDAAbb.CrearABB[int, int](comp_int)
	claves := []int{6515, 8849153, 518, 1358, 351, 1861, 5, 313}
	valores := []int{0, 43, 2153, 2, 1, 571, 412, 3}
	claves_ordenadas_rango := []int{518, 1358, 1861}
	for i := range claves {
		abb.Guardar(claves[i], valores[i])
	}
	a := 500
	b := 2000
	desde := &a
	hasta := &b
	lista_claves := []int{}
	c := 0
	datos := &c
	f := func(clave int, dato int) bool {
		lista_claves = append(lista_claves, clave)
		*datos = *datos + dato
		return true
	}
	abb.IterarRango(desde, hasta, f)
	require.EqualValues(t, lista_claves, claves_ordenadas_rango)
	require.EqualValues(t, *datos, 2153+2+571)

}

func TestIteradorInternoRangoLeerVarios(t *testing.T) {
	t.Log("Comprueba que el Iterador interno por rango en un abb con varios elementos itera inorder en los rangos indicados y cumple la condicon de corte.")
	abb := TDAAbb.CrearABB[int, int](comp_int)
	claves := []int{6515, 8849153, 518, 1358, 351, 1861, 5, 313}
	valores := []int{0, 43, 2153, 2, 1, 571, 412, 3}
	claves_ordenadas_rango := []int{518, 1358, 1861}
	for i := range claves {
		abb.Guardar(claves[i], valores[i])
	}
	a := 500
	desde := &a

	lista_claves2 := []int{}
	d := 0
	datos2 := &d
	f2 := func(clave int, dato int) bool {
		if clave > 2000 {
			return false
		}
		lista_claves2 = append(lista_claves2, clave)
		*datos2 = *datos2 + dato
		return true
	}
	abb.IterarRango(desde, nil, f2)
	require.EqualValues(t, lista_claves2, claves_ordenadas_rango)
	require.EqualValues(t, *datos2, 2153+2+571)

}

func TestVolumen(t *testing.T) {
	t.Log("Test de volumen de guardados y borrados.")
	abb := TDAAbb.CrearABB[int, int](comp_int)
	slice := make([]int, TAM_VOLUMEN)
	for i := range slice {
		slice[i] = i + 1
	}
	rand.Shuffle(len(slice), func(i int, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
	for i := range slice {
		abb.Guardar(i, i*10)
	}
	for i := range slice {
		require.EqualValues(t, abb.Obtener(i), i*10)
		require.True(t, abb.Pertenece(i))

	}

	for i := range slice {
		require.EqualValues(t, abb.Borrar(i), i*10)
		require.EqualValues(t, abb.Cantidad(), len(slice)-i-1)
		require.False(t, abb.Pertenece(i))
		require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener(i) })
		require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(i) })
	}
}

func comp_string(x string, y string) int {
	return len(x) - len(y)
}

func TestStrings(t *testing.T) {
	t.Log("Hago operaciones basicas de guardado y borrado con un abb de strings.")
	abb := TDAAbb.CrearABB[string, string](comp_string)
	abb.Guardar("hola", "mundo")
	abb.Guardar("chau", "tierra")
	abb.Guardar("buenos", "dias")
	require.EqualValues(t, "tierra", abb.Obtener("chao"))
	require.EqualValues(t, "dias", abb.Obtener("buenxs"))
	require.EqualValues(t, "tierra", abb.Borrar("1234"))
	require.False(t, abb.Pertenece("chau"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener("abcd") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar("efgh") })
}

type persona struct {
	nombre string
	edad   int
}

func comp_struct(x persona, y persona) int {
	return x.edad - y.edad
}
func TestStructs(t *testing.T) {
	t.Log("Compruebo varios operaciones de un abb de structs con una funcion de comparacion especifica.")
	abb := TDAAbb.CrearABB[persona, int](comp_struct)
	edades := []int{1, 13, 15, 14, 8, 7, 80}
	nombres := []string{"Lucas", "Juan", "Pedro", "Analia", "Carmen", "Lautaro", "Luz"}
	edades_ordenado := []int{1, 7, 8, 13, 14, 15, 80}
	nombres_ordenado := []string{"Lucas", "Lautaro", "Carmen", "Juan", "Analia", "Pedro", "Luz"}
	personas := make([]persona, 7)
	personas_ordenado := make([]persona, 7)
	for i := range personas {
		personas[i].nombre = nombres[i]
		personas[i].edad = edades[i]
		personas_ordenado[i].nombre = nombres_ordenado[i]
		personas_ordenado[i].edad = edades_ordenado[i]
		abb.Guardar(personas[i], personas[i].edad+3)
		require.True(t, abb.Pertenece(personas[i]))
		require.EqualValues(t, abb.Obtener(personas[i]), personas[i].edad+3)
	}
	lista := []persona{}
	a := 0
	suma := &a
	abb.Iterar(func(clave persona, dato int) bool {
		lista = append(lista, clave)

		*suma = *suma + dato
		return true
	})
	require.EqualValues(t, lista, personas_ordenado)
	require.EqualValues(t, 1+7+8+13+14+15+80+3*7, a)

	for i := range personas {
		require.EqualValues(t, abb.Borrar(personas_ordenado[i]), personas_ordenado[i].edad+3)
		require.EqualValues(t, abb.Cantidad(), len(personas_ordenado)-1-i)
		require.False(t, abb.Pertenece(personas_ordenado[i]))
		require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener(personas_ordenado[i]) })
		require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(personas_ordenado[i]) })
	}
}
