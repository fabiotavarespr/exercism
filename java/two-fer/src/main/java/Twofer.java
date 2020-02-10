public class Twofer {
    public String twofer(String name) {
        name = (name == null || name.trim().isEmpty()) ? name = "you" : name;
        return "One for " + name + ", one for me.";
    }
}