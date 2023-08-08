package funciones

import "tp1/votos"

const (
	longitud_dni = 8
)

func RadixSort(lista []int) []int {
	for i := 0; i < longitud_dni; i++ {
		lista = countingSort(lista, i)
	}
	return lista
}

func countingSort(lista []int, digito int) []int {
	frecuencias := make([]int, 10)
	for _, valor := range lista {
		frecuencias[(valor/(potencia(valor, digito)))%10]++
	}
	acum := make([]int, 10)
	for i := 1; i < 10; i++ {
		acum[i] = acum[i-1] + frecuencias[i-1]
	}
	final := make([]int, len(lista))
	for _, valor := range lista {
		final[acum[(valor/(potencia(valor, digito)))%10]] = valor
		acum[(valor/(potencia(valor, digito)))%10]++
	}
	return final
}

func potencia(valor, digito int) int {
	pot := 1
	for i := 0; i < digito; i++ {
		pot *= 10
	}
	return pot
}

func BusquedaBinaria(lista []votos.PersonaPadron, valor int) (bool, int) {
	return _busqueda_binaria(lista, valor, 0, len(lista)-1)
}

func _busqueda_binaria(lista []votos.PersonaPadron, valor, inicio, fin int) (bool, int) {
	if inicio > fin {
		return false, -1
	}
	medio := (inicio + fin) / 2
	if lista[medio].LeerDNI() == valor {
		return true, medio
	} else if lista[medio].LeerDNI() > valor {
		return _busqueda_binaria(lista, valor, inicio, medio-1)
	} else {
		return _busqueda_binaria(lista, valor, medio+1, fin)
	}
}
