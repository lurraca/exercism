import java.util.*;

public class DNA {

    private String string;

    private static List<Character> nucleotidesArray =  Arrays.asList('A' ,'C', 'T', 'G');

    public DNA(String s) {
        string = s;
    }


    public Integer count(char nucleotide) {
        if (!nucleotidesArray.contains(nucleotide)) throw new IllegalArgumentException();
        return countMatches(string, nucleotide);
    }

    private int countMatches(String string, char c) {
        return string.replaceAll("[^"+ c +"]", "").length();
    }

    public Map<Character, Integer> nucleotideCounts() {
        Map<Character, Integer> count = new HashMap<>();
        nucleotidesArray.forEach(nucleotide -> {
            count.put(nucleotide, countMatches(string, nucleotide));
        });
        return count;
    }
}
