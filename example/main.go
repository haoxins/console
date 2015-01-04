package main

import "../../console"
import "fmt"

func main() {
	fmt.Printf("1: InterceptChar - %s\n", console.InterceptChar())
	fmt.Printf("2: interceptLine - %s\n", console.InterceptLine())

	fmt.Printf("3: ReadChar - %s\n", console.ReadChar())
	fmt.Printf("4: ReadLine - %s\n", console.ReadLine())
}
