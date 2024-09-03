// Crea un programa que solicite al usuario ingresar una frase
// y luego la muestre en pantalla con un espacio entre cada letra.
// Es importante, almacenar esta nueva palabra con espacios
// en una variable destinada a dicho fin. Por ejemplo:

Algoritmo EspaciarLetras
	Definir indiceCaracter Como Entero
	Definir frase, fraseEspaciada, letraFrase Como Caracter
	
	fraseEspaciada = ""
	
	Escribir "Ingrese la frase que desea espaciar"
	Leer frase
	
	Para indiceCaracter = 0 Hasta Longitud(frase) - 1 Hacer
		fraseEspaciada = Concatenar(fraseEspaciada, Subcadena(frase, indiceCaracter, indiceCaracter))
		fraseEspaciada = Concatenar(fraseEspaciada, " ")
	FinPara
	
	Escribir fraseEspaciada
FinAlgoritmo
