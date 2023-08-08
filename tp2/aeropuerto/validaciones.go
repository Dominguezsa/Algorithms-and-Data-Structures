package aeropuerto

import (
	"os"
	"strconv"
	"strings"
	"time"
	"tp2/errores"
)

const (
	VUELO_X = "indistinto" //sirve cuando quiero comparar fechas a las que no me interesa el codigo de vuelo
)

func validarAgregarVuelos(parametros []string) (string, error) {
	if len(parametros) != 2 {
		err := errores.ErrorAgregarArchivo{}
		return "", err
	}
	archivo, err := os.Open(parametros[1])
	if err != nil {
		return "", err
	}
	archivo.Close()
	return parametros[1], nil
}
func validarVerTablero(parametros []string) (int, error) {
	if len(parametros) != 5 {
		err := errores.ErrorVerTablero{}
		return 0, err
	}
	num, err := strconv.Atoi(parametros[1])
	if err != nil {
		err := errores.ErrorVerTablero{}
		return 0, err
	}
	return num, nil
}

func validarInfoVuelo(parametros []string) error {
	if len(parametros) != 2 {
		err := errores.ErrorInfoVuelo{}
		return err
	}
	return nil
}

func validarPrioridadVuelos(parametros []string) (int, error) {
	if len(parametros) != 2 {
		err := errores.ErrorPrioridadVuelos{}
		return -1, err
	}
	num, err := strconv.Atoi(parametros[1])
	if err != nil || num < 0 {
		err := errores.ErrorPrioridadVuelos{}
		return -1, err
	}
	return num, nil
}
func validarSiguienteVuelo(parametros []string) (string, error) {
	if len(parametros) != 4 {
		err := errores.ErrorSiguienteVuelo{}
		return "", err
	}
	ciudades := parametros[1] + "-" + parametros[2]
	return ciudades, nil
}
func validarBorrarVuelos(parametros []string) error {
	if len(parametros) != 3 {
		return errores.ErrorBorrar{}
	}
	return nil
}

func CmpPrioridadVuelos(v1 *vuelo_datos, v2 *vuelo_datos) int {
	prioridad1, _ := strconv.Atoi(v1.obtenerPrioridad()) //no puede existir el error porque el vuelo siempre deberia ser un int lo cual se verifica antes
	prioridad2, _ := strconv.Atoi(v2.obtenerPrioridad())
	if prioridad1 > prioridad2 {
		return 1
	}
	if prioridad1 < prioridad2 {
		return -1
	}
	//se ordena si empatan por codigo de vuelo
	if v1.obtenerCodigo() > v2.obtenerCodigo() {
		return -1
	}
	if v1.obtenerCodigo() < v2.obtenerCodigo() {
		return 1
	}
	return 0
}

func CmpFechas(a, b string) int {
	layout := "2006-01-02T15:04:05" // Formato de fecha esperado
	fecha1, _ := time.Parse(layout, a)
	fecha2, _ := time.Parse(layout, b)
	if fecha1.Before(fecha2) {
		return -1
	} else if fecha1.After(fecha2) {
		return 1
	}
	return 0
}

func CmpFechasParaAbbVuelos(a, b string) int {
	c, d := strings.Split(a, " "), strings.Split(b, " ")
	if c[1] == d[1] {
		return 0
	}
	comp := CmpFechas(c[0], d[0])
	if comp != 0 && c[1] != d[1] {
		return comp
	}
	if comp == 0 && (d[1] == VUELO_X || c[1] == VUELO_X) {
		return 0
	}
	if c[1] > d[1] {
		return 1
	}
	return -1
}