#!/usr/bin/python3
import sys
from tdas.administrador import AdministradorAeropuertos
from aeropuerto.constantes import *
from aeropuerto.archivos import completar_info_aeropuertos

def main(flycomby: AdministradorAeropuertos):
    kml = []#list of flights ej: ["ASF","SSD", "LKS"]
    for input in sys.stdin:
        try: 
            ingresado = input.rstrip("\n").split(" ",1)
            if len(ingresado) <INPUT_MINIMO:
                continue
            comando = ingresado[POS_COMANDO]
            parametros = ingresado[POS_PARAMETROS]
            res = ""
            if comando == CAMINO_MAS:
                res = flycomby.camino_mas(parametros)
                kml = res
                print(SEPARADOR_FLECHAS.join(res))
            elif comando == CAMINO_ESCALAS:
                res = flycomby.camino_mas(parametros)
                kml = res
                print(SEPARADOR_FLECHAS.join(res))
            elif comando == CENTRALIDAD:
                res = flycomby.obtener_k_centrales(parametros)
                print(SEPARADOR_COMA.join(res))
            elif comando == ITINERARIOS:
                res = flycomby.itinerario(parametros)
                for i in range(len(res)):
                    if i == POSICION_CIUDADES:
                        print(SEPARADOR_COMA.join(res[i]))
                    else:
                        print(SEPARADOR_FLECHAS.join(res[i]))
            elif comando == NUEVA_AEROLINEA:
                flycomby.nueva_aerolinea(parametros)
                print(OK)
            elif comando == EXPORTAR_KML:
                flycomby.exportar_kml(parametros,kml)
                print(OK)
            else:
                continue
        except Exception as e:
            print("Error:", e)

if __name__ == '__main__':
    if len(sys.argv) != ENTRADA_MINIMA:
        print("Error en los parametros de entrada")
    else:
        flycomby = AdministradorAeropuertos()
        archivo_aeropuertos  = sys.argv[POS_AEROPUERTOS]
        archivo_vuelos = sys.argv[POS_ARISTAS]
        completar_info_aeropuertos(flycomby, archivo_aeropuertos, archivo_vuelos)
        main(flycomby)

