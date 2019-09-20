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
