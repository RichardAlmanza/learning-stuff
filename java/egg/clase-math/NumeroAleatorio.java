public class NumeroAleatorio {
    public static void main(String[] args) {
        final int MINIMUM_NUMBER = 1;
        final int MAXIMUM_NUMBER = 355;

        double randomDouble = Math.random();
        double randomDoubleInRange = randomDouble * (MAXIMUM_NUMBER - MINIMUM_NUMBER) + MINIMUM_NUMBER;
        double randomDoubleInRangeIncluingMax = randomDouble * (MAXIMUM_NUMBER - MINIMUM_NUMBER + 1) + MINIMUM_NUMBER;

        int randomInt = (int) (randomDouble * (MAXIMUM_NUMBER - MINIMUM_NUMBER) + MINIMUM_NUMBER);
        int randomIntIncludingMax = (int) (randomDouble * (MAXIMUM_NUMBER - MINIMUM_NUMBER + 1) + MINIMUM_NUMBER);

        System.out.println(String.format("""
                Numero aleatorio generado:                  \t %3$f
                Numero minimo:                              \t %1$d
                Numero maximo:                              \t %2$d
                Numero aleatorio en rango [%1$d, %2$d):         \t %4$f
                Numero entero aleatorio en rango [%1$d, %2$d):  \t %6$d
                Numero aleatorio en rango [%1$d, %2$d]:         \t %5$f
                Numero entero aleatorio en rango [%1$d, %2$d]:  \t %7$d
                    """, MINIMUM_NUMBER, MAXIMUM_NUMBER, randomDouble, randomDoubleInRange,
                randomDoubleInRangeIncluingMax, randomInt, randomIntIncludingMax));
    }
}
