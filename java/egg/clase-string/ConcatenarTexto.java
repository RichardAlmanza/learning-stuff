import java.util.Scanner;

public class ConcatenarTexto {

    public static void main(String[] args) {
        String name;
        String surname;
        String fullName;
        Scanner sc = new Scanner(System.in);

        System.out.println("Ingrese su primer nombre");
        name = sc.nextLine();

        System.out.println("Ingrese su primer apellido");
        surname = sc.nextLine();

        fullName = name.concat(" ").concat(surname);

        System.out.println("Nombre completo: " + fullName);

        sc.close();
    }
}
