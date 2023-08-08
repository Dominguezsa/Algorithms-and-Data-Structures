from tdas.grafo import Grafo
from aeropuerto.constantes import OK
from tdas.aeropuerto import Aeropuerto
from aeropuerto.biblioteca import _dfs

def crear_grafo_desde_archivo(archivo):
    '''crear_grafo_desde_archivo crea un grafo de vertices ciudades(ej: Oklahoma") y aristas vuelos(ej: Oklahoma, Atlanta, 1)'''
    grafo = Grafo(dirigido=True)
    try:
        with open (archivo, "r") as archivo:
            for linea in archivo:
                linea = linea.rstrip('\n').split(',')
                if len(linea) > 2:
                    for i in linea:
                        grafo.agregar_vertice(i)
                elif len(linea) == 2:
                    origen, destino = linea[0], linea[1]
                    grafo.agregar_arista(origen, destino,1)
        return grafo
    except:
        raise Exception("Error al leer el archivo")

def escribir_ruta_optima(ruta, grafo_mst: Grafo, grafo_grande: Grafo):
    '''Se escribe en un archivo las conexiones entre todos los aeropuertos respetantando 
    que sea lo mas barato posible para recorrer todos los aeropuertos'''
    lista, visitados = [], set()
    origen = grafo_grande.obtener_vertice_aleatorio()
    _dfs(grafo_mst, lista, visitados, origen, grafo_grande)
    try:
        with open(ruta, "w") as archivo:
            for linea in lista: archivo.write(f"{linea[0]},{linea[1]},{linea[2]},{linea[3]},{linea[4]}\n")
    except:
        raise Exception("Error al escribir el archivo")

def escribir_kml(ruta, aeropuertos):
    '''escribir_kml se encarga de escribir un archivo kml con la informacion de los aeropuertos y
    las lineas que los conectan'''
    try:
        with open (ruta, "w") as archivo:
            archivo.write('<?xml version="1.0" encoding="UTF-8"?>\n')
            archivo.write('<kml xmlns="http://earth.google.com/kml/2.1">\n')
            archivo.write("    <Document>\n")
            archivo.write("        <name>KML</name>\n")
            #escribir aeropuertos
            for aeropuerto in aeropuertos:
                archivo.write("        <Placemark>\n")
                archivo.write("            <name>"+aeropuerto.codigo+"</name>\n")
                archivo.write("            <Point>\n")
                archivo.write("                <coordinates>"+str(aeropuerto.longitud)+", "+str(aeropuerto.latitud)+"</coordinates>\n")
                archivo.write("            </Point>\n")
                archivo.write("        </Placemark>\n\n")
            #escribir lineas
            for i in range(len(aeropuertos)-1):
                archivo.write("        <Placemark>\n")
                archivo.write("            <LineString>\n")
                archivo.write("                <coordinates>"+str(aeropuertos[i].longitud)+", "+str(aeropuertos[i].latitud)+" "+str(aeropuertos[i+1].longitud)+", "+str(aeropuertos[i+1].latitud)+"</coordinates>\n")
                archivo.write("            </LineString>\n")
                archivo.write("        </Placemark>\n")
            archivo.write("    </Document>\n")
            archivo.write("</kml>\n")
    except:
        raise Exception("Error al escribir el archivo")


def completar_info_aeropuertos(flycombi,archivo_aeropuertos, archivo_vuelos):
    '''completar_info_aeropuertos recibe dos rutas de archivos, uno con informacion de aeropuertos y otro con informacion
    de vuelos, para guardar esa información y poder reutilizarla en el futuro.'''
    try:
        with open (archivo_aeropuertos, "r") as archivo:
            for linea in archivo:
                linea = linea.rstrip('\n').split(',')
                ciudad,codigo,latitud,longitud=linea[0], linea[1], linea[2], linea[3]
                aero = Aeropuerto(ciudad,codigo,latitud,longitud)
                flycombi.conexiones_aeropuertos.agregar_vertice(aero)
                flycombi.codigos_aeropuertos[codigo]=aero            
                if ciudad not in flycombi.ciudades_aeropuertos:
                    flycombi.ciudades_aeropuertos[ciudad]=[]
                flycombi.ciudades_aeropuertos[ciudad].append(aero)
        with open (archivo_vuelos, "r") as archivo:
            for linea in archivo:
                linea = linea.rstrip('\n').split(',')
                codigo_aeropuerto_1,codigo_aeropuerto_2,tiempo_promedio,precio,cant_vuelos= linea
                aeropuerto_i = flycombi.codigos_aeropuertos[codigo_aeropuerto_1]
                aeropuerto_j = flycombi.codigos_aeropuertos[codigo_aeropuerto_2]
                info_peso = [int(tiempo_promedio),int(precio),int(cant_vuelos)]
                flycombi.conexiones_aeropuertos.agregar_arista(aeropuerto_i, aeropuerto_j,info_peso)#type:ignore
    except FileNotFoundError:
        raise Exception("Error al leer el archivo")
            

