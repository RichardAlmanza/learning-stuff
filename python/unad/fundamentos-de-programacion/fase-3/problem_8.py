# Problema 8: La familia de Pepito ha celebrado sus cumpleaños desde el
# primer año de vida, y en el año 2026 cumplirá 6 años, por lo que
# nuevamente organizarán una fiesta con familiares y amigos. Los padres
# del niño son muy cuidadosos con sus finanzas y llevan un control
# detallado de sus gastos.
# El año anterior tuvieron un inconveniente con el presupuesto destinado
# a la compra de helados, ya que no lograron llevar un registro exacto.
# Para evitar que esto vuelva a ocurrir, le han solicitado desarrollar un
# programa que permita gestionar la compra y el control del dinero
# disponible.
#
# Requisitos del programa
# 1. Entrada inicial:
# - Solicitar la cantidad de dinero disponible para realizar la
# compra de helados.
#
# 2. Proceso de compra:
# - El usuario podrá comprar la cantidad de helados que desee,
# siempre que el dinero alcance.
# - Cada helado se selecciona según la siguiente Información:
# Tipo      Descripción         Valor
# 1         vaso de Helado      $2.500
# 2         choco cono          $2.000
# 3         paleta Drácula      $3.800
# 4         polet               $4.500
# 5         platillo            $1.800
# - Los valores deben almacenarse en un arreglo unidimensional,
# donde la posición 0 corresponde al primer tipo de helado, la
# posición 1 al segundo, y así sucesivamente:
# Posición  0           1           2           3           4
#           $2.500      $2.000      $3.800      $4.500      $1.800
# - Después de cada compra, el programa debe mostrar el dinero
# restante para que el usuario tenga control del presupuesto.
#
# 3. Condiciones para finalizar el ciclo de compra:
# - Cuando el dinero restante no alcance para comprar ningún tipo
# de helado (el programa debe mostrar un mensaje indicando esta
# situación).
# - O cuando el usuario decida no continuar comprando.
# Al terminar el proceso, el programa debe mostrar:
# • Cantidad total de helados comprados.
# • Cantidad de helados por cada tipo.
# • Valor total gastado.
# • Valor promedio por helado.
# • Dinero sobrante.

# Nombre:           Richard A. Vega Almanza
# Grupo:            213022_154
# Programa:         Ingenieria Electronica
# Codigo Fuente:    Autoria Propia

from problem_3 import ask, try_int, natural_validator


def format_price(value: int|float) -> str:
    """Format number to currency"""
    return f"${value:,.0f}".replace(',','.')


def format_shopping_options(
        name_list: list[str],
        price_list: list[float|int],
        headers: list[str]=None) -> str:
    """Format a list for shopping options based on item names and their prices"""

    if not name_list:
        return ""

    output = []

    max_length_in_names = len(max(name_list))

    if headers:
        padding = (max_length_in_names - len(headers[1])) * " "
        output.append(f"{headers[0]}\t{headers[1]}{padding}\t\t{headers[2]}")

    for i, (name, price) in enumerate(zip(name_list, price_list)):
        padding = (max_length_in_names - len(name)) * " "
        output.append(f"{i + 1}.\t{name}{padding}\t\t{format_price(price)}")

    return "\n".join(output)


def filter_available_shopping_options(
        name_list: list[str],
        price_list: list[float|int],
        budget: int|float):
    """Filter the shopping options based on budget"""

    mapped = []
    new_name_list = []
    new_price_list = []

    for i, (name, price) in enumerate(zip(name_list, price_list)):
        if price <= budget:
            mapped.append(i)
            new_name_list.append(name)
            new_price_list.append(price)

    return [mapped, new_name_list, new_price_list]


def option_validator_creator(min_value: int, max_value: int):
    """
    Return a validator function for the user input,
    to validate if the selected option is inside of the range described
    """

    if max_value < min_value:
        max_value, min_value = min_value, max_value

    def option_validator(user_input):
        option, ok = try_int(user_input)

        if min_value <= option and option <= max_value and ok:
            return [option, True]

        return [0, False]

    return option_validator


def calculate_bill(prices: list[int|float], quantities: list[int]):
    """Sum total cost of the items"""

    total = 0
    for price, quantity in zip(prices, quantities):
        total += price * quantity

    return total


if __name__ == "__main__":
    ORIGINAL_BUDGET = ask(
        "Por favor ingrese el dinero disponible para comprar? ",
        "Ingrese un numero entero positivo, sin centavos, ni simbolos o separadores",
        natural_validator
    )
    budget = ORIGINAL_BUDGET

    fields = ["Tipo", "Descripcion", "Precio"]
    ice_cream_available = ["Vaso de helado", "Chococono", "Paleta Dracula", "Polet", "Platillo"]
    ice_cream_prices = [2500, 2000, 3800, 4500, 1800]
    quantities = [ 0 for _ in ice_cream_available ]

    # Buy process
    while True:
        # Calculate current budget
        budget = ORIGINAL_BUDGET - calculate_bill(ice_cream_prices, quantities)
        print(f"Dinero disponible {format_price(budget)}")

        # Generate item list for items available
        mapped, filtered_ic, filtered_prices = filter_available_shopping_options(
            ice_cream_available, ice_cream_prices, budget)

        if not filtered_ic:
            print("No tiene suficiente dinero disponible para seguir comprando")
            break

        # Ask for the user to select a item to buy
        MAX_OPTION = len(filtered_ic) + 1
        shopping_options = format_shopping_options(filtered_ic, filtered_prices, fields)
        selected_option = ask(
            f"{shopping_options}\n{MAX_OPTION}.\tFinalizar la compra\n",
            "Por favor ingrese el numero relacionado con la opcion.",
            option_validator_creator(1, MAX_OPTION)
        )

        # Breaks the loop based on user selection
        if selected_option == MAX_OPTION:
            break

        # Get original list index
        index = mapped[selected_option - 1]

        # Buy multiple ice creams
        MAX_ITEMS = int(budget / ice_cream_prices[index])
        SELECTED_ICE_CREAM = ice_cream_available[index]
        quantity = ask(
            f"Cuantos {SELECTED_ICE_CREAM} desea comprar? con el dinero actual puede un maximo de {MAX_ITEMS}: ",
            f"Por favor ingrese numeros enteros desde 0 (para cancelar la compra) hasta {MAX_ITEMS}",
            option_validator_creator(0, MAX_ITEMS)
        )

        quantities[index] += quantity

    # Resume
    # • Cantidad total de helados comprados.
    total = sum(quantities)
    print(f"Total de helados comprados: {total}")
    # • Cantidad de helados por cada tipo.
    for ice_cream, quantity in zip(ice_cream_available, quantities):
        if quantity > 0:
            print(f"Se compraron {quantity} helados de {ice_cream}")
    # • Valor total gastado.
    bill = calculate_bill(ice_cream_prices, quantities)
    print(f"Se gasto en total: {format_price(bill)}")
    # • Valor promedio por helado.
    avg = bill / (total if total != 0 else 1)
    print(f"El valor promedio por helado es de {format_price(avg)}")
    # • Dinero sobrante.
    budget = ORIGINAL_BUDGET - bill
    print(f"El dinero sobrante es de {format_price(budget)}")
