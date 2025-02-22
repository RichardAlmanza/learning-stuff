public class ContandoPares {
    public static void main(String[] args) {
        final int SIZE = 10;
        final int MINIMUM_NUMBER = 0;
        final int MAXIMUM_NUMBER = 99;

        int evenCounter = 0;
        int[] numbers = new int[SIZE];
        StringBuilder strBuilder = new StringBuilder();

        for (int i = 0; i < SIZE; i++) {
            numbers[i] = (int) (Math.random() * (MAXIMUM_NUMBER - MINIMUM_NUMBER + 1) + MINIMUM_NUMBER);
        }

        for (int number : numbers) {
            if (number % 2 == 0) {
                evenCounter++;
            }
        }

        strBuilder.append('[');

        for (int i = 0; i < SIZE; i++) {
            strBuilder.append(numbers[i]);
            strBuilder.append(", ");
        }

        if (SIZE > 0) {
            strBuilder.setLength(strBuilder.length() - 2);
        }
        strBuilder.append(']');

        System.out.println(strBuilder.toString());
        System.out.println("Y la cantidad de numero pares en el lista es de " + evenCounter);
    }
}
