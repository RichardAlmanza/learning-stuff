// Desarrolla un programa que rellene un vector de tamaño N,
// con valores numéricos ingresados por el usuario.
// A continuación, se deberá crear una función que reciba
// el vector y devuelva el valor más grande del arreglo.

Algoritmo MaximoNumeroDelVector
	Definir tamanoVector Como Entero
	Definir vectorNumeros Como Real
	
	Repetir
		Escribir "Ingrese el tamano del vector de numeros"
		Leer tamanoVector
	Hasta Que tamanoVector > 0
	
	Dimension vectorNumeros[tamanoVector]
	
	llenarVectorManualmente(vectorNumeros, tamanoVector, "Ingrese el numero real para el elemento del vector.")
	mostrarVector(vectorNumeros, tamanoVector)
	Escribir ""
	
	Escribir "El numero maximo del vector es ", obtenerMaximo(vectorNumeros, tamanoVector)
FinAlgoritmo

Funcion maximo <- obtenerMaximo(vector, tamano)
	Definir indice Como Entero
	Definir maximo Como Real
	
	maximo = vector[0]
	
	Para indice = 0 Hasta tamano - 1 Con Paso 1 Hacer
		Si vector[indice] > maximo Entonces
			maximo = vector[indice]
		FinSi
	FinPara
FinFuncion

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
		
		Si indice < ultimoIndice Entonces
			Escribir Sin Saltar ", "
		FinSi
	FinPara
	
	Escribir Sin Saltar "]"
FinSubProceso
