package diccionario

import (
	"fmt"
)

type estado int

const (
	//Estados posibles de los elementos del hash
	ESTADOVACIO estado = iota
	ESTADOOCUPADO
	ESTADOBORRADO
)

const (
	//Razon maxima y minima que sirven de tope para redimensionar el tamaño del hash hacia arriba o hacia abajo respectivamente
	FACTORMAXIMO float64 = 0.6
	FACTORMINIMO float64 = 0.2
)

const (
	//Cantidad que aumenta o disminuye el tamaño del hash
	PROPORCION int = 2
)

// Cantidad minima de elementos que puede contener el hash
const TAMANOMINIMO int = 20

type hashCerrado[K comparable, V any] struct {
	tabla    []elemHashCerrado[K, V]
	cant     int
	borrados int
}

type elemHashCerrado[K comparable, V any] struct {
	clave      K
	valor      V
	estadoElem estado
}

func crearTabla[K comparable, V any](tamano int) []elemHashCerrado[K, V] {
	return make([]elemHashCerrado[K, V], tamano)
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	hash := new(hashCerrado[K, V])
	hash.tabla = crearTabla[K, V](TAMANOMINIMO)
	return hash
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

// PRIMITIVAS DEL ELEMENTO DEL HASH
func (e *elemHashCerrado[K, V]) LeerClaveValor() (K, V) {
	return e.clave, e.valor
}

func (e *elemHashCerrado[K, V]) LeerEstado() estado {
	return e.estadoElem
}

func (e *elemHashCerrado[K, V]) CambiarEstado(estadoElem estado) {
	e.estadoElem = estadoElem
}

// Funcion de hashing extraida de: https://golangprojectstructure.com/hash-functions-go-code/
func fvnHash(data []byte) uint64 {
	var hash uint64 = 0xcbf29ce484222325
	var uint64Prime uint64 = 0x00000100000001b3

	for _, b := range data {
		hash ^= uint64(b)
		hash *= uint64Prime
	}

	return hash
}

// PRIMITIVAS DE HASH
func (h *hashCerrado[K, V]) Guardar(clave K, dato V) {
	pos := h.obtenerPos(clave)
	if h.tabla[pos].estadoElem == ESTADOVACIO {
		h.cant++
		h.tabla[pos].estadoElem = ESTADOOCUPADO
		h.tabla[pos].clave = clave
	}
	h.tabla[pos].valor = dato
	if (float64(h.cant)/float64(len(h.tabla))) > FACTORMINIMO && float64((h.cant+h.borrados))/float64(len(h.tabla)) > FACTORMAXIMO {
		h.redimension(len(h.tabla) * PROPORCION)
	}
}

func (h *hashCerrado[K, V]) Pertenece(clave K) bool {
	pos := h.obtenerPos(clave)
	return h.tabla[pos].LeerEstado() == ESTADOOCUPADO
}

func (h *hashCerrado[K, V]) Obtener(clave K) V {
	pos := h.obtenerPos(clave)
	h.validarPos(pos)
	_, valor := h.tabla[pos].LeerClaveValor()
	return valor
}

func (h *hashCerrado[K, V]) Borrar(clave K) V {
	pos := h.obtenerPos(clave)
	h.validarPos(pos)
	_, valor := h.tabla[pos].LeerClaveValor()
	h.tabla[pos].CambiarEstado(ESTADOBORRADO)
	h.cant--
	h.borrados++
	if float64(h.cant)/float64(len(h.tabla)) < FACTORMINIMO && len(h.tabla) > TAMANOMINIMO {
		h.redimension(len(h.tabla) / PROPORCION)
	}
	return valor
}

func (h *hashCerrado[K, V]) Cantidad() int {
	return h.cant
}

func (h *hashCerrado[K, V]) Iterar(f func(clave K, dato V) bool) {
	for i := 0; i < len(h.tabla); i++ {
		if h.tabla[i].LeerEstado() == ESTADOOCUPADO {
			clave, dato := h.tabla[i].LeerClaveValor()
			if !f(clave, dato) {
				return
			}
		}
	}
}

func (h *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	iter := new(iterDiccionario[K, V])
	iter.hash = h
	for pos := 0; pos < len(h.tabla); pos++ {
		if h.tabla[pos].estadoElem == ESTADOOCUPADO {
			break
		}
		iter.pos++
	}
	return iter
}

func (h *hashCerrado[K, V]) obtenerPos(clave K) int {
	bytes := convertirABytes(clave)
	pos_inicial := int(fvnHash(bytes) % uint64(len(h.tabla)))
	condicion := false
	tope := len(h.tabla)
	var pos_resultado int
	for pos := pos_inicial; pos < tope; pos++ {

		if h.tabla[pos].estadoElem == ESTADOVACIO || (h.tabla[pos].clave == clave && h.tabla[pos].estadoElem != ESTADOBORRADO) {
			pos_resultado = pos
			break
		}

		if pos == tope-1 && !condicion {
			tope = pos_inicial
			pos = 0
			condicion = true
		}
	}
	return pos_resultado
}

func (h *hashCerrado[K, V]) validarPos(pos int) {
	if h.tabla[pos].LeerEstado() == ESTADOVACIO {
		panic("La clave no pertenece al diccionario")
	}
}

func (h *hashCerrado[K, V]) redimension(tamanoNuevo int) {
	tabla_anterior := h.tabla
	nuevaTabla := crearTabla[K, V](tamanoNuevo)
	h.tabla = nuevaTabla
	h.cant = 0
	h.borrados = 0
	for i := 0; i < len(tabla_anterior); i++ {
		if tabla_anterior[i].estadoElem == ESTADOOCUPADO {
			clave := tabla_anterior[i].clave
			valor := tabla_anterior[i].valor
			pos := h.obtenerPos(clave)
			h.tabla[pos].clave = clave
			h.tabla[pos].valor = valor
			h.tabla[pos].estadoElem = ESTADOOCUPADO
			h.cant++
		}
	}
}

// PRIMITIVAS DEL ITERADOR

type iterDiccionario[K comparable, V any] struct {
	hash *hashCerrado[K, V]
	pos  int
}

func (i *iterDiccionario[K, V]) HaySiguiente() bool {
	return i.pos < len(i.hash.tabla)
}

func (i *iterDiccionario[K, V]) Siguiente() {
	i.validarHaySiguiente()
	i.pos++
	for i.pos < len(i.hash.tabla) {
		if i.hash.tabla[i.pos].LeerEstado() == ESTADOOCUPADO {
			return
		}
		i.pos++
	}
}

func (i *iterDiccionario[K, V]) VerActual() (K, V) {
	i.validarHaySiguiente()
	return i.hash.tabla[i.pos].clave, i.hash.tabla[i.pos].valor
}

func (i *iterDiccionario[K, V]) validarHaySiguiente() {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
}
