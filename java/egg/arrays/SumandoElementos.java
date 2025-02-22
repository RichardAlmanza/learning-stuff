public class SumandoElementos {

    public static void main(String[] args) {
        int[] numbers = { (int) (Math.random() * 100), (int) (Math.random() * 100), (int) (Math.random() * 100) };

        System.out.println(String.format("""
                Array[0] = %d
                Array[1] = %d
                Array[2] = %d

                La suma generada esta es %d""", numbers[0], numbers[1], numbers[2],
                numbers[0] + numbers[1] + numbers[2]));
    }
}
