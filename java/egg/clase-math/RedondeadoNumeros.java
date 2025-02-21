import java.util.Scanner;

public class RedondeadoNumeros {

    public static void main(String[] args) {
        double number;
        Scanner sc = new Scanner(System.in);

        System.out.println("Ingrese un numero decimal para redondear");
        number = sc.nextDouble();

        System.out.println("El valor redondeado es " + Math.round(number));
        sc.close();
    }
}
