package main

import "fmt"

func main() {
	music := make(map[string]string)
	music["guitarra"] = "jazz"
	music["piano"] = "classica"
	music["bateria"] = "rock"
	fmt.Println(music)
	fmt.Println(len(music))

	tech := make(map[string]string)
	tech["pc"] = "windows"
	tech["celular"] = "android"
	tech["tablet"] = "ios"
	fmt.Println(tech)
	fmt.Println(len(tech))

	delete(music, "guitarra")
	fmt.Println(music)
	fmt.Println(len(music))

	fmt.Println(tech["pc"])

	content, ok := music["piano"]
	fmt.Println(content, ok)
}
