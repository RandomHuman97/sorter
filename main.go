package main

import (
	"errors"
	"fmt"
	"mime"
	"os"
	"strconv"
	"strings"
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

		dot_split := strings.LastIndex(v.Name(), ".")
		extension := v.Name()[dot_split+1:]
		mime := mime.TypeByExtension("." + extension)
		if mime == "" {
			fmt.Println("skipped unknown")
			continue
		}
		mime_split := strings.SplitAfterN(mime, "/", 3)
		mime = mime_split[0]
		// only split into first 1, we dont carer abt rest
		fmt.Println(mime)

		new_dir := "./" + mime
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
