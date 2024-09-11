// Desarrolla un programa que defina un vector de tama�o n,
// que almacene n�meros, determinando por el usuario el
// tama�o que tendr� dicho arreglo.
// Este tama�o debe ser solicitado al usuario por teclado,
// y almacenado en una variable para dicho fin. 

// Haciendo uso de la variable creada para ese fin,
// dimensionar el arreglo, y llenar cada una de sus
// posiciones con valores aleatorios entre 1 y 25. 

// Luego, se solicitar� al usuario que ingrese un
// n�mero para buscar dentro del arreglo.
// El programa buscar� el elemento dentro del arreglo y
// mostrar� la posici�n donde se encuentra.

// Si el n�mero se encuentra dentro del arreglo, 
// se imprimir�n todas las posiciones donde se 
// encuentra ese valor, en caso de que estuviera repetido.

// Si el n�mero a buscar no est� dentro del arreglo,
// se mostrar� un mensaje indic�ndolo.

Algoritmo BuscarElemento
	Definir LIMITE_INFERIOR, LIMITE_SUPERIOR Como Entero
	Definir tamanoVector, vectorNumeros, tamanoPosiciones, vectorPosiciones Como Entero
	Definir numeroABuscar Como Entero
	Definir seEncontro Como Logico
	
	LIMITE_INFERIOR = 1
	LIMITE_SUPERIOR = 25
	
	Repetir
		Escribir "Ingrese el tamano del vector de numeros"
		Leer tamanoVector
	Hasta Que tamanoVector > 0
	
	Dimension vectorNumeros[tamanoVector], vectorPosiciones[tamanoVector]
	
	llenarVectorAleatoriamente(vectorNumeros, tamanoVector, LIMITE_INFERIOR, LIMITE_SUPERIOR)
	mostrarVector(vectorNumeros, tamanoVector)
	Escribir ""
	
	Escribir "Ingrese el numero entero a buscar"
	Leer numeroABuscar
	
	seEncontro = buscarEnVector(numeroABuscar, vectorNumeros, vectorPosiciones, tamanoVector, tamanoPosiciones)
	
	Si seEncontro Entonces
		Escribir "Se encontro el numero ", numeroABuscar, " en ", tamanoPosiciones, " posicion/es"
	SiNo
		Escribir "No se encontro el numero ", numeroABuscar, " en el vector generado"
	FinSi
	
	Escribir Sin Saltar "Posiciones: "
	mostrarVector(vectorPosiciones, tamanoPosiciones)
	Escribir ""
FinAlgoritmo

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
