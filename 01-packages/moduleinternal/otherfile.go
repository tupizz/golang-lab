package moduleinternal

import "fmt"

func OtherPublicFn() {
	fmt.Println("OtherPublicFn logging...")
}

func privateFuncForPackage() {
	fmt.Println("writing within a private function")
}
