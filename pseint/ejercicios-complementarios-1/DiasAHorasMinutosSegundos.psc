// A partir de una conocida cantidad de días que el usuario ingresa a través del teclado,
// escribe un programa para convertir los días en horas, en minutos y en segundos.
// Por ejemplo: 1 día = 24 horas = 1440 minutos = 86400 segundos.

Algoritmo DiasAHorasMinutosSegundos
	Definir dias, horas, minutos, segundos_ Como Real
	
	Escribir "Ingrese el numero de dias que desea ver representado en horas, minutos o dias"
	Leer dias
	
	horas = 24 * dias
	minutos = 60 * horas
	segundos_ = 60 * minutos
	
	Escribir "Las equivalencias son las siguientes"
	Escribir dias, " Dias = ", horas, " Horas"
	Escribir dias, " Dias = ", minutos, " Minutos"
	Escribir dias, " Dias = ", segundos_, " Segundos"
FinAlgoritmo
