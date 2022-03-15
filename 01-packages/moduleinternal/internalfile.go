package moduleinternal

import "fmt"

func PublicFunction() {
	fmt.Println("writing in aux package")
	privateFuncForPackage()
}
