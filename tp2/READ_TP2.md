To execute this program:  ./algueiza.go

To execute this program tests (placed at pruebas_algueiza): ./pruebas.sh ../algueiza  

El tp2 consiste en un programa el cual opera la informacion de los vuelos de un aeropuerto el cual recibe un archivo de texto con datos de vuelos, crea una estructura de vuelo y los agrega a una serie de TDAs para que luego se puedan realizar consultas y modificaciones sobre los mismos. Dentro de dicha estructura se contiene toda la informacion del vuelo.

Dichos TDA los guardamos dentro de la estructura de un TDA aeropuerto, que es el que tiene las primitivas para poder realizar las acciones solicitadas por los comandos.

Los TDA dentro del TDA aeropuerto son:
1. Un diccionario con clave el codigo de vuelo y valor la estructura del vuelo, 
2. Un Arbol Binario de Busqueda con clave la fecha (por el cual se ordenaran) y el codigo (para diferenciar aquellos con la misma fecha), y valores la estructura del vuelo.
3. Un diccionario con clave las ciudades de origen y destino de todos los vuelos, y valor un Arbol Binario de Busqueda que tiene como clave la fecha y como valor la estructura del vuelo.

El TDA aeropuerto tiene las siguientes acciones:

-Devolver informacion de un vuelo: dado su codigo, esto se realiza en tiempo constante O(1) utilizando un diccionario que tiene como claves los codigos de los vuelos y como valores la informacion de los mismos.

-Ver el proximo vuelo que conecta dos aeropuertos: Dado los aeropuertos y una fecha a partir de la cual debe hallarse el vuelo que los conecta, se realiza en tiempo log(cantidad de vuelos que conectan los aeropuertos), utilizando un diccionario que tiene como claves los codigos de los aeropuertos y como valores un abb que tiene como claves las fechas de los vuelos y como valores los codigos de los vuelos que conectan los aeropuertos en esa fecha. Es decir, encontrar todos los vuelos que conectan una ciudad con otra es O(1), y luego encontrar el proximo vuelo es en tiempo logaritmico, ya que se utiliza un abb ordenado por fecha para almacenar los vuelos que conectan dos aeropuertos en especifico. A pesar de que a primera vista pareceria mas simple utilizar una lista, ordenando los vuelos por fecha, esto no es posible ya que al agregar un vuelo se debe agregar a la lista ordenada, lo cual es O(n) y no O(log(n)) como en el abb.

-Ver los k vuelos con mayor prioridad: para esto, primero se itera el diccionario de vuelos (siempre se actualiza al agregar vuelos o borrarlos), que tiene complejidad O(n), obteniendo la informacion de los vuelos para despues crear un heap con los mismos (esto tambien es O(n)). Luego se desencolan los K vuelos y se los muestra por pantalla, lo cual es
O(k log n). Por lo tanto la complejidad total es O(n + k log n).

-Agregar vuelos: para esto se itera el archivo csv (siendo V los vuelos a agregar, es con complejidad O(V)) con los datos y se va agregando a las estructuras correspondientes. 
Primero, el guardado en un abb de vuelos ordenados por fecha, teniendo tambien el codigo unico para diferenciar aquellos con la misma fecha y hora. Asi, entonces, cada vuelo se guarda en O(log n) por ser abb, y se pisa el anterior si tienen el mismo codigo. 
Luego, en el diccionario de clave el codigo y valor los datos del vuelo, lo cual es siempre O(1), ya sea que el vuelo no existiera previamente o lo pisara con nueva informacion. 
Finalmente, obteniendo el ABB de las ciudades de origen y destino del diccionario de ciudades y ABB en O(1) (que corresponde al vuelo a agregar), se agrega a ese ABB en O(log F), siendo F la cantidad de fechas diferentes en las que se puede hacer el viaje entre las dos ciudades. Por lo tanto, la complejidad es O(V logF), siendo F bastante menor a n.
Entonces, la complejidad total es O(V log n), siendo V los vuelos a agregar, dado principalmente por el tiempo de agregar al abb.

-Borrar vuelos: Primero, en base a un rango de fechas, se itera el abb en dicho rango, que en un caso promedio es O(log(n)) por cada clave en el rango. 
Con cada clave, primero se busca con el codigo de vuelo en el diccionario de vuelos, y se borra en O(1) K veces por lo que es O(K).
Luego guarda para despues de terminar la iteracion, borrar en el abb esos K vuelos. Este borrado del abb posterior lo hace en O(K log n). 
Finalmente, por cada vuelo, obtiene las ciudades de origen y destino, y busca el abb correspondiente en el diccionario de ABBs. Esto es O(1). Con el ABB correspondiente, borra el vuelo en O(log F), siendo F la cantidad de fechas diferentes en las que se puede hacer el viaje de las ciudades de origen y destino correspondientes. Por lo tanto, la complejidad es O(K logF), siendo F bastante menor a n.
Por lo tanto, la complejidad total es O(K log n), dado principalmente por el tiempo de borrar en el abb.

-Ver tablero: para esto se utiliza el abb, el cual tiene como clave los vuelos ordenados por fecha, y valores los codigos de vuelo. Para esto se utiliza el iterador de rango que en un caso promedio es O(log(n)). En el caso que se quieran ver las fechas de todos los vuelos, se utiliza el iterador de abb, el cual es O(n). Ademas, si se quiere ver la informacion de forma des, se utiliza una pila para dar vuelta las fechas, lo cual es O(k), siendo k la cantidad de vuelos pedidos.

En conclusion, el programa no podria funcionar con las mismas complejidades sin los TDAS ABB, Diccionario y Heap, los cuales se consideran como imprescindibles para una considerable mejora en cuanto a las complejidades de los comandos. 