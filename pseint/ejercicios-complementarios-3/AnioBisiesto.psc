// Desarrolla un programa que, dado un año, determine si es bisiesto o no.
// Un año es bisiesto si cumple las siguientes condiciones: debe ser divisible por 4 pero no por 100,
// a menos que también sea divisible por 400. Utiliza la función mod de PseInt para esta tarea.

Algoritmo AnioBisiesto
	Definir anio Como Entero
	Definir esBisiesto Como Logico
	
	esBisiesto = Falso
	
	Escribir "Ingrese el año que desea saber si es bisiesto"
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
		Escribir "El año ", anio, " es Bisiesto"
	SiNo
		Escribir "El año ", anio, " NO es Bisiesto"
	FinSi
FinAlgoritmo
