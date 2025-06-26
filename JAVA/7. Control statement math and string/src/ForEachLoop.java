public class ForEachLoop {
    public static void main(String[] args) {
        String[] array = new String[] {
            "Java", "Python", "C++", "JavaScript", "Ruby"
        };
        printArrayForEach(array);
        printArrayFor(array);
    }

    public static void printArrayForEach(String[] array) {
        // Using for-each loop to iterate through the array
        for (String element : array) {
            System.out.print(element + " ");
        }
        System.out.println();
    }

    public static void printArrayFor(String[] array) {
        // Using for loop to iterate through the array
       for (int i = 0; i < array.length; i++) {
            System.out.print(array[i] + " ");
        }
        System.out.println();
    }
}

