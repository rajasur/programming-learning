import java.util.Scanner;

public class Switch {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);
        System.out.println("Welcome to the day of the week checker\n");
        System.out.println("Please enter day of number: ");
        int day = input.nextInt();
        oldSwitch(day);
        newSwitch(day);
    }
    // This program demonstrates the use of switch statements in Java.
    public static void newSwitch(int day) {
       String dayStr = switch (day) {
           case 1 -> "Monday";
           case 2 -> "Tuesday";
           case 3 -> "Wednesday";
           case 4 -> "Thursday";
           case 5 -> "Friday";
           case 6 -> "Saturday";
           case 7 -> "Sunday";
           default -> "Invalid day number. Please enter a number between 1 and 7.";
       };
        System.out.println("New Switch day: "+dayStr);
    }

    // This program demonstrates the use of switch statements in Java.
    public static void oldSwitch(int day) {
        switch(day) {
            case 1:
                System.out.println("Monday");
                break;
            case 2:
                System.out.println("Tuesday");
                break;
            case 3:
                System.out.println("Wednesday");
                break;
            case 4:
                System.out.println("Thursday");
                break;
            case 5:
                System.out.println("Friday");
                break;
            case 6:
                System.out.println("Saturday");
                break;
            case 7:
                System.out.println("Sunday");
                break;
            default:
                System.out.println("Invalid day number. Please enter a number between 1 and 7.");
        }
    }
}
