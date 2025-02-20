import java.util.Scanner;

public class DiasSemana {

    public static void main(String[] args) {
        int day;
        String strDay;
        Scanner sc = new Scanner(System.in);

        System.out.println("Ingrese el dia de la semana de 1 al 7");
        day = sc.nextInt();

        strDay = switch (day) {
            case 1 -> "Lunes";
            case 2 -> "Martes";
            case 3 -> "Miercoles";
            case 4 -> "Jueves";
            case 5 -> "Viernes";
            case 6 -> "Sabado";
            case 7 -> "Domingo";
            default -> null;
        };

        if (strDay != null) {
            System.out.println("El dia ingresado es " + strDay);
        } else {
            System.out.println("No ingreso un numero valido!");
        }

        sc.close();
    }
}
