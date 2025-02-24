import java.util.InputMismatchException;
import java.util.Scanner;

public class Fibonacci {
    public static void main(String[] args) {
        int position;

        int a = 0;
        int b = 1;

        try (Scanner sc = new Scanner(System.in)) {
            do {
                System.out.println("Ingrese un numero natural correspondiente la posicion en la secuencia fibonacci");
                position = sc.nextInt();
            } while (position < 0);
        } catch (InputMismatchException e) {
            System.out.println("Ingreso incorrectamente el numero");
            return;
        }

        for (int i = 0; i < position; i++) {
            int aux = a;
            a = b;
            b += aux;
        }

        System.out.println(String.format("Finbonacci(%d) = %d", position, a));
    }
}
