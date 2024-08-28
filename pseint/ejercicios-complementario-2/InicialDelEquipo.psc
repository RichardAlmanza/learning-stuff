// Escribe un programa que solicite el nombre de un equipo de fútbol y valide si la primera letra de ese nombre es una 'A'.
// Si la primera letra es una 'A', se imprimirá un mensaje que diga "CORRECTO"; de lo contrario, se imprimirá "INCORRECTO".
// Se sugiere investigar la función Subcadena de PSeInt para realizar esta tarea de manera eficiente.

Algoritmo InicialDelEquipo
	Definir nombreEquipo Como Caracter
	
	Escribir "Ingrese el nombre del equipo de futbol"
	Leer nombreEquipo
	
	nombreEquipo = Mayusculas(nombreEquipo)
	
	Si Subcadena(nombreEquipo, 0, 0) == "A" Entonces
		Escribir "CORRECTO"
	SiNo
		Escribir "INCORRECTO"
	FinSi
FinAlgoritmo
