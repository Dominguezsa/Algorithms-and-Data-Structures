package errores

import "fmt"

func ImprimirError(err error) {
	fmt.Println(err.Error())
}

type ErrorLeerArchivo struct{}

func (e ErrorLeerArchivo) Error() string {
	return "ERROR: Lectura de archivos"
}

type ErrorParametros struct{}

func (e ErrorParametros) Error() string {
	return "ERROR: Faltan parámetros"
}

type ErrorDatosPadron struct{}

func (e ErrorDatosPadron) Error() string {
	return "ERROR: Hay padrones cargados que no son validos"
}

type ErrorDatosCandidatos struct{}

func (e ErrorDatosCandidatos) Error() string {
	return "ERROR: Hay Candidatos cargados que no son validos"
}

type DNIError struct{}

func (e DNIError) Error() string {
	return "ERROR: DNI incorrecto"
}

type DNIFueraPadron struct{}

func (e DNIFueraPadron) Error() string {
	return "ERROR: DNI fuera del padrón"
}

type FilaVacia struct{}

func (e FilaVacia) Error() string {
	return "ERROR: Fila vacía"
}

type ErrorVotanteFraudulento struct {
	Dni int
}

func (e ErrorVotanteFraudulento) Error() string {
	return fmt.Sprintf("ERROR: Votante FRAUDULENTO: %d", e.Dni)
}

type ErrorTipoVoto struct{}

func (e ErrorTipoVoto) Error() string {
	return "ERROR: Tipo de voto inválido"
}

type ErrorAlternativaInvalida struct{}

func (e ErrorAlternativaInvalida) Error() string {
	return "ERROR: Alternativa inválida"
}

type ErrorNoHayVotosAnteriores struct{}

func (e ErrorNoHayVotosAnteriores) Error() string {
	return "ERROR: Sin voto a deshacer"
}

type ErrorComandoInvalido struct{}

func (e ErrorComandoInvalido) Error() string {
	return "ERROR: Comando inválido"
}

type ErrorCiudadanosSinVotar struct{}

func (e ErrorCiudadanosSinVotar) Error() string {
	return "ERROR: Ciudadanos sin terminar de votar"
}

type ErrorComandoDesconocido struct{}

func (e ErrorComandoDesconocido) Error() string {
	return "ERROR: Comando desconocido"
}
