package hpepkg

import "fmt"

const version = "0.1.0"

func Greeting(name string) {
	fmt.Printf("Hello %s, Greetings from HPE package version v%s", name, version)
}
