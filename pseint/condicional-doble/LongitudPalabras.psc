// Desarrolla un programa que solicite al usuario ingresar un nombre para su competencia, el cual debe constar de una frase o palabra de exactamente 6 caracteres.
// Si el usuario ingresa una frase o palabra de 6 caracteres, el programa imprimirá por pantalla el mensaje "LONGITUD  CORRECTA".
// En caso contrario, se imprimirá "LONGITUD INCORRECTA". Se sugiere investigar la función Longitud() de PSeInt para realizar esta verificación de manera eficiente.
// Recuerda nombrar y guardar tu algoritmo.

Algoritmo LongitudPalabras
	Definir frase Como Caracter
	
	Escribir "Ingrese la frase o palabra"
	Leer frase
	
	Si Longitud(frase) == 6 Entonces
		Escribir "LONGITUD CORRECTA"
	SiNo
		Escribir "LONGITUD INCORRECTA"
	FinSi
	
FinAlgoritmo
