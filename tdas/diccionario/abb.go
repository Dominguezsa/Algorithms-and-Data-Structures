package diccionario

import (
	TDACola "tdas/cola"
)

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	dict := new(abb[K, V])
	dict.cmp = funcion_cmp
	return dict
}

func crearNodoAbb[K comparable, V any](clave K, dato V) *nodoAbb[K, V] {
	nodo := new(nodoAbb[K, V])
	nodo.clave = clave
	nodo.dato = dato
	return nodo
}

type abb[K comparable, V any] struct {
	cmp      func(K, K) int
	raiz     *nodoAbb[K, V]
	cantidad int
}

type nodoAbb[K comparable, V any] struct {
	clave    K
	dato     V
	hijo_izq *nodoAbb[K, V]
	hijo_der *nodoAbb[K, V]
}

func (ab *abb[K, V]) Guardar(clave K, dato V) {
	if !ab.tieneRaiz() {
		ab.raiz = crearNodoAbb(clave, dato)
		ab.cantidad++
		return
	}
	padre := ab.raiz.buscarPadre(clave, ab.cmp)
	nuevoElemento := padre.devolverNodo(clave, ab.cmp)
	if nuevoElemento == nil {
		nuevo_nodo := crearNodoAbb(clave, dato)
		if es_mayor(clave, padre.clave, ab.cmp) {
			padre.hijo_der = nuevo_nodo
		} else {
			padre.hijo_izq = nuevo_nodo
		}
		ab.cantidad++
	} else {
		nuevoElemento.clave = clave
		nuevoElemento.dato = dato
	}
}

func (ab *abb[K, V]) Pertenece(clave K) bool {
	if !ab.tieneRaiz() {
		return false
	}
	if son_iguales(clave, ab.raiz.clave, ab.cmp) {
		return true
	}
	nodo_buscado := ab.raiz.devolverNodo(clave, ab.cmp)
	return nodo_buscado != nil
}

func (ab *abb[K, V]) Cantidad() int {
	return ab.cantidad
}

func (ab *abb[K, V]) Obtener(clave K) V {
	if !ab.tieneRaiz() {
		panic("La clave no pertenece al diccionario")
	}
	nodo_buscado := ab.raiz.devolverNodo(clave, ab.cmp)
	if nodo_buscado == nil {
		panic("La clave no pertenece al diccionario")
	}
	return nodo_buscado.dato
}

func (ab *abb[K, V]) Borrar(clave K) V {
	if !ab.Pertenece(clave) {
		panic("La clave no pertenece al diccionario")
	}
	if son_iguales(clave, ab.raiz.clave, ab.cmp) {
		return ab.borrarRaiz()
	}
	ab.cantidad--
	padre := ab.raiz.buscarPadre(clave, ab.cmp)
	elementoAEliminar := padre.devolverNodo(clave, ab.cmp)
	if elementoAEliminar.no_tiene_hijos() {
		return ab.borrarSinHijos(padre, elementoAEliminar, clave)
	} else if elementoAEliminar.solo_hijo_der() {
		return ab.borrarHijoDer(padre, elementoAEliminar, clave)
	} else if elementoAEliminar.solo_hijo_izq() {
		return ab.borrarHijoIzq(padre, elementoAEliminar, clave)
	}
	return ab.borrarCon2Hijos(padre, elementoAEliminar, clave)
}

func (ab *abb[K, V]) Iterar(f func(clave K, dato V) bool) {
	ab.IterarRango(nil, nil, f)
}


type iterDiccionarioOrdenado[K comparable, V any] struct {
	cola_nodos TDACola.Cola[*nodoAbb[K, V]]
	actual     *nodoAbb[K, V]
}

func (ab *abb[K, V]) Iterador() IterDiccionario[K, V] {
	iter := ab.IteradorRango(nil, nil)
	return iter
}

func (i *iterDiccionarioOrdenado[K, V]) HaySiguiente() bool {
	return !i.cola_nodos.EstaVacia()
}

func (i *iterDiccionarioOrdenado[K, V]) Siguiente() {
	i.validarHaySiguiente()
	i.cola_nodos.Desencolar()
	if !i.cola_nodos.EstaVacia() {
		i.actual = i.cola_nodos.VerPrimero()
	} else {
		i.actual = nil
	}
}

func (i *iterDiccionarioOrdenado[K, V]) VerActual() (K, V) {
	i.validarHaySiguiente()
	return i.actual.clave, i.actual.dato
}

func (i *iterDiccionarioOrdenado[K, V]) validarHaySiguiente() {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
}

func (ab *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	if ab.tieneRaiz() {
		iterarInorder(ab.raiz, desde, hasta, visitar, ab.cmp, true)
	}
}

func (ab *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iter := new(iterDiccionarioOrdenado[K, V])
	iter.cola_nodos = TDACola.CrearColaEnlazada[*nodoAbb[K, V]]()
	iteradorInorder(ab.raiz, iter.cola_nodos, desde, hasta, ab.cmp)
	if !iter.cola_nodos.EstaVacia() {
		iter.actual = iter.cola_nodos.VerPrimero()
	}
	return iter
}

func iteradorInorder[K comparable, V any](nodo *nodoAbb[K, V], cola TDACola.Cola[*nodoAbb[K, V]],
	desde *K, hasta *K, cmp func(K, K) int) {
	//devuelve una cola con los elementos del arbol in orden
	if nodo == nil {
		return
	}
	if hasta != nil && cmp(nodo.clave, *hasta) > 0 {
		iteradorInorder(nodo.hijo_izq, cola, desde, hasta, cmp)
		return
	}
	if desde != nil && cmp(nodo.clave, *desde) < 0 {
		iteradorInorder(nodo.hijo_der, cola, desde, hasta, cmp)
		return
	}
	iteradorInorder(nodo.hijo_izq, cola, desde, hasta, cmp)
	if (desde == nil || cmp(nodo.clave, *desde) >= 0) && (hasta == nil || cmp(nodo.clave, *hasta) <= 0) {
		cola.Encolar(nodo)
	}
	iteradorInorder(nodo.hijo_der, cola, desde, hasta, cmp)
}

func iterarInorder[K comparable, V any](nodo *nodoAbb[K, V],
	desde *K, hasta *K, visitar func(clave K, dato V) bool, cmp func(K, K) int, condicion bool) bool {
	if nodo == nil {
		return true
	}
	if hasta != nil && cmp(nodo.clave, *hasta) > 0 {
		return iterarInorder(nodo.hijo_izq, desde, hasta, visitar, cmp, condicion)
	}
	if desde != nil && cmp(nodo.clave, *desde) < 0 {
		return iterarInorder(nodo.hijo_der, desde, hasta, visitar, cmp, condicion)
	}
	if !iterarInorder(nodo.hijo_izq, desde, hasta, visitar, cmp, condicion) {
		return false
	}
	if (desde == nil || cmp(nodo.clave, *desde) >= 0) && (hasta == nil || cmp(nodo.clave, *hasta) <= 0) {
		condicion = visitar(nodo.clave, nodo.dato)
	}
	if !iterarInorder(nodo.hijo_der, desde, hasta, visitar, cmp, condicion) {
		return false
	}
	return condicion
}

func (nodo *nodoAbb[K, V]) buscar_mas_izq_y_anterior(nodo_anterior *nodoAbb[K, V]) (*nodoAbb[K, V], *nodoAbb[K, V]) {
	//Busca el nodo mas a la izquierda y el anterior a ese.
	if nodo.hijo_izq == nil {
		return nodo, nodo_anterior
	}
	return nodo.hijo_izq.buscar_mas_izq_y_anterior(nodo)
}

func (ab *abb[K, V]) borrarRaiz() V {
	dato := ab.raiz.dato
	if ab.raiz.no_tiene_hijos() {
		ab.raiz = nil
	} else if ab.raiz.solo_hijo_izq() {
		ab.raiz = ab.raiz.hijo_izq
	} else if ab.raiz.solo_hijo_der() {
		ab.raiz = ab.raiz.hijo_der
	} else {
		ab.borrarCon2Hijos(ab.raiz, ab.raiz, ab.raiz.clave)
	}
	ab.cantidad--
	return dato
}

func (ab *abb[K, V]) tieneRaiz() bool {
	return ab.raiz != nil
}

func (nodo *nodoAbb[K, V]) devolverNodo(clave K, comp func(K, K) int) *nodoAbb[K, V] {
	//Busca el nodo de la clave dada
	if nodo == nil {
		return nil
	}
	if son_iguales(clave, nodo.clave, comp) {
		return nodo
	}
	if es_mayor(clave, nodo.clave, comp) {
		return nodo.hijo_der.devolverNodo(clave, comp)

	} else {
		return nodo.hijo_izq.devolverNodo(clave, comp)
	}
}

func (nodo *nodoAbb[K, V]) buscarPadre(clave K, comp func(a K, b K) int) *nodoAbb[K, V] {
	//busca el nodo que es padre (o lo seria) de la clave dada, ya sea que la misma exista o no.
	if es_mayor(clave, nodo.clave, comp) {
		if nodo.hijo_der == nil || son_iguales(clave, nodo.hijo_der.clave, comp) {
			return nodo
		}
		return nodo.hijo_der.buscarPadre(clave, comp)
	} else {
		if nodo.hijo_izq == nil || son_iguales(clave, nodo.hijo_izq.clave, comp) {
			return nodo
		}
		return nodo.hijo_izq.buscarPadre(clave, comp)
	}
}

func es_mayor[K comparable](clave1 K, clave2 K, cmp func(K, K) int) bool {
	return cmp(clave1, clave2) > 0
}

func son_iguales[K comparable](clave1 K, clave2 K, cmp func(K, K) int) bool {
	return cmp(clave1, clave2) == 0
}

func (ab *nodoAbb[K, V]) no_tiene_hijos() bool {
	return ab.hijo_izq == nil && ab.hijo_der == nil
}

func (ab *nodoAbb[K, V]) solo_hijo_izq() bool {
	return ab.hijo_izq != nil && ab.hijo_der == nil
}

func (ab *nodoAbb[K, V]) solo_hijo_der() bool {
	return ab.hijo_izq == nil && ab.hijo_der != nil
}

func (ab *abb[K, V]) borrarCon2Hijos(padre *nodoAbb[K, V], nodo *nodoAbb[K, V], clave K) V {
	nodo_reemp, padre_reemp := nodo.hijo_der.buscar_mas_izq_y_anterior(nodo)
	if padre_reemp != nodo {
		padre_reemp.hijo_izq = nodo_reemp.hijo_der
	} else {
		nodo.hijo_der = nodo_reemp.hijo_der
	}
	dato := nodo.dato
	nodo.clave = nodo_reemp.clave
	nodo.dato = nodo_reemp.dato
	return dato
}

func (ab *abb[K, V]) borrarSinHijos(padre *nodoAbb[K, V], nodo *nodoAbb[K, V], clave K) V {
	if es_mayor(clave, padre.clave, ab.cmp) {
		padre.hijo_der = nil
	} else {
		padre.hijo_izq = nil
	}
	return nodo.dato
}

func (ab *abb[K, V]) borrarHijoDer(padre *nodoAbb[K, V], nodo *nodoAbb[K, V], clave K) V {
	if es_mayor(clave, padre.clave, ab.cmp) {
		padre.hijo_der = nodo.hijo_der
	} else {
		padre.hijo_izq = nodo.hijo_der
	}
	return nodo.dato
}

func (ab *abb[K, V]) borrarHijoIzq(padre *nodoAbb[K, V], nodo *nodoAbb[K, V], clave K) V {
	if es_mayor(clave, padre.clave, ab.cmp) {
		padre.hijo_der = nodo.hijo_izq
	} else {
		padre.hijo_izq = nodo.hijo_izq
	}
	return nodo.dato
}
