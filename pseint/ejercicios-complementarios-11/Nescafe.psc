// Una distribuidora de Nescafé cuenta con 4 representantes que recorren
// toda Argentina para ofrecer sus productos.
// Para la gestión administrativa, el país está dividido en cinco zonas:
// Norte, Sur, Este, Oeste y Centro.
// Mensualmente, la empresa registra los datos de ventas de los representantes
// en cada zona y recopila diversas estadísticas sobre su desempeño.

// Se requiere un programa que lea el monto de las ventas de los representantes
// en cada zona y realice los siguientes cálculos:

// - Total de ventas de una zona especificada por el usuario.

// - Total de ventas de un representante seleccionado por el
//   usuario en cada una de las zonas.

// - Total de ventas de todos los representantes.

Algoritmo Nescafe
	Definir ZONAS, REPRESENTANTES, MESES Como Entero
	Definir LIMITE_INFERIOR, LIMITE_SUPERIOR Como Entero
	Definir matrizVentas Como Entero
	
	ZONAS = 5
	REPRESENTANTES = 4
	MESES = 12
	
	LIMITE_INFERIOR = 0
	LIMITE_SUPERIOR = 17
	
	Dimension matrizVentas[ZONAS, REPRESENTANTES, MESES]
	
	llenarMatriz3DAleatoriamenteEnteros(matrizVentas, ZONAS, REPRESENTANTES, MESES, LIMITE_INFERIOR, LIMITE_SUPERIOR)
	
	menu(matrizVentas, ZONAS, REPRESENTANTES, MESES)
FinAlgoritmo

SubProceso menu(matrizVentas, ZONAS, REPRESENTANTES, MESES)
	Definir OPCION_ZONAS, OPCION_REPRESENTANTES, OPCION_TOTAL Como Entero
	Definir opcionUsuario Como Entero
	Definir continuar Como Logico
	
	OPCION_ZONAS = 1
	OPCION_REPRESENTANTES = 2
	OPCION_TOTAL = 3
	
	continuar = Verdadero
	
	Mientras continuar Hacer
		Escribir "Ingrese la opcion del menu correspondiente"
		Escribir OPCION_ZONAS, ". Total de ventas de una zona especificada."
		Escribir OPCION_REPRESENTANTES, ". Total de ventas de un representante."
		Escribir OPCION_TOTAL, ". Total de ventas de todos los representantes."
		Escribir "9. Salir."
		Leer opcionUsuario
		
		Segun opcionUsuario
			OPCION_ZONAS:
				menuZonas(matrizVentas, ZONAS, REPRESENTANTES, MESES)
			OPCION_REPRESENTANTES:
				menuRepresentante(matrizVentas, ZONAS, REPRESENTANTES, MESES)
			OPCION_TOTAL:
				Escribir sumaRangoMatriz3D(matrizVentas, 0, ZONAS - 1, 0, REPRESENTANTES - 1, 0, MESES - 1)
			9:
				continuar = Falso
			De Otro Modo:
				Escribir "Opciona no valida."
		FinSegun
	FinMientras
FinSubProceso
	
SubProceso menuZonas(matrizVentas, ZONAS, REPRESENTANTES, MESES)
	Definir ZONA_NORTE, ZONA_SUR, ZONA_ESTE, ZONA_OESTE, ZONA_CENTRO Como Entero
	Definir nombreZona Como Caracter
	Definir opcionUsuario Como Entero
	Definir continuar Como Logico
	
	ZONA_NORTE = 1
	ZONA_SUR = 2
	ZONA_ESTE = 3
	ZONA_OESTE = 4
	ZONA_CENTRO = 5
	
	continuar = Verdadero
	nombreZona = ""
	
	Mientras continuar Hacer
		Escribir "Ingrese la opcion de la zona que desea conocer las ventas"
		Escribir ZONA_NORTE, ". Zona Norte."
		Escribir ZONA_SUR, ". Zona Sur."
		Escribir ZONA_ESTE, ". Zona Este."
		Escribir ZONA_OESTE, ". Zona Oeste."
		Escribir ZONA_CENTRO, ". Zona Centro."
		Escribir "9. Salir."
		Leer opcionUsuario
		
		Segun opcionUsuario
			ZONA_NORTE:
				nombreZona = "Zona Norte"
			ZONA_SUR:
				nombreZona = "Zona Sur"
			ZONA_ESTE:
				nombreZona = "Zona Este"
			ZONA_OESTE:
				nombreZona = "Zona Oeste"
			ZONA_CENTRO:
				nombreZona = "Zona Centro"
			9:
				continuar = Falso
			De Otro Modo:
				Escribir "Opciona no valida."
		FinSegun
		
		Si nombreZona <> "" Entonces
			Escribir "El total de ventas de la ", nombreZona, " es ", sumaRangoMatriz3D(matrizVentas, opcionUsuario - 1, opcionUsuario - 1, 0, REPRESENTANTES - 1, 0, MESES - 1)
			nombreZona = ""
		FinSi
	FinMientras
FinSubProceso

SubProceso menuRepresentante(matrizVentas, ZONAS, REPRESENTANTES, MESES)
	Definir opcionUsuario Como Entero
	
	Repetir
		Escribir "Ingrese el ID de uno de los representantes [1, ", REPRESENTANTES, "]"
		Leer opcionUsuario
	Hasta Que 1 <= opcionUsuario Y opcionUsuario <= REPRESENTANTES
	
	Escribir "El total de ventas del representante ", opcionUsuario, " es ", sumaRangoMatriz3D(matrizVentas, 0, ZONAS - 1, opcionUsuario - 1, opcionUsuario - 1, 0, MESES - 1)
FinSubProceso

Funcion sumaTotal <- sumaRangoMatriz3D(matriz, filaInicial, filaFinal, columnaInicial, columnaFinal, fondoInicial, fondoFinal)
	Definir columna, fila, fondo Como Entero
	Definir sumaTotal Como Real
	
	sumaTotal = 0
	
	Para fila = filaInicial Hasta filaFinal Con Paso 1 Hacer		
		Para columna = columnaInicial Hasta columnaFinal Con Paso 1 Hacer
			Para fondo = fondoInicial Hasta fondoFinal Con Paso 1 Hacer
				sumaTotal = sumaTotal + matriz[fila, columna, fondo]
			FinPara
		FinPara
	FinPara
FinFuncion

SubProceso llenarMatriz3DAleatoriamenteEnteros(matriz, tamanoFila, tamanoColumna, tamanoFondo, limiteInferior, limiteSuperior)
	Definir columna, fila, fondo Como Entero
	
	Para fila = 0 Hasta tamanoFila - 1 Con Paso 1 Hacer		
		Para columna = 0 Hasta tamanoColumna - 1 Con Paso 1 Hacer
			Para fondo = 0 Hasta tamanoFondo - 1 Con Paso 1 Hacer
				matriz[fila, columna, fondo] = Aleatorio(limiteInferior, limiteSuperior)
			FinPara
		FinPara
	FinPara
FinSubProceso
