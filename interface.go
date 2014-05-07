package main

import (
	"fmt"
	"os"
	"strings"
)

var test_data = `
this is the test data
札幌でまたガ
  スボンベ爆発事件`

func main() {
	r := strings.NewReader(test_data)
	buf, err := NewBuffer(r)
	if err != nil {
		fmt.Printf("couldn't read buffer: %v\n", err)
		return
	}
	w := NewWindow(buf)
	w.Print(os.Stdout)

	fmt.Printf("\n\nmove c0 by 5, 20")
	w.Move(0, 40, 5)
	w.Print(os.Stdout)

	f, err := os.Create("/tmp/test")
	if err != nil {
		fmt.Printf("couldn't create file: %v\n", err)
	}
	err = w.Buf.Write(f)
	if err != nil {
		fmt.Printf("couldn't write file: %v\n", err)
	}
}
