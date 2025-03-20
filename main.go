package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

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
		new_path := new_dir + "/" + v.Name()
		if _, err := os.Stat(new_path); err == nil {
			i := 0
			// 100 is absurd but we ball
			for i = 0; i < 100; i++ {

				if _, err := os.Stat(new_path + strconv.Itoa(i)); errors.Is(err, os.ErrNotExist) {
					break
				}
			}
			if i > 100 {
				continue // just skip the file to be sure
			}
			new_path = new_path + strconv.Itoa(i)
		}

		os.Rename("./"+v.Name(), new_path)

	}

}
