import java.util.Scanner;

public class Recursion {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        System.out.println("Enter a number to calculate its factorial: ");
        int number = sc.nextInt();
        long fact = factorial(number);
        System.out.println("The factorial of " + number + " is: " + fact);
    }
   public static int factorial(int number) {
       if (number == 1 || number == 0) {
           return 1;
       } else {
          return number * factorial(number - 1);
       }
   }
}