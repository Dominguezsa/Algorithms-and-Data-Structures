class Aeropuerto:
    '''El Aeropuerto es un objeto que funciona como un vertice del grafo, el cual
       tiene como atributos la ciudad, el codigo, la latitud y la longitud.'''
    def __init__(self, ciudad,codigo, latitud, longitud):
        self.ciudad = ciudad
        self.codigo = codigo
        self.latitud = latitud
        self.longitud = longitud