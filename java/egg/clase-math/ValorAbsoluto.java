import java.util.Scanner;

public class ValorAbsoluto {

    public static void main(String[] args) {
        double number;
        Scanner sc = new Scanner(System.in);

        System.out.println("Ingrese un numero para obtener su valor absoluto");
        number = sc.nextDouble();

        System.out.println("El valor absoluto es " + Math.abs(number));
        sc.close();
    }
}