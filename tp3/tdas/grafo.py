import random

class Grafo:
    def __init__(self, dirigido=False):
        self.dirigido = dirigido
        self.aristas = {} # {A: {B: 1, C: 2}, B: {A: 1}, C: {A: 2}

    def agregar_vertice(self, v):   
        '''Agrega un vertice al grafo.'''
        if not v in self.aristas: self.aristas[v] = {}
        else: raise ValueError(f"El vertice {v} ya existe en el grafo.")

    def borrar_vertice(self, v):
        '''Borra un vertice del grafo.'''
        self._error_vertice_no_existe(v)
        if not self.dirigido:
            for w in self.aristas[v]: del self.aristas[w][v]
        elif self.dirigido:
            for w in self.aristas:
                if v in self.aristas[w]: del self.aristas[w][v]
        del self.aristas[v]

    def agregar_arista(self, v, u, peso=1):
        '''Agrega una arista al grafo.'''
        self._error_vertice_no_existe(v)
        self._error_vertice_no_existe(u)
        self.aristas[v][u] = peso
        if not self.dirigido: self.aristas[u][v] = peso

    def borrar_arista(self, v, u):
        '''Borra una arista del grafo.'''
        self._error_vertice_no_existe(v)
        self._error_vertice_no_existe(u)
        del self.aristas[v][u]
        if not self.dirigido: del self.aristas[u][v]

    def esta_unido(self, v, u):
        '''Devuelve un bool si los vertices estan unidos, error en caso de que no exista alguno'''
        self._error_vertice_no_existe(v)
        self._error_vertice_no_existe(u)
        return u in self.aristas[v]

    def adyacentes(self, v):
        '''Devuelve los vertices adyacentes al vertice pasado por parametro'''
        self._error_vertice_no_existe(v)
        return set(self.aristas[v].keys())

    def peso(self, v, u):
        '''Devuelve el peso de la arista que une los vertices pasados por parametro'''
        self._error_vertice_no_existe(v)
        self._error_vertice_no_existe(u)
        if not self.esta_unido(v, u):
            raise ValueError(f"El vertice {v} no está relacionado con {u}")
        return self.aristas[v][u]

    def obtener_vertice_aleatorio(self):
        '''Devuelve un vertice aleatorio del grafo'''
        if len(self.aristas)==0:
            raise ValueError("El grafo esta vacío.")
        return random.choice(list(self.aristas.keys()))

    def obtener_vertices(self):
        '''Devuelve una lista con todos los vertices del grafo'''
        return list(self.aristas.keys())

    def obtener_aristas(self):
        '''Devuelve una lista con todas las aristas del grafo'''
        aristas=[]
        for v in self.aristas:
            for w in self.aristas[v]: aristas.append((v,w,self.aristas[v][w]))
        return aristas

    def __len__(self):
        '''Devuelve la cantidad de vertices del grafo'''
        return len(self.aristas)

    def _error_vertice_no_existe(self,v):
        if v not in self.aristas:
            raise IndexError(f"El vertice {v} no existe en el grafo.")