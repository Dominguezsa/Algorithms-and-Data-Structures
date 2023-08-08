package aeropuerto

import (
	"bufio"
	"os"
	"strings"
	TDACola_prioridad "tdas/cola_prioridad"
	TDADict "tdas/diccionario"
	errores "tp2/errores"
)

const (
	ASCENDENTE  = "asc"
	DESCENDENTE = "desc"
)

func CrearAeropuerto() Aeropuerto {
	aero := new(aeropuerto)
	aero.ciudadesConFechas = TDADict.CrearHash[string, TDADict.DiccionarioOrdenado[string, *vuelo_datos]]()
	aero.dataVuelos = TDADict.CrearHash[string, *vuelo_datos]()
	aero.fechasVuelos = TDADict.CrearABB[string, *vuelo_datos](CmpFechasParaAbbVuelos)
	return aero
}

type aeropuerto struct {
	ciudadesConFechas TDADict.Diccionario[string, TDADict.DiccionarioOrdenado[string, *vuelo_datos]] //guarda ciudades cuyos valores son (fechas:vuelos)
	dataVuelos        TDADict.Diccionario[string, *vuelo_datos]                                      //guarda codigos de vuelo con su info correspondiene
	fechasVuelos      TDADict.DiccionarioOrdenado[string, *vuelo_datos]                              //guarda fechas con sus vuelos
}

func (a *aeropuerto) InfoVuelo(parametros []string) ([]string, error) {
	//Devuelve la informacion del vuelo y un posible mensaje de error
	err := validarInfoVuelo(parametros)
	if err != nil {
		return nil, err
	}
	codigo := parametros[1]
	if !a.dataVuelos.Pertenece(codigo) {
		return nil, errores.ErrorInfoVuelo{}
	}
	res := make([]string, 1)
	vuelo := a.dataVuelos.Obtener(codigo)
	info := vuelo.mostrarDatos()
	res[0] = info
	return res, nil
}

func (a *aeropuerto) PrioridadVuelos(parametros []string) ([]string, error) {
	// ejemplo res:["10 - 1234","03 - 1624","03 - 325",OK]
	kVuelos, err := validarPrioridadVuelos(parametros)
	if err != nil {
		return nil, err
	}
	sliceVuelosInicial := make([]*vuelo_datos, 0)
	sliceVuelosResultado := make([]string, 0)
	for iter := a.dataVuelos.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		_, vuelo := iter.VerActual()
		sliceVuelosInicial = append(sliceVuelosInicial, vuelo)
	}
	heapTemporal := TDACola_prioridad.CrearHeapArr(sliceVuelosInicial, CmpPrioridadVuelos)
	for i := 0; i < kVuelos; i++ {
		if heapTemporal.EstaVacia() {
			break
		}
		vuelo := heapTemporal.Desencolar()
		prioridad := vuelo.obtenerPrioridad()
		sliceVuelosResultado = append(sliceVuelosResultado, prioridad+" - "+vuelo.obtenerCodigo())
	}
	return sliceVuelosResultado, nil
}

func (a *aeropuerto) VerTablero(parametros []string) ([]string, error) {
	k_vuelos, err := validarVerTablero(parametros)
	if err != nil {
		return nil, err
	}
	orden := parametros[2]
	//La clave del abb es "<fecha> <codigo>". La funcion de cmp del abb revisa primero si el codigo de vuelo
	//es igual en ambas claves a comparar.
	//Si no lo son, compara por fechas. Por eso, agregamos este codigo falso para que lo compare, siempre es
	//diferente, y luego revise con la fecha que es lo que nos interesa recorrer.
	claveFecha1 := parametros[3] + " " + VUELO_X
	claveFecha2 := parametros[4] + " " + VUELO_X
	if orden == ASCENDENTE || orden == DESCENDENTE {
		return ver_tablero(k_vuelos, orden, claveFecha1, claveFecha2, a.fechasVuelos)
	}
	return nil, errores.ErrorVerTablero{}
}

func ver_tablero(n int, orden string, fecha1, fecha2 string, abb TDADict.DiccionarioOrdenado[string, *vuelo_datos]) ([]string, error) {
	// ejemplo res:2["018-10-10T08:51:32 - 1234", ..., "OK"]
	preRespuesta := make([]string, 0)
	res := make([]string, 0)
	for iter := abb.IteradorRango(&fecha1, &fecha2); iter.HaySiguiente(); iter.Siguiente() {
		_, vuelo := iter.VerActual()
		fraseParaImprimir := vuelo.obtenerFecha() + " - " + vuelo.obtenerCodigo()
		preRespuesta = append(preRespuesta, fraseParaImprimir)
	}
	if orden == ASCENDENTE {
		for i := 0; i < n && i < len(preRespuesta); i++ {
			res = append(res, preRespuesta[i])
		}
	} else {
		for i := len(preRespuesta) - 1; i >= 0 && i >= len(preRespuesta)-n; i-- {
			res = append(res, preRespuesta[i])
		}
	}
	return res, nil
}

func (a *aeropuerto) SiguienteVuelo(parametros []string) ([]string, error) {
	claveCiudades, err := validarSiguienteVuelo(parametros)
	if err != nil {
		return nil, err
	}
	ciudades := strings.Split(claveCiudades, "-")
	if !a.ciudadesConFechas.Pertenece(claveCiudades) {
		frase := "No hay vuelo registrado desde " + ciudades[0] + " hacia " + ciudades[1] + " desde " + parametros[3]
		res := make([]string, 0)
		res = append(res, frase)
		return res, nil
	}
	res := make([]string, 1)
	abb := a.ciudadesConFechas.Obtener(claveCiudades)
	desdeFecha := parametros[3]
	for iter := abb.IteradorRango(&desdeFecha, nil); iter.HaySiguiente(); iter.Siguiente() {
		claveFecha, valor := iter.VerActual()
		if CmpFechas(claveFecha, desdeFecha) > 0 {
			res[0] = valor.mostrarDatos()
			return res, nil
		}
	}
	frase := "No hay vuelo registrado desde " + ciudades[0] + " hacia " + ciudades[1] + " desde " + parametros[3]
	res[0] = frase
	return res, nil
}

func (a *aeropuerto) BorrarVuelos(parametros []string) ([]string, error) {
	err := validarBorrarVuelos(parametros)
	if err != nil {
		return nil, err
	}
	desde := parametros[1]
	hasta := parametros[2]
	abb := a.fechasVuelos
	res := make([]string, 0)
	fechasBorrar := make([]string, 0)
	//explicacion de uso VUELO_X en verTablero
	desde += " " + VUELO_X
	hasta += " " + VUELO_X
	for iter := abb.IteradorRango(&desde, &hasta); iter.HaySiguiente(); iter.Siguiente() {
		fecha, vuelo := iter.VerActual()
		ciudades := vuelo.obtenerCiudades()
		a.dataVuelos.Borrar(vuelo.obtenerCodigo())
		abb := a.ciudadesConFechas.Obtener(ciudades)
		fechaCiudades := strings.Split(fecha, " ")[0]
		if abb.Pertenece(fechaCiudades) {
			abb.Borrar(fechaCiudades)
		}
		fechasBorrar = append(fechasBorrar, fecha)
		res = append(res, vuelo.mostrarDatos())
	}
	for _, clave := range fechasBorrar {
		abb.Borrar(clave)
	}
	return res, nil
}

func (a *aeropuerto) AgregarVuelos(parametros []string) error {
	path, err := validarAgregarVuelos(parametros)
	if err != nil {
		return errores.ErrorAgregarArchivo{}
	}
	informacion, err := leer_archivo(path)
	if err != nil {
		return errores.ErrorAgregarArchivo{}
	}
	for _, data := range informacion {
		datos := strings.Split(data, ",")
		vuelo := crearVuelo(datos)
		fecha := vuelo.obtenerFecha()
		codigo := vuelo.obtenerCodigo()
		ciudades := vuelo.obtenerCiudades()
		claveAbbVuelo := fecha + " " + codigo

		if a.dataVuelos.Pertenece(codigo) {
			vueloAnterior := a.dataVuelos.Obtener(codigo)
			fechaAnterior := vueloAnterior.obtenerFecha()
			claveAnterior := fechaAnterior + " " + codigo
			ciudadesAnterior := vueloAnterior.obtenerCiudades()

			a.fechasVuelos.Borrar(claveAnterior)
			abbCiudadesAnterior := a.ciudadesConFechas.Obtener(ciudadesAnterior)
			abbCiudadesAnterior.Borrar(fechaAnterior)
		}

		a.dataVuelos.Guardar(codigo, vuelo)
		a.fechasVuelos.Guardar(claveAbbVuelo, vuelo)
		if !a.ciudadesConFechas.Pertenece(ciudades) {
			a.ciudadesConFechas.Guardar(ciudades, TDADict.CrearABB[string, *vuelo_datos](CmpFechas))
		}
		abb := a.ciudadesConFechas.Obtener(ciudades)
		abb.Guardar(fecha, vuelo)

	}
	return nil
}

func leer_archivo(path string) ([]string, error) {
	archivo, _ := os.Open(path)
	var err error
	res := []string{}
	dict := TDADict.CrearHash[string, string]()
	defer archivo.Close()
	s := bufio.NewScanner(archivo)
	for s.Scan() {
		linea := s.Text()
		datos := strings.Split(linea, ",")
		if len(datos) != 10 {
			return nil, errores.ErrorAgregarArchivo{}
		}
		dict.Guardar(datos[0], linea)
	}
	err = s.Err()
	if err != nil {
		return nil, errores.ErrorAgregarArchivo{}
	}
	for iter := dict.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		_, v := iter.VerActual()
		res = append(res, v)
	}
	return res, nil
}
