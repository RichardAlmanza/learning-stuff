// Diseña un algoritmo que contenga un condicional anidado. El mismo se ocupará de tomar la orden de bebida para un desayuno. Debe contemplar:

// Si el usuario quiere tomar té o café. 
//   - Si elige café, preguntar si lo quiere solo o cortado.
//       - Si prefiere cortado, preguntar si lo prefiere con leche animal o vegetal.
//   + Si elige té, preguntar si desea rodajas de limón.
// Ten en cuenta las posibles opciones y muestra en pantalla el mensaje final acorde a la selección del usuario. Recuerda nombrar y guardar tu algoritmo.

Algoritmo CafeTe
	Definir opcionNumero Como Entero
	Definir error Como Logico
	Definir errorMensaje, mensajeSeleccion, orden Como Caracter
	
	error = Falso
	errorMensaje = "No selecciono entre las opciones disponible"
	
	mensajeSeleccion = "Seleccione una opcion del menu escribiendo solamente el numero correspondiente del menu"
	orden = ""
	
	Escribir mensajeSeleccion
	Escribir "Que desea?"
	Escribir "1. Cafe"
	Escribir "2. Te"
	Leer opcionNumero
	
	Segun opcionNumero Hacer
		1:
			orden = "Cafe"
			
			Escribir mensajeSeleccion
			Escribir "Como lo desea?"
			Escribir "1. Solo"
			Escribir "2. Cortado"
			Leer opcionNumero
			
			Segun opcionNumero Hacer
				1:
					orden = Concatenar(orden, " solo")
				2:
					orden = Concatenar(orden, " con leche")
					
					Escribir mensajeSeleccion
					Escribir "Con que tipo de leche?"
					Escribir "1. Animal"
					Escribir "2. Vegetal"
					Leer opcionNumero
					
					Segun opcionNumero Hacer
						1:
							orden = Concatenar(orden, " animal")
						2:
							orden = Concatenar(orden, " vegetal")
						De Otro Modo:
							error = Verdadero
					FinSegun
				De Otro Modo:
					error = Verdadero
			FinSegun
		2:
			orden = "Te"
			
			Escribir mensajeSeleccion
			Escribir "Desea rodajas de limon?"
			Escribir "1. Si"
			Escribir "2. No"
			Leer opcionNumero
			
			Segun opcionNumero Hacer
				1:
					orden = Concatenar(orden, " con rodajas de limon")
				2:
				De Otro Modo:
					error = Verdadero
			FinSegun
		De Otro Modo:
			error = Verdadero
	FinSegun
	
	Si error Entonces
		Escribir errorMensaje
	SiNo
		Escribir "Usted desea un ", orden
	FinSi
FinAlgoritmo
