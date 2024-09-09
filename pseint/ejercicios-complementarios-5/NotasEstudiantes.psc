// Desarrolla un programa que solicite al usuario el número 
// de estudiantes en un curso y para cada estudiante, 
// pida su nombre, edad y tres calificaciones. 
// Luego, calcula el promedio de calificaciones de cada 
// estudiante y muestra un mensaje indicando si aprobaron 
// o reprobaron el curso. 
// Emplea estructuras anidadas para manejar los datos de múltiples estudiantes.

Algoritmo NotasEstudiantes
	Definir NUMERO_NOTAS, NOTA_APROBATORIA Como Entero
	
	Definir numeroEstudiantes, iteradorNotas, iteradorEstudiantes, edadEstudiante Como Entero
	Definir promedio, sumatoria, nota Como Real
	Definir nombreEstudiante Como Caracter
	
	NUMERO_NOTAS = 3
	NOTA_APROBATORIA = 6
	
	Escribir "Ingrese el numero de Estudiantes"
	Leer numeroEstudiantes
	
	Para iteradorEstudiantes = 1 Hasta numeroEstudiantes Hacer
		Escribir "Ingrese el nombre del estudiante"
		Leer nombreEstudiante
		
		Escribir "Ingrese la edad de ", nombreEstudiante
		Leer edadEstudiante
		
		sumatoria = 0
		Para iteradorNotas = 1 Hasta NUMERO_NOTAS Hacer
			Escribir "Cual es la nota " iteradorNotas
			Leer nota
			
			sumatoria = sumatoria + nota
		FinPara
		
		promedio = sumatoria / NUMERO_NOTAS
		
		Escribir Sin Saltar "El estudiante ", nombreEstudiante
		Escribir " con edad ", edadEstudiante
		Escribir Sin Saltar "Tiene un promedio de ", promedio
		
		Si promedio >= NOTA_APROBATORIA Entonces
			Escribir " y aprobo el curso"
		SiNo
			Escribir " y reprobo el curso"
		FinSi
	FinPara
FinAlgoritmo
