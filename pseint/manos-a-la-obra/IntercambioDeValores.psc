// Escribe un programa que calcule el precio promedio de un producto.
// El precio promedio se debe calcular a partir del precio del mismo producto en tres establecimientos distintos.

// Nota: Aseg�rate de solicitar al usuario que ingrese tres valores del producto, 
// los cuales ser�n almacenados en tres variables previamente definidas.

// Posteriormente, podr�s llevar a cabo la operaci�n correspondiente.
// Recuerda nombrar y guardar tu algoritmo.

Algoritmo IntercambioDeValores
	Definir num1, num2, auxiliar Como Entero
	
	Escribir "Ingrese el valor de la variable entera num1"
	Leer num1
	
	Escribir "Ingrese el valor de la variable entera num2"
	Leer num2
	
	auxiliar = num1
	num1 = num2
	num2 = auxiliar
	
	Escribir "El valor de la variable num1 es ", num1
	Escribir "El valor de la variable num2 es ", num2
FinAlgoritmo
