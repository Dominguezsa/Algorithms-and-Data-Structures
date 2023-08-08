package aeropuerto

type Aeropuerto interface {

	//InfoVuelo devuelve la informacion del vuelo en base al codigo pasado por parametro.
	InfoVuelo([]string) ([]string, error)

	//Agregar_vuelos carga la informacion de vuelos de un archivo pasado por parametro
	//en los tdas internos del aeropuerto, para su posterior consulta / modificacion.
	AgregarVuelos([]string) error

	//Ver tablero devuelve info de los k vuelos correspondientes a las fechas y k cantidad indicadas, ordenados por fecha.
	VerTablero([]string) ([]string, error)

	//PrioridadVuelos devuelve los k vuelos ordenados por prioridad.
	PrioridadVuelos([]string) ([]string, error)

	//SiguienteVuelo devuelve el siguiente vuelo que sale de una ciudad en una fecha.
	SiguienteVuelo([]string) ([]string, error)

	//Borrar elimina los vuelos que parten desde una fecha (incluida para borrar) hasta otra (incluida para borrar).
	BorrarVuelos([]string) ([]string, error)
}

type Vuelo interface {
	//mostrar datos devuelve un string con todos los datos del vuelo separados por un espacio.
	mostrarDatos() string
}
