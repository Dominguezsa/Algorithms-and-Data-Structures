package votos

import (
	"strconv"
)

type partidoImplementacion struct {
	Nombre       string
	Candidatos   [CANT_VOTACION]string
	VotosPorTipo [CANT_VOTACION]int
}

type partidoEnBlanco struct {
	VotosPorTipo [CANT_VOTACION]int
}

func CrearPartidoEnBlanco() Partido {
	blanco := new(partidoEnBlanco)
	return blanco
}

func CrearPartido(nombre string, candidatos [int(CANT_VOTACION)]string) Partido {
	partido := new(partidoImplementacion)
	partido.Nombre = nombre
	partido.Candidatos = candidatos
	return partido
}

func CrearVotosEnBlanco() Partido {
	blanco := new(partidoEnBlanco)
	return blanco
}

func (partido *partidoImplementacion) VotadoPara(tipo TipoVoto) {
	partido.VotosPorTipo[tipo]++
}

func (partido partidoImplementacion) ObtenerResultado(tipo TipoVoto) string {
	votosEnString := strconv.Itoa(partido.VotosPorTipo[tipo])
	if partido.VotosPorTipo[tipo] == 1 {
		return partido.Nombre + " - " + partido.Candidatos[tipo] + ": 1 voto"
	} else {
		return partido.Nombre + " - " + partido.Candidatos[tipo] + ": " + votosEnString + " votos"
	}
}

func (blanco *partidoEnBlanco) VotadoPara(tipo TipoVoto) {
	blanco.VotosPorTipo[tipo]++
}

func (blanco partidoEnBlanco) ObtenerResultado(tipo TipoVoto) string {
	votosEnString := strconv.Itoa(blanco.VotosPorTipo[tipo])
	if blanco.VotosPorTipo[tipo] == 1 {
		return "Votos en Blanco: 1 voto"
	} else {
		return "Votos en Blanco: " + votosEnString + " votos"
	}
}
