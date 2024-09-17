// Crea un vector que contenga 100 notas de 100 supuestos estudiantes,
// con valores entre 0 y 10 generadas aleatoriamente..

// Luego, de acuerdo a las notas almacenadas en el arreglo,
// el programa debe indicar cuántos estudiantes son:

// Deficientes: 0-3
// Regulares: 4-6
// Buenos: 7-8
// Excelentes: 9-10

Algoritmo ContadorCalificacionEstudiantes
	Definir LIMITE_INFERIOR, LIMITE_SUPERIOR, TAMANO_VECTOR Como Entero
	Definir indice, vectorNotas Como Entero
	Definir contadorDeficientes, contadorRegulares, contadorBuenos, contadorExcelentes Como Entero
	
	TAMANO_VECTOR = 100
	LIMITE_INFERIOR = 0
	LIMITE_SUPERIOR = 10
	
	Dimension vectorNotas[TAMANO_VECTOR]
	
	llenarVectorAleatoriamente(vectorNotas, TAMANO_VECTOR, LIMITE_INFERIOR, LIMITE_SUPERIOR)
	
	Escribir "Vector generado:"
	mostrarVector(vectorNotas, TAMANO_VECTOR)
	Escribir ""
	
	contadorDeficientes = 0
	contadorRegulares = 0
	contadorBuenos = 0
	contadorExcelentes = 0
	
	Para indice = 0 Hasta TAMANO_VECTOR - 1 Con Paso 1 Hacer
		Si 0 <= vectorNotas[indice] Y vectorNotas[indice] <= 3 Entonces
			contadorDeficientes = contadorDeficientes + 1
		FinSi
		
		Si 4 <= vectorNotas[indice] Y vectorNotas[indice] <= 6 Entonces
			contadorRegulares = contadorRegulares + 1
		FinSi
		
		Si 7 <= vectorNotas[indice] Y vectorNotas[indice] <= 8 Entonces
			contadorBuenos = contadorBuenos + 1
		FinSi
		
		Si 9 <= vectorNotas[indice] Y vectorNotas[indice] <= 10 Entonces
			contadorExcelentes = contadorExcelentes + 1
		FinSi
	FinPara
	
	Escribir "La cantidad de estudiantes con notas deficientes son: ", contadorDeficientes
	Escribir "La cantidad de estudiantes con notas regulaes son: ", contadorRegulares
	Escribir "La cantidad de estudiantes con notas buenos son: ", contadorBuenos
	Escribir "La cantidad de estudiantes con notas excelentes son: ", contadorExcelentes
	
FinAlgoritmo

SubProceso llenarVectorAleatoriamente(vector, tamano, limiteInferior, limiteSuperior)
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
