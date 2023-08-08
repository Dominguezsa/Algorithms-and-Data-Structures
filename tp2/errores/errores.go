package errores

import (
	"fmt"
	"os"
)

func ImprimirError(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err.Error())
}

const error_cmd string = "Error en comando "

type ErrorAgregarArchivo struct{}

func (e ErrorAgregarArchivo) Error() string {
	return error_cmd + "agregar_archivo"
}

type ErrorVerTablero struct{}

func (e ErrorVerTablero) Error() string {
	return error_cmd + "ver_tablero"
}

type ErrorInfoVuelo struct{}

func (e ErrorInfoVuelo) Error() string {
	return error_cmd + "info_vuelo"
}

type ErrorPrioridadVuelos struct{}

func (e ErrorPrioridadVuelos) Error() string {
	return error_cmd + "prioridad_vuelos"
}

type ErrorSiguienteVuelo struct{}

func (e ErrorSiguienteVuelo) Error() string {
	return error_cmd + "siguiente_vuelo"
}

type ErrorBorrar struct{}

func (e ErrorBorrar) Error() string {
	return error_cmd + "borrar"
}
