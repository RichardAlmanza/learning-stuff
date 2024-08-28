// Elabora un programa que solicite al usuario ingresar un d�a de la semana y, tras un an�lisis,
// determine si es un d�a de entrenamiento o no (los d�as de entrenamiento son de lunes a jueves).
// Por ahora, no es necesario considerar validaciones de entrada de datos, como may�sculas o min�sculas,
// asumiendo que el usuario ingresar� el d�a de la semana en may�sculas(cada uno de sus caracteres). Recuerda nombrar y guardar tu algoritmo.

// FUNCIONES EN PSEINT
// Las funciones son herramientas proporcionadas por PSeInt que te ayudan a resolver ciertos problemas de manera m�s eficiente.

// Por ejemplo, si necesitas calcular la ra�z cuadrada de un n�mero, PSeInt ofrece una funci�n que, al proporcionarle un n�mero,
// devuelve su ra�z cuadrada. Este resultado puede asignarse a una variable o concatenarse con la instrucci�n "escribir" para
// mostrarlo directamente sin la necesidad de una variable adicional.

// Adem�s, las funciones pueden utilizarse dentro de cualquier expresi�n o estructura. 
// Cuando eval�as la expresi�n, la funci�n se reemplaza por su resultado correspondiente.

// Existen dos tipos de funciones en PSeInt: las funciones matem�ticas y las funciones de cadenas de texto.
// Las funciones matem�ticas reciben un �nico par�metro num�rico y devuelven un �nico valor num�rico.
// Por otro lado, las funciones de cadenas de texto reciben un �nico par�metro de tipo cadena,
// pudiendo devolver un valor tanto de tipo cadena como num�rico, dependiendo de la funci�n utilizada.

Algoritmo DiaSemana
	Definir dia Como Caracter
	
	Escribir "Ingrese el dia de la semana"
	Leer dia
	
	dia = Mayusculas(dia)
	
	Si dia == "LUNES" O dia == "MARTES" O dia == "MIERCOLES" O dia == "JUEVES" Entonces
		Escribir "Es dia de entrenamiento"
	FinSi
FinAlgoritmo
