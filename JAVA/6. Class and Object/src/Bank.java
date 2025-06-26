public class Bank {
    String name;
    String address;
    String phoneNumber;
    Bank(String name, String address, String phoneNumber) {
        this.name = name;
        this.address = address;
        this.phoneNumber = phoneNumber;
    }
    public void displayBankInfo() {
        System.out.println("Bank Name: " + name);
        System.out.println("Address: " + address);
        System.out.println("Phone Number: " + phoneNumber);
    }
}

