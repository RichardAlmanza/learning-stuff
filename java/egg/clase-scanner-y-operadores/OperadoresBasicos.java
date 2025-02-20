import java.util.Scanner;

public class OperadoresBasicos {

    public static void main(String[] args) {
        int a = 342;
        int b = 121;
        Scanner sc = new Scanner(System.in);

        System.out.println(String.format("La variable a contiene el numero %d, y b tiene %d", a, b));
        System.out.println("sus operaciones siguen el orden a OPERARION b");
        System.out.println(String.format("La suma es %d", a + b));
        System.out.println(String.format("La resta es %d", a - b));
        System.out.println(String.format("La multiplicacion es %d", a * b));
        System.out.println(String.format("La division entera es %d", a / b));
        System.out.println(String.format("el modulo es %d", a % b));

        System.out.println("Ingrese el nuevo valor de a");
        a = sc.nextInt();

        System.out.println("Ingrese el nuevo valor de b");
        b = sc.nextInt();

        System.out.println(String.format("La suma es %d", a + b));

        sc.close();
    }
}
