// Escribe un programa que solicite al usuario ingresar tres notas y determine si un alumno aprueba o reprueba un curso.
// El criterio para aprobar es que el promedio de las tres notas sea igual o mayor a 70.
// Si el promedio cumple con este criterio, el programa imprimirá "El alumno aprueba el curso"; de lo contrario, imprimirá "El alumno reprueba el curso".

Algoritmo AprobarCurso
	Definir nota1, nota2, nota3, promedio Como Real
	
	Escribir "Ingrese la nota 1"
	Leer  nota1
	
	Escribir "Ingrese la nota 2"
	Leer  nota2
	
	Escribir "Ingrese la nota 3"
	Leer  nota3
	
	promedio = (nota1 + nota2 + nota3) / 3
	
	Si promedio >= 70 Entonces
		Escribir "El alumno aprueba el curso"
	SiNo
		Escribir "El alumno reprueba el curso"
	FinSi
FinAlgoritmo
