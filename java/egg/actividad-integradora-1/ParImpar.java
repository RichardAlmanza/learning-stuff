import java.util.Scanner;

public class ParImpar {

    public static void main(String[] args) {
        int number;
        String resultado;
        Scanner sc = new Scanner(System.in);

        System.out.print("Ingrese el numero entero a evalular si es par o impar: ");
        number = sc.nextInt();

        resultado = number % 2 == 0 ? "Par" : "Impar";
        System.out.println(String.format("El numero %d es %s", number, resultado));

        sc.close();
    }
}
