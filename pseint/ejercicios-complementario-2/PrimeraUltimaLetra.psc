// Continuando con el ejercicio anterior, ahora se solicitar� el nombre del equipo y se validar� si la primera letra del nombre es igual a la �ltima letra del nombre.
// Si ambas letras son iguales, se imprimir� un mensaje que diga "CORRECTO"; de lo contrario, se imprimir� "INCORRECTO".

Algoritmo PrimeraUltimaLetra
	Definir nombreEquipo Como Caracter
	Definir final Como Entero
	
	Escribir "Ingrese el nombre del equipo"
	Leer nombreEquipo
	
	nombreEquipo = Minusculas(nombreEquipo)
	final = Longitud(nombreEquipo) - 1
	
	Si Subcadena(nombreEquipo, 0, 0) == Subcadena(nombreEquipo, final, final) Entonces
		Escribir "CORRECTO"
	SiNo
		Escribir "INCORRECTO"
	FinSi
FinAlgoritmo
