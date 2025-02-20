import java.util.InputMismatchException;
import java.util.Scanner;

public class ExcepcionTipo {

    public static void main(String[] args) {
        int a;
        int b;

        try (Scanner sc = new Scanner(System.in);) {
            System.out.println("Ingrese el numero entero que sera restado");
            a = sc.nextInt();

            System.out.println("Ingrese el numero entero que sera el que resta");
            b = sc.nextInt();

            System.out.println(String.format("La resta de %d - %d es %d", a, b, a - b));
        } catch (InputMismatchException e) {
            System.out.println("No ingreso un numero entero");
            System.out.println(e.getMessage());
        }

    }
}
