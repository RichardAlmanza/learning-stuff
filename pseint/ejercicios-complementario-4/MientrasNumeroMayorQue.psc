// Escribe un programa que solicite al usuario ingresar un número.
// Mientras ese número sea mayor que 10, se le pedirá al 
// usuario que ingrese el número nuevamente.
// Utiliza la estructura repetitiva mientras para resolver esta actividad.

Algoritmo MientrasNumeroMayorQue
	Definir NUMERO_CONDICION Como Real
	Definir nuevoNumero Como Real
	
	NUMERO_CONDICION = 10
	
	Repetir		
		Escribir "Ingresar un numero entero positivo para sumar"
		Leer nuevoNumero
	Mientras Que nuevoNumero > NUMERO_CONDICION
	
FinAlgoritmo
