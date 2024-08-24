// Escribe un programa que permita al usuario ingresar el valor de dos variables numéricas de tipo entero.
// Posteriormente, el programa debe intercambiar los valores de ambas variables y mostrar el resultado final por pantalla.

// Por ejemplo, si el usuario ingresa los valores num1 = 9 y num2 = 3, la salida a del programa deberá mostrar: num1 = 3 y num2 = 9.

// Ayuda: Para intercambiar los valores de dos variables se debe utilizar una variable auxiliar.

Algoritmo PrecioPromedio
	Definir precio1, precio2, precio3, promedio Como Real
	
	Escribir "Ingrese el valor del precio de la primera tienda"
	Leer precio1
	
	Escribir "Ingrese el valor del precio de la segunda tienda"
	Leer precio2
	
	Escribir "Ingrese el valor del precio de la tercera tienda"
	Leer precio3
	
	promedio = (precio1 + precio2 + precio3) / 3
	
	Escribir "El precio promedio es: ", promedio
FinAlgoritmo
