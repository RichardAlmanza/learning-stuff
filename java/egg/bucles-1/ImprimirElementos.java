public class ImprimirElementos {

    public static void main(String[] args) {
        final int SIZE = 20;
        final int MINIMUM_NUMBER = 1;
        final int MAXIMUM_NUMBER = 355;

        int[] numbers = new int[SIZE];

        for (int i = 0; i < SIZE; i++) {
            numbers[i] = (int) (Math.random() * (MAXIMUM_NUMBER - MINIMUM_NUMBER + 1) + MINIMUM_NUMBER);
        }

        for (int i = 0; i < SIZE; i++) {
            System.out.println(String.format("[%d] -> %d", i, numbers[i]));
        }
    }
}
