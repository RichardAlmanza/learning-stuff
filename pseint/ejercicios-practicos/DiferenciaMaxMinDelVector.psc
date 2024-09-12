// Desarrolla un programa que inicialice un arreglo de números,
// permitiéndote elegir el tipo y tamaño del arreglo.
// Puedes generar datos aleatorios para poblar el arreglo o
// asignar valores manualmente. Posteriormente, crea una función
// que calcule y devuelva la diferencia entre el valor más pequeño
// y el valor más grande de este arreglo. Para cumplir con las necesidades
// de la actividad, se sugiere dividir el proceso en subprocesos o funciones,
// esto te permitirá tener un código más modular y fácil de entender,
// cumpliendo con las necesidades de la actividad y facilitando futuras
// modificaciones o expansiones del programa

Algoritmo DiferenciaMaxMinDelVector
	Definir TIPO_NUMERO_REAL, TIPO_NUMERO_ENTERO Como Entero
	
	Definir indice, tamanoVector, tipoVector Como Entero
	Definir vectorNumerosReales, diferencia Como Real
	Definir vectorNumerosEntero Como Entero
	
	TIPO_NUMERO_REAL = 1
	TIPO_NUMERO_ENTERO = 2
	
	Repetir
		Escribir "Ingrese el tamano del vector"
		Leer tamanoVector
	Hasta Que tamanoVector > 0
	
	Dimension vectorNumerosEntero[tamanoVector], vectorNumerosReales[tamanoVector]
	
	Repetir
		Escribir "Ingrese el numero correspondiente a la opcion que desea realizar"
		Escribir TIPO_NUMERO_REAL, ". Crear vector tipo Real"
		Escribir TIPO_NUMERO_ENTERO, ". Crear vector tipo Entero"
		Leer tipoVector
	Hasta Que tipoVector == TIPO_NUMERO_REAL O tipoVector == TIPO_NUMERO_ENTERO
	
	Segun tipoVector
		TIPO_NUMERO_REAL:
			menuLLenado(vectorNumerosReales, tamanoVector, TIPO_NUMERO_REAL)
			diferencia = obtenerDiferenciaMaxMin(vectorNumerosReales, tamanoVector)
		TIPO_NUMERO_ENTERO:
			menuLLenado(vectorNumerosEntero, tamanoVector, TIPO_NUMERO_ENTERO)
			diferencia = obtenerDiferenciaMaxMin(vectorNumerosEntero, tamanoVector)
	FinSegun
	
	Escribir "La diferencia es ", diferencia
	
FinAlgoritmo

SubProceso menuLLenado(vector, tamanoVector, tipoVector)
	Definir LIMITE_INFERIOR, LIMITE_SUPERIOR Como Entero
	Definir TIPO_NUMERO_REAL, TIPO_NUMERO_ENTERO Como Entero
	Definir LLENADO_MANUAL, LLENADO_ALEATORIO Como Entero
	Definir FACTOR_DECIMAL Como Real
	Definir opcionUsuario Como Entero
	
	TIPO_NUMERO_REAL = 1
	TIPO_NUMERO_ENTERO = 2
	LLENADO_MANUAL = 1
	LLENADO_ALEATORIO = 2
	
	LIMITE_INFERIOR = -500
	LIMITE_SUPERIOR = 500
	FACTOR_DECIMAL = 0.001
	
	Repetir
		Escribir "Ingrese el numero correspondiente a la opcion que desea realizar"
		Escribir LLENADO_MANUAL, ". Llenar manualmente"
		Escribir LLENADO_ALEATORIO, ". Llenar automaticamente"
		Leer opcionUsuario
	Hasta Que opcionUsuario == LLENADO_MANUAL O opcionUsuario == LLENADO_ALEATORIO
	
	Segun opcionUsuario
		LLENADO_MANUAL:
			llenarVectorManualmente(vector, tamanoVector, "Ingrese el numero para el elemento del vector")
		LLENADO_ALEATORIO:
			Segun tipoVector
				TIPO_NUMERO_REAL:
					llenarVectorAleatoriamenteReales(vector, tamanoVector, LIMITE_INFERIOR, LIMITE_SUPERIOR, FACTOR_DECIMAL)
				TIPO_NUMERO_ENTERO:
					llenarVectorAleatoriamenteEnteros(vector, tamanoVector, LIMITE_INFERIOR, LIMITE_SUPERIOR)
			FinSegun
	FinSegun
	
	mostrarVector(vector, tamanoVector)
	Escribir ""
FinSubProceso

Funcion diferencia <- obtenerDiferenciaMaxMin(vector, tamano)
	Definir diferencia Como Real
	
	diferencia = obtenerMaximo(vector, tamano) - obtenerMinimo(vector, tamano)
FinFuncion

Funcion minimo <- obtenerMinimo(vector, tamano)
	Definir indice Como Entero
	Definir minimo Como Real
	
	minimo = vector[0]
	
	Para indice = 0 Hasta tamano - 1 Con Paso 1 Hacer
		Si vector[indice] < minimo Entonces
			minimo = vector[indice]
		FinSi
	FinPara
FinFuncion

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

SubProceso llenarVectorAleatoriamenteReales(vector, tamano, limiteInferior, limiteSuperior, factorDecimal)
	Definir indice Como Entero
	
	Para indice = 0 Hasta tamano - 1 Con Paso 1 Hacer
		vector[indice] = Aleatorio(limiteInferior, limiteSuperior) + Aleatorio(limiteInferior, limiteSuperior) * factorDecimal
	FinPara
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
