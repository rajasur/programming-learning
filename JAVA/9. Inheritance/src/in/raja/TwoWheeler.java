package in.raja;

import in.raja.inheritance.Vehicle;

public class TwoWheeler extends Vehicle {
    TwoWheeler() {
        setNumberOfTires(2);
        this.noOfSeats = 5;
    }

    public void balance() {
        System.out.println("I can balance myself with 2 tires");
        System.out.println("Number of seats: " + noOfSeats);
    }
}
