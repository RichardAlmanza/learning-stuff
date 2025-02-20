import java.util.Scanner;

public class OperadoresComparacion {

    public static void main(String[] args) {
        int a, b, c, d;
        Scanner sc = new Scanner(System.in);

        System.out.print("Ingrese el primer numero entero a comparar: ");
        a = sc.nextInt();

        System.out.print("Ingrese el segundo numero entero a comparar: ");
        b = sc.nextInt();

        System.out.println(String.format("El numero %d es mayor que %d ? = %b", a, b, a > b));
        System.out.println(String.format("El numero %d es distinto que %d ? = %b", a, b, a != b));
        System.out.println(String.format("El numero %d es divisible entre 2 ? = %b", a, a % 2 == 0));

        System.out.print("Ingrese un nuevo numero entero: ");
        c = sc.nextInt();

        System.out.print("Ingrese un nuevo numero entero: ");
        d = sc.nextInt();

        System.out.println(String.format("El numero %d es mayor que %d Y el numero %d es mayor que %d ? = %b", a, b, c, d, a > b && c > d));
        System.out.println(String.format("El numero %d es mayor que %d O el numero %d es mayor que %d ? = %b", a, b, c, d, a > b || c > d));

        sc.close();
    }
}
