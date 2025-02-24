import java.util.Arrays;

public class Calificaciones {
    public static void main(String[] args) {
        final int MINIMUM_LIMIT = 1;
        final int MAXIMUM_LIMIT = 10;
        final int APPROVED_THRESHOLD = 4;
        final int EXCELLENT_THRESHOLD = 10;
        final int SIZE = 100;

        String[] scores = new String[SIZE];
        String[] excellentScores = null;
        String[] approvedScores = null;
        String[] disapprovedScores = null;

        for (int i = 0; i < SIZE; i++) {
            scores[i] = String.valueOf(Math.random() * (MAXIMUM_LIMIT - MINIMUM_LIMIT + 1) + MINIMUM_LIMIT);
        }

        excellentScores = Arrays.stream(scores).filter(str -> Double.parseDouble(str) >= EXCELLENT_THRESHOLD)
                .map(str -> String.valueOf((int) Double.parseDouble(str)))
                .toArray(size -> new String[size]);
        approvedScores = Arrays.stream(scores).filter(
                str -> APPROVED_THRESHOLD <= Double.parseDouble(str) && Double.parseDouble(str) < EXCELLENT_THRESHOLD)
                .map(str -> String.valueOf((int) Double.parseDouble(str)))
                .toArray(size -> new String[size]);
        disapprovedScores = Arrays.stream(scores).filter(str -> Double.parseDouble(str) < APPROVED_THRESHOLD)
                .map(str -> String.valueOf((int) Double.parseDouble(str)))
                .toArray(size -> new String[size]);

        System.out.println("Calificaciones: " + scores.length);
        System.out.println(Arrays.toString(scores));
        System.out.println("Calificaciones desaprobadas: " + disapprovedScores.length);
        System.out.println(Arrays.toString(disapprovedScores));
        System.out.println("Calificaciones aprobadas: " + approvedScores.length);
        System.out.println(Arrays.toString(approvedScores));
        System.out.println("Calificaciones Excelentes: " + excellentScores.length);
        System.out.println(Arrays.toString(excellentScores));

    }
}
