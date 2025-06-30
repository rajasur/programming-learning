package in.raja.inheritance;

import in.raja.TwoWheeler;

public class InheritanceTest {
    public static void main(String[] args) {
        Vehicle vehicle = new Vehicle();
        vehicle.commute();


        TwoWheeler two = new TwoWheeler();
        two.commute();
        two.balance();

        MotorCycle motor = new MotorCycle();
        motor.commute();
        motor.balance();
        motor.star();
    }
}
