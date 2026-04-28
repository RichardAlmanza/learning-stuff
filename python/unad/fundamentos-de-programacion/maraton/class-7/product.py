class Product:
    def __init__(self, uid, name, price):
        self.uid = uid
        self.name = name
        self.__price = price

    def __str__(self):
        return ("{"
        f"\"uid\": \"{self.uid}\", "
        f"\"name\": \"{self.name}\", "
        f"\"price\": {self.price}"
        "}")

    @property
    def price(self):
        return self.__price

    @price.setter
    def price(self, new_price):
        if new_price < 0:
            # raise ValueError("the price cannot be negative")
            print("el precio no puede ser menor a 0")
            return

        self.__price = new_price

class ProductDiscount(Product):
    def __init__(self, uid, name, price, discount):
        super().__init__(uid, name, price)
        self.discount = discount

    def __str__(self):
        return (f"{super().__str__()[:-1]}, "
        f"\"discount\": {self.discount}"
        "}")

    @property
    def price(self):
        base_price = super().price
        return base_price * (1 - self.discount)

product_1 = Product(156,"laptop", 1200)
product_1.price = 3000


product_2 = ProductDiscount(12, "cookies", 3_000, 0.5)
print(product_2)
