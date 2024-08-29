// Escribe un programa que valide si una nota está entre 0 y 10.
// Si la nota no está dentro de este rango, se solicitará al usuario
// que ingrese la nota nuevamente hasta que sea correcta.
// Recuerda nombrar y guardar tu algoritmo.

Algoritmo ValidarNumeroDentroRango
	Definir LIMITE_INFERIOR, LIMITE_SUPERIOR Como Entero
	Definir nota Como Real
	
	LIMITE_INFERIOR = 0
	LIMITE_SUPERIOR = 10
	
	nota = -1
	
	Mientras nota < LIMITE_INFERIOR O nota > LIMITE_SUPERIOR Hacer
		Escribir "Ingrese la nota dentro del rango [ ", LIMITE_INFERIOR, ", ", LIMITE_SUPERIOR, " ]"
		Leer nota
	FinMientras
FinAlgoritmo
