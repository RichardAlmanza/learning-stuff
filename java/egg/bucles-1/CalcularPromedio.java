public class CalcularPromedio {

    public static void main(String[] args) {
        final int SIZE = 20;
        final int MINIMUM_NUMBER = 0;
        final int MAXIMUM_NUMBER = 99;

        double sum = 0;
        double[] numbers = new double[SIZE];
        StringBuilder strBuilder = new StringBuilder();

        for (int i = 0; i < SIZE; i++) {
            numbers[i] = Math.random() * (MAXIMUM_NUMBER - MINIMUM_NUMBER + 1) + MINIMUM_NUMBER;
        }

        for (double number : numbers) {
            sum += number;
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
        System.out.println("Y el promedio de la lista es de " + sum / SIZE);
    }

}