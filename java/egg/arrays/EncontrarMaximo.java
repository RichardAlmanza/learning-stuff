import java.util.InputMismatchException;
import java.util.Scanner;

public class EncontrarMaximo {

    public static void main(String[] args) {
        int maxNumber;
        int[] numbers = new int[5];

        try (Scanner sc = new Scanner(System.in)) {
            System.out.println("Ingrese el primer elemento del array de entero");
            numbers[0] = sc.nextInt();

            System.out.println("Ingrese el segundo elemento del array de entero");
            numbers[1] = sc.nextInt();

            System.out.println("Ingrese el tercero elemento del array de entero");
            numbers[2] = sc.nextInt();

            System.out.println("Ingrese el cuarto elemento del array de entero");
            numbers[3] = sc.nextInt();

            System.out.println("Ingrese el quinto elemento del array de entero");
            numbers[4] = sc.nextInt();

            maxNumber = numbers[0];

            if (numbers[1] > maxNumber) {
                maxNumber = numbers[1];
            }

            if (numbers[2] > maxNumber) {
                maxNumber = numbers[2];
            }

            if (numbers[3] > maxNumber) {
                maxNumber = numbers[3];
            }

            if (numbers[4] > maxNumber) {
                maxNumber = numbers[4];
            }

            System.out.println("El numero maximo es " + maxNumber);

        } catch (InputMismatchException e) {
            System.out.println("Ingreso incorrectamente lo que debe ser un numero entero");
        }
    }
}
