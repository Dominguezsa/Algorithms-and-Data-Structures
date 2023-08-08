package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	TDACola "tdas/cola"
	"tp1/errores"
	"tp1/funciones"
	"tp1/votos"
)

const (
	POS_VERBO        = 0
	opcion_ingresar  = "ingresar"
	opcion_votar     = "votar"
	opcion_deshacer  = "deshacer"
	opcion_fin_votar = "fin-votar"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println(errores.ErrorParametros{}.Error())
		return
	}
	//crea los partidos que contienen sus candidatos
	Partidos, errBoletas := funciones.CargarCandidatos(os.Args[1])
	if errBoletas != nil {
		errores.ImprimirError(errBoletas)
		return
	}
	//crea el padron
	Padron, errPadron := funciones.CargarPadron(os.Args[2])
	if errPadron != nil {
		errores.ImprimirError(errPadron)
		return
	}
	cola := TDACola.CrearColaEnlazada[votos.Votante]()
	votosImpugnados := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		ingresasdo := strings.Split(input, " ")
		comando := ingresasdo[POS_VERBO]
		var err error
		if comando == opcion_ingresar {
			err = funciones.Ingresar(cola, ingresasdo, Padron)
		} else if comando == opcion_votar {
			err = funciones.Votar(cola, ingresasdo, len(Partidos), Padron)
		} else if comando == opcion_deshacer {
			err = funciones.Deshacer(cola, ingresasdo, Padron)
		} else if comando == opcion_fin_votar {
			err = funciones.FinVotar(cola, Partidos, ingresasdo, Padron, &votosImpugnados)
		} else {
			err = errores.ErrorComandoInvalido{}
		}
		if err != nil {
			errores.ImprimirError(err)
		} else {
			fmt.Println("OK")
		}
	}
	if !cola.EstaVacia() {
		fmt.Println(errores.ErrorCiudadanosSinVotar{}.Error())
	}
	funciones.ImprimirResultados(Partidos, votosImpugnados)
}
