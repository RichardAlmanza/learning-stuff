// En esta actividad, el objetivo es escribir una función
// recursiva que calcule la suma de los primeros N enteros.
// El programa principal solicitará al usuario un número,
// y la función recursiva se encargará de calcular la suma
// hasta ese número de manera recursiva.
// Por ejemplo, si el usuario ingresa 5,
// el programa calculará la suma de 1 + 2 + 3 + 4 + 5
// utilizando recursividad.

Algoritmo SumaNPrimerosNumeros
	Definir numeroUsuario Como Entero
	
	Repetir
		Escribir "Ingrese el valor de N, para calcular la suma de los primeros N naturales"
		Leer numeroUsuario
	Hasta Que numeroUsuario > 0 
	
	Escribir "La suma es " SumaNPrimerosNumerosRecursivo(numeroUsuario)
FinAlgoritmo

Funcion resultado <- SumaNPrimerosNumerosRecursivo (numeroFinal)
	Definir resultado Como Entero
	
	Si numeroFinal == 1 Entonces
		resultado = 1
	SiNo
		resultado = numeroFinal + SumaNPrimerosNumerosRecursivo(numeroFinal - 1)
	FinSi
FinFuncion
	