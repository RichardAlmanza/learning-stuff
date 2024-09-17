// Crea un programa que solicite al usuario
// ingresar el tamaño deseado para dos vectores.
// El primero almacenará nombres de personas como cadenas,
// mientras que el segundo vector contendrá la longitud
// de cada uno de los nombres.
// Luego, se mostrarán en pantalla los nombres
// junto con su respectiva longitud.

Algoritmo NombresLongitudes
	Definir tamanoVectores, indice, vectorLongitudes Como Entero
	Definir vectorNombres Como Caracter
	
	Repetir
		Escribir "Ingrese el tamano de la lista de nombres a ingresar"
		Leer tamanoVectores
	Hasta Que tamanoVectores > 0
	
	Dimension vectorNombres[tamanoVectores], vectorLongitudes[tamanoVectores]
	
	llenarVectorManualmente(vectorNombres, tamanoVectores, "Ingrese el nombre")
	llenarLongitudes(vectorNombres, vectorLongitudes, tamanoVectores)
	
	Para indice = 0 Hasta tamanoVectores - 1 Con Paso 1 Hacer
		Escribir vectorNombres[indice], ", ", vectorLongitudes[indice]
	FinPara
FinAlgoritmo

SubProceso llenarLongitudes(vectorCadenas, vectorLongitudes, tamanoVectores)
	Definir indice Como Entero
	
	Para indice = 0 Hasta tamanoVectores - 1 Con Paso 1 Hacer
		vectorLongitudes[indice] = Longitud(vectorCadenas[indice])
	FinPara
FinSubProceso

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