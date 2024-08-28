// Solicita al usuario que ingrese un valor del 1 al 7.
// El programa mostrará un mensaje que indique a qué día de la semana corresponde dicho valor.
// Por ejemplo, el número 1 corresponde al "Lunes", el número 2 al "Martes" y así sucesivamente. Recuerda nombrar y guardar tu algoritmo.

Algoritmo DiaSemana
	Definir dia Como Entero
	Definir diaTexto Como Caracter
	
	Escribir "Ingrese el dia de la semana en numero"
	Leer dia
	
	Segun dia Hacer
		1:
			diaTexto = "Lunes"
		2:
			diaTexto = "Martes"
		3:
			diaTexto = "Miercoles"
		4:
			diaTexto = "Jueves"
		5:
			diaTexto = "Viernes"
		6:
			diaTexto = "Sabado"
		7:
			diaTexto = "Domingo"
		De Otro Modo:
			diaTexto = "No es dia de la semana"
	FinSegun
	
	Escribir diaTexto
FinAlgoritmo
