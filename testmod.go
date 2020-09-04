package testmod

import "fmt"

//HelloFunc returns string
func HelloFunc(name string) string {

	return fmt.Sprintf("Hello, %s", name)
}
