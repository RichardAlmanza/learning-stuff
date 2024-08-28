// Elabora un programa que solicite al usuario ingresar un día de la semana y, tras un análisis,
// determine si es un día de entrenamiento o no (los días de entrenamiento son de lunes a jueves).
// Por ahora, no es necesario considerar validaciones de entrada de datos, como mayúsculas o minúsculas,
// asumiendo que el usuario ingresará el día de la semana en mayúsculas(cada uno de sus caracteres). Recuerda nombrar y guardar tu algoritmo.

// FUNCIONES EN PSEINT
// Las funciones son herramientas proporcionadas por PSeInt que te ayudan a resolver ciertos problemas de manera más eficiente.

// Por ejemplo, si necesitas calcular la raíz cuadrada de un número, PSeInt ofrece una función que, al proporcionarle un número,
// devuelve su raíz cuadrada. Este resultado puede asignarse a una variable o concatenarse con la instrucción "escribir" para
// mostrarlo directamente sin la necesidad de una variable adicional.

// Además, las funciones pueden utilizarse dentro de cualquier expresión o estructura. 
// Cuando evalúas la expresión, la función se reemplaza por su resultado correspondiente.

// Existen dos tipos de funciones en PSeInt: las funciones matemáticas y las funciones de cadenas de texto.
// Las funciones matemáticas reciben un único parámetro numérico y devuelven un único valor numérico.
// Por otro lado, las funciones de cadenas de texto reciben un único parámetro de tipo cadena,
// pudiendo devolver un valor tanto de tipo cadena como numérico, dependiendo de la función utilizada.

Algoritmo DiaSemana
	Definir dia Como Caracter
	
	Escribir "Ingrese el dia de la semana"
	Leer dia
	
	dia = Mayusculas(dia)
	
	Si dia == "LUNES" O dia == "MARTES" O dia == "MIERCOLES" O dia == "JUEVES" Entonces
		Escribir "Es dia de entrenamiento"
	FinSi
FinAlgoritmo
