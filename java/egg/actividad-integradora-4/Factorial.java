import java.util.InputMismatchException;
import java.util.Scanner;

public class Factorial {
    public static void main(String[] args) {
        int userNumber;
        int number;

        long result = 1;

        try (Scanner sc = new Scanner(System.in)) {
            do {
                System.out.println("Ingrese un numero natural para sacar su factorial");
                userNumber = sc.nextInt();
            } while (userNumber < 0);
        } catch (InputMismatchException e) {
            System.out.println("Ingreso incorrectamente el numero");
            return;
        }

        number = userNumber == 0? 1 : userNumber;
        do {
            result *= number;
            number--;
        } while (number > 1);

        System.out.println(String.format("%d! = %d", userNumber, result));
    }
}
