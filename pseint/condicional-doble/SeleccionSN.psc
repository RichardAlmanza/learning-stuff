// Realiza un programa que solicite al usuario un caracter.
// Si el caracter ingresado es S? o ?N?, se deberá de imprimir un mensaje por pantalla que diga "CORRECTO",
// en caso contrario, se deberá imprimir "INCORRECTO". Recuerda nombrar y guardar tu algoritmo

Algoritmo SeleccionSN
	Definir caracter_ Como Caracter
	
	Escribir "Ingrese un caracter"
	Leer caracter_
	
	Si caracter_ == "S" O caracter_ == "N" Entonces
		Escribir "CORRECTO"
	SiNo
		Escribir "INCORRECTO"
	FinSi
FinAlgoritmo
