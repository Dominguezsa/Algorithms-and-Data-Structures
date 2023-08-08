package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	aero "tp2/aeropuerto"
	"tp2/errores"
)

const (
	POS_VERBO        = 0
	Agregar_archivo  = "agregar_archivo"
	Ver_tablero      = "ver_tablero"
	Info_vuelo       = "info_vuelo"
	Prioridad_vuelos = "prioridad_vuelos"
	Siguiente_vuelo  = "siguiente_vuelo"
	Borrar           = "borrar"
	error_comando    = "Error en comando "
	TodoBien         = "OK"
)

func main() {
	algueiza := aero.CrearAeropuerto()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		ingresado := strings.Split(input, " ")
		comando := ingresado[POS_VERBO]
		var res []string
		var err error
		if comando == Agregar_archivo {
			err = algueiza.AgregarVuelos(ingresado)
		} else if comando == Ver_tablero {
			res, err = algueiza.VerTablero(ingresado)
		} else if comando == Info_vuelo {
			res, err = algueiza.InfoVuelo(ingresado)
		} else if comando == Prioridad_vuelos {
			res, err = algueiza.PrioridadVuelos(ingresado)
		} else if comando == Siguiente_vuelo {
			res, err = algueiza.SiguienteVuelo(ingresado)
		} else if comando == Borrar {
			res, err = algueiza.BorrarVuelos(ingresado)
		}
		if err != nil {
			errores.ImprimirError(err)
		} else {
			for _, s := range res {
				fmt.Println(s)
			}
			fmt.Println(TodoBien)
		}
	}
}
