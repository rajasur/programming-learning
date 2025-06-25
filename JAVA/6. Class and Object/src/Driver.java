public class Driver {
    String name;
    String dateOfLicense;
    static int minAgeForDriving = 18;
    public static void main(String[] args) {
        // Create a Car object
        Car myCar = new Car();

        // Set properties of the car
        myCar.noOfWheels = 4;
        myCar.color = "Red";
        myCar.maxSpeed = 200.0f;
        myCar.currentFuelInLiters = 50.0f;
        myCar.noOfSeats = 5;

        // Drive the car
        myCar.drive();

        // Add fuel to the car
        myCar.addFuel(20.0f);

        // Get current fuel level
        float fuelLevel = myCar.getCurrentFuelLevel();
        System.out.println("Current fuel level: " + fuelLevel + " liters");

        Driver myDriver = new Driver();
        System.out.println("minAgeForDriving: " + Driver.minAgeForDriving);

    }
}
