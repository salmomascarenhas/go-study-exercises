package main

import (
	"fmt"
)

type Greeter interface {
	Greet()
}

type Person struct {
	Name string
}

func (p Person) Greet() {
	fmt.Println("Hello, my name is", p.Name)
}

func main() {
	var person Greeter
	person = Person{Name: "Salmo Mascarenhas"}
	person.Greet()
}

// -> Interfaces in Go <-
//
// In Go, an interface is a contract that specifies a set of methods that a type must
// implement. It is a way to define a behavior that a type should have.
//
// An interface has no implementation itself, but it defines a set of method signatures
// that a struct or other type must implement.
//
// Here's an example of an interface in Go:
// type Greeter interface {
// 	Greet() string
// }
// The Greeter interface requires a type to implement a method called Greet that
// returns a string.
//
// Here's an example of a struct that implements the Greeter interface:
// type Person struct {
// 	Name string
// }
// func (p Person) Greet() string {
// 	return fmt.Sprintf("Hello, my name is %s", p.Name)
// }
// In this example, the Person struct implements the Greeter interface by providing
// an implementation for the Greet method.
//
// Interfaces can be used to achieve polymorphism in Go. You can declare a variable
// of an interface type and assign it a value of a struct that implements the interface.
//
// Here's an example:
// func main() {
// 	var greeter Greeter
// 	greeter = Person{Name: "Salmo"}
// 	fmt.Println(greeter.Greet())
// }
// In this example, we declare a variable of type Greeter and assign it a value of a
// Person struct. We can then call the Greet method on the greeter variable without
// knowing the concrete type of the value at runtime.

// -> Structs in Go
//
// In Go, a struct is a composite data type that groups together zero or more
// fields. It is used to define custom types that can encapsulate related data.
//
// Here's an example of a struct in Go:
// type Person struct {
// 	Name string
// 	Age  int
// 	Email string
// }
// In this example, the Person struct has three fields: Name, Age, and Email.
//
// Structs can be instantiated using the new keyword or by directly declaring a
// variable of the struct type.
//
// Here's an example of instantiating a Person struct:
// func main() {
// 	person := Person{Name: "John Doe", Age: 30, Email: "john@example.com"}
// 	fmt.Println(person)
// }
// In this example, we declare a variable person of type Person and initialize it
// with values for its fields.
//
// Structs can also be accessed using dot notation. For example, to access the
// Name field of the person struct, we can use person.Name.
//
// Structs can have methods associated with them. We can define methods on structs
// using a receiver, which is the struct type followed by a pointer, value, or
// interface type.
//
// Here's an example of a method on a struct:
// func (p Person) Greet() string {
// 	return fmt.Sprintf("Hello, my name is %s", p.Name)
// }
// In this example, we define a Greet method on the Person struct. The receiver is
// of type Person, which means we can call this method on a Person struct value.
//
// To call the Greet method on a Person struct value, we can use dot notation:
// person.Greet()
//
// Structs are mutable by default, which means their fields can be modified.
//
// Here's an example of modifying a struct field:
// person.Age = 31
// fmt.Println(person)
//
// Structs can also be compared for equality. Two struct values are equal if all
// their corresponding fields are equal.
//
// Here's an example of comparing two Person struct values for equality:
// person1 := Person{Name: "John Doe", Age: 30, Email: "john@example.com"}
// person2 := Person{Name: "John Doe", Age: 30, Email: "john@example.com"}
// fmt.Println(person1 == person2)
//

// -> Struct Tags
//
// In Go, struct tags are a way to attach metadata to struct fields. They are used
// to provide additional information about the field, such as how it should be
// serialized, validated, or even used in templates.
//
// Struct tags are defined using the format `key:"value"`. They are placed after
// the field type and before the field name.
//
// For example, let's define a struct with a field that has a struct tag:
// type User struct {
// 	ID uint64 `json:"id"`
// 	Name string `json:"name"`
// 	Birthdate time.Time `json:"birthdate"`
// 	IsAdmin bool `json:"is_admin"`
// 	Email string `json:"email"`
// 	Phone string `json:"phone"`
// }
// In this example, each field has a struct tag that specifies the JSON key that
// should be used when serializing the struct to JSON.
//
// The `json:"id"` tag specifies that the field should be serialized as "id" in
// JSON. The other fields have similar tags.
//
// Struct tags can be used with various packages and libraries, such as
// encoding/json, encoding/xml, and validator.

// -> Pointer Receiver Methods
// In Go, a pointer receiver allows you to modify the value of the struct that
// the method is called on. This is useful when you want to modify the state
// of the struct within the method.
//
// A pointer receiver is denoted by a pointer to the struct type. For example,
// (*User).PrintID() is a pointer receiver method for the User struct.
//
// Here's an example to illustrate the difference between a value receiver and a
// pointer receiver:
//
// 