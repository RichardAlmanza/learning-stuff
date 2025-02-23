package wrappers;

import java.util.Scanner;

public class NumeroInteger {
    public static void main(String[] args) {
        Integer number;
        String numberString;

        try (Scanner sc = new Scanner(System.in)) {
            System.out.println("Ingrese el numbero entero a convertir");
            numberString = sc.nextLine();

            number = Integer.valueOf(numberString);

            System.out.println("El numero fue convertido correctamente: " + number);

        } catch (NumberFormatException e) {
            System.out.println("Ingreso incorrectamente el numero");
        }
    }

}