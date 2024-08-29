// Escribe un programa para obtener el grado de eficiencia de un operario de una f�brica de tornillos,
// de acuerdo a las siguientes dos condiciones que se le imponen para un per�odo de prueba:

// - Producir menos de 200 tornillos defectuosos.

// - Producir m�s de 10000 tornillos sin defectos.

// El grado de eficiencia se determina de la siguiente manera:
	
// - GRADO 5: Si no cumple ninguna de las condiciones.
// - GRADO 6: Si s�lo cumple la primera condici�n
// - GRADO 7: Si s�lo cumple la segunda condici�n.
// - GRADO 8: Si cumple las dos condiciones.

// Muestra un mensaje acorde al grado de eficiencia de un operario, luego de evaluar las condiciones. 
// Recuerda nombrar y guardar tu algoritmo.

// Nota: Para abordar este ejercicio de manera ordenada y eficiente, se recomienda probar cada inciso
// por separado en lugar de intentar resolver todos los puntos simult�neamente.
// Esto permite identificar y corregir posibles errores de forma m�s sencilla, ya que abordar todas las 
// tareas a la vez puede dificultar la identificaci�n de problemas espec�ficos

Algoritmo PeriodoPrueba
	Definir tornillosDefectuosos, tornillosPerfectos, banderas, grado Como Entero
	
	banderas = 0
	grado = 5
	
	// Lo siguiente no es recomendado por complicar la lectura, pero estaba aburrido
	
	// Recuerden Facil de entender/leer es mejor que `tener menos lineas`
	
	Escribir "Ingrese la cantidad de tornillos defectuosos"
	Leer tornillosDefectuosos
	
	Escribir "Ingrese la cantidad de tornillos perfectos"
	Leer tornillosPerfectos
	
	Si tornillosDefectuosos < 200 Entonces
		banderas = banderas + 1
	FinSi
	
	Si tornillosPerfectos > 1000 Entonces
		banderas = banderas + 2
	FinSi
	
	grado = grado + banderas
	
	Escribir "El grado es ", grado
FinAlgoritmo
