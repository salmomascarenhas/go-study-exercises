package main

// Type parameters, also known as generics, are a way to allow a function or struct
// to work with multiple types. They are a way to abstract away the type of data
// a function or struct works with, allowing it to work with any type that
// satisfies the constraints specified by the type parameter.
//
// The syntax for type parameters is as follows:
// func MyFunction[T any](x T) T {
// 	// ...
// }
// or
// type MyStruct[T any] struct {
// 	x T
// }
//
// In the above examples, `T` is the type parameter, and `any` is the type
// constraint. The type constraint is a way to specify what types the type
// parameter can be. In this case, `any` is a built-in type constraint that allows
// `T` to be any type.
//
// The type parameter can also be a struct or interface type. For example:
// type MyStruct[T any] struct {
// 	x T
// 	y T
// }
// or
// type MyStruct[T interface {
// 	~string | ~int
// }] struct {
// 	x T
// 	y T
// }
//
// Type parameters can also have multiple constraints. For example:
// type MyStruct[T interface {
// 	~string | ~int
// 	~float64
// }] struct {
// 	x T
// 	y T
// }
//
// Type parameters can also have default values. For example:
// type MyStruct[T ~string | ~int] struct {
// 	x T
// 	y T
// }
//
// Type parameters can also be used in functions. For example:
// func MyFunction[T ~string | ~int](x T) T {
// 	// ...
// 	return x
// }
//
// And here's an example of using type parameters with a struct:
// type MyStruct[T ~string | ~int] struct {
// 	x T
// 	y T
// }
//
// func main() {
// 	s := MyStruct[string]{
// 		x: "foo",
// 		y: "bar",
// 	}
// 	t := MyStruct[int]{
// 		x: 1,
// 		y: 2,
// 	}
// }

// The comparable constraint is a predeclared constraint that matches any type whose
// values can be compared using the == and != operators. It is satisfied by all
// types that support comparison, including integers, floats, strings, booleans,
// pointers, arrays, slices, structs, and interfaces.
//
// For example:
// func equals[T comparable](a, b T) bool {
// 	return a == b
// }
//
// The any constraint is a predeclared constraint that matches any type. It is
// equivalent to the empty interface, interface{}.
//
// For example:
// func print[T any](x T) {
// 	fmt.Println(x)
// }
