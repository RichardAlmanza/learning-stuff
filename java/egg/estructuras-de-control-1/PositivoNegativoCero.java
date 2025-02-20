import java.util.Scanner;

public class PositivoNegativoCero {

    public static void main(String[] args) {
        double number;
        Scanner sc = new Scanner(System.in);

        System.out.println("Ingrese el numero a evaluar");
        number = sc.nextDouble();

        if (number > 0) {
            System.out.println("El numero es Positivo");
        } else if (number < 0) {
            System.out.println("El numero es negativo");
        } else {
            System.out.println("El numero es 0");
        }

        sc.close();
    }
}
