// Desarrolla un programa que, dado un a�o, determine si es bisiesto o no.
// Un a�o es bisiesto si cumple las siguientes condiciones: debe ser divisible por 4 pero no por 100,
// a menos que tambi�n sea divisible por 400. Utiliza la funci�n mod de PseInt para esta tarea.

Algoritmo AnioBisiesto
	Definir anio Como Entero
	Definir esBisiesto Como Logico
	
	esBisiesto = Falso
	
	Escribir "Ingrese el a�o que desea saber si es bisiesto"
	Leer anio
	
	// esBisiesto = (anio mod 400 == 0) O (anio mod 4 == 0 Y anio mod 100 <> 0)
	
	Si anio mod 400 == 0 Entonces
		esBisiesto = Verdadero
	SiNo		
		Si anio mod 4 == 0 Y anio mod 100 <> 0 Entonces
			esBisiesto = Verdadero
		FinSi
	FinSi
	
	Si esBisiesto Entonces
		Escribir "El a�o ", anio, " es Bisiesto"
	SiNo
		Escribir "El a�o ", anio, " NO es Bisiesto"
	FinSi
FinAlgoritmo
