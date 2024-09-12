// Desarrolla un programa que ordene un vector lleno
// de números enteros aleatorios de menor a mayor.
// La dimensión del vector debe ser solicitada al usuario.
// Puedes investigar el método de ordenamiento burbuja
// para implementarlo.
// Para obtener más información sobre el ordenamiento burbuja,
// puedes consultar el siguiente enlace: Ordenamiento Burbuja.

Algoritmo OrdenamientoBurbuja
	Definir LIMITE_INFERIOR, LIMITE_SUPERIOR Como Entero
	Definir vectorNumeros, tamanoVector, indice Como Entero
	
	LIMITE_INFERIOR = -999
	LIMITE_SUPERIOR = 999
	
	Repetir
		Escribir "Ingrese el tamano del vector"
		Leer tamanoVector
	Hasta Que tamanoVector > 0
	
	Dimension vectorNumeros[tamanoVector]
	
	llenarVectorAleatoriamenteEnteros(vectorNumeros, tamanoVector, LIMITE_INFERIOR, LIMITE_SUPERIOR)
	
	Escribir "El vector sin ordenar es:"
	mostrarVector(vectorNumeros, tamanoVector)
	Escribir ""
	
	ordenarVectorPorBurbuja(vectorNumeros, tamanoVector)
	
	Escribir "El vector ordenado es:"
	mostrarVector(vectorNumeros, tamanoVector)
	Escribir ""
FinAlgoritmo

SubProceso ordenarVectorPorBurbuja(vector, tamano)
	Definir indice, indiceComprobacion Como Entero
	
	Para indice = 0 Hasta tamano - 1 Con Paso 1 Hacer
		Para indiceComprobacion = 1 Hasta tamano - 1 Con Paso 1 Hacer
			Si vector[indiceComprobacion - 1] > vector[indiceComprobacion] Entonces
				intercambiarNumeros(vector[indiceComprobacion - 1], vector[indiceComprobacion])
			FinSi
		FinPara
	FinPara
FinSubProceso

SubProceso intercambiarNumeros(numeroA Por Referencia, numeroB Por Referencia)
	Definir auxiliar Como Entero
	
	auxiliar = numeroA
	numeroA = numeroB
	numeroB = auxiliar	
FinSubProceso

SubProceso llenarVectorAleatoriamenteEnteros(vector, tamano, limiteInferior, limiteSuperior)
	Definir indice Como Entero
	
	Para indice = 0 Hasta tamano - 1 Con Paso 1 Hacer
		vector[indice] = Aleatorio(limiteInferior, limiteSuperior)
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
