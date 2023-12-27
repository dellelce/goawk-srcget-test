package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/benhoyt/goawk/interp"
	"github.com/benhoyt/goawk/parser"
	"io"
	"net/http"
)

// later: awk/firstaftermatchword-trim.latest.awk
// go:embed awk/github.latest.awk
var f1 []byte

func download(url string) ([]byte, error) {
	resp, err := http.Get(url)
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return bodyBytes, nil
	}

	return nil, err
}

func main() {
	body, err := download("https://github.com/landley/toybox/tags")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("debug: github content:%s\n", body[:100])

	source := f1
	if len(source) == 0 {
		fmt.Println("did not read awk source correctly")
		return
	}

	fmt.Println("debug: awk file:%s", source[:100])

	prog, err := parser.ParseProgram([]byte(source), nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	config := &interp.Config{
		Stdin: bytes.NewBufferString(string(body)),
	}

	_, err = interp.ExecProgram(prog, config)
	if err != nil {
		fmt.Println(err)
		return
	}

}
