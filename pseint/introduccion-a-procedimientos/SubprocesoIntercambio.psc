// Realiza un procedimiento que permita intercambiar el 
// valor de dos variables de tipo entero.
// La variable A debe terminar con el valor de la
// variable B, y viceversa.
// Este cambio debe ser de forma permanente,
// es decir, los valores deben ser sobre escritos.
// Recuerda nombrar y guardar tu algoritmo.

// Nota: Ten presente el paso por referencia, 
// el cual te permite modificar los valores de 
// variables declaradas en el algoritmo principal

Algoritmo SubprocesoIntercambio
	Definir numero1, numero2 Como Entero
	
	Escribir "Ingrese el primer numero"
	Leer numero1
	
	Escribir "Ingrese el segundo numero"
	Leer numero2
	
	intercambiar(numero1, numero2)
	
	Escribir "Ahora el primer numero es ", numero1
	Escribir "y el segundo numero es ", numero2
FinAlgoritmo

SubProceso intercambiar (numeroA Por Referencia, numeroB Por Referencia)
	Definir auxiliar Como Entero
	
	auxiliar = numeroA
	numeroA = numeroB
	numeroB = auxiliar
FinSubProceso

