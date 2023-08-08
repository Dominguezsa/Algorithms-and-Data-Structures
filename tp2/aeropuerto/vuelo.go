package aeropuerto

import (
	"fmt"
	"strconv"
	"strings"
)

type vuelo_datos struct {
	datos []string
}

func crearVuelo(datos []string) *vuelo_datos {
	vuelo := new(vuelo_datos)
	vuelo.datos = datos
	return vuelo
}

func (v vuelo_datos) mostrarDatos() string {
	prioridad := v.datos[5]
	prioridadInt, _ := strconv.Atoi(prioridad)
	v.datos[5] = fmt.Sprint(prioridadInt)
	departureDelay := v.datos[7]
	intDelay, _ := strconv.Atoi(departureDelay)
	v.datos[7] = fmt.Sprint(intDelay)
	res := strings.Join(v.datos, " ")
	return res
}

func (v vuelo_datos) obtenerFecha() string {
	return v.datos[6]
}

func (v vuelo_datos) obtenerCiudades() string {
	return v.datos[2] + "-" + v.datos[3]
}

func (v vuelo_datos) obtenerCodigo() string {
	return v.datos[0]
}

func (v vuelo_datos) obtenerPrioridad() string {
	return v.datos[5]
}
