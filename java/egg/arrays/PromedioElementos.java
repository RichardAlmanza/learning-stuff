import java.util.InputMismatchException;
import java.util.Scanner;

public class PromedioElementos {

    public static void main(String[] args) {
        int sum;
        double average;
        int[] numbers = new int[4];

        try (Scanner sc = new Scanner(System.in)) {
            System.out.println("Ingrese el primer elemento del array de entero");
            numbers[0] = sc.nextInt();

            System.out.println("Ingrese el segundo elemento del array de entero");
            numbers[1] = sc.nextInt();

            System.out.println("Ingrese el tercero elemento del array de entero");
            numbers[2] = sc.nextInt();

            System.out.println("Ingrese el cuarto elemento del array de entero");
            numbers[3] = sc.nextInt();

            sum = numbers[0] + numbers[1] + numbers[2] + numbers[3];
            average = sum / 4.0;

            System.out.println(String.format("""
                La suma de los elemento es %d
                El promedio es %f
                """, sum, average));
        } catch (InputMismatchException e) {
            System.out.println("Ingreso incorrectamente lo que debe ser un numero entero");
        }
    }
}
