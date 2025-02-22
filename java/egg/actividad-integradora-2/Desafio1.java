import java.util.InputMismatchException;
import java.util.Scanner;

public class Desafio1 {

    public static void main(String[] args) {
        double degrees;
        double result;
        String unit;
        String unitTarget;

        try (Scanner sc = new Scanner(System.in)) {
            System.out.print("Ingrese la temperatura (solo el numero): ");
            degrees = sc.nextDouble();
            sc.nextLine();

            System.out.print("Ingrese solo la unidad de medida ingresada (C/F): ");
            unit = sc.nextLine();

            unitTarget = switch (unit) {
                case "c", "C" -> "F";
                case "f", "F" -> "C";
                default -> throw new Exception("No ingreso correctamente una de las unidades posibles");
            };

            if (unitTarget.equals("C")) {
                result = (degrees - 32) * 5 / 9;
            } else {
                result = degrees * 9 / 5 + 32;
            }

            System.out.println(String.format("%f%s equivalen a %f%s", degrees, unit, result, unitTarget));
        } catch (InputMismatchException e) {
            System.out.println("Ingreso incorrentamente la temperatura");
        } catch (Exception e) {
            System.out.println(e.getMessage());
        }
    }
}
