package main

import (
	"fmt"
	"package/moduleinternal"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("hi motherfucker")
	moduleinternal.PublicFunction()
	moduleinternal.OtherPublicFn()

	err := checkmail.ValidateFormat("tadeu.tupiz@gmail.com")
	if err != nil {
		fmt.Println("email is not valid")
	} else {
		fmt.Println("email is valid")
	}
}
