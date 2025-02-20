import java.util.Scanner;

public class LongitudCadena {

    public static void main(String[] args) {
        String userText;
        Scanner sc = new Scanner(System.in);

        System.out.println("Ingrese un texto");
        userText = sc.nextLine();

        System.out.println(String.format("El texto \"%s\" tiene %d caracteres", userText, userText.length()));

        sc.close();
    }
}
