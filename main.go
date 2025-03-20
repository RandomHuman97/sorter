package main

import (
	"fmt"
	"os"

	"github.com/h2non/filetype"
)

func main() {
	entries, err := os.ReadDir(".")
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, v := range entries {
		if !v.Type().IsRegular() || v.Name() == "sorter" {
			continue
		}

		buf, _ := os.ReadFile(v.Name())
		kind, _ := filetype.Match(buf)
		if kind.MIME.Type == "" {
			fmt.Println("skipped unknown")
			continue
		}
		new_dir := "./" + kind.MIME.Type

		err := os.MkdirAll(new_dir, os.ModePerm)
		if err != nil {
			fmt.Println(err.Error())
		}
		os.Rename("./"+v.Name(), new_dir+"/"+v.Name())

	}

}
