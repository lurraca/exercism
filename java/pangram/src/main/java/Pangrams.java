//TODO: Look into a more functional way of doing this.
public class Pangrams {

    public final static char[] alphabet = "abcdefghijklmnopqrstuvwxyz".toCharArray();;

    public static Boolean isPangram(String sentence) {
        String trimmedSentence = cleanUp(sentence.trim());
        Boolean result = false;

        for(char letter : alphabet){
            if (trimmedSentence.contains(Character.toString(letter))) {
                result = true;
            } else {
                return false;
            }
        }
        return result;
    }

    private static String cleanUp(String sentence) {
        String cleanSentence = sentence.toLowerCase().replaceAll("[^\\w\\s]", "");

        return cleanSentence;
    }
}
