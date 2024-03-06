package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func saveData(id string, data any) error {
	if id == "" {
		return &validationError{"Validaton error"}
	}

	if id != "Eko" {
		return &notFoundError{"Data Not Found"}
	}

	// ok

	return nil

}

func main() {

	err := saveData("Eko", nil)
	if err != nil {
		// terjadi error

		// if validationErr, ok := err.(*validationError); ok{
		// 	fmt.Println("Validation: error", validationErr.Error())
		// } else if notFoundErr, ok := err.(*notFoundError); ok {
		// 	 fmt.Println("not found error: ", notFoundErr.Error())
		// } else{
		// 	fmt.Println("unknown error: ", err.Error())
		// }

		switch finalError := err.(type){
		case *validationError:
			fmt.Println("Validation: error", finalError.Error())
		case *notFoundError:
			fmt.Println("not found error: ", finalError.Error())
		default:
			fmt.Println("unknown error: ", finalError.Error())
		}		

	} else {
		fmt.Println("Sukses")
	}

}