package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func uniq(reader io.Reader, writer io.Writer) error {
	in := bufio.NewScanner(reader)
	var prev string
	for in.Scan() {
		txt := in.Text()
		if txt == prev {
			continue
		}
		if txt < prev {
			return fmt.Errorf("file no sorted")
		}
		prev = txt
		_, err := fmt.Fprintln(writer, txt)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	err := uniq(os.Stdin, os.Stdout)
	if err != nil {
		panic(err.Error())
	}
}
