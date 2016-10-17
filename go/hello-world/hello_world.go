package hello

import "fmt"


const testVersion = 2

// HelloWorld returns a string greeting `n` name, or "world" if a name isn't provided.
func HelloWorld(n string) string {
	if n == "" {
		return "Hello, World!"
	}
	s := fmt.Sprintf("Hello, %s!", n)
	return  s
}