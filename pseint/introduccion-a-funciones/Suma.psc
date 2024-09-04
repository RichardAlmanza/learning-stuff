// Diseña una función calcularSuma que calcule la suma de dos números.
// En el programa principal, solicita al usuario los dos números y pásalos a la función.
// La función debe calcular la suma y devolver el resultado para que 
// se imprima en el programa principal.
// Recuerda nombrar y guardar tu algoritmo.

Algoritmo Suma
	Definir numA, numB Como Entero
	
	Escribir "Ingrese el primer numero entero a sumar"
	Leer numA
	
	Escribir "Ingrese el segundo numero entero a sumar"
	Leer numB
	
	Escribir "El resultado de la suma es ", calcularSuma(numA, numB)
FinAlgoritmo

Funcion resultado <- calcularSuma (a, b)
	Definir resultado Como Entero
	
	resultado = a + b
FinFuncion
