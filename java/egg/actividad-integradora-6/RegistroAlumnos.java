import java.util.InputMismatchException;
import java.util.Scanner;

public class RegistroAlumnos {
    public static void main(String[] args) {
        final int REGISTER_STUDENT_OPTION = 1;
        final int SHOW_STUDENTS_OPTION = 2;
        final int SHOW_AVERAGE_SCORE_OPTION = 3;
        final int FIND_STUDENT_BY_NAME_OPTION = 4;
        final int CHANGE_SCORE_BY_NAME_OPTION = 5;
        final int REMOVE_STUDENT_BY_NAME_OPTION = 6;
        final int EXIT_OPTION = 7;

        final float MINIMUM_SCORE = (float) 0.0;
        final float MAXIMUM_SCORE = (float) 10.0;

        final int STUDENTS_MAX_SIZE = 100;

        int userOption;
        int studentCount = 0;
        String[] studentNames = new String[STUDENTS_MAX_SIZE];
        float[] studentScores = new float[STUDENTS_MAX_SIZE];

        try (Scanner sc = new Scanner(System.in)) {
            do {
                System.out.println(String.format("""
                        %d. Registrar alumno.
                        %d. Mostrar todos los alumnos.
                        %d. Mostrar promedio de notas.
                        %d. Buscar alumno por nombre.
                        %d. Modificar nota por nombre.
                        %d. Eliminar alumno por nombre.
                        %d. Salir
                        """, REGISTER_STUDENT_OPTION, SHOW_STUDENTS_OPTION, SHOW_AVERAGE_SCORE_OPTION,
                        FIND_STUDENT_BY_NAME_OPTION, CHANGE_SCORE_BY_NAME_OPTION, REMOVE_STUDENT_BY_NAME_OPTION,
                        EXIT_OPTION));

                do {
                    System.out.println("Ingrese una opcion del menu");
                    userOption = sc.nextInt();
                    sc.nextLine();
                } while (userOption < REGISTER_STUDENT_OPTION || EXIT_OPTION < userOption);

                if (!(userOption == REGISTER_STUDENT_OPTION || userOption == EXIT_OPTION) && studentCount == 0) {
                    System.out.println("No tiene ningun estudiante registrado");
                    continue;
                }

                switch (userOption) {
                    case REGISTER_STUDENT_OPTION -> {

                        boolean isNameValid = false;
                        boolean isScoreValid = false;
                        String studentName;
                        float studentScore;

                        do {
                            System.out.println("Ingrese el nombre del alumno");
                            studentName = sc.nextLine().trim();

                            isNameValid = !studentName.isEmpty();
                            isNameValid &= studentName.chars()
                                    .filter(c -> !(Character.isWhitespace(c) || Character.isLetter(c)))
                                    .toArray().length == 0;

                            if (!isNameValid) {
                                System.out.println(
                                        "El nombre del alumno solo puede contener letras y espacios");
                            }
                        } while (!isNameValid);

                        do {
                            System.out.println("Ingrese la nota del alumno");
                            studentScore = sc.nextFloat();
                            sc.nextLine();

                            isScoreValid = MINIMUM_SCORE <= studentScore && studentScore <= MAXIMUM_SCORE;

                            if (!isScoreValid) {
                                System.out.println(
                                        String.format(
                                                "La nota del alumno debe estar en el siguiente rango [%.2f, %.2f]",
                                                MINIMUM_SCORE, MAXIMUM_SCORE));
                            }
                        } while (!isScoreValid);

                        studentNames[studentCount] = studentName;
                        studentScores[studentCount] = studentScore;
                        studentCount++;
                    }
                    case SHOW_STUDENTS_OPTION -> {
                        StringBuilder sbStudents = new StringBuilder();

                        sbStudents.append("[");

                        for (int i = 0; i < studentCount; i++) {
                            sbStudents.append(String.format("\n  {\n    name: \"%s\",\n    score: %.2f\n  }",
                                    studentNames[i], studentScores[i]));
                            sbStudents.append(',');
                        }

                        sbStudents.setLength(sbStudents.length() - 1);
                        sbStudents.append("\n]\n");

                        System.out.println(sbStudents.toString());
                    }
                    case SHOW_AVERAGE_SCORE_OPTION -> {

                        float avgScore;
                        float sum = 0;

                        for (int i = 0; i < studentCount; i++) {
                            sum += studentScores[i];
                        }

                        avgScore = sum / studentCount;

                        System.out.println("La nota promedio es " + avgScore);
                    }
                    case FIND_STUDENT_BY_NAME_OPTION -> {

                        boolean studentFound = false;
                        boolean isNameValid = false;
                        String studentName;
                        float studentScore = (float) -1.0;

                        do {
                            System.out.println("Ingrese el nombre del alumno");
                            studentName = sc.nextLine().trim();

                            isNameValid = !studentName.isEmpty();
                            isNameValid &= studentName.chars()
                                    .filter(c -> !(Character.isWhitespace(c) || Character.isLetter(c)))
                                    .toArray().length == 0;

                            if (!isNameValid) {
                                System.out.println(
                                        "El nombre del alumno solo puede contener letras y espacios");
                            }
                        } while (!isNameValid);

                        for (int i = 0; i < studentCount; i++) {
                            if (studentName.equals(studentNames[i])) {
                                studentScore = studentScores[i];
                                studentFound = true;
                                break;
                            }
                        }

                        if (studentFound) {
                            System.out.println(
                                    String.format("La nota del alumno %s es %.2f", studentName, studentScore));
                        } else {
                            System.out
                                    .println(String.format("No se encontro el alumno %s en el registro", studentName));
                        }
                    }
                    case CHANGE_SCORE_BY_NAME_OPTION -> {

                        boolean studentFound = false;
                        boolean isNameValid = false;
                        boolean isScoreValid = false;
                        int studentIndex = -1;
                        float studentScore;
                        String studentName;

                        do {
                            System.out.println("Ingrese el nombre del alumno");
                            studentName = sc.nextLine().trim();

                            isNameValid = !studentName.isEmpty();
                            isNameValid &= studentName.chars()
                                    .filter(c -> !(Character.isWhitespace(c) || Character.isLetter(c)))
                                    .toArray().length == 0;

                            if (!isNameValid) {
                                System.out.println(
                                        "El nombre del alumno solo puede contener letras y espacios");
                            }
                        } while (!isNameValid);

                        for (int i = 0; i < studentCount; i++) {
                            if (studentName.equals(studentNames[i])) {
                                studentFound = true;
                                studentIndex = i;
                                break;
                            }
                        }

                        if (studentFound) {
                            do {
                                System.out.println("Ingrese la nota del alumno");
                                studentScore = sc.nextFloat();
                                sc.nextLine();

                                isScoreValid = MINIMUM_SCORE <= studentScore && studentScore <= MAXIMUM_SCORE;

                                if (!isScoreValid) {
                                    System.out.println(
                                            String.format(
                                                    "La nota del alumno debe estar en el siguiente rango [%.2f, %.2f]",
                                                    MINIMUM_SCORE, MAXIMUM_SCORE));
                                }
                            } while (!isScoreValid);

                            studentScores[studentIndex] = studentScore;
                        } else {
                            System.out
                                    .println(String.format("No se encontro el alumno %s en el registro", studentName));
                        }
                    }
                    case REMOVE_STUDENT_BY_NAME_OPTION -> {

                        boolean studentFound = false;
                        boolean isNameValid = false;
                        int studentIndex = -1;
                        String studentName;

                        do {
                            System.out.println("Ingrese el nombre del alumno");
                            studentName = sc.nextLine().trim();

                            isNameValid = !studentName.isEmpty();
                            isNameValid &= studentName.chars()
                                    .filter(c -> !(Character.isWhitespace(c) || Character.isLetter(c)))
                                    .toArray().length == 0;

                            if (!isNameValid) {
                                System.out.println(
                                        "El nombre del alumno solo puede contener letras y espacios");
                            }
                        } while (!isNameValid);

                        for (int i = 0; i < studentCount; i++) {
                            if (studentName.equals(studentNames[i])) {
                                studentFound = true;
                                studentIndex = i;
                                break;
                            }
                        }

                        if (studentFound) {

                            for (int i = studentIndex; i < studentCount - 1; i++) {
                                studentNames[i] = studentNames[i + 1];
                                studentScores[i] = studentScores[i + 1];
                            }

                            studentCount--;

                            System.out.println(String.format("El estudiante %s fue eliminado", studentName));
                        } else {
                            System.out
                                    .println(String.format("No se encontro el alumno %s en el registro", studentName));
                        }
                    }
                    case EXIT_OPTION -> {
                        System.out.println("Gracias vuelva pronto");
                    }
                    default -> System.out.println("F");
                }
            } while (userOption != EXIT_OPTION);

        } catch (InputMismatchException e) {
            System.out.println("Solo se permiten numeros reales!");
        } catch (Exception e) {
            System.out.println("Hizo KABOOM");
            e.printStackTrace(System.out);
        }

    }
}