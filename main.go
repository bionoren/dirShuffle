package main

import (
	"fmt"
	"io/fs"
	"math/rand"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// main takes a directory to search, followed by a list of allowed file extensions, and prints a list of filtered files in the directory, shuffled by directory.
// The original goal was to shuffle music by album, but there may be other uses.
//
// sample usage:
// go build .
// # print a list of go files in the parent directory (and all subdirectories), shuffled by directory
// ./dirShuffle $(dirname $(pwd)) ".go"
func main() {
	if len(os.Args) < 2 {
		panic("must specify an absolute path for the root directory to shuffle")
	}

	root := os.Args[1]
	fmt.Println("reading " + root)

	allowedExtensions := make(map[string]struct{}, len(os.Args)-2)
	for _, ext := range os.Args[2:] {
		// use "." to indicate files with no extension
		if ext == "." {
			ext = ""
		}
		allowedExtensions[ext] = struct{}{}
	}
	files := make(map[string][]string)
	err := fs.WalkDir(os.DirFS(root), ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if path[0] == '.' || strings.Contains(path, "/.") { // skip hidden directories and files
			return nil
		}

		dir := filepath.Dir(path)
		if !d.IsDir() {
			ext := strings.ToLower(filepath.Ext(path))
			if _, ok := allowedExtensions[ext]; ok {
				files[dir] = append(files[dir], path)
			}
		}

		return nil
	})
	if err != nil {
		panic(err)
	}

	for _, dir := range files {
		sort.Strings(dir)
	}

	keys := make([]string, 0, len(files))
	for k := range files {
		keys = append(keys, k)
	}

	rand.Shuffle(len(keys), func(i, j int) {
		keys[i], keys[j] = keys[j], keys[i]
	})

	for _, dir := range keys {
		for _, filename := range files[dir] {
			fmt.Println(filename)
		}
	}
}
