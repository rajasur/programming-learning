// File: Vehicle.java
// Package declaration
package in.raja.inheritance;

/**
 * Base class representing a general Vehicle.
 */
public class Vehicle {

    // Private: Accessible only within this class
    private int numberOfTires;

    // Protected: Accessible within the same package or by subclasses in other packages
    protected int noOfSeats;

    // Public getter for private member
    public int getNumberOfTires() {
        return numberOfTires;
    }

    // Public setter for private member
    public void setNumberOfTires(int numberOfTires) {
        this.numberOfTires = numberOfTires;
    }

    // Public method to describe commuting using the vehicle
    public void commute() {
        System.out.printf("I am going from Place A to Place B using %d tires\n", numberOfTires);
    }
}
