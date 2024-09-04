Algoritmo VarianteReciclajeDesafioGrupal
	Definir USUARIO, CONTRASENA Como Caracter
	Definir MAXIMOS_INTENTOS Como Entero
	Definir PESO_MINIMO_BOTELLA, PESO_MAXIMO_BOTELLA Como Entero
	Definir PESO_BOTELLA_CATEGORIA_1, PESO_BOTELLA_CATEGORIA_2 Como Entero
	Definir VALOR_BOTELLA_CATEGORIA_1, VALOR_BOTELLA_CATEGORIA_2, VALOR_BOTELLA_CATEGORIA_3 Como Real
	
	
	Definir saldo Como Real
	Definir estaAutenticado Como Logico
	
	USUARIO = "Albus_D"
	CONTRASENA = "caramelosDeLimon"
	MAXIMOS_INTENTOS = 3
	
	PESO_MINIMO_BOTELLA = 100
	PESO_MAXIMO_BOTELLA = 3000
	PESO_BOTELLA_CATEGORIA_1 = 500
	PESO_BOTELLA_CATEGORIA_2 = 1500
	VALOR_BOTELLA_CATEGORIA_1 = 50.0
	VALOR_BOTELLA_CATEGORIA_2 = 125.0
	VALOR_BOTELLA_CATEGORIA_3 = 200.0
	
	saldo = 0
	estaAutenticado = Falso
	
	// Manejo de credenciales
	Definir contadorIntentos Como Entero
	Definir usuarioIngresado, contrasenaIngresada Como Caracter
	
	contadorIntentos = 0
	
	Mientras contadorIntentos < MAXIMOS_INTENTOS Y NO estaAutenticado Hacer
		contadorIntentos = contadorIntentos + 1
		
		Escribir "Ingrese el usuario"
		Leer usuarioIngresado
		
		Escribir "Ingrese la contrasena"
		Leer contrasenaIngresada
		
		estaAutenticado = USUARIO == usuarioIngresado Y CONTRASENA == contrasenaIngresada
	FinMientras
	// Fin Manejo de credenciales
	
	// Menu
	Definir opcionIngresada Como Entero
	
	Definir numeroBotellas, iteracionBotellas, pesoBotella Como Entero
	Definir sumatoriaValor, valorBotella Como Real
	Definir confirmacionUsuario Como Caracter
	
	Mientras estaAutenticado Hacer
		Escribir "Ingrese el numero correspondiente de la opcion de menu que desea utilizar"
		Escribir "1. Ingresar Botellas"
		Escribir "2. Consultar Saldo"
		Escribir "3. Salir"
		Leer opcionIngresada
		
		Segun opcionIngresada
			1:
				// Opcion ingresar botellas
				Repetir
					Escribir "Ingrese el numero de botellas a ingresar"
					Leer numeroBotellas
				Hasta Que numeroBotellas > 0
				
				sumatoriaValor = 0
				
				Para iteracionBotellas <- 1 Hasta numeroBotellas Hacer
					pesoBotella = Aleatorio(PESO_MINIMO_BOTELLA, PESO_MAXIMO_BOTELLA)
					
					Si pesoBotella <= PESO_BOTELLA_CATEGORIA_1 Entonces
						valorBotella = VALOR_BOTELLA_CATEGORIA_1
					SiNo
						Si pesoBotella <= PESO_BOTELLA_CATEGORIA_2 Entonces
							valorBotella = VALOR_BOTELLA_CATEGORIA_2
						SiNo
							valorBotella = VALOR_BOTELLA_CATEGORIA_3
						FinSi
					FinSi
					
					sumatoriaValor = sumatoriaValor + valorBotella
				FinPara
				
				Escribir "El total por las botellas ingresada es ", sumatoriaValor
				
				Repetir
					Escribir "Quiere confirmar la transaccion? S/N"
					Leer confirmacionUsuario
					
					confirmacionUsuario = Mayusculas(confirmacionUsuario)
				Hasta Que confirmacionUsuario == "S" O confirmacionUsuario == "N"
				
				Si confirmacionUsuario == "S" Entonces
					saldo = saldo + sumatoriaValor
					
					Escribir "Su nuevo saldo es ", saldo
				SiNo
					Escribir "Devolviendo material"
				FinSi
			2:
				// Opcion Consultar Saldo
				Escribir "Su saldo actual es ", saldo
			3:
				// Opcion Salir (Desconectar usuario)
				estaAutenticado = Falso
			De Otro Modo:
				Escribir "Ingreso una opcion invalida, vuelva a intentarlo"
		FinSegun
	FinMientras
	// Fin menu
FinAlgoritmo
