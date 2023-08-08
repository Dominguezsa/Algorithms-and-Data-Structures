import aeropuerto.validaciones as validaciones
from tdas.grafo import Grafo
import aeropuerto.archivos as archivos
from aeropuerto.biblioteca import(camino_mas,encontrar_camino_escalas,centralidad,
orden_topologico_bfs, mst_precios, devolver_camino)
import heapq
from aeropuerto.constantes import *

class AdministradorAeropuertos:
    '''El Administrador_aeropuertos es un objeto que se encarga de administrar las conexiones entre distintos
    aeropuertos como las caracteristicas de los mismos.'''
    def __init__(self):
        self.conexiones_aeropuertos = Grafo()
        self.ciudades_aeropuertos = {} #{"Mar del Plata": [Aeropuerto{"MDQ"}, Aeropuerto{"MDQ"}]}
        self.codigos_aeropuertos = {}    #{"MDQ": Aeropuerto("Mar del Plata", "MDQ", "37° 58' 60S", "57° 32' 60W")}

    def obtener_aeropuertos(self, ciudad):
        '''obtener_aeropuertos recibe una ciudad y devuelve una lista con los aeropuertos que se encuentran en la misma.'''
        if ciudad not in self.ciudades_aeropuertos:
            return []
        return self.ciudades_aeropuertos[ciudad]
    
    def camino_mas(self,parametros):
        '''camino_mas recibe una lista de parametros y devuelve el camino mas rapido, barato, o con menos escalas
        de un origen a un destino, los cuales se encuentras en los parametros ingresados.'''
        parametros = validaciones.validar_camino_mas_o_escalas(parametros)
        if len(parametros) == 3:
           tipo_camino=parametros[0]
           origen=parametros[1]
           destino=parametros[2]
           es_escalas = False
        else:
            origen = parametros[0]
            destino = parametros[1]
            es_escalas = True
        aeropuertos_origen, aeropuertos_destino = self.obtener_aeropuertos(origen), self.obtener_aeropuertos(destino)
        camino, acumulado = [], float("inf")
        for aero_origen in aeropuertos_origen:
            if not es_escalas:
                distancia, padres = camino_mas(self.conexiones_aeropuertos,tipo_camino,aero_origen) #type: ignore
            else:
                distancia, padres = encontrar_camino_escalas(self.conexiones_aeropuertos, aero_origen)
            for aero_destino in aeropuertos_destino:
                if distancia[aero_destino] < acumulado:
                    acumulado = distancia[aero_destino]
                    camino = devolver_camino(padres,aero_destino)
        return camino

    def itinerario(self, parametros):
        '''itinerario recibe una lista de parametros y devuelve el orden topologico de las ciudades entre las mismas.'''
        validaciones.validar_itinerario(parametros)
        grafo_visita_ciudades = archivos.crear_grafo_desde_archivo(parametros)
        ciudades_ordenadas = orden_topologico_bfs(grafo_visita_ciudades)
        caminos_minimos_ciudades= self.obtener_caminos_minimos(ciudades_ordenadas)
        res_imprimir = [ciudades_ordenadas]
        for camino in caminos_minimos_ciudades:
            res_imprimir.append(camino)
        return res_imprimir
    
    def obtener_caminos_minimos(self, ciudades_ordenadas):
        '''obtener_caminos_minimos recibe una lista de ciudades y devuelve una lista con
        los caminos mas rapidos entre las mismas.'''
        caminos_minimos_ciudades = []
        for i in range(len(ciudades_ordenadas)-1):
            ciudad_origen = ciudades_ordenadas[i]
            ciudad_destino = ciudades_ordenadas[i+1]
            aeropuertos_origen = self.obtener_aeropuertos(ciudad_origen)
            aeropuertos_destino = self.obtener_aeropuertos(ciudad_destino)
            camino = ""
            acumulado = float("inf")
            for aero_origen in aeropuertos_origen:
                distancia, padres = camino_mas(self.conexiones_aeropuertos,RAPIDO,aero_origen)
                for aero_destino in aeropuertos_destino:
                    if distancia[aero_destino] < acumulado:
                        acumulado = distancia[aero_destino]
                        camino = devolver_camino(padres,aero_destino)
            caminos_minimos_ciudades.append(camino)
        return caminos_minimos_ciudades

    def obtener_k_centrales(self, k):
        '''obtener_k_centrales recibe un numero k y devuelve los k aeropuertos mas centrales, utilizando
        el algoritmo de centralidad Betweenness.'''
        k = validaciones.validar_centralidad(k)
        centralidades = centralidad(self.conexiones_aeropuertos)
        vertices_centrales_heap = [(centralidades[aerop] * -1,aerop.codigo,aerop) for aerop in centralidades.keys()]
        heapq.heapify(vertices_centrales_heap)
        k_centrales=[]
        for _ in range(k):#type: ignore
            _,_,aeropuerto=heapq.heappop(vertices_centrales_heap)
            k_centrales.append(aeropuerto.codigo)
        k_centrales=", ".join(k_centrales)
        return k_centrales
            
    def exportar_kml(self, archivo,kml):
        '''exportar_kml recibe un archivo y una lista de codigos de aeropuertos y devuelve un archivo kml
        que contienen a los aeropuertos que se usaron en el ultimo comando del camino más rápido, barato
        o con menos escalas.'''
        validaciones.validar_exportar_kml(archivo)
        aeropuertos = [] 
        for clave in kml:
            aeropuertos.append(self.codigos_aeropuertos[clave])
        archivos.escribir_kml(archivo,aeropuertos)
    
    def nueva_aerolinea(self, parametros):
        '''nueva_aerolinea recibe la ruta de un archivo para escribir en el las conexiones mas baratas entre
        todos los aeropuertos que se registraron.'''
        validaciones.validar_nueva_aerolinea(parametros)
        mst = mst_precios(self.conexiones_aeropuertos)
        archivos.escribir_ruta_optima(parametros,mst, self.conexiones_aeropuertos)


