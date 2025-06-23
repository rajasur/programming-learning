# 🧾 Go Interfaces and Design Principles — Detailed Notes

This document explores important Go concepts based on the program you provided, including **interfaces**, **composition**, **Open/Closed Principle**, and real-world payment gateway abstractions.

---

## 📦 Code Summary

```go
package main

import "fmt"

type paymenter interface {
	pay(amount float32)
	refund(amount float32, account string)
}

type payment struct {
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("making payment using razorpay", amount)
}

type fakepayment struct{}

func (f fakepayment) pay(amount float32) {
	fmt.Println("making payment using fake gateway for testing purpose")
}

type paypal struct{}

func (p paypal) pay(amount float32) {
	fmt.Println("making payment using paypal", amount)
}

func (p paypal) refund(amount float32, account string) {
	// refund logic here
}

func main() {
	paypalGw := paypal{}
	newPayment := payment{
		gateway: paypalGw,
	}
	newPayment.makePayment(100)
}
```

---

## 📌 1. Interface in Go

An `interface` in Go defines a set of method signatures. Any type that implements all those methods **implicitly** satisfies the interface.

### ✅ Syntax:

```go
type paymenter interface {
	pay(amount float32)
	refund(amount float32, account string)
}
```

This defines a contract that **any type** must satisfy to become a `paymenter`.
In this case:

* `paypal` implements both `pay` and `refund`, so it satisfies `paymenter`.

---

## 📌 2. Struct with Interface Field (Composition)

```go
type payment struct {
	gateway paymenter
}
```

The `payment` struct has a field `gateway` of type `paymenter`. This allows **any struct that implements the interface** to be assigned.

### 🚀 Benefit:

This allows `payment` to work with any payment gateway **without modification**, following the **Open/Closed Principle** (OCP).

---

## 📌 3. Open/Closed Principle (OCP)

> “Software entities (classes, modules, functions, etc.) should be open for extension, but closed for modification.”

In this code:

* You can **add new gateways** (`stripe`, `paypal`, `fakepayment`, etc.) without changing the logic of `makePayment`.

### ✅ OCP in Action:

```go
func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}
```

It does not care *which* gateway is used, as long as it satisfies the `paymenter` interface.

---

## 📌 4. Concrete Structs Implementing Interface

Each payment gateway implements the `pay` method:

### 🧱 Example: Razorpay

```go
type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("making payment using razorpay", amount)
}
```

### 🧱 Example: Paypal

```go
type paypal struct{}

func (p paypal) pay(amount float32) {
	fmt.Println("making payment using paypal", amount)
}

func (p paypal) refund(amount float32, account string) {
	// implementation
}
```

Only `paypal` implements both `pay` and `refund`, so only `paypal` satisfies the `paymenter` interface.

---

## 📌 5. Struct Not Fully Satisfying Interface

### ❌ This will cause a compile error if used:

```go
type fakepayment struct{}

func (f fakepayment) pay(amount float32) {
	fmt.Println("making payment using fake gateway")
}
```

`fakepayment` **does not implement** `refund()`, so it **does not satisfy** `paymenter`.

To fix:

```go
func (f fakepayment) refund(amount float32, account string) {
	fmt.Println("fake refund issued")
}
```

---

## 📌 6. Interface Usage in `main`

### 🧪 Example:

```go
paypalGw := paypal{}
newPayment := payment{
	gateway: paypalGw,
}
newPayment.makePayment(100)
```

This creates a `paypal` object, wraps it in a `payment` struct, and calls `makePayment`.

Output:

```
making payment using paypal 100
```

---

## 📌 7. Commented Out (Unused) Code

The original code shows how to extend the solution:

```go
// type stripe struct{}
// func (s stripe) pay(amount float32) {
//     fmt.Println("making payment using stripe", amount)
// }
```

If needed, just implement `refund()` as well and plug it into the system without changing any core logic.

---

## 🧵 Summary Table

| Concept                    | Used? | Description                                       |
| -------------------------- | ----- | ------------------------------------------------- |
| Interface declaration      | ✅     | `paymenter` with `pay`, `refund`                  |
| Interface implementation   | ✅     | `paypal`, `razorpay`                              |
| Struct composition         | ✅     | `payment` includes `paymenter`                    |
| Open/Closed Principle      | ✅     | Can extend with new gateways                      |
| Interface polymorphism     | ✅     | `makePayment()` works on any gateway              |
| Partial interface impl.    | ✅     | `fakepayment` lacks `refund()`                    |
| Testing with fake gateway  | ✅     | Use `fakepayment` to simulate payments            |
| Interface-based decoupling | ✅     | `payment` is decoupled from actual implementation |

---

## 📄 Final Thoughts

This design pattern using interfaces is excellent for **testability**, **maintainability**, and **extensibility**. You can add new gateways (like `Stripe`, `Paytm`, etc.) without changing the main business logic.