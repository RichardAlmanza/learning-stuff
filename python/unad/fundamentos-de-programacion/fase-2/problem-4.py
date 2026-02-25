# Problema 4: Diseñe un programa para calcular el valor final de un producto según su categoría.
# Requisitos:
# - Definir una constante IVA = 0.19.
# - Solicitar el precio base del producto.
#   ✓ Si el precio ≤ 0, mostrar error y finalizar.
# - Solicitar el código de categoría:
#   ✓ Producto básico: no paga IVA.
#   ✓ Producto estándar: paga IVA del 19%.
#   ✓ Producto de lujo: paga IVA del 19% + recargo del 5%.
# - Si el código es inválido, mostrar error y finalizar.
# - Calcular y mostrar el valor final según la categoría.

IVA = 0.19
EXTRA = 0.05

base_price = float(input("Ingrese el valor del producto base: "))
total = base_price

if base_price > 0:
    product_category = int(input("""Ingrese la opcion correspondiente (numero entero) a la categoria del producto:
1. Basico
2. Estandar
3. Lujo
"""))
    
    if 1 <= product_category and product_category <= 3:
        if product_category == 1:
            print("No paga IVA")
        elif product_category == 2:
            print(f"Paga el IVA de {IVA*100: .0f}%")
            total += base_price * IVA
        elif product_category == 3:
            print(f"Paga el IVA de {IVA*100: .0f}% + recargo del {EXTRA*100: .0f}%")
            total += base_price * (IVA + EXTRA)
        else:
            print("Este mensaje nunca debio aparecer!")
        
        print(f"El total a pagar es {total}")
    else:
        print("Categoria invalida")
    
else:
    print("Valor invalido")
