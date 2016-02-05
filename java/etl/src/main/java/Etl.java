import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class Etl {
    public Map<String, Integer> transform(Map<Integer, List<String>> old) {
        Map<String, Integer> transformed = new HashMap<>();
        old.forEach( (k,v) -> v.forEach(value -> transformed.put(value.toLowerCase(), k)));
        return transformed;
    }
}
