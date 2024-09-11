// Desarrolla un programa con un men� de opciones que permita
// al usuario realizar diversas operaciones con vectores hasta
// que elija la opci�n 'Salir'.

// El men� contempla las siguientes opciones:
// - Llenar Vector A: Este vector, de tama�o N,
//   se llena de manera aleatoria utilizando la
//   funci�n Aleatorio(valorMin, valorMax) de PSeInt.

// - Llenar Vector B: Este vector, de tama�o N,
//   se llena de manera aleatoria utilizando la
//   funci�n Aleatorio(valorMin, valorMax) de PSeInt.

// - Llenar Vector C con la suma de los vectores A y B:
//   La suma se realiza elemento a elemento
//   (es decir, C[i] = A[i] + B[i]).
//   Esto quiere decir, por ejemplo, que la posici�n 1 del
//   vector C ser� el n�mero resultante de la suma de la
//   posici�n 1 del vector A y la posici�n 1 del vector B.

// - Llenar Vector C con la resta de los vectores B y A:
//   La resta se realiza elemento a elemento
//   (es decir, C[i] = B[i] - A[i]).
//   Esto quiere decir, por ejemplo, que la posici�n 1 del
//   vector C ser� el n�mero resultante de la resta de la
//   posici�n 1 del vector A y la posici�n 1 del vector B.

// - Mostrar: Esta opci�n permite al usuario decidir qu�
//   vector desea mostrar: Vector A, B o C.

// - Salir.

// Tener en cuenta:
// - El rango de los n�meros aleatorios para los vectores ser�
//   de -100 a 100.
//   La longitud para todos los vectores ser� la misma,
//   por lo tanto, esa informaci�n solo se solicitar� una vez.

// - Utiliza funciones o subprocesos para mejorar la
//   reutilizaci�n de c�digo.

Algoritmo SumaYRestaDeVectores
	Definir vectorA, vectorB, vectorC, tamanoVectores Como Entero
	
	Repetir
		Escribir "Ingrese el tamano de los vectores"
		Leer tamanoVectores
	Hasta Que tamanoVectores > 0
	
	Dimension vectorA[tamanoVectores], vectorB[tamanoVectores], vectorC[tamanoVectores]
	
	menu(vectorA, vectorB, vectorC, tamanoVectores)
FinAlgoritmo

SubProceso menu(vectorA, vectorB, vectorC, tamanoVectores)
	Definir LIMITE_INFERIOR, LIMITE_SUPERIOR Como Entero
	
	Definir menuActivo, vectorALlenado, vectorBLlenado, vectorCLlenado Como Logico
	Definir opcionUsuario Como Entero
	
	LIMITE_INFERIOR = -100
	LIMITE_SUPERIOR = 100
	
	menuActivo = Verdadero
	vectorALlenado = Falso
	vectorBLlenado = Falso
	vectorCLlenado = Falso
	
	Mientras menuActivo Hacer
		Escribir "Ingrese el numero correspondiente a la opcion que desea realizar"
		Escribir "1. Llenar Vector A"
		Escribir "2. Llenar Vector B"
		Escribir "3. Llenar Vector C con la suma de los vectores A y B"
		Escribir "4. Llenar Vector C con la resta de los vectores A y B"
		Escribir "5. Mostrar"
		Escribir "9. Salir"
		Leer opcionUsuario
		
		Segun opcionUsuario
			1:
				llenarVectorAleatoriamente(vectorA, tamanoVectores, LIMITE_INFERIOR, LIMITE_SUPERIOR)
				vectorALlenado = Verdadero
				Escribir "Vector A ha sido (re)llenado"
			2:
				llenarVectorAleatoriamente(vectorB, tamanoVectores, LIMITE_INFERIOR, LIMITE_SUPERIOR)
				vectorBLlenado = Verdadero
				Escribir "Vector B ha sido (re)llenado"
			3:
				Si vectorALlenado Y vectorBLlenado Entonces
					sumarVectores(vectorA, vectorB, vectorC, tamanoVectores)
					vectorCLlenado = Verdadero
					Escribir "Vector C ha sido (re)llenado con la suma de los vectores A y B"
				SiNo
					Escribir "Debe llenar los vectores A y B antes de poder llenar C"
				FinSi
			4:
				Si vectorALlenado Y vectorBLlenado Entonces
					restarVectores(vectorA, vectorB, vectorC, tamanoVectores)
					vectorCLlenado = Verdadero
					Escribir "Vector C ha sido (re)llenado con la resta de los vectores A y B"
				SiNo
					Escribir "Debe llenar los vectores A y B antes de poder llenar C"
				FinSi
			5:
				Escribir "Ingresando al menu de mostrar vectores"
				Escribir ""
				Esperar Tecla
				menuMostrar(vectorA, vectorB, vectorC, tamanoVectores, vectorALlenado, vectorBLlenado, vectorCLlenado)
				Escribir "Volviendo al menu pricipal"
			9:
				menuActivo = Falso
				Escribir "Saliendo"
			De Otro Modo:
				Escribir "Ingreso un opcion invalida"
		FinSegun
		
		Escribir ""
		Esperar Tecla
	FinMientras
FinSubProceso

SubProceso menuMostrar(vectorA, vectorB, vectorC, tamanoVectores, vectorALlenado, vectorBLlenado, vectorCLlenado)
	Definir menuActivo Como Logico
	Definir opcionUsuario Como Entero
	
	menuActivo = Verdadero
	
	Mientras menuActivo Hacer
		Escribir "Ingrese el numero correspondiente a la opcion que desea realizar"
		Escribir "1. Mostrar Vector A"
		Escribir "2. Mostrar Vector B"
		Escribir "3. Mostrar Vector C"
		Escribir "9. Salir"
		Leer opcionUsuario
		
		Segun opcionUsuario
			1:
				Escribir "El contenido del Vector A es "
				
				Si vectorALlenado Entonces
					mostrarVector(vectorA, tamanoVectores)
				SiNo
					mostrarVector(vectorA, 0)
				FinSi
				
				Escribir ""
			2:
				Escribir "El contenido del Vector B es "
				
				Si vectorBLlenado Entonces
					mostrarVector(vectorB, tamanoVectores)
				SiNo
					mostrarVector(vectorB, 0)
				FinSi
				
				Escribir ""
			3:
				Escribir "El contenido del Vector C es "
				
				Si vectorCLlenado Entonces
					mostrarVector(vectorC, tamanoVectores)
				SiNo
					mostrarVector(vectorC, 0)
				FinSi
				
				Escribir ""
			9:
				menuActivo = Falso
				Escribir "Saliendo del menu mostrar vectores"
			De Otro Modo:
				Escribir "Ingreso un opcion invalida"
		FinSegun
		
		Escribir ""
		Esperar Tecla
	FinMientras
FinSubProceso

SubProceso restarVectores(vectorA, vectorB, vectorC, tamano)
	Definir indice Como Entero
	
	Para indice = 0 Hasta tamano - 1 Con Paso 1 Hacer
		vectorC[indice] = vectorA[indice] - vectorB[indice]
	FinPara
FinSubProceso

SubProceso sumarVectores(vectorA, vectorB, vectorC, tamano)
	Definir indice Como Entero
	
	Para indice = 0 Hasta tamano - 1 Con Paso 1 Hacer
		vectorC[indice] = vectorA[indice] + vectorB[indice]
	FinPara
FinSubProceso

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
