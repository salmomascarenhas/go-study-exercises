package main

import (
	"errors"
	"fmt"
	"math"
)





type User struct {
	foo string
}

func (u User) Foo() {
	fmt.Println(u.foo)
}

// Capturando erros.
func NewUser(wantError bool) (*User, error) {
	if wantError {
		return nil, fmt.Errorf("error")
	}
	return &User{foo: "bar"}, nil
}


// Criando os próprios erros.

type SqrtError struct {msg string}

func (s SqrtError) Error() string {return s.msg}

func raizQuadrada(x float64) (float64, error) {
	if x < 0 {
		return 0, SqrtError{msg: "raiz quadrada de um número negativo"}
	}
	// math lib does not return errors, but return NaN (not a number). Calable sentinel values.
	resultado := math.Sqrt(x)
	return resultado, nil
}

// func main() {
// 	x := 4.0
// 	resultado, err := raizQuadrada(x)

// 	if err != nil {
// 		fmt.Println(err)
// 		return 
// 	}
// 	 fmt.Println("Raiz quadrada de ", x, " = ", resultado)
// }

// Comparando tipos de erros (Errors.is e Errors.As).
var ErrorNotFound = errors.New("not found")

// func foo() error { return ErrorNotFound}
	
// Para o caso de error criado a partir do método New de errors. 
// func main() {
// 	err := foo()
// 	if err != nil && errors.Is(err, ErrorNotFound) {
// 		fmt.Println("error not found")
// 		return
// 	}
// 	fmt.Println("Caiu fora do if de error not found")
// }

//  Para o caso de error criado a partir da implementação da interface Error.


func foo() error { return SqrtError{msg: "foo"} }
 func main() {
	err := foo()
	if err != nil && errors.As(err, &SqrtError{}) {
		fmt.Println("error sqrt error")
		return
	}
	fmt.Println("Caiu fora do if de error sqrt error")
 }


// Wrapping errors
// Go provides a way to wrap errors with more information. It is done using the fmt.Errorf function.
// The fmt.Errorf function takes a format string and a list of arguments, much like fmt.Printf.
// The format string is used to create a new error message.
// The ...interface{} are the arguments to the format string.
// The errors.Unwrap function can be used to extract the original error.

// func raizQuadrada(x float64) (float64, error) {
// 	if x < 0 {
// 		return 0, fmt.Errorf("não é possível calcular raiz quadrada de um número negativo: %w", fmt.Errorf("número negativo: %g", x))
// 	}
// 	// math lib does not return errors, but return NaN (not a number). Calable sentinel values.
// 	resultado := math.Sqrt(x)
// 	return resultado, nil
// }

// func main() {
// 	_, err := raizQuadrada(-10)
// 	if err != nil {
// 		fmt.Println(err)
// 		if errors.Is(errors.Unwrap(err), ErrNegativeSqrt) {
// 			fmt.Println("O erro é o de raiz quadrada de um número negativo")
// 		}
// 	}
// }

// var ErrNegativeSqrt = errors.New("não é possível calcular raiz quadrada de um número negativo")


// Explain the errors.Join method
// Go provides a function called errors.Join that helps to join multiple errors into a single error.
// The errors.Join function takes a slice of errors and returns a new error that contains all the
// errors joined together.
//
// The errors.Join function creates a new error using fmt.Errorf and joins the errors together
// using the ": " separator. It also adds a newline character before each error message except
// for the first one.
//
// Here's an example of using errors.Join:
// err1 := errors.New("error 1")
// err2 := errors.New("error 2")
// err3 := errors.New("error 3")
//
// err := errors.Join(err1, err2, err3)
// fmt.Println(err)
// Output:
// error 1
// : error 2
// : error 3
//
// The errors.Join function can be used to create a new error that contains multiple errors
// from different sources. It can be useful when you want to collect and report multiple errors
// at once.
//

