# Object-Oriented Programming (OOP) Principles

Object-Oriented Programming is a programming paradigm based on the concept of "objects", which can contain data and code: data in the form of fields, and code in the form of procedures.

OOP is centered around **4 main principles**:

---

## 1. Encapsulation

### 📌 Definition:
Encapsulation is the process of wrapping data (variables) and methods (functions) together as a single unit (class). It restricts direct access to some of the object's components.

### 🎯 Purpose:
- Protect internal object state
- Hide implementation details
- Prevent unwanted modifications

### 🔒 Access Modifiers:
- `private` – accessible within the class
- `protected` – accessible within the class and subclasses
- `public` – accessible everywhere

### ✅ Example in Java:
```java
class BankAccount {
    private double balance;

    public void deposit(double amount) {
        if (amount > 0) {
            balance += amount;
        }
    }

    public double getBalance() {
        return balance;
    }
}
````

---

## 2. Abstraction

### 📌 Definition:

Abstraction is hiding the internal implementation details and showing only the necessary functionality.

### 🎯 Purpose:

* Reduce complexity
* Increase code readability
* Separate "what" from "how"

### 🔧 Achieved using:

* Abstract classes
* Interfaces

### ✅ Example in Java:

```java
abstract class Animal {
    abstract void makeSound();
}

class Dog extends Animal {
    void makeSound() {
        System.out.println("Bark");
    }
}
```

---

## 3. Inheritance

### 📌 Definition:

Inheritance is a mechanism where a new class (subclass/child) inherits the properties and behaviors of an existing class (superclass/parent).

### 🎯 Purpose:

* Code reuse
* Establish relationship between base and derived class

### 🧩 Syntax:

```java
class Parent {
    void display() {
        System.out.println("Parent class");
    }
}

class Child extends Parent {
    void show() {
        System.out.println("Child class");
    }
}
```

### ✅ Output:

```
Parent class
Child class
```

---

## 4. Polymorphism

### 📌 Definition:

Polymorphism means "many forms". A single action can behave differently based on the object that invokes it.

### 🎯 Purpose:

* Implement flexibility and reusability
* Decouple code components

### Types:

#### 1. Compile-time Polymorphism (Method Overloading)

```java
class Calculator {
    int add(int a, int b) {
        return a + b;
    }

    double add(double a, double b) {
        return a + b;
    }
}
```

#### 2. Run-time Polymorphism (Method Overriding)

```java
class Animal {
    void makeSound() {
        System.out.println("Some sound");
    }
}

class Cat extends Animal {
    void makeSound() {
        System.out.println("Meow");
    }
}
```

### ✅ Output:

```
Meow
```

---

## 🧱 Other Key OOP Concepts

### Association

* Relationship between two classes.
* Example: A student and a college.

### Aggregation

* "Has-a" relationship (weak relationship)
* Example: A department has teachers, but teachers can exist without the department.

### Composition

* Stronger "has-a" relationship.
* Example: A house has rooms; if the house is destroyed, rooms are destroyed.

---

## 🔠 SOLID Principles (Advanced OOP Design)

| Principle                     | Description                                                                   |
| ----------------------------- | ----------------------------------------------------------------------------- |
| **S** - Single Responsibility | A class should have one and only one reason to change                         |
| **O** - Open/Closed           | Software entities should be open for extension, but closed for modification   |
| **L** - Liskov Substitution   | Subtypes must be substitutable for their base types                           |
| **I** - Interface Segregation | Many client-specific interfaces are better than one general-purpose interface |
| **D** - Dependency Inversion  | Depend on abstractions, not concretions                                       |

---

## 🧠 Summary Table

| Principle     | Purpose                        | Mechanism                         |
| ------------- | ------------------------------ | --------------------------------- |
| Encapsulation | Hide internal state            | Access Modifiers, Getters/Setters |
| Abstraction   | Hide implementation complexity | Abstract Classes, Interfaces      |
| Inheritance   | Reuse code                     | `extends` keyword                 |
| Polymorphism  | Dynamic behavior               | Method Overloading/Overriding     |

---

## 📚 Visual Diagram

```text
         +-------------------------+
         |      Class (Object)     |
         +-------------------------+
         | - Fields / Attributes   |
         | - Methods / Behaviors   |
         +-------------------------+
                  /\
      Encapsulation || Inheritance
                  ||
     ------------------------
    |          Polymorphism        |
    |          Abstraction         |
     ------------------------
```

---

## ✅ Real-world Analogy

* **Encapsulation**: A capsule of medicine hides its inner contents.
* **Abstraction**: You drive a car without knowing how the engine works.
* **Inheritance**: A child inherits traits from parents.
* **Polymorphism**: A person behaves differently in different situations (as a student, as a friend, etc.)
