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

//go:embed awk/github.latest.awk
var source []byte

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
	if len(source) == 0 {
		fmt.Println("did not read awk source correctly")
		return
	}

	prog, err := parser.ParseProgram([]byte(source), nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	output := new(bytes.Buffer)

	config := &interp.Config{
		Stdin:  bytes.NewBuffer(body),
		Output: output,
	}

	_, err = interp.ExecProgram(prog, config)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s\n", output)

}
