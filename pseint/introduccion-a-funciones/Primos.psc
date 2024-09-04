// Crea una función que determine si un número ingresado por el usuario es primo.
// Un número es primo si solo es divisible por 1
// y por sí mismo (por ejemplo: 2, 3, 5, 7, 11, 13, 17, etc.).
// Recuerda nombrar y guardar tu algoritmo.

Algoritmo Primos
	Definir iter, numeroDesde, numeroHasta Como Entero
	
	Escribir "Para generar un lista de numeros primos debe ingresar un numero de inicio y otro final"
	
	Repetir
		Escribir "Ingrese el numero inicial mayor que 1"
		Leer numeroDesde
	Hasta Que numeroDesde > 1
	
	Repetir
		Escribir "Ingrese el numero final mayor o igual que ", numeroDesde
		Leer numeroHasta
	Hasta Que numeroHasta >= numeroDesde
	
	imprimirNumerosPrimos(numeroDesde, numeroHasta)
FinAlgoritmo

Funcion _esPrimo <- esPrimo(numero_)
	Definir _esPrimo, esDivisible Como Logico
	Definir divisor, maximoDivisor Como Entero
	
	maximoDivisor = trunc(raiz(numero_))
	
	esDivisible = Falso
	
	Si numero_ mod 2 == 0 Y numero_ <> 2 Entonces
		esDivisible = Verdadero
	SiNo
		Para divisor <- 3 Hasta maximoDivisor Con Paso 2 Hacer
			Si NO esDivisible Entonces
				esDivisible = numero_ mod divisor == 0
			FinSi
		FinPara
	FinSi
	
	_esPrimo = NO esDivisible
FinFuncion

/// numeroInicial debe ser mayor que 1
/// numeroFinal debe ser mayor o igual que numeroInicial
Funcion imprimirNumerosPrimos(numeroInicial, numeroFinal)
	Definir iter, contador Como Entero
	
	contador =0
	
	Escribir "Lista de numeros desde ", numeroInicial " primos hasta ", numeroFinal
	
	Para iter <- numeroInicial Hasta numeroFinal Con Paso 1 Hacer
		Si esPrimo(iter) Entonces
			Escribir iter
			contador = contador + 1
		FinSi
	FinPara
	
	Escribir "Con un total de ", contador, " numeros primos"
FinFuncion
	