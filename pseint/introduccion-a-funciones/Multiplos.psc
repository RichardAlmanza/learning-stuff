// Crea una funci�n esMultiplo que reciba dos n�meros proporcionados por
// el usuario y valide si el primer n�mero es m�ltiplo del segundo.
// La funci�n debe devolver True si el primer n�mero es m�ltiplo del segundo,
// y False en caso contrario.
// Recuerda nombrar y guardar tu algoritmo.

Algoritmo Multiplos
	Definir iter, multiploUsuario, baseUsuario Como Entero
	
	Para iter <- 1 Hasta 5 Hacer
		Escribir "Ingresara los numeros que desea saber si A es multiplo de B"
		Escribir "Ingrese un numero entero A"
		Leer multiploUsuario
		
		Escribir "Ingrese un numero entero B"
		Leer baseUsuario
		
		Escribir "El resultado es ", esMultiplo(multiploUsuario, baseUsuario)
	FinPara
FinAlgoritmo

Funcion resultado <- esMultiplo (mutiplo, base)
	Definir  resultado Como Logico
	
	resultado = mutiplo mod base == 0
FinFuncion
	