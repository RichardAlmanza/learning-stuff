# Problema 3: Un nuevo operador de televisión por cable desea ofrecer nuevos servicios en su ciudad.
#
# Para esto, se requiere saber en una muestra de 50 personas:
# - Cuántas horas a la semana invierten en ver televisión.
# - Qué tipo de canal prefieren ver: deportivo, cultural, de noticias o de películas.
# - Cuántas personas están dispuestas a pagar más de 50 mil pesos por el servicio de televisión.
# - La edad de la persona
#
# El programa debe mostrar:
# - El promedio de horas semanales que invierten los encuestados en ver televisión.
# - La cantidad de personas interesadas en cada canal: deportivo, cultural, de noticias o de películas.
# - La cantidad de personas que están dispuestas a pagar más de 50 mil pesos por el servicio.
# - El promedio de edades de los encuestados

# Nombre:           Richard A. Vega Almanza
# Grupo:            213022_154
# Programa:         Ingenieria Electronica
# Codigo Fuente:    Autoria Propia

import random

def stats(responses, functions):
    """Calls the functions and pass the columns as parameters"""

    result = []
    transposed_responses = list(zip(*responses))

    for stat_func, answers in zip(functions, transposed_responses):
        result.append(stat_func(answers))

    return result


def count_elements(values):
    """Count how many times every element is the list"""

    elements = []
    counters = []

    for value in values:
        if value not in elements:
            elements.append(value)
            counters.append(1)
        else:
            index = elements.index(value)
            counters[index] += 1

    return list(zip(elements, counters))


def mean(values):
    """Calculate the arithmetic average of a number list"""

    return sum(values) / len(values)


def questionnaire(questions):
    """
    Ask the questionnaire to the users, where `questions` is defined as follows:
    [
        [question, error_message, validator_function],
        [question, error_message, validator_function],
        ...
    ]
    """

    responses = []

    for question_set in questions:
        responses.append(ask(*question_set))

    return responses


def ask(question, error_message, validator):
    """
    Ask constantly until the user inserts a valid input

    question: string
    validator: function(string) -> [any, bool]
    """

    while True:
        user_input = input(question)

        response, is_valid = validator(user_input)
        if is_valid:
            return response

        print(error_message)


def try_int(user_input):
    """Cast int"""

    try:
        int_number = int(user_input)
        return [int_number, True]
    except ValueError:
        return [0, False]


def natural_validator(user_input):
    """Validate if the user input is a valid natural number"""

    n, _ = try_int(user_input)
    if n > 0:
        return [n, True]

    return [0, False]


def age_validator(user_input):
    """Validate if the user input is a valid age"""

    age, ok = try_int(user_input)
    if not ok:
        return [0, False]

    if 18 <= age and age <= 90:
        return [age, True]

    return [0, False]


def yn_validator(user_input):
    """Validate if the user input is a valid Y/N response"""

    response = user_input.lower()
    is_yes = response == "s"

    if is_yes or response == "n":
        return [is_yes, True]

    return ["", False]


def category_validator(user_input):
    """Validate if the user input is a valid option"""

    try:
        option = int(user_input)
    except ValueError:
        option = 0

    if 1 <= option and option <= 4:
        return [option, True]

    return [0, False]


def hours_validator(user_input):
    """Validate if the user input is valid"""

    hours, ok = try_int(user_input)
    if not ok:
        return [0, False]

    if 0 <= hours and hours <= 168: # 24 * 7 = 168
        return [hours, True]

    return [0, False]


# Simulate Questionnaire

def generate_responses(n=50):
    """Generate random responses for the problem"""

    responses = []

    for i in range(n):
        age = random.randint(18, 90)
        hours = random.randint(0,168)
        category = random.randint(1, 4)
        yn_answer = random.randint(0,1) == 1

        responses.append([age, hours, category, yn_answer])

    return responses


if __name__ == "__main__":
    categories = ["Deportivo", "Cultural", "Noticias", "Peliculas"]
    stats_functions = [mean, mean, count_elements, count_elements]
    question_sets = [
        [
            "Que edad tiene? ",
            (
                "Solo personas con 18 años cumplidos pueden continuar este cuestionario, "
                "ingrese la edad cumplida."
            ),
            age_validator,
        ],
        [
            "Cuantas horas a la semana invierte en ver television? ",
            "Por favor ingrese numeros entero para las horas",
            hours_validator,
        ],
        [
            (
                "Que tipo de canal prefiere ver? \n"
                "1. Deportivo \n"
                "2. Cultural \n"
                "3. Noticias \n"
                "4. Peliculas \n"
            ),
            "Por favor un numero correspondiente a alguna de las opciones.",
            category_validator,
        ],
        [
            "Esta dispuesto a pagar mas de 50 mil pesos por el servicio de television? S/N ",
            "Ingrese solamente la letra S, si esta dispuesta, si no lo esta ingrese la letra N.",
            yn_validator,
        ],
    ]

    simulate = ask(
        "Quiere generar los datos automaticamente? S/N ",
        "Ingrese solamente la letra S, si quiere generar las respuestas automaticamente",
        yn_validator
    )

    if simulate:
        response_list = generate_responses()
    else:
        n_respondents = ask(
            "Cuantas personas se van a encuestar? ",
            "Ingrese un numero natural que indique el numero de personas",
            natural_validator
        )

        response_list = []
        for i in range(n_respondents):
            print(f"Entrevistado {i + 1}")
            response_list.append(questionnaire(question_sets))

    stats_results = stats(response_list, stats_functions)

    # Print responses
    print("Edad\t\tHoras invertidas\t\tCategoria\t\tDisposicion de pagar")
    for response in response_list:
        age = response[0]
        hours = response[1]
        category = categories[response[2] - 1]
        yn = "Si" if response[3] else "No"

        print(f"{age}\t\t{hours}\t\t\t\t{category}\t\t{yn}")

    print()

    # Print stats
    print(f"El promedio de edad de la encuesta es {stats_results[0]} años")
    print(f"El promedio de horas invertida por semana {stats_results[1]} horas")

    for category_result in stats_results[2]:
        category = categories[category_result[0] - 1]
        count = category_result[1]

        print(f"El numero de encuestados que respondieron {category} es de {count}")

    for yn_answer_result in stats_results[3]:
        yn = "Si" if yn_answer_result[0] else "No"
        count = yn_answer_result[1]

        print(f"El numero de encuestados que respondieron {yn} es de {count}")
