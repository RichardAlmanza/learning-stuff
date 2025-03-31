import java.util.Scanner;

public class TablaMultiplicar2 {
    private static Scanner sc = new Scanner(System.in);

    public static void main(String[] args) {
        printTable(readManyInts(3));
    }

    private static int readInt() {
        int number = 0;

        try {
            System.out.println("Ingrese un numero entero");
            number = sc.nextInt();
            sc.nextLine();
        } catch (Exception e) {
            // TODO: handle exception
            System.out.println("Kaboom! por no ingresar un numero entero");
        }

        return number;
    }

    private static int[] readManyInts(int n){
        int[] numbers = new int[n];

        for (int i = 0; i < n; i++) {
            numbers[i] = readInt();
        }

        return numbers;
    }

    private static void printTable(int number) {
        for (int i = 1; i <= 10; i++) {
            System.out.println(String.format("%d X %d = %d", number, i, i * number));
        }
    }

    private static void printTable(int[] numbers) {
        for (int number : numbers) {
            System.out.println("numero " + number);
            printTable(number);
        }
    }
}
