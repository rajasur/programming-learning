public class StringTest {
    public static void main(String[] args) {
        StringBuilder sb = new StringBuilder();
        sb.append(45);
        sb.append(", now this is the");
        sb.append(" 78.80");
        sb.toString();
        System.out.println(sb);
    }
}
