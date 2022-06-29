package main

import (
	"fmt"
	"os"
)

func copyFile(srcName, dstName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	written, err = io.Copy(dst,src)
	return
}

func main(){
	written, err := copyFile("myfile.txt", "copyFile.txt")
	fmt.Println(written, err)
}
