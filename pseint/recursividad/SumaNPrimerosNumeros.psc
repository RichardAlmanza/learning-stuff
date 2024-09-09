// En esta actividad, el objetivo es escribir una funci�n
// recursiva que calcule la suma de los primeros N enteros.
// El programa principal solicitar� al usuario un n�mero,
// y la funci�n recursiva se encargar� de calcular la suma
// hasta ese n�mero de manera recursiva.
// Por ejemplo, si el usuario ingresa 5,
// el programa calcular� la suma de 1 + 2 + 3 + 4 + 5
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
	