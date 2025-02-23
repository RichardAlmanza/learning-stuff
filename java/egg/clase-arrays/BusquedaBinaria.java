import java.util.Arrays;
import java.util.Scanner;

public class BusquedaBinaria {
    public static void main(String[] args) {
        int userNumber;
        int foundIndex;
        int[] numbers = new int[50];

        for (int i = 0; i < numbers.length; i++) {
            numbers[i] = (int) (Math.random() * 100 + 1);
        }

        Arrays.sort(numbers);
        System.out.println(Arrays.toString(numbers));

        try (Scanner sc = new Scanner(System.in)) {
            System.out.println("Ingrese a buscar en el array");
            userNumber = sc.nextInt();

            foundIndex = Arrays.binarySearch(numbers, userNumber);

            if (foundIndex < 0) {
                System.out.println("El numero no se encuentra en el array");
            } else {
                System.out.println("EL indice donde se encuentra el numero en el arreglo es " + foundIndex);
            }

        } catch (Exception e) {
            System.out.println("Ingreso incorrectamente el numero");
        }
    }
}
