package funciones

import (
	"fmt"
	"tp1/votos"
	TDACola "tdas/cola"
)

const (
	cargo_presidente string = "Presidente"
	cargo_gobernador string = "Gobernador"
	cargo_intendente string = "Intendente"
)

func Ingresar(cola TDACola.Cola[votos.Votante], input []string, padron []votos.PersonaPadron) error {
	dni, err := validarIngresar(cola, input, padron)
	if err != nil {
		return err
	}
	votante := votos.CrearVotante(int(dni))
	cola.Encolar(votante)
	return nil
}

func Votar(cola TDACola.Cola[votos.Votante], input []string, cant_partidos int, padron []votos.PersonaPadron) error {
	cargo, num, err := validarVotar(cola, input, cant_partidos, padron)
	if err != nil {
		return err
	}
	votante := cola.VerPrimero()
	if cargo == cargo_presidente {
		votante.Votar(votos.PRESIDENTE, int(num))
	} else if cargo == cargo_gobernador {
		votante.Votar(votos.GOBERNADOR, int(num))
	} else {
		votante.Votar(votos.INTENDENTE, int(num))
	}
	return nil
}

func Deshacer(cola TDACola.Cola[votos.Votante], input []string, padron []votos.PersonaPadron) error {
	err := validarDeshacer(cola, input, padron)
	if err != nil {
		return err
	}
	return nil
}

func FinVotar(cola TDACola.Cola[votos.Votante],
	Boletas []votos.Partido, input []string, padron []votos.PersonaPadron, votosImpugnados *int) error {
	votante, pos, err := validarFinVotar(cola, Boletas, input, padron, votosImpugnados)
	if err != nil {
		return err
	}
	voto := votante.FinVoto()
	if voto.Impugnado {
		*votosImpugnados++
		return nil
	}
	lista_de_votos := voto.Devolver_votos()
	for i, v := range lista_de_votos {
		if i == int(votos.PRESIDENTE) {
			Boletas[v].VotadoPara(votos.PRESIDENTE)
		} else if i == int(votos.GOBERNADOR) {
			Boletas[v].VotadoPara(votos.GOBERNADOR)
		} else {
			Boletas[v].VotadoPara(votos.INTENDENTE)
		}
	}
	padron[pos].CambiarFraude()
	return nil
}

func ImprimirResultados(Partidos []votos.Partido, votosImpugnados int) {
	fmt.Printf("%v:\n", cargo_presidente)
	for _, partido := range Partidos {
		fmt.Println(partido.ObtenerResultado(votos.PRESIDENTE))
	}

	fmt.Printf("\n%v:\n", cargo_gobernador)
	for _, partido := range Partidos {
		fmt.Println(partido.ObtenerResultado(votos.GOBERNADOR))
	}

	fmt.Printf("\n%v:\n", cargo_intendente)
	for _, partido := range Partidos {
		fmt.Println(partido.ObtenerResultado(votos.INTENDENTE))
	}

	if votosImpugnados == 1 {
		fmt.Println("\nVotos Impugnados: 1 voto")
	} else {
		fmt.Println("\nVotos Impugnados:", votosImpugnados, "votos")
	}
}
