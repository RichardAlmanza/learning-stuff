import java.util.Arrays;

public class ArreglosIguales {
    public static void main(String[] args) {
        boolean areEquals = false;
        int[] numbersA = new int[3];
        int[] numbersB = new int[3];

        while (!areEquals) {
            for (int i = 0; i < numbersA.length; i++) {
                numbersA[i] = (int) (Math.random() * 3 + 1);
                numbersB[i] = (int) (Math.random() * 3 + 1);
            }

            areEquals = Arrays.equals(numbersA, numbersB);

            System.out.println(Arrays.toString(numbersA));
            System.out.println(Arrays.toString(numbersB));

            System.out.println("ArrayA es igual a ArrayB? " + (areEquals ? "Si" : "No"));

        }
    }
}
