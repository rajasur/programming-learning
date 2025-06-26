public class Car {
    // Instance variables
    int noOfWheels;
    String color;
    float maxSpeed;
    float currentFuelInLiters;
    int noOfSeats;

    // Default constructor
    Car(String color) {
        this.noOfWheels = 4;
        this.color = color;
        this.maxSpeed = 120.0f;
        this.currentFuelInLiters = 50.0f;
        this.noOfSeats = 5;
    }
    // Constructor Chaining
    Car() {
        this("Black");
    }
    //Instance methods
    public void drive() {
        if (currentFuelInLiters == 0) {
            System.out.println("Cannot drive, fuel tank is empty.");
        } else if (currentFuelInLiters < 5) {
            System.out.println("Warning: Low fuel level. Please refuel soon.");
            currentFuelInLiters--;
        } else {
            System.out.println("Car is driving");
            currentFuelInLiters--;
        }
    }

    public void addFuel(float currentFuelInLiters) {
        this.currentFuelInLiters += currentFuelInLiters;
    }

    public float getCurrentFuelLevel() {
        return currentFuelInLiters;
    }
}
