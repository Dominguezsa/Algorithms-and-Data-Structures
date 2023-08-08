from aeropuerto.constantes import *

def validar_camino_mas_o_escalas(parametros):
    parametros = parametros.split(",")
    if parametros[0] in [BARATO,RAPIDO] and len(parametros)==3:
        return parametros
    elif len(parametros)==2:
        return parametros
    raise Exception(ERROR_PARAMETROS)

def validar_centralidad(parametros):
    if not parametros.isnumeric():
        raise Exception(ERROR_CENTRALIDAD)
    return int(parametros)

def validar_nueva_aerolinea(parametros):
    parametros = parametros.split(".")
    if len(parametros)!=2 and parametros[1] !="csv":
        raise Exception(ERROR_NUEVA_AEROLINEA)

def validar_itinerario(parametros):
    parametros = parametros.split(".")
    if len(parametros)>2 and parametros[len(parametros)-1].lower() =="csv":
        raise Exception(ERROR_ITINERARIO)
    
def validar_exportar_kml(parametros):
    #si no termina en .kml devuelve error
    parametros = parametros.split(".")
    if len(parametros)!=2 and parametros[1] !="kml":
        raise Exception(ERROR_KML)
