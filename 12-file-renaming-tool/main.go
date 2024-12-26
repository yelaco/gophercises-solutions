package main

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Must[T any](x T, err error) T {
	if err != nil {
		panic(err)
	}
	return x
}

func main() {
	pwd := Must(os.Getwd())

	err := filepath.Walk(pwd, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		oldName := info.Name()
		idx := strings.IndexRune(oldName, '_')
		if idx < 0 {
			return nil
		}
		fileExtensionIdx := strings.Index(oldName, ".txt")
		newName := oldName[idx+1:fileExtensionIdx] + " - " + oldName[:idx] + oldName[fileExtensionIdx:]
		return os.Rename(path, filepath.Join(filepath.Dir(path), newName))
	})
	if err != nil {
		log.Fatal(err)
	}
}
