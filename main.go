package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func fileinfoDemo(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(path, info.Size())
	return nil
}

func main() {

	// https://yourbasic.org/golang/list-files-in-directory/
	// I think a lot of this is builtin syntax
	// https://golang.org/pkg/path/filepath/#Walk
	// https://golang.org/pkg/path/filepath/#WalkFunc
	// path := "/Users/sherchowdhury/github-public/codingbee/golang-samples/gsg_get_filepath_list_from_directory"
	path, err := os.Getwd()

	err = filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			fmt.Println(path, info.Size())
			return nil
		})
	if err != nil {
		log.Println(err)
	}

}
