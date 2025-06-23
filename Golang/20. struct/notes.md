# 📘 Go Structs — Concepts Explained from Code

This document provides an in-depth explanation of **Go struct concepts** as demonstrated in the provided Go program.

---

## 📌 1. Basic Struct Declaration

Structs in Go are used to group related data together.

### 🧱 Example:

```go
type customer struct {
	name  string
	phone string
}
```

This defines a struct named `customer` with two fields:

* `name` of type `string`
* `phone` of type `string`

---

## 📌 2. Struct Embedding (Composition)

Go does **not support inheritance**, but it supports **composition** via **struct embedding**.

### 🧱 Example:

```go
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
	customer
}
```

Here, `customer` is embedded into `order`, which means:

* Fields of `customer` (like `name` and `phone`) can be accessed directly on an `order` instance.

#### ✅ Access Example:

```go
newOrder.customer.name = "robin"
```

Or simply:

```go
newOrder.name = "robin" // because of struct promotion
```

---

## 📌 3. Struct Instantiation (Initialization)

You can initialize a struct either **field-wise** or **inline embedded struct-wise**.

### 🧱 Example:

```go
newOrder := order{
	id:     "1",
	amount: 30,
	status: "received",
	customer: customer{
		name:  "john",
		phone: "1234567890",
	},
}
```

This creates an `order` and initializes its embedded `customer` in one go.

---

## 📌 4. Field Update

You can change a struct’s field even after creation:

```go
newOrder.customer.name = "robin"
```

If the field is **embedded**, you can also access it like this:

```go
newOrder.name = "robin"
```

---

## 📌 5. Anonymous Struct

Go allows defining and using **anonymous structs** (inline struct definition without a type name).

### 🧱 Example (commented in code):

```go
language := struct {
	name   string
	isGood bool
}{"golang", true}

fmt.Println(language)
```

This creates a one-time-use struct variable.

---

## 📌 6. Struct Methods (Receiver Functions)

You can associate methods with structs using **receiver functions**.

### 🔧 Pointer Receiver (for modifying original struct):

```go
func (o *order) changeStatus(status string) {
	o.status = status
}
```

Use a pointer `*order` to change the value of the original struct.

### 🧪 Example Usage:

```go
myOrder.changeStatus("confirmed")
```

---

### 🔎 Value Receiver (read-only access):

```go
func (o order) getAmount() float32 {
	return o.amount
}
```

This copies the struct and does not affect the original.

---

## 📌 7. Zero Values

When you create a struct without assigning values, its fields get **zero values**:

| Type   | Zero Value             |
| ------ | ---------------------- |
| int    | 0                      |
| float  | 0.0                    |
| string | ""                     |
| bool   | false                  |
| struct | all fields zero-valued |

### 🧪 Example:

```go
myOrder := order{}
fmt.Println(myOrder.status) // prints ""
```

---

## 📌 8. `time.Time` in Struct

The `time.Time` type from the `time` package stores timestamp values.

### 🧱 Example:

```go
createdAt time.Time
```

To set it:

```go
myOrder.createdAt = time.Now()
```

---

## 🧵 Summary

| Concept                | Used In Program? | Description                 |
| ---------------------- | ---------------- | --------------------------- |
| Basic struct           | ✅                | `customer`, `order`         |
| Struct embedding       | ✅                | `customer` inside `order`   |
| Struct initialization  | ✅                | With inline values          |
| Anonymous struct       | ✅ (commented)    | Ad-hoc structs              |
| Methods (receiver)     | ✅ (commented)    | `changeStatus`, `getAmount` |
| Pointer vs value recv. | ✅ (commented)    | To mutate or read-only      |
| Zero values            | ✅                | Shown as comments           |
| time.Time usage        | ✅                | For timestamp fields        |
