import java.util.Scanner;

public class UsandoScanner {

    public static void main(String[] args) {
        int age;
        String name;
        Scanner sc = new Scanner(System.in);

        System.out.println("Ingrese su edad");
        age = sc.nextInt();
        sc.nextLine();

        System.out.println("Ingrese su nombre");
        name = sc.nextLine();

        System.out.println(String.format("Su nombre es %s y su edad es %d", name, age));

        sc.close();
    }
}