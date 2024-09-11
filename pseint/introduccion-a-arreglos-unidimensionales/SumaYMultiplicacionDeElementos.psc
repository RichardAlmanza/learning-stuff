// Realiza un programa que lea 10 números reales por teclado,
// los almacene en un arreglo y muestre por pantalla la suma
// y multiplicación de todos los números ingresados al arreglo.

Algoritmo SumaYMultiplicacionDeElementos
	Definir TAMANO_VECTOR Como Entero
	Definir vectorNumeros Como Real
	
	TAMANO_VECTOR = 10
	
	Dimension vectorNumeros[TAMANO_VECTOR]
	
	llenarVectorManualmente(vectorNumeros, TAMANO_VECTOR, "Ingrese el numero real para el elemento del vector.")
	
	mostrarVector(vectorNumeros, TAMANO_VECTOR)
	Escribir ""
	Escribir "La suma de los elementos del vector es ", sumaElementosVectorNumeros(vectorNumeros, TAMANO_VECTOR)
	Escribir "La multiplicacion de los elementos del vector es ", multiplicacionElementosVectorNumeros(vectorNumeros, TAMANO_VECTOR)
FinAlgoritmo

Funcion total <- multiplicacionElementosVectorNumeros(vector, tamanoVector)
	Definir indice Como Entero
	Definir total Como Real
	
	total = 1
	
	Para indice = 0 Hasta tamanoVector - 1 Con Paso 1 Hacer
		total = total * vector[indice]
	FinPara
FinFuncion

Funcion total <- sumaElementosVectorNumeros(vector, tamanoVector)
	Definir indice Como Entero
	Definir total Como Real
	
	total = 0
	
	Para indice = 0 Hasta tamanoVector - 1 Con Paso 1 Hacer
		total = total + vector[indice]
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
		
		Si indice == ultimoIndice Entonces
			Escribir Sin Saltar "]"
		SiNo
			Escribir Sin Saltar ", "
		FinSi
	FinPara	
FinSubProceso
