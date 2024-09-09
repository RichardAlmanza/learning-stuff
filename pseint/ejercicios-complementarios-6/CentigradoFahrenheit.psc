// Crea una función que convierta una temperatura de 
// grados Celsius a grados Fahrenheit.
// Luego, en el programa principal,
// solicita al usuario una temperatura en Celsius y
// utiliza la función para convertirla y mostrar el resultado.
// Fórmula de conversión: F = C * 9/5 + 32

Algoritmo CentigradoFahrenheit
	Definir iter Como Entero
	Definir centigrados Como Real
	
	Para iter <- 1 Hasta 5 Hacer
		Escribir "Ingrese los grados centigrados que desee convertir a fahrenheit"
		Leer centigrados
		
		Escribir "Son ", centigradoAFahrenheit(centigrados), " Fahrenheit"
	FinPara
FinAlgoritmo

Funcion fahrenheit <- centigradoAFahrenheit (centigrados)
	Definir fahrenheit Como Real
	
	fahrenheit = centigrados * 9/5 + 32
FinFuncion
	