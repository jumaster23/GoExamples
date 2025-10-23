package main

import "fmt"

func main() {
	//var flags [3]string
	//flags[0] = "-v"
	//flags[1] = "-a"
	//flags[2] = "-b"

	flags := [...]string{"-v", "-a", "-b"}
	fmt.Println(flags)
}
