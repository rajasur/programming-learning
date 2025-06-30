package in.raja;

public class AccessTest {
    public static void main(String[] args) {
        Car car = new Car();
        car.color = "Red"; // Accessing public field
        car.model = "Sedan";
        car.costOfPurchase// Accessing public field

        System.out.println("Car color: " + car.color);
        System.out.println("Car model: " + car.model);

    }
}
