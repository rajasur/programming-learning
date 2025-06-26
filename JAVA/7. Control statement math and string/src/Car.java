public class Car {
    int noOfWheels;
    int noOfDoors;
    int maxSpeed;
    String name;
    String model;
    String company;

    public Car(int noOfWheels, int noOfDoors, int maxSpeed, String name, String model, String company) {
        this.noOfWheels = noOfWheels;
        this.noOfDoors = noOfDoors;
        this.maxSpeed = maxSpeed;
        this.name = name;
        this.model = model;
        this.company = company;
    }

    @Override
    public String toString() {
        return "Car{" +
                "noOfWheels=" + noOfWheels +
                ", noOfDoors=" + noOfDoors +
                ", maxSpeed=" + maxSpeed +
                ", name='" + name + '\'' +
                ", model='" + model + '\'' +
                ", company='" + company + '\'' +
                '}';
    }

    public static void main(String[] args) {
        Car swift = new Car(4, 4, 180, "Swift", "VXI", "Maruti Suzuki");
        System.out.println(swift.toString());
    }
}
