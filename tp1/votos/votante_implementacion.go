package votos

import (
	"tp1/errores"
	TDAPila "tdas/pila"
)

type PersonaPadron interface {
	LeerDNI() int
	EsFraude() bool
	CambiarFraude()
}
type personaPadronImp struct {
	DNI     int
	ya_voto bool
}

func CrearPersonaPadron(dni int) PersonaPadron {
	persona := new(personaPadronImp)
	persona.DNI = dni
	persona.ya_voto = false
	return persona
}

func (persona *personaPadronImp) CambiarFraude() {
	persona.ya_voto = true
}

func (persona personaPadronImp) EsFraude() bool {
	return persona.ya_voto
}

func (persona personaPadronImp) LeerDNI() int {
	return persona.DNI
}

type votanteImplementacion struct {
	DNI       int
	Historial TDAPila.Pila[Voto]
}

func CrearVotante(dni int) Votante {
	votante := new(votanteImplementacion)
	votante.DNI = dni
	pila := TDAPila.CrearPilaDinamica[Voto]()
	votante.Historial = pila
	return votante
}

func (votante votanteImplementacion) LeerDNI() int {
	return votante.DNI
}

func (votante *votanteImplementacion) Votar(tipo TipoVoto, alternativa int) {
	if alternativa == 0 || (!votante.Historial.EstaVacia() && votante.Historial.VerTope().Impugnado) {
		votante.Historial.Apilar(Voto{[CANT_VOTACION]int{}, true})
	} else {
		var voto Voto
		if votante.Historial.EstaVacia() {
			voto = CrearVoto()
		} else {
			voto = votante.Historial.VerTope()
		}
		voto.Votar(tipo, alternativa)
		votante.Historial.Apilar(voto)
	}
}

func (votante *votanteImplementacion) Deshacer() error {
	if votante.Historial.EstaVacia() {
		return errores.ErrorNoHayVotosAnteriores{}
	}
	votante.Historial.Desapilar()
	return nil
}

func (votante *votanteImplementacion) FinVoto() Voto {
	voto_blanco := Voto{}
	if votante.Historial.EstaVacia() {
		return voto_blanco
	}
	voto := votante.Historial.Desapilar()
	return voto
}
