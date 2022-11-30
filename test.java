import java.util.List;
import java.util.Map;

class FirstTestFunctions {
    public static void testFunction() {
        Test test = (new Test());
        String testAgain = (new Test()).pos.x;
    }
}

class Point {
    String x;
    Long y;

}
class Test {
    Point pos;
    String value;
    Long number;
    Map<Long, List<Long>> lookup;
    List<List<Long>> things;
    public Long test() {
        number = (42L + 2L);
        if ((things.get(Long(1)).get(Long(1)) == lookup.get(Long(1)).get(Long(2)))) {
            number = Long(22);

        } else {
            number = Long(12);
        }
        return things.get(Long(0)).get(Long(1));
    }
}