// Crea un programa que solicite al usuario ingresar 5 valores.
// Los valores ingresados deben ser almacenados en un �nico arreglo.
// y debe mostrarlos posteriormente por pantalla. Sigue estos pasos:
//
// Declara el tipo de dato que almacenar� el vector.
//
// Define la dimensi�n del arreglo.
//
// Utiliza un bucle para recorrer el arreglo reci�n creado,
// posici�n por posici�n, y solicitar al usuario que introduzca un dato.
// Puedes emplear una estructura de bucle "Para" para esta tarea. 
//
// Puntos importantes:
//
// - Definici�n del vector: Se define un vector llamado vectorNumeros de
//   tipo entero con una dimensi�n de 5 elementos. Esto significa que el 
//   vector puede almacenar 5 n�meros enteros.
//
// - Bucle de entrada de datos: Se utiliza un bucle "Para" que va desde 0
//   hasta 4 para recorrer cada posici�n del vector. En cada iteraci�n del bucle, 
//   se solicita al usuario que ingrese un n�mero para esa posici�n del vector.
//
// - �ndices del vector: Los �ndices del vector (i) se utilizan para
//   acceder a cada posici�n del vector, y se manejan de 0 a 4 debido
//   a que los arreglos en muchos lenguajes de programaci�n comienzan
//   en la posici�n 0.

Algoritmo CrearVector
	Definir TAMANO Como Entero
	Definir vectorNumeros Como Entero
	
	TAMANO = 5
	
	Dimension vectorNumeros[TAMANO]
	
	llenarVectorManualmente(vectorNumeros, TAMANO, "Ingrese el numero para el elemento del vector.")
	
	mostrarVector(vectorNumeros, TAMANO)
	Escribir ""
FinAlgoritmo

SubProceso llenarVectorManualmente(vector, tamano, mensaje)
	Definir indice Como Entero
	
	Para indice = 0 Hasta tamano - 1 Con Paso 1 Hacer
		Escribir mensaje
		Escribir "Posicion ", indice + 1, " de ", tamano
		Leer vector[indice]
	FinPara
FinSubProceso

SubProceso mostrarVector(vector, tamano)
	Definir indice, ultimoIndice Como Entero
	
	ultimoIndice = tamano - 1
	
	Escribir Sin Saltar "["
	
	Para indice = 0 Hasta ultimoIndice Con Paso 1 Hacer
		Escribir Sin Saltar vector[indice]
		
		Si indice == ultimoIndice Entonces
			Escribir Sin Saltar "]"
		SiNo
			Escribir Sin Saltar ", "
		FinSi
	FinPara	
FinSubProceso
