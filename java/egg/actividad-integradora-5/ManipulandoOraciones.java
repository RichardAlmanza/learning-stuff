import java.util.Arrays;
import java.util.InputMismatchException;
import java.util.Scanner;

public class ManipulandoOraciones {

    public static void main(String[] args) {
        final int CREATE_DELETE_OPTION = 1;
        final int COUNT_CHARS_OPTION = 2;
        final int COUNT_WORDS_OPTION = 3;
        final int PRINT_SORTED_WORDS_OPTION = 4;
        final int GET_WORD_AT_OPTION = 5;
        final int FIND_WORD_OPTION = 6;
        final int CHANGE_WORD_OPTION = 7;
        final int APPEND_OPTION = 8;
        final int EXIT_OPTION = 9;

        int userOption;
        String sentenceAction = null;
        String userSentence = null;
        String[] userSentenceWords = null;
        String[] sortedWords = null;

        try (Scanner sc = new Scanner(System.in)) {
            do {
                sentenceAction = userSentence == null ? "Crear" : "Borrar";

                System.out.println(String.format("""
                        %1$d. %10$s oracion
                        %2$d. Cantidad de caracteres de la oracion
                        %3$d. Cantidad de palabras de la oracion
                        %4$d. Mostrar palabras ordenadas alfabeticamente
                        %5$d. Ingresar un numero y devolver la palabra correspondiente
                        %6$d. Buscar palabra dentro de la oracion
                        %7$d. Modificar palabra dentro de la oracion
                        %8$d. Agregar contenido a la oracion
                        %9$d. Salir
                        """, CREATE_DELETE_OPTION, COUNT_CHARS_OPTION, COUNT_WORDS_OPTION, PRINT_SORTED_WORDS_OPTION,
                        GET_WORD_AT_OPTION, FIND_WORD_OPTION, CHANGE_WORD_OPTION, APPEND_OPTION, EXIT_OPTION,
                        sentenceAction));

                do {
                    System.out.println("Ingrese una opcion del menu");
                    userOption = sc.nextInt();
                    sc.nextLine();
                } while (userOption < CREATE_DELETE_OPTION || EXIT_OPTION < userOption);

                if (!(userOption == CREATE_DELETE_OPTION || userOption == EXIT_OPTION) && userSentence == null) {
                    System.out.println("No tiene ninguna oracion almacenada");
                    continue;
                }

                switch (userOption) {
                    case CREATE_DELETE_OPTION -> {
                        boolean isNewSentenceValid = false;

                        if (userSentence == null) {
                            do {
                                System.out.println("Ingrese la palabra u oracion");
                                userSentence = sc.nextLine().trim();

                                isNewSentenceValid = !userSentence.isEmpty();
                                isNewSentenceValid &= userSentence.chars()
                                        .filter(c -> !(Character.isWhitespace(c) || Character.isLetter(c)))
                                        .toArray().length == 0;

                                if (!isNewSentenceValid) {
                                    System.out.println(
                                            "La nueva palabra o frase solo puede contener letras y/o espacios");
                                }
                            } while (!isNewSentenceValid);

                            userSentenceWords = userSentence.trim().split("\\s");
                            sortedWords = Arrays.copyOf(userSentenceWords, userSentenceWords.length);
                            Arrays.sort(sortedWords);
                        } else {
                            userSentence = null;
                            System.out.println("La oracion ha sido borrada");
                        }
                    }
                    case COUNT_CHARS_OPTION -> {

                        System.out.println(String.format("La oracion tiene %d caracteres", userSentence.length()));
                    }
                    case COUNT_WORDS_OPTION -> {

                        System.out.println(String.format("La oracion tiene %d palabras", userSentenceWords.length));
                    }
                    case PRINT_SORTED_WORDS_OPTION -> {

                        System.out.println(Arrays.toString(sortedWords));
                    }
                    case GET_WORD_AT_OPTION -> {

                        int wordIndex;

                        do {
                            System.out.println(
                                    String.format("Ingrese un numero entero entre 1 y %d:", userSentenceWords.length));
                            wordIndex = sc.nextInt() - 1;
                            sc.nextLine();
                        } while (wordIndex < 0 || userSentenceWords.length <= wordIndex);

                        System.out.println(String.format("La palabra en la posicion %d es \"%s\"", wordIndex,
                                userSentenceWords[wordIndex]));
                    }
                    case FIND_WORD_OPTION -> {

                        String wordToFind;
                        int wordIndex = 0;
                        boolean foundWord = false;
                        boolean tryAgain = false;

                        do {
                            System.out.println("Ingrese la palabra que desee buscar");
                            wordToFind = sc.next();
                            sc.nextLine();

                            for (; wordIndex < userSentenceWords.length; wordIndex++) {
                                if (wordToFind.toLowerCase().equals(userSentenceWords[wordIndex].toLowerCase())) {
                                    foundWord = true;
                                    break;
                                }
                            }

                            if (foundWord) {
                                System.out.println(String
                                        .format("La palabra se encuentra en la posicion %d de la oracion",
                                                wordIndex + 1));
                            } else {
                                System.out.println("La palabra no se encuentra en la oracion");

                                System.out.println("Desea intentar otra vez? Si/No");
                                tryAgain = sc.next().toLowerCase().equals("si");
                                sc.nextLine();
                            }
                        } while (tryAgain);
                    }
                    case CHANGE_WORD_OPTION -> {

                        String wordToFind;
                        String newSentence = null;
                        int wordIndex = 0;
                        boolean isNewSentenceValid = false;
                        boolean foundWord = false;
                        boolean tryAgain = false;

                        do {
                            System.out.println("Ingrese la palabra que desee buscar para cambiar");
                            wordToFind = sc.next();
                            sc.nextLine();

                            for (; wordIndex < userSentenceWords.length; wordIndex++) {
                                if (wordToFind.toLowerCase().equals(userSentenceWords[wordIndex].toLowerCase())) {
                                    foundWord = true;
                                    break;
                                }
                            }

                            if (foundWord) {
                                System.out.println(String
                                        .format("La palabra \"%s\" se encuentra en la posicion %d de la oracion",
                                                userSentenceWords[wordIndex], wordIndex + 1));

                                do {

                                    System.out.println("Ingrese la palabra o frase por la cual desea reemplazar");
                                    newSentence = sc.nextLine().trim();

                                    isNewSentenceValid = !newSentence.isEmpty();
                                    isNewSentenceValid &= newSentence.chars()
                                            .filter(c -> !(Character.isWhitespace(c) || Character.isLetter(c)))
                                            .toArray().length == 0;

                                    if (!isNewSentenceValid) {
                                        System.out.println(
                                                "La nueva palabra o frase solo puede contener letras y/o espacios");
                                    }
                                } while (!isNewSentenceValid);

                            } else {
                                System.out.println("La palabra no se encuentra en la oracion");

                                System.out.println("Desea intentar otra vez? Si/No");
                                tryAgain = sc.next().toLowerCase().equals("si");
                                sc.nextLine();
                            }
                        } while (tryAgain);

                        if (isNewSentenceValid) {
                            userSentenceWords[wordIndex] = newSentence;
                            userSentence = String.join(" ", userSentenceWords);
                            userSentenceWords = userSentence.trim().split("\\s");
                            sortedWords = Arrays.copyOf(userSentenceWords, userSentenceWords.length);
                            Arrays.sort(sortedWords);

                            System.out.println(String.format("La nueva oracion es \"%s\"", userSentence));
                        }
                    }
                    case APPEND_OPTION -> {
                        String newSentence = null;
                        boolean isNewSentenceValid = false;

                        do {
                            System.out.println("Ingrese la palabra o frase a agregar al final de la oracion");
                            newSentence = sc.nextLine().trim();

                            isNewSentenceValid = !newSentence.isEmpty();
                            isNewSentenceValid &= newSentence.chars()
                                    .filter(c -> !(Character.isWhitespace(c) || Character.isLetter(c)))
                                    .toArray().length == 0;

                            if (!isNewSentenceValid) {
                                System.out.println(
                                        "La nueva palabra o frase para agregar solo puede contener letras y/o espacios");
                            }
                        } while (!isNewSentenceValid);

                        userSentence = userSentence.concat(" ").concat(newSentence);
                        userSentenceWords = userSentence.trim().split("\\s");
                        sortedWords = Arrays.copyOf(userSentenceWords, userSentenceWords.length);
                        Arrays.sort(sortedWords);

                        System.out.println(String.format("La nueva oracion es \"%s\"", userSentence));
                    }
                    case EXIT_OPTION -> {
                        System.out.println("Gracias vuelva pronto");
                    }
                    default -> System.out.println("F");
                }
            } while (userOption != EXIT_OPTION);

        } catch (InputMismatchException e) {
            System.out.println("Solo se permiten numeros enteros!");
        } catch (Exception e) {
            System.out.println("Hizo KABOOM");
        }

    }
}