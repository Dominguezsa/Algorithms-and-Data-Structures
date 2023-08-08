from tdas.grafo import Grafo
from collections import deque
import heapq
from aeropuerto.constantes import *

def camino_mas(grafo: Grafo ,tipo_camino: str, origen): 
    '''camino_mas se encarga de devolver el camino mas rapido, barato o frecuente segun el tipo 
    de camino que recibe por parametro'''
    dist, padre,q = {}, {}, [] 
    for v in grafo.obtener_vertices(): 
        dist[v] = float('inf')
    dist[origen], padre[origen] = 0, None
    if tipo_camino == RAPIDO:     
        pos_peso = POS_TIEMPO_PROMEDIO
    elif tipo_camino == BARATO: 
        pos_peso = POS_PRECIO
    elif tipo_camino == FRECUENTE:
        pos_peso = POS_FRECUENCIA
    else:
        raise Exception("Tipo de camino invalido")
    heapq.heappush(q,(dist[origen],origen.codigo,origen))
    while len(q)>0:
        _,_,aeropuerto = heapq.heappop(q)
        for aerop_ady in grafo.adyacentes(aeropuerto):
            info_peso = grafo.peso(aeropuerto,aerop_ady) 
            peso = int(info_peso[pos_peso])
            if tipo_camino == FRECUENTE: peso = 1/peso
            if dist[aeropuerto] + peso < dist[aerop_ady]:
                dist[aerop_ady] = dist[aeropuerto] + peso
                padre[aerop_ady] = aeropuerto
                heapq.heappush(q,(dist[aerop_ady],aerop_ady.codigo,aerop_ady))
    return dist, padre

def devolver_camino(padre: dict,destino):
    '''devolver_camino se encarga de devolver cada vertice por el que debe pasar
    el vertice que recibe por parametro hasta llegar al origen'''
    res = []
    while destino != None:
        res.append(destino.codigo)
        destino = padre[destino]
    return res[::-1]

def encontrar_camino_escalas(grafo,origen):
    '''encontrar_camino_escalas se encarga de devolver el camino con menos escalas posibles'''
    dist,padre,visitados,q = {},{},set(), deque()
    dist[origen],padre[origen] = 0,None
    q.appendleft(origen)
    visitados.add(origen)
    while len(q)>0:
        v = q.pop()
        for ady in grafo.adyacentes(v):
            if ady not in visitados:
                visitados.add(ady)
                padre[ady]=v 
                dist[ady]=dist[v]+1
                q.appendleft(ady)
    return dist, padre

def centralidad(grafo):
    '''centralidad se encarga de devolver los vertices con sus resepctivos valores como centrales'''
    cent = {}
    vertices = grafo.obtener_vertices()
    for v in vertices:
        cent[v] = 0
    for v in vertices:
        distancia, padre = camino_mas(grafo, FRECUENTE, origen=v)
        centr_aux = {}
        for w in vertices:
            centr_aux[w] = 0
        vertices_ordenados = ordenar_vertices(vertices, distancia)
        for w in vertices_ordenados:
            if padre[w] is not None:
                centr_aux[padre[w]] += 1 + centr_aux[w]
        for w in vertices:
            if w == v: continue
            cent[w] += centr_aux[w]
    return cent

def ordenar_vertices(vertices,distancia):
    '''ordenar_vertices ordena los vertices de menor a mayor segun la distancia que recibe por parametro'''
    return sorted(vertices, key=lambda x: int(distancia[x]))

def orden_topologico_bfs(grafo: Grafo):
    '''orden_topologico_bfs se encarga de devolver un orden topologico del grafo que recibe por parametro'''
    grados, res, q = {}, [], deque()
    vertices = grafo.obtener_vertices()
    for v in vertices: grados[v] = 0
    for v in vertices:
        for w in grafo.adyacentes(v): grados[w] += 1
    for v in vertices:
        if grados[v] == 0: q.appendleft(v)
    while len(q) > 0:
        v = q.pop()
        res.append(v)
        for w in grafo.adyacentes(v):
            grados[w] -= 1
            if grados[w] == 0: q.appendleft(w)
    return res

def _dfs(grafo, lista, visitados, v, grafo_grande):
    '''_dfs se encarga de recorrer el grafo y agregar a la lista la informacion de cada aeropuerto'''
    visitados.add(v)
    for w in grafo.adyacentes(v):
        if w not in visitados:
            peso = grafo_grande.peso(v, w)
            tiempo, precio, cant_vuelos = peso[0], peso[1], peso[2]
            lista.append((v.codigo, w.codigo, tiempo, precio, cant_vuelos))
            _dfs(grafo, lista, visitados, w, grafo_grande)

def mst_precios(grafo):
    '''mst_precios devuelve el arbol de tendido minimo del grafo que recibe por parametro'''
    vertices = grafo.obtener_vertices()
    conjuntos = UnionFind(len(vertices))
    mst = Grafo(dirigido=False)
    aristas = grafo.obtener_aristas()
    aristas.sort(key=lambda x: int(x[POS_PESO_ARISTA][POS_PRECIO]))
    indice, i = {}, 0
    for v in vertices: 
        indice[v] = i 
        i += 1 
        mst.agregar_vertice(v)
    for a in aristas:
        v, w, peso = a
        peso = int(peso[POS_PRECIO])
        if conjuntos.find(indice[v]) != conjuntos.find(indice[w]):
            conjuntos.union(indice[v], indice[w])
            mst.agregar_arista(v, w, peso)
    return mst

class UnionFind:
    '''UnionFind es una clase auxiliar que se encarga de crear conjuntos disjuntos
    para aplicar el algoritmo de Kruskal'''
    def __init__(self, n):
        self.groups = list(range(n))

    def find(self, v):
        if self.groups[v] != v:
            self.groups[v] = self.find(self.groups[v])
        return self.groups[v]

    def union(self, u, v):
        new_group, other = self.find(u), self.find(v)
        if new_group != other:
            self.groups[new_group] = other
