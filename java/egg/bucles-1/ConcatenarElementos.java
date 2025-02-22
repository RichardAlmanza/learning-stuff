public class ConcatenarElementos {
    public static void main(String[] args) {
        String sentence = "";
        String[] sentences = { "Hola mundo!", "Como estan?", "Aqui no hay nada nuevo!", "Buggs!" };

        for (String aSentence : sentences) {
            sentence = sentence.concat(" ").concat(aSentence);
        }

        System.out.println(sentence);
    }
}
