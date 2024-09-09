// Crea un procedimiento que calcule la temperatura media de
// un d�a a partir de la temperatura m�xima y m�nima.
// Luego, desarrolla un programa principal que, utilizando
// este procedimiento, solicite la temperatura m�xima y
// m�nima de n d�as y muestre la media de cada d�a.
// El programa pedir� al usuario el n�mero de d�as a introducir.
// Recuerda nombrar y guardar tu algoritmo.

Algoritmo SubprocesoCalcularMedia
	Definir temperaturaMedia, temperaturaMinima, temperaturaMaxima Como Real
	Definir cantidadDias, dia Como Entero
	
	Repetir
		Escribir "Ingrese la cantidad de dias que va a ingresar"
		Leer cantidadDias
	Hasta Que cantidadDias > 0
	
	Para dia <- 1 Hasta cantidadDias Con Paso 1 Hacer
		Escribir "Ingrese la temperatura minima del dia ", dia
		Leer temperaturaMinima
		
		Escribir "Ingrese la temperatura maxima del dia ", dia
		Leer temperaturaMaxima
		
		calcularMedia(temperaturaMedia, temperaturaMinima, temperaturaMaxima)
		
		Escribir "La temperatura media del dia ", dia," fue de ", temperaturaMedia
	FinPara
FinAlgoritmo

SubProceso calcularMedia (media Por Referencia, minima, maxima)
	media = minima + maxima
	media = media / 2
FinSubProceso
	