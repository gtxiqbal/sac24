package helper

import "fmt"

func LogIfError(err error) {
	if err != nil {
		fmt.Println(err)
		err = nil
	}
}

func PanicIfError(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func RecoverIfError() {
	if err := recover(); err != nil {
		fmt.Println("panic occured:", err)
	}
}
