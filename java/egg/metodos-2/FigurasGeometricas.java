import java.util.Scanner;

public class FigurasGeometricas {
    final static int CIRCLE_OPTION = 1;
    final static int TRIANGLE_OPTION = 2;
    final static int PARALLELOGRAM_OPTION = 3;

    final static int PERIMETER_OPTION = 1;
    final static int AREA_OPTION = 2;

    private static Scanner sc = new Scanner(System.in);

    public static void main(String[] args) {
        int figureOption = figureMenu();
        int operationOption = operationMenu();
        double result = -1;

        switch (figureOption) {
            case CIRCLE_OPTION -> {
                switch (operationOption) {
                    case PERIMETER_OPTION -> {
                        result = askForCirclePerimeter();
                    }
                    case AREA_OPTION -> {
                        result = askForCircleArea();
                    }
                }
            }
            case TRIANGLE_OPTION -> {
                switch (operationOption) {
                    case PERIMETER_OPTION -> {
                        result = askForTrianglePerimeter();
                    }
                    case AREA_OPTION -> {
                        result = askForTriangleArea();
                    }
                }
            }
            case PARALLELOGRAM_OPTION -> {
                switch (operationOption) {
                    case PERIMETER_OPTION -> {
                        result = askForParallelogramPerimeter();
                    }
                    case AREA_OPTION -> {
                        result = askForParallelogramArea();
                    }
                }
            }

            default -> System.out.println("No deberias estar aqui");
        }

        System.out.println(formatResult(figureOption, operationOption, result));
    }

    private static String formatResult(int figureOption, int operationOption, double result) {
        final String RESULT_SCHEME = "El %s del %s es %.5f";

        String figure;
        String operation;

        figure = switch (figureOption) {
            case CIRCLE_OPTION:
                yield "circulo";
            case TRIANGLE_OPTION:
                yield "triangulo";
            case PARALLELOGRAM_OPTION:
                yield "paralelogramo";
            default:
                yield "<< Ups! no deberias estar viendo esto >>";
        };

        operation = switch (operationOption) {
            case PERIMETER_OPTION:
                yield "perimetro";
            case AREA_OPTION:
                yield "area";
            default:
                yield "<< Ups! no deberias estar viendo esto >>";
        };

        return String.format(RESULT_SCHEME, operation, figure, result);
    }

    private static double askForCirclePerimeter() {
        String[] questions = new String[] { "Cual es el radio del circulo?" };
        double[] answers = new double[questions.length];

        askFor(questions, answers);

        return circlePerimeter(answers[0]);
    }

    private static double askForCircleArea() {
        String[] questions = new String[] { "Cual es el radio del circulo?" };
        double[] answers = new double[questions.length];

        askFor(questions, answers);

        return circleArea(answers[0]);
    }

    private static double askForTrianglePerimeter() {
        String questionBase = "Cuanto mide el lado %s del triangulo?";
        String[] questions = new String[] {
                String.format(questionBase, "A"),
                String.format(questionBase, "B"),
                String.format(questionBase, "C")
        };
        double[] answers = new double[questions.length];

        askFor(questions, answers);

        return trianglePerimeter(answers[0], answers[1], answers[2]);
    }

    private static double askForTriangleArea() {
        String questionBase = "Cuanto mide la %s del triangulo?";
        String[] questions = new String[] {
                String.format(questionBase, "base"),
                String.format(questionBase, "altura")
        };
        double[] answers = new double[questions.length];

        askFor(questions, answers);

        return triangleArea(answers[0], answers[1]);
    }

    private static double askForParallelogramPerimeter() {
        String questionBase = "Cuanto mide el lado abyacente %s del paralelogramo?";
        String[] questions = new String[] {
                String.format(questionBase, "A"),
                String.format(questionBase, "B")
        };
        double[] answers = new double[questions.length];

        askFor(questions, answers);

        return parallelogramPerimeter(answers[0], answers[1]);
    }

    private static double askForParallelogramArea() {
        String questionBase = "Cuanto mide la %s del paralelogramo?";
        String[] questions = new String[] {
                String.format(questionBase, "base"),
                String.format(questionBase, "altura")
        };
        double[] answers = new double[questions.length];

        askFor(questions, answers);

        return parallelogramArea(answers[0], answers[1]);
    }

    private static int figureMenu() {
        int userOption = -1;

        System.out.println(String.format("""
                %d. Circulo.
                %d. Triangulo.
                %d. Paralelogramo.
                """, CIRCLE_OPTION, TRIANGLE_OPTION, PARALLELOGRAM_OPTION));

        do {
            System.out.println("Ingrese una opcion del menu");
            userOption = readInt();
        } while (userOption < CIRCLE_OPTION || PARALLELOGRAM_OPTION < userOption);

        return userOption;
    }

    private static int operationMenu() {
        int userOption = -1;

        System.out.println(String.format("""
                %d. Perimetro.
                %d. Area.
                """, PERIMETER_OPTION, AREA_OPTION));

        do {
            System.out.println("Ingrese una opcion del menu");
            userOption = readInt();
        } while (userOption < PERIMETER_OPTION || AREA_OPTION < userOption);

        return userOption;
    }

    private static void askFor(String[] questions, double[] answers) {
        for (int i = 0; i < answers.length; i++) {
            System.out.println(questions[i]);
            answers[i] = readNumber();
        }
    }

    private static int readInt() {
        int number = 0;

        try {
            System.out.println("Ingrese un numero entero");
            number = sc.nextInt();
            sc.nextLine();
        } catch (Exception e) {
            // TODO: handle exception
            System.out.println("Kaboom! por no ingresar un numero entero");
        }

        return number;
    }

    private static double readNumber() {
        double number = 0;

        try {
            System.out.println("Ingrese un numero");
            number = sc.nextDouble();
            sc.nextLine();
        } catch (Exception e) {
            // TODO: handle exception
            System.out.println("Kaboom! por no ingresar un numero");
        }

        return number;
    }

    private static double circlePerimeter(double radius) {
        return 2 * Math.PI * radius;
    }

    private static double circleArea(double radius) {
        return Math.PI * Math.pow(radius, 2);
    }

    private static double trianglePerimeter(double sideA, double sideB, double sideC) {
        return sideA + sideB + sideC;
    }

    private static double triangleArea(double base, double height) {
        return  base * height / 2;
    }

    private static double parallelogramPerimeter(double sideA, double sideB) {
        return 2 * (sideA + sideB);
    }

    private static double parallelogramArea(double base, double height) {
        return base * height;
    }
}
