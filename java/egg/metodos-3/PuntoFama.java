import java.text.Normalizer;
import java.util.Scanner;

public class PuntoFama {
    private static String selectedWord;
    private static char[] userLetters;

    private static Scanner sc = new Scanner(System.in);

    public static void main(String[] args) {
        final int MAXIMUM_TRIES = 20;

        int triesCounter = 0;
        String userGuess;

        selectedWord = normalizeString(getRandomWord());
        initUserLetters(selectedWord.length());

        System.out.println(userLettersToString());

        do {
            System.out.println("Ingrese un letra para ver si se encuentra en la palabra");
            do {
                System.out.println("Solo le permiten letras (No simbolos, no numeros, no espacios)");
                System.out.println("Si ingresa varias letras, contara cada una un intento");
                System.out.println("En caso que sea la palabra completa y correcta se tomara como un solo intento");

                userGuess = readLine();
            } while (!isWordValid(userGuess));

            if (userGuess.equals(selectedWord)) {
                revealWord();
                break;
            }

            for (int i = 0; i < userGuess.length(); i++) {
                writeLetter(userGuess.charAt(i));
                triesCounter++;

                if (triesCounter >= MAXIMUM_TRIES) {
                    break;
                }
            }

            System.out.println(userLettersToString());
            System.out.println(String.format("Intentos %d/%d\n\n", triesCounter, MAXIMUM_TRIES));

        } while (triesCounter < MAXIMUM_TRIES && !isCompleted());

        if (isCompleted()) {
            System.out.println(String.format("Ganaste con %d intentos de %d", triesCounter, MAXIMUM_TRIES));
        } else {
            System.out.println("Has perdido");
            revealWord();
        }

        System.out.println(userLettersToString());

    }

    private static boolean isCompleted() {
        return String.valueOf(userLetters).equals(selectedWord);
    }

    private static String userLettersToString() {
        StringBuilder sbUserLetter = new StringBuilder();

        sbUserLetter.append("|| ");

        for (char letter : userLetters) {
            sbUserLetter.append(letter);
            sbUserLetter.append(' ');
        }

        sbUserLetter.append("||");

        return sbUserLetter.toString();
    }

    private static void initUserLetters(int size) {
        userLetters = new char[size];

        for (int i = 0; i < size; i++) {
            userLetters[i] = '_';
        }
    }

    private static boolean isWordValid(String userWord) {
        boolean isValid;

        isValid = !userWord.isEmpty();
        isValid &= userWord.chars().allMatch(Character::isLetter);

        return isValid;
    }

    private static void revealWord() {
        for (int i = 0; i < userLetters.length; i++) {
            userLetters[i] = selectedWord.charAt(i);
        }
    }

    private static void writeLetter(char letter) {
        for (int i = 0; i < userLetters.length; i++) {
            if (selectedWord.charAt(i) == letter) {
                userLetters[i] = letter;
            }
        }
    }

    private static String readLine() {
        String userInput = "";

        try {
            userInput = sc.nextLine().trim();
            userInput = normalizeString(userInput);
        } catch (Exception e) {
            System.out.println("Kaboom! por que habra explotado?");
            e.printStackTrace();
        }

        return userInput;
    }

    private static String getRandomWord() {
        String[] words = new String[] {
                "caracol", "babosa", "lombriz", "marisco", "molusco", "lagarto", "serpiente", "nuez", "almendra",
                "castaña", "arroz", "avena", "cebada", "trigo", "verdura", "mañana", "amanecer", "mediodía", "tarde",
                "anochecer", "noche", "lunes", "martes", "miércoles", "viento", "trueno", "rayo", "tormenta", "cielo",
                "este", "oeste", "calor", "agua", "hielo", "vapor", "fuego", "gas", "equipo", "autoridad", "cargo",
                "campaña", "club", "comisión", "congreso", "consejo", "partido", "país", "nación", "almacén", "hotel",
                "fábrica", "cuenta", "boleto", "entrada", "dinero", "dormitorio", "habitación", "cuarto", "oficina",
                "panel", "puerta", "ventana", "entrada", "arma", "escultura", "libro", "revista", "cuadro", "camiseta",
                "zapatilla", "cordones", "abrigo", "calle", "carretera", "autopista", "avenida", "papel", "carta",
                "comunicación", "expresión", "voz", "texto", "periodismo", "blanco", "negro", "gris", "rojo", "arte",
                "cine", "dibujo", "pintura", "música", "religión", "dios", "artículo", "educación", "escuela",
                "instituto", "colegio", "cero", "uno", "dos", "tres", "cuatro", "cinco", "seis", "siete", "ocho",
                "nueve", "ancho", "mayoría", "minoría", "aumento", "reducción", "crecimiento", "cosa", "aspecto",
                "contenido", "objeto", "parte", "sector", "palabra", "nombre", "código", "secreto", "formalidad",
                "presente", "pasado", "futuro", "ocasión", "vez", "acción", "actividad", "acto", "programa", "proyecto",
                "obra", "acuerdo", "actitud", "atención"
        };

        return words[randomIntRange(0, words.length - 1)];
    }

    private static String normalizeString(String str) {
        return Normalizer.normalize(str, Normalizer.Form.NFD)
                .toLowerCase().replaceAll("[^\\p{ASCII}]", "");
    }

    private static int randomIntRange(int minimumNumber, int maximumNumber) {
        return (int) (Math.random() * (maximumNumber - minimumNumber + 1) + minimumNumber);
    }

}
