package wrappers;

import java.util.Scanner;

public class VerificarLetras {
    public static void main(String[] args) {
        String userInput;

        try (Scanner sc = new Scanner(System.in)) {
            System.out.println("Ingrese una frase");
            userInput = sc.nextLine();

            for (Character letter : userInput.toCharArray()) {
                if (!Character.isLetter(letter)) {
                    System.out.println(String.format("En la frase \"%s\", \"%s\" no es una letra", userInput,
                            Character.toString(letter)));
                    return;
                }
            }
            System.out.println("La frase ingresada solo contiene Letras");

        } catch (Exception e) {
            System.out.println("Ha ocurrido un error! si tu ni yo sabemos cual :)");
        }
    }

}