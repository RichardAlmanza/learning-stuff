import java.util.Scanner;

public class ExtraerSubcadena {

    public static void main(String[] args) {
        int beginIndex;
        int endIndex;
        String sentence;


        try (Scanner sc = new Scanner(System.in)) {
            System.out.println("Ingrese un texto o frase");
            sentence = sc.nextLine();

            System.out.println("Ingrece el indice inicial de la subcadena");
            beginIndex = sc.nextInt();

            System.out.println("Ingrece el indice final de la subcadena");
            endIndex = sc.nextInt();

            System.out.println(String.format("El texto es \"%s\" y su subcadena es \"%s\"", sentence, sentence.substring(beginIndex, endIndex)));
        } catch (IndexOutOfBoundsException e) {
            System.out.println("Ingreso incorrectamente los indice");
        }
    }
}
