import java.util.Scanner;

public class Calificacion {

    public static void main(String[] args) {
        int score;
        String result;
        Scanner sc = new Scanner(System.in);

        System.out.println("Ingrese un puntaje del 1 al 5, numero entero");
        score = sc.nextInt();

        result = switch (score) {
            case 1 -> "Muy deficiente";
            case 2 -> "Deficiente";
            case 3 -> "Suficiente";
            case 4 -> "Notable";
            case 5 -> "Sobresaliente";
            default -> null;
        };

        if (result != null) {
            System.out.println("La calificacion es " + result);
        } else {
            System.out.println("No ingreso un numero valido!");
        }

        sc.close();
    }
}
