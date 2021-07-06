package golearn

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("hi, %v. welcom!", name)
	return message
}
