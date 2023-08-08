package funciones

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	TDACola "tdas/cola"
	"tp1/errores"
	"tp1/votos"
)

func validarDNI(cadena string) bool {
	num_dni, err := strconv.ParseInt(cadena, 10, 64)
	return len(cadena) <= 8 && len(cadena) > 0 && err == nil && num_dni > 0 && !strings.Contains(cadena, "+")
}

func validarCandidatos(cadena string) bool {
	slice_cadenas := strings.Split(cadena, ",")
	if len(slice_cadenas) != 4 {
		return false
	}
	for _, value := range slice_cadenas {
		if strings.TrimSpace(value) == "" {
			return false
		}
	}
	return true
}

func CargarPadron(path string) ([]votos.PersonaPadron, error) {
	datos := make([]int, 0)

	archivo, err := os.Open(path)
	if err != nil {
		return nil, errores.ErrorLeerArchivo{}
	}
	defer archivo.Close()
	s := bufio.NewScanner(archivo)
	for s.Scan() {
		linea := s.Text()
		if validarDNI(linea) {
			num, _ := strconv.ParseInt(linea, 10, 64)
			datos = append(datos, int(num))
		} else {
			return nil, errores.ErrorDatosPadron{}
		}
	}
	err = s.Err()
	if err != nil {
		return nil, errores.ErrorDatosPadron{}
	}
	padronOrdenado := RadixSort(datos)
	Padron := []votos.PersonaPadron{}
	for _, dni := range padronOrdenado {
		nuevo_dni := votos.CrearPersonaPadron(dni)
		Padron = append(Padron, nuevo_dni)
	}
	return Padron, nil
}

func CargarCandidatos(path string) ([]votos.Partido, error) {
	datos := make([][]string, 0)

	archivo, err := os.Open(path)
	if err != nil {
		return nil, errores.ErrorLeerArchivo{}
	}
	defer archivo.Close()
	s := bufio.NewScanner(archivo)
	for s.Scan() {
		linea := s.Text()
		if validarCandidatos(linea) {
			datos = append(datos, strings.Split(linea, ","))
		} else {
			return nil, errores.ErrorDatosCandidatos{}
		}
	}
	err = s.Err()
	if err != nil {
		return nil, errores.ErrorDatosCandidatos{}
	}
	partidos := make([]votos.Partido, 0)
	partido_blanco := votos.CrearPartidoEnBlanco()
	partidos = append(partidos, partido_blanco)
	for _, value := range datos {
		partido := votos.CrearPartido(value[0], [3]string{value[1], value[2], value[3]})
		partidos = append(partidos, partido)
	}
	return partidos, nil
}

func validarDNIEnPadron(dni string, padron []votos.PersonaPadron) bool {
	num_dni, _ := strconv.ParseInt(dni, 10, 64)
	bool, _ := BusquedaBinaria(padron, int(num_dni))
	return bool
}

func validarIngresar(cola TDACola.Cola[votos.Votante], input []string, padron []votos.PersonaPadron) (int, error) {
	if len(input) != 2 {
		return 0, errores.ErrorParametros{}
	}
	dni_string := input[1]
	dni_64, err := strconv.ParseInt(dni_string, 10, 64)
	if err != nil || int(dni_64) <= 0 {
		return 0, errores.DNIError{}
	}
	dni := int(dni_64)
	encontrado, _ := BusquedaBinaria(padron, dni)
	if !encontrado {
		return 0, errores.DNIFueraPadron{}
	}
	return dni, nil
}

func validarVotar(cola TDACola.Cola[votos.Votante], input []string, cant_partidos int, padron []votos.PersonaPadron) (string, int, error) {
	if cola.EstaVacia() {
		return "", 0, errores.FilaVacia{}
	}
	if len(input) != 3 {
		return "", 0, errores.ErrorParametros{}
	}
	cargo := input[1]
	if cargo != cargo_gobernador && cargo != cargo_intendente && cargo != cargo_presidente {
		return "", 0, errores.ErrorTipoVoto{}
	}
	num, err := strconv.ParseInt(input[2], 10, 64)
	if err != nil || num < 0 || int(num) >= cant_partidos {
		return "", 0, errores.ErrorAlternativaInvalida{}
	}
	votante := cola.VerPrimero()
	_, pos := BusquedaBinaria(padron, votante.LeerDNI())
	if padron[pos].EsFraude() {
		cola.Desencolar()
		return "", 0, errores.ErrorVotanteFraudulento{Dni: votante.LeerDNI()}
	}
	return cargo, int(num), nil
}

func validarDeshacer(cola TDACola.Cola[votos.Votante], input []string, padron []votos.PersonaPadron) error {
	if cola.EstaVacia() {
		return errores.FilaVacia{}
	}
	votante := cola.VerPrimero()
	_, pos := BusquedaBinaria(padron, votante.LeerDNI())
	if padron[pos].EsFraude() {
		cola.Desencolar()
		return errores.ErrorVotanteFraudulento{Dni: votante.LeerDNI()}
	}
	err := votante.Deshacer()
	return err
}

func validarFinVotar(
	cola TDACola.Cola[votos.Votante],
	Boletas []votos.Partido,
	input []string,
	padron []votos.PersonaPadron,
	votosImpugnados *int) (votos.Votante, int, error) {
	if cola.EstaVacia() {
		return nil, 0, errores.FilaVacia{}
	}
	votante := cola.Desencolar()
	_, pos := BusquedaBinaria(padron, votante.LeerDNI())
	if padron[pos].EsFraude() {
		return nil, 0, errores.ErrorVotanteFraudulento{Dni: votante.LeerDNI()}
	}
	return votante, pos, nil
}
