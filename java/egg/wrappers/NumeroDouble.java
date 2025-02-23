package wrappers;

import java.util.Scanner;

public class NumeroDouble {
    public static void main(String[] args) {
        Double number;
        String numberString;

        try (Scanner sc = new Scanner(System.in)) {
            System.out.println("Ingrese el numbero decimal a convertir");
            numberString = sc.nextLine();

            number = Double.valueOf(numberString);

            System.out.println("El numero fue convertido correctamente: " + number);

        } catch (NumberFormatException e) {
            System.out.println("Ingreso incorrectamente el numero");
        }
    }

}