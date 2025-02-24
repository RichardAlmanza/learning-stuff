import java.util.Arrays;
import java.util.InputMismatchException;
import java.util.Scanner;

public class Primos {
    public static void main(String[] args) {
        int position;
        int[] primes;

        int number = 3;

        try (Scanner sc = new Scanner(System.in)) {
            do {
                System.out.println("Ingrese la cantidad de los primeros numeros primos que quiera monstrar");
                position = sc.nextInt();
            } while (position < 0);
        } catch (InputMismatchException e) {
            System.out.println("Ingreso incorrectamente el numero");
            return;
        }

        primes = new int[position];

        if (position > 0) {
            primes[0] = 2;
        }

        for (int i = 1; i < primes.length;) {
            boolean isPrime = true;

            for (int j = 0; j < i; j++) {
                if (Math.sqrt(number) < primes[j]) {
                    break;
                }

                if (number % primes[j] == 0) {
                    isPrime = false;
                    break;
                }
            }

            if (isPrime) {
                primes[i] = number;
                i++;
            }

            number += 2;
        }

        System.out.println(Arrays.toString(primes));
    }
}
