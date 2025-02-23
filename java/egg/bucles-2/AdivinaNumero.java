import java.util.Scanner;

public class AdivinaNumero {
    public static void main(String[] args) {
        final int randomNumber = (int) (Math.random() * (21 - 1) + 1);

        int userGuess;

        try (Scanner sc = new Scanner(System.in)) {
            do {
                System.out.println("Ingrese un numero entero del 1 al 20");
                userGuess = sc.nextInt();
            } while (userGuess != randomNumber);

            System.out.println("Adivino el numero!");
        } catch (Exception e) {
            System.out.println("Ingreso incorrectamente el numero");
        }
    }
}
