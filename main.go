package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello world")

	b := []byte("This is a test text\n")
	saveData("/Users/shalvinkumar/Desktop/personal_workspace/my-db/test.txt", b)
}

func saveData(path string, data []byte) error {
	fp, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)
	if err != nil {
		return err
	}

	defer fp.Close()

	_, err = fp.Write(data)
	if err != nil {
		return err
	}

	return fp.Sync()
}
