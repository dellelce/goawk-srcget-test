package main

import _ "embed"

//go:embed awk/firstaftermatchword-trim.latest.awk
var f1 []byte

func embedtest() {
	print(string(f1))
	// awk/github.latest.awk
}

func main() {
	embedtest()
}
