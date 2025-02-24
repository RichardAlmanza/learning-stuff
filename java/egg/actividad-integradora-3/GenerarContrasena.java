import java.util.InputMismatchException;
import java.util.Scanner;

public class GenerarContrasena {
    public static void main(String[] args) {
        final int MINIMUM_PASSWORD_SIZE = 8;
        final int LOWER_CASE = 0;
        final int UPPER_CASE = 1;
        final int DIGITS = 2;
        final int SYMBOLS = 3;

        int randomSelector;
        int passwordSize = 0;
        int[] counters = {
                0, // Lower case letters
                0, // Upper case letters
                0, // Digits
                0, // Symbols
        };

        StringBuilder sb = new StringBuilder();

        try (Scanner sc = new Scanner(System.in)) {
            do {
                System.out.println("Ingrese el tamano de la contrasena, debe ser minimo 8");
                passwordSize = sc.nextInt();
            } while (passwordSize < MINIMUM_PASSWORD_SIZE);
        } catch (InputMismatchException e) {
            System.out.println("Ingreso incorrectamente el tamano de la cotrasena");
            return;
        }

        while (sb.length() < passwordSize) {
            randomSelector = (int) (Math.random() * 4);

            if (passwordSize - sb.length() < 5) {
                for (int i = 0; i < counters.length; i++) {
                    if (counters[i] == 0) {
                        randomSelector = i;
                        break;
                    }
                }
            }

            switch (randomSelector) {
                case LOWER_CASE -> sb.append((char) (Math.random() * ('z' - 'a' + 1) + 'a'));
                case UPPER_CASE -> sb.append((char) (Math.random() * ('Z' - 'A' + 1) + 'A'));
                case DIGITS -> sb.append((char) (Math.random() * ('9' - '0' + 1) + '0'));
                case SYMBOLS -> sb.append((char) (Math.random() * ('.' - '!' + 1) + '!'));
            }

            counters[randomSelector]++;
        }

        System.out.println(String.format("Contrasena generada: `%s`", sb.toString()));
    }
}
