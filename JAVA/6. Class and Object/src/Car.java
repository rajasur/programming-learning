public class Car {
    // Instance variables
    int noOfWheels;
    String color;
    float maxSpeed;
    float currentFuelInLiters;
    int noOfSeats;

    //Instance methods
    public void drive() {
        if (currentFuelInLiters == 0) {
            System.out.println("Cannot drive, fuel tank is empty.");
            return;
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
