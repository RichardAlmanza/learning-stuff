import java.util.Scanner;

public class DivisionSegura {

    public static void main(String[] args) {
        int a;
        int b;

        try (Scanner sc = new Scanner(System.in);) {
            System.out.println("Ingrese el numero entero que sera el dividendo");
            a = sc.nextInt();

            System.out.println("Ingrese el numero entero que sera el divisor");
            b = sc.nextInt();

            System.out.println(String.format("La division entera de %d / %d es %d", a, b, a / b));
        } catch (ArithmeticException e) {
            System.out.println("No se pueden realizar divisiones por 0");
            System.out.println(e.getMessage());
        }

    }
}
