// Diseña un programa que trabaje con un vector
// de datos lógicos de tamaño 5.
// El programa debe realizar las siguientes tareas:

// - Llenar el vector con valores lógicos 
//   (verdadero o falso) ingresados por el usuario.

// - Mostrar en pantalla el vector original.

// - Determinar y mostrar un mensaje en pantalla
//   indicando si todos los elementos del vector
//   son verdaderos.

// - Determinar y mostrar un mensaje en pantalla
//   indicando si al menos uno de los elementos
//   del vector es verdadero, junto con la posición
//   en la que se encuentra.

// - Determinar y mostrar un mensaje en pantalla
//   indicando si todos los elementos del vector
//   son falsos.

// - Determinar y mostrar un mensaje en pantalla
//   indicando si al menos uno de los elementos
//   del vector es falso, junto con la posición
//   en la que se encuentra.

Algoritmo VectorLogico
	Definir TAMANO_VECTOR Como Entero
	Definir vectorNumeros, indice Como Entero
	Definir vectorLogicos Como Logico
	
	TAMANO_VECTOR = 5
	
	Dimension vectorLogicos[TAMANO_VECTOR], vectorNumeros[TAMANO_VECTOR]
	
	llenarVectorManualmente(vectorNumeros, TAMANO_VECTOR, "Ingrese 0 para un valor logico Falso, cualquier otro numero se tomara como Verdadero")
	
	Para indice = 0 Hasta TAMANO_VECTOR - 1 Con Paso 1 Hacer
		vectorLogicos[indice] = vectorNumeros[indice] <> 0
	FinPara
	
	mostrarVector(vectorLogicos, TAMANO_VECTOR)
	Escribir ""
	
	determinarEstadoVector(Verdadero, vectorLogicos, TAMANO_VECTOR)
	determinarEstadoVector(Falso, vectorLogicos, TAMANO_VECTOR)
FinAlgoritmo

SubProceso determinarEstadoVector(estado, vector, tamanoVector)
	Definir vectorPosiciones, tamanoPosiciones Como Entero
	
	Dimension vectorPosiciones[tamanoVector]
	
	Si buscarEnVector(estado, vector, vectorPosiciones, tamanoVector, tamanoPosiciones) Entonces
		Si tamanoPosiciones == tamanoVector Entonces
			Escribir "Todos los elementos del vector son ", estado
		SiNo
			Escribir "El valor de ", estado, " se en cuentra en las siguientes posiciones"
			mostrarVector(vectorPosiciones, tamanoPosiciones)
			Escribir ""
		FinSi
	FinSi
FinSubProceso

Funcion encontrado <- buscarEnVector(elementoBuscado, vectorABuscar, vectorPosiciones, tamanoVector, tamanoPosiciones Por Referencia)
	Definir indice Como Entero
	Definir encontrado Como Logico
	
	encontrado = Falso
	tamanoPosiciones = 0
	
	Para indice = 0 Hasta tamanoVector - 1 Con Paso 1 Hacer
		Si vectorABuscar[indice] == elementoBuscado Entonces
			vectorPosiciones[tamanoPosiciones] = indice
			tamanoPosiciones = tamanoPosiciones + 1
		FinSi
	FinPara
	
	Si tamanoPosiciones > 0 Entonces
		encontrado = Verdadero
	FinSi
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
