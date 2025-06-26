import java.util.Scanner;

public class TernaryOperator {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);
        System.out.println("Welcome to number checker\n");
        System.out.println("Please enter your first number: ");
        int num1 = input.nextInt();
        System.out.println("Please enter your second number: ");
        int num2 = input.nextInt();

        // Using if-else to find the greater number
        int greaterNumber;
        if (num1 > num2) {
            greaterNumber = num1;
        } else {
           greaterNumber = num2;
        }
        System.out.println("The greater number is: " + greaterNumber);

        // Using ternary operator to find the greater number
        int greaterNumberTernary = (num1 > num2) ? num1 : num2;
        System.out.println("The greater number using ternary operator is: " + greaterNumberTernary);
    }
}
