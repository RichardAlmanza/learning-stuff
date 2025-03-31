import java.util.Scanner;

public class TablaMultiplicar {

    public static void main(String[] args) {
        printTable(readInt());
    }

    private static int readInt() {
        int number = 0;

        try (Scanner sc = new Scanner(System.in)) {
            System.out.println("Ingrese un numero entero");
            number = sc.nextInt();
            sc.nextLine();
        } catch (Exception e) {
            // TODO: handle exception
            System.out.println("Kaboom! por no ingresar un numero entero");
        }

        return number;
    }

    private static void printTable(int number) {
        for (int i = 1; i <= 10; i++) {
            System.out.println(String.format("%d X %d = %d", number, i, i * number));
        }
    }
}