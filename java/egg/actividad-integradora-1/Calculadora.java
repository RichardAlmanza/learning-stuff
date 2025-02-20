import java.util.Scanner;

public class Calculadora {

    public static void main(String[] args) {
        final String RESULT_MSG_TEMPLATE = "El resultado es: %.2f %s %.2f = %.2f";

        double a;
        double b;
        String operator;
        String result;
        Scanner sc = new Scanner(System.in);

        System.out.println("Ingrese el operador que desea realizar");
        System.out.println("""
                + sumar
                - restar
                * mutiplicar
                / dividir
                """);
        operator = sc.nextLine();

        System.out.println("Ingrese el primer numero que desea operar");
        a = sc.nextDouble();

        System.out.println("Ingrese el segundo numero que desea operar");
        b = sc.nextDouble();

        result = operator.equals("+") ? String.format(RESULT_MSG_TEMPLATE, a, operator, b, a + b) :
                 operator.equals("-") ? String.format(RESULT_MSG_TEMPLATE, a, operator, b, a - b) :
                 operator.equals("*") ? String.format(RESULT_MSG_TEMPLATE, a, operator, b, a * b) :
                 operator.equals("/") ? String.format(RESULT_MSG_TEMPLATE, a, operator, b, a / b) :
                 "No ingreso una opcion valida";

        System.out.println(result);
        
        sc.close();
    }
}
