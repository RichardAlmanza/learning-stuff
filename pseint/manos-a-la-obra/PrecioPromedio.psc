// Escribe un programa que calcule el precio promedio de un producto.
// El precio promedio se debe calcular a partir del precio del mismo producto en tres establecimientos distintos.

// Nota: Aseg�rate de solicitar al usuario que ingrese tres valores del producto,
// los cuales ser�n almacenados en tres variables previamente definidas.

// Posteriormente, podr�s llevar a cabo la operaci�n correspondiente.
// Recuerda nombrar y guardar tu algoritmo.

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
