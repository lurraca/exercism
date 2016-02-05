import java.util.Arrays;
import java.util.Hashtable;
import java.util.List;
import java.util.Map;

public class WordCount {

    public Map<String,Integer> phrase(String phrase){
        List<String> arrayPhrase = Arrays.asList(phrase.split(" "));
        Map<String,Integer> wordCount = new Hashtable<>();
        arrayPhrase.forEach((word) -> {
            String cleanWord = cleanUp(word);
            if (cleanWord.isEmpty()) return;
            wordCount.put(cleanWord, (wordCount.get(cleanWord) == null ? 0 : wordCount.get(cleanWord)) + 1);
        });
        return wordCount;
    }

    private String cleanUp(String word) {
        String cleanWord = word.toLowerCase().replaceAll("[^\\w\\s]", "");

        return cleanWord;
    }
}
