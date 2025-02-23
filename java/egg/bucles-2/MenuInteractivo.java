import java.util.InputMismatchException;
import java.util.Scanner;

public class MenuInteractivo {
    public static void main(String[] args) {
        final String menuOptions = """
                1. Comprar producto.
                2. Realizar devolución.
                3. Ver mis pedidos.
                4. Preguntas frecuentes.
                5. Salir.
                                """;

        int option = -1;

        try (Scanner sc = new Scanner(System.in)) {
            while (option != 5) {
                System.out.println(menuOptions);
                option = sc.nextInt();

                switch (option) {
                    case 1 -> System.out.println("Compre su producto");
                    case 2 -> System.out.println("Realice su devolucion");
                    case 3 -> System.out.println("Estos son sus pedidos");
                    case 4 -> System.out.println("Preguntas frecuentes");
                    case 5 -> System.out.println("Vuelva relativamente pronto");
                    default -> System.out.println("No es una opcion valdia de la lista");
                }
            }
        } catch (InputMismatchException e) {
            System.out.println("Ingreso incorrectamente la opcion, solo se permiten numeros enteros");
        }

    }
}
