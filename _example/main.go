package main

import "github.com/goware/pp"

func main() {
	pp.Red("hihi %d", 1).Green("yes").Red("ok").White("hello").Println()
	pp.Red("hihi %d", 2).Blue("welcome").Println()

}
