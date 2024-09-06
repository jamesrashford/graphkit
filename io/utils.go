package io

import (
	"os"
	"path/filepath"
	"strings"
)

func GetExamples(ext string) ([]string, error) {
	var paths []string
	err := filepath.Walk("../examples",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.HasSuffix(path, ext) {
				paths = append(paths, path)
			}
			return nil
		})

	return paths, err
}
