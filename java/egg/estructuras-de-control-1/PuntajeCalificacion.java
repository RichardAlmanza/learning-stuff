import java.util.Scanner;

public class PuntajeCalificacion {

    public static void main(String[] args) {
        final int MINIMUM_SCORE = 0;
        final int MAXIMUM_SCORE = 100;

        int score;
        String result;
        Scanner sc = new Scanner(System.in);

        System.out.println(String.format("Ingrese el puntaje del estudiante entre valores [%d, %d]", MINIMUM_SCORE, MAXIMUM_SCORE));
        score = sc.nextInt();

        if (score >= 90) {
            result = "A";
        } else if (score >= 80) {
            result = "B";
        } else if (score >= 70) {
            result = "C";
        } else if (score >= 60) {
            result = "D";
        } else {
            result = "F";
        }

        if (MINIMUM_SCORE <= score && score <= MAXIMUM_SCORE) {
            System.out.println(String.format("Con un puntaje de %d su calificacion es %s.", score, result));
        } else {
            System.out.println("Ingreso un numero fuera del rango");
        }
        
        sc.close();
    }
}